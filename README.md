# reflectui

Dynamic UI generation for [Go](http://www.golang.org/) based on reflection and templates 

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/infutur/reflectui/rui)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/infutur/reflectui/master/LICENSE)

## Overview
reflectui has currently one package **rui**. Package **rui** provides template functions, which dynamically convert go data structures into HTML. 

Within a go template, a call to a rui-template function like 

```go
{{ rui MyDataStructure }}
```
 
 for a variable of type
 
```go
type MyDataType struct {
	Field1 Type1
	Field2 Type2
	...
}
```
 
 will analyze the variable by go reflection, generate an (internal) abstract tree representation of the data and recursively execute template functions in a top-down, breadth-first way along the internal representation, returning a HTML partial e.g. like 
 
```html
<div class="MyDataStructureClass">
	<label>Field1</label><div> .... </div>
	<label>Field2</label><div> .... </div>
	....
</div>
```

## Releases / Status

The first version of the project is currently under progress

## Installation
To use reflectui install the packages by

```sh
$ go get github.com/infutur/reflectui/rui
```

and then import with the import path

```go
import "gitbuh.com/infutur/reflectui/rui
```

## Getting started

**Initialization**

```go
// the first step is to create a new rui - system object.
// By using "rui" as name parameter the function will return
// the predefined rui-system
ruiSys = rui.New("rui")

// next you load the rui template function into the application
// templates, optionally with building templates.  
// This has to be done before parsing templates which use the
// rui template functions
myTemplates = ruiSys.Load(appTmpl, useBuiltIn)

// you can change the default settings for BreakType
ruiSys.AddBTypes(myArrayofTypeNames)

// or any application data
ruiSys.SetData(myApplicationData)

// Finally, after setting up all application templates
// the rui system has to be attached to the final set of
// templates, which shall be accessible within the rui
// template execution
ruiSys.Attach(myFinalTemplate)
```

## Reflection & template execution

The internal representation is as follows

```go
// The structure of the internal representation
type Node struct {
	// Name of a struct field
	// or type name for an embedded type
	// or key of a map, as string
	// or index of a array/slice element, as string
	// or name of the wrapping root element
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
```

Template names for execution are build by

- a prefix & `Style` & (`Scheme` or `Type` or `Kind`)

if for a given `Style` no template with postfix `Scheme` exists, a template with postfix `Type` is looked up, if not exists a template with postfix `Kind`, a.s.o. If no template for a `Style` exists at all, the same procedure with `"std"` as `Style` is repeated.   

## Tags

The reflection and template lookup process can be influenced by go struct tags

```go
type MyDataType struct {
	Field1 Type1 `rui:ignore, scheme=form, break`
	...
}
```

`ignore`: to omit a struct field from the internal structure  
`break`: to stop deeper reflection and treat Type1 as a unstructured (scalar) type  
`scheme=< name >`: to set Scheme to < name >  (default: equal to `Type`) 

## Data access within templates



## BuiltIn templates & CSS Files

## Restrictions
Recursive/circular data structures can not be handled and will currently generate a stack overflow. To avoid this the struct-tag ignore can be used (see tags)

## References