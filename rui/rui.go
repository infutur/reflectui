//go:generate vfsgendev -source="github.com/infutur/reflectui/rui".Assets

// Package reflectui implements a UI generator based on golang reflection and
// golang templates
//
// reflectui transforms golang data structures to an internal representation and
// schedules template execution based on this representation
package rui

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// The structure of the internal representation
type Node struct {
	// Name of a struct field
	// or type name for an embedded type
	// or key of a map, as string
	// or index of a array/slice element, as string
	// or rootname for the wrapping root element
	Name string
	// Label of the node, by default the Name, can be set by struct tag
	// `rui:"label=nicer Name",...`  ,
	// used as human readable label e.g. in a form ....
	Label string
	// Scheme, Type and Kind are used to select the template
	// correspnding to the node.
	//
	// Scheme has highest precedence, can be set by
	// struct tag `rui:"scheme=BlogList",...`
	Scheme string
	// Type is the reflected value / struct field has no type name
	// e.g.   SubNodes []*Node
	// Type will be set to Name
	Type string
	// Kind is the golang kind of the value, i.e. struct, ptr, slice, ...., uint8
	Kind string
	// Style is used to select templates. Style is set by the call of the
	// refelctui template functions
	Style string
	// Each node gets a "unique" id, usable for html ids. The id generation
	// is done in a topdown traveral of the structure. The id generation
	// function can be set via seetings
	Id string
	// Each node has a level - relative index, useable for id generation
	Idx int
	// The golang value correspnding to the node
	Value interface{}
	// The subnodes of a struct or the elemens of an array / slice
	SubNodes []*Node
	// Each subnode is mapped by it's name (jey) into the field map, i.e.
	// each element can be accessed in a template by .Fields.<key>
	Fields map[string]interface{}
	// Application speficic data, accessible in every template by
	// .App.<xxx>
	App interface{}
	// The reflectui system structure
	Rui interface{}
	// The "final" TemplateName after template resolution, corresponding to the
	// node.
	TemplateName string
}

//--------------------------------------------------

const (
	_ruiBuiltin  = "rui"
	_ruiRootKind = "root"
	_ruiStdStyle = "std"
)

// setting for a reflectui system
type RuiSettings struct {
	// prefix for all template names used by the system
	TemplatePrefix string
	// tag for struct tags
	TagPrefix string
	// prefix for unique id="" generation
	IdPrefix string
	// template start and end delimeter to avoid conflicz with other
	// templateing systems, e.g. vuejs
	StartDelim string
	EndDelim   string
	// Names for reflectUI Template functions
	RuiFuncName   string
	NodeFuncName  string
	NodesFuncName string
	// set of type names which are "builtin", i.e. reflection will stop
	// as the type level.
	// Example:
	// golang's time.Time is a struct but as BType not further reflected,
	// i.e. the sub structure is hidden
	BTypes []string
	// generator function for unique Id's of node to be used in HTML
	// parts of templates.
	IdGenerator RuiIdGenerator
}

// a id generator function takes a prefix and a node and generates an id for the
// node. The ig generator function is triggered top down, depth first
// Any node field can be used, each node has a unique index relative to it's encosing
// struct, map, slice, ..
type RuiIdGenerator func(string, *Node) string

// Default id generation function, generates unique id's like
// <prefix>_0, <prefix>_1_3, <prefix>_1_3_0, ....
func _ruiIdGenerator(prefix string, n *Node) string {
	return prefix + "_" + strconv.Itoa(n.Idx)
}

// Default settings for a new rui system. To be copied and adapted
var RuiDefaultSettings = RuiSettings{
	TemplatePrefix: "rui",            // prefix used by builtin templates
	TagPrefix:      "rui",            // tag used for struct tags
	IdPrefix:       "rui",            // used for unique div, .. id="" generation
	StartDelim:     "%{",             // to avoid conflicts with e.g. vuejs
	EndDelim:       "}%",             // to avoid conflicts with e.g. vuejs
	RuiFuncName:    "rui",            // template function for reflection
	NodeFuncName:   "ruinode",        // template function 1 with "rui templates"
	NodesFuncName:  "ruinodes",       // template function 2 with "rui templates"
	BTypes:         []string{"Time"}, // time.Time, ... are builtin
	IdGenerator:    _ruiIdGenerator,  // standard id generator function
}

type RuiSystem struct {
	name string
	RuiSettings
	templates *template.Template
	data      interface{}
}

// the builtin RUI system for access by RUI function
var _ruiSystems = make(map[string]*RuiSystem)
var _ *RuiSystem = New(_ruiBuiltin)

// Init sequence
// New(myName)
// .Settings(mySettings)
// .AddBTypes(myBuiltInTypes)
// .SetData(myApplicationData)
//
// To add reflectui functions and builtin templates and using (attaching) the
// resulting template in (to) reflectui, has to be done before any further
// template using reflectui functions is loaded
// Load(InMyTemplates, use-BuiltIn
//
// to use other templates in the refecltui system
// Attach(myTemplate)

// Predefined special types
// RuiStyleData can be used to "Style" Data with builtn template
// rui_std_RuiStyleData: -> ruinode .Fields.Style .Fields.Data
type RuiStyleData struct {
	Style string
	Data  interface{}
}

// New creates a new reflectui system with name <name>, default settings and
// returns it. The system can be access by template functions
//
// - rui                   within using regular templates
// - ruinode, ruinodes     within templates executed by the reflectui system
//
// if there is already a system with <name>, no changes are applied
// and the existing system is returned
func New(name string) *RuiSystem {
	if _, exists := _ruiSystems[name]; !exists {
		rui := &RuiSystem{name, RuiDefaultSettings, nil, nil}
		_ruiSystems[name] = rui
	}
	return _ruiSystems[name]
}

func (ruisys *RuiSystem) Setting(settings RuiSettings) *RuiSystem {
	ruisys.RuiSettings = settings
	return ruisys
}

func (ruisys *RuiSystem) AddBTypes(bTypes []string) *RuiSystem {
	ruisys.BTypes = append(ruisys.BTypes, bTypes...)
	return ruisys
}

func (ruisys *RuiSystem) SetData(data interface{}) *RuiSystem {
	ruisys.data = data
	return ruisys
}

func (ruisys *RuiSystem) Load(appTmpl *template.Template, useBuiltIn bool) *template.Template {
	if appTmpl == nil {
		appTmpl = template.New(ruisys.name).Delims(ruisys.StartDelim, ruisys.EndDelim)
	}

	appTmpl = appTmpl.Funcs(template.FuncMap{
		ruisys.RuiFuncName: _rui,
	})

	if useBuiltIn {
		if f, err := Assets.Open("/templates"); err != nil {
			log.Println(err)
		} else {
			defer f.Close()
			if dirFiles, err := f.Readdir(0); err != nil {
				log.Println(err)
			} else {
				for _, dirFile := range dirFiles {
					if !dirFile.IsDir() {
						if fTmpl, err := Assets.Open("/templates/" + dirFile.Name()); err != nil {
							log.Println(err)
						} else {
							defer fTmpl.Close()
							buf := new(bytes.Buffer)
							if _, err = buf.ReadFrom(fTmpl); err != nil {
								log.Println(err)
							} else {
								if _, err = appTmpl.Parse(buf.String()); err != nil {
									log.Println(err)
								}
							}
						}
					} // else ignore
				}
			}
		}
	}

	ruisys.templates = appTmpl
	return appTmpl
}

func (ruisys *RuiSystem) Attach(tmpl *template.Template) *template.Template {
	ruisys.templates = tmpl
	return tmpl
}

// Template functions
//---------------------------------------------------------------------------------
func _ruiRootExt(system, style string, val interface{}) template.HTML {
	if val == nil {
		log.Println("reflectui/rui: execution of nil value")
		return template.HTML("")
	}

	ruisys, found := _ruiSystems[system]
	if !found {
		log.Println("reflectui/rui: execution of unkwown system:", system)
		return template.HTML("")
	}

	if style == "" {
		style = _ruiStdStyle
	}

	// the root node is a wrapper node for the actual value
	rootNode := &Node{
		Name:   ruisys.name + "_" + _ruiRootKind,
		Type:   ruisys.name + "_" + _ruiRootKind,
		Label:  ruisys.name,
		Kind:   _ruiRootKind,
		Fields: make(map[string]interface{}),
		Scheme: ruisys.name + "_" + _ruiRootKind,
		Style:  style,
	}

	rootNode.SubNodes = ruisys.ruiReflectNode(
		reflect.StructField{Name: rootNode.Name},
		reflect.ValueOf(val))

	ruisys.ruiNodeMap(0, ruisys.IdPrefix, rootNode)

	rootNode.TemplateName = ruisys.ruiTemplateLookup(style, rootNode)

	var b bytes.Buffer
	if err := ruisys.templates.ExecuteTemplate(&b, rootNode.TemplateName, rootNode); err != nil {
		return template.HTML(err.Error())
	}
	return template.HTML(b.String())
}
func _ruiRootSys(system, style string, v interface{}) template.HTML {
	return _ruiRootExt(system, style, v)
}
func _ruiRoot(style string, v interface{}) template.HTML {
	return _ruiRootExt(_ruiBuiltin, style, v)
}

func _ruiNodeSys(system, style string, nodes ...*Node) template.HTML {
	if nodes == nil || len(nodes) == 0 {
		return template.HTML("")
	}

	ruisys, found := _ruiSystems[system]
	if !found {
		log.Panic("templates for", system, "not found")
	}

	if style == "" {
		style = "std"
	}

	var res string = ""
	for _, node := range nodes {
		node.Style = style
		node.TemplateName = ruisys.ruiTemplateLookup(style, node)

		var tmp bytes.Buffer
		if err := ruisys.templates.ExecuteTemplate(&tmp, node.TemplateName, node); err != nil {
			return template.HTML(err.Error())
		}
		res = res + tmp.String()
	}

	return template.HTML(res)
}
func _ruiNodes(style string, nodes []*Node) template.HTML {
	return _ruiNodeSys(_ruiBuiltin, style, nodes...)
}
func _ruiNode(style string, node *Node) template.HTML {
	return _ruiNodeSys(_ruiBuiltin, style, node)
}

func _rui(style string, val interface{}) template.HTML {
	if val == nil {
		log.Println("reflectui/rui: execution of nil value")
		return template.HTML("")
	}

	if nval, ok := val.(*Node); ok {
		return _ruiNode(style, nval)
	} else if nval, ok := val.([]*Node); ok {
		return _ruiNodes(style, nval)
	} else {
		return _ruiRoot(style, val)
	}
}

func (ruisys *RuiSystem) _isBType(t string) bool {
	for _, bType := range ruisys.BTypes {
		if bType == t {
			// log.Println("Break Type", field.Kind, field.Type)
			return true
		}
	}
	return false
}

type ruiTags struct {
	Ignore  bool
	NoEmbed bool
	Label   string
	Scheme  string
	Break   bool
}

func (ruisys *RuiSystem) NodeTags(sf reflect.StructField) (rtags ruiTags) {
	if tagString := sf.Tag.Get(ruisys.TagPrefix); tagString == "" {
		return
	} else {
		tags := strings.Split(tagString, ",")
		for _, tag := range tags {
			// remove leading or trailing blanks
			ttag := strings.TrimSpace(tag)

			// ignore discard the node from the reflection tree
			if ttag == "ignore" {
				rtags.Ignore = true
				return // return no further tag processing
			}

			// label=name set the label of the node to name
			if strings.HasPrefix(ttag, "label=") {
				rtags.Label = strings.Split(ttag, "=")[1]
				continue
			}

			// scheme=name set the scheme of the node to name
			if strings.HasPrefix(ttag, "scheme=") {
				rtags.Scheme = strings.Split(ttag, "=")[1]
				continue
			}

			// embed - all subnodes of the will be added at the same level as
			// node and node will be discarded from the tree for embedded types
			if ttag == "noembed" {
				rtags.NoEmbed = true
				continue
			}

			// break sets the kind of the node to the type name of the node and
			// does not reflect the substructure, has precedence over
			if ttag == "break" {
				rtags.Break = true
				continue
			}
		}
	}

	return

}

func (ruisys *RuiSystem) ruiReflectNode(sf reflect.StructField, v reflect.Value) []*Node {
	// if StructField has ignore tag return no nodes
	tags := ruisys.NodeTags(sf)

	if tags.Ignore {
		return []*Node{}
	}

	// Ptr and Interface elements
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			// omit nil - ptr or  nil interfaces
			return []*Node{}
		} else {
			// handle same struct element with *value
			return ruisys.ruiReflectNode(sf, v.Elem())
		}
	}

	// non Ptr or Interface Elements
	t := v.Type()
	tName := t.Name()
	if tName == "" {
		tName = sf.Name
	}

	node := Node{
		Name:   sf.Name,
		Type:   tName,
		Fields: make(map[string]interface{}),
		Label:  sf.Name, // default
		Kind:   v.Kind().String(),
		Scheme: tName, // type name is default
	}

	if v.CanInterface() {
		node.Value = v.Interface()
	}

	// override default woth Tags if given
	if tags.Label != "" {
		node.Label = tags.Label
	}

	if tags.Scheme != "" {
		node.Scheme = tags.Scheme
	}

	if ruisys._isBType(node.Type) || tags.Break {
		node.Kind = node.Type
		return []*Node{&node}
	}

	// log.Println("_ruiReflectNode:", sf.Name, tName, v.Kind())
	switch v.Kind() {
	case reflect.Struct:
		subNodes := []*Node{}
		for i := 0; i < v.NumField(); i++ {
			subNodes = append(subNodes, ruisys.ruiReflectNode(t.Field(i), v.Field(i))...)
		}

		if sf.Anonymous && !tags.NoEmbed {
			return subNodes
		}

		node.SubNodes = subNodes
	case reflect.Map:
		if keys := v.MapKeys(); len(keys) != 0 {
			// log.Println("reflect.Map: number of keys:", len(keys), sf.Name, tName)
			for _, key := range keys {
				keyVal := v.MapIndex(key)
				// the key is mapped to the Field Name
				if key.CanInterface() {
					keyString := fmt.Sprint(key.Interface())
					keyField := reflect.StructField{Name: keyString}
					// log.Println("reflect.Map: key", keyString, keyVal)
					node.SubNodes = append(node.SubNodes, ruisys.ruiReflectNode(keyField, keyVal)...)
				}
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			keyString := fmt.Sprint(i)
			keyVal := v.Index(i)
			keyField := reflect.StructField{Name: keyString}
			node.SubNodes = append(node.SubNodes, ruisys.ruiReflectNode(keyField, keyVal)...)
		}
	case reflect.Bool, reflect.Int, reflect.Uint, reflect.String:
	default:
		node.Kind = "std" // until supported
	}

	return []*Node{&node}
}

func (ruisys *RuiSystem) ruiNodeMap(idx int, id string, node *Node) {
	node.App = ruisys.data
	node.Idx = idx

	node.Id = ruisys.IdGenerator(id, node)

	for i, _ := range node.SubNodes {
		snodePtr := node.SubNodes[i]
		node.Fields[snodePtr.Name] = snodePtr
		ruisys.ruiNodeMap(i, node.Id, snodePtr)
	}
}

func (ruisys *RuiSystem) ruiTemplateLookup(style string, node *Node) (templateName string) {
	// a small wrapper for loggin puposes
	templateName = ruisys._ruiTemplateLookup(style, node)

	lookUpLog := false
	if lookUpLog {
		log.Println(fmt.Sprintf("%s_%s_%s >> %s", ruisys.TemplatePrefix, style, node.Scheme, templateName))
	}

	return
}

func (ruisys *RuiSystem) _ruiTemplateLookup(style string, node *Node) (templateName string) {
	// <system>_<style>_<scheme>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, style, node.Scheme)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<style>_<type>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, style, node.Type)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<style>_<kind>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, style, node.Kind)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<style>_default
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, style, "default")
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<std>_<scheme>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, "std", node.Scheme)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<std>_<type>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, "std", node.Type)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<std>_<kind>
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, "std", node.Kind)
	if ruisys.templates.Lookup(templateName) != nil {
		return
	}

	// <system>_<std>_default
	templateName = fmt.Sprintf("%s_%s_%s", ruisys.TemplatePrefix, "std", "default")
	// no further lookup
	return
}
