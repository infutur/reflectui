// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package rui

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 6, 13, 28, 17, 964859438, time.UTC),
		},
		"/css": &vfsgen۰DirInfo{
			name:    "css",
			modTime: time.Date(2018, 11, 6, 13, 28, 17, 963729271, time.UTC),
		},
		"/css/reflectui.css": &vfsgen۰CompressedFileInfo{
			name:             "reflectui.css",
			modTime:          time.Date(2018, 12, 30, 9, 56, 43, 12550669, time.UTC),
			uncompressedSize: 428,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x50\xcd\x0a\xc2\x30\x0c\x3e\xaf\x4f\x91\xa3\x8e\xd5\xe9\x75\x03\x5f\x42\xef\xa3\xda\xa8\x85\xda\x49\x9a\x4e\xa4\xec\xdd\x65\x45\x98\xb0\x21\x33\x39\xe5\xfb\x09\x1f\x5f\x99\xcb\xc5\x03\x07\xd6\xf0\x5b\x92\x97\x20\x36\x14\x4c\xe3\x59\x37\xd4\xb6\x5c\x8c\xa7\x67\x0a\x67\x86\x28\xb2\xbb\xa2\xab\x71\x15\xec\xb6\x90\xb6\x16\xbd\x18\x75\x1a\x2f\x2a\x58\x06\x6d\x3a\x88\xb3\x8c\x55\x27\xb4\x51\x64\x4f\xa3\xf9\x56\x41\xa7\x68\x25\xe5\x20\x48\x44\x93\xe0\xf5\xfc\x53\xe3\x1e\x21\x65\x98\x78\x13\xf3\xed\xfd\xab\x9a\x23\x21\x2e\xed\x86\x09\xf1\xd3\x46\x31\x41\x60\x0f\xc1\x0e\x01\xad\xf1\x2c\x3d\xbf\x2c\x56\xae\x75\x58\x8b\xfe\x1d\x00\x00\xff\xff\x95\x10\xc6\xff\xac\x01\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2018, 12, 30, 9, 38, 12, 384207211, time.UTC),
		},
		"/templates/rui_form.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rui_form.tmpl",
			modTime:          time.Date(2018, 12, 30, 9, 40, 17, 225986308, time.UTC),
			uncompressedSize: 1040,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xcf\x6a\xe3\x30\x10\xc6\xcf\xd6\x53\xcc\x9a\xd5\x31\xce\x3d\x38\x3a\x2d\x81\x85\x36\x85\x24\xf4\x1a\xc6\xd6\xa4\x15\xc8\x7f\x90\xac\x80\x31\x7e\xf7\x22\xd9\x4e\x48\x1b\x97\x40\x7c\x92\xbf\x99\xd1\xfc\xf4\xcd\xa4\x7f\x16\x0b\xd8\xbc\xed\x5e\xe1\x40\x45\xad\xb1\x21\x0b\x8b\xe7\x3e\xc1\x18\xef\x40\xd2\x49\x95\x04\xb1\x71\xea\x78\xaa\x4c\x71\x34\x55\xd5\xc4\x3d\x67\xa9\x54\x67\xc8\x35\x5a\xbb\xfe\x16\x14\x2c\xe2\x1d\xfc\xb5\x4d\xab\x09\x56\x6b\x48\xf6\xe1\xd4\xf3\xa0\x1b\x2c\x3f\x08\x92\xbd\xcb\xb6\x95\x24\x1b\xe4\xa0\x3b\x35\xd5\xdc\x06\x79\x07\x54\xca\x9e\x47\x2c\x5d\x4a\x75\x16\x6c\x12\xee\xf3\xd9\xc6\xb8\x3c\x10\xce\x40\xdc\x05\x1f\xab\x04\x8b\xd2\x4c\xf0\x0e\x92\x2d\x16\x3e\x3b\x5d\x66\x62\xe5\xff\x0f\x6d\x7d\x7d\xc2\x7d\xd4\x88\x01\x3c\x84\x58\x60\x3d\xf2\x35\xe3\xb0\x7e\xf2\x43\x02\x43\xca\x6f\x4f\xd5\x2a\xa7\xd9\x59\x0c\x51\x71\x21\x1e\x1d\xb8\x25\x7e\x88\x17\x8d\xc1\x76\xb6\xcf\x10\x0d\x77\x3c\xd7\x46\xd2\x09\x9d\x9e\x5f\xae\x29\xee\x87\x64\x6b\x2c\x05\x8b\x00\x00\xa2\x54\x63\x46\x7a\x3e\x9d\x77\xc9\x8b\xcf\xf0\xd3\x0c\xa9\x97\x42\x55\xd6\xae\x99\x2d\x04\x25\xd7\x31\xef\x92\xff\xb2\xe7\x31\xd4\x1a\x73\xfa\xac\xb4\x24\x13\x54\xbf\x10\x5e\x3f\xa3\x76\xe4\x15\x48\xde\xfd\x11\xbc\x68\x08\x65\x55\xea\x16\x04\xf3\x7d\x96\x03\xee\x43\x2e\xec\x9c\x0a\x0e\xfe\xc3\x06\xc7\x2d\x09\xb6\x6e\x14\x69\x69\x07\x77\xc7\x4e\x93\xe6\x53\x83\xcd\xd7\x9b\xbf\x02\x00\x00\xff\xff\x76\x9f\x58\x9e\x10\x04\x00\x00"),
		},
		"/templates/rui_internal.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rui_internal.tmpl",
			modTime:          time.Date(2018, 11, 26, 10, 7, 9, 666917867, time.UTC),
			uncompressedSize: 645,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\x3f\x6b\xc3\x30\x10\xc5\x67\xfb\x53\x5c\x93\x1c\xb4\xd0\xc4\xbb\x51\x0c\x1d\x3a\x84\x06\x0f\xb5\xe9\x1a\xe4\xea\xd2\x0a\x14\x35\x38\x56\x70\x30\xfa\xee\x45\x7f\x4c\x63\xaf\xf5\xe2\xc7\xfd\xde\x3d\x49\x77\xec\x61\xbd\x86\x5d\x59\xbf\xbe\x97\x2f\x7b\xa8\xe9\x74\x56\xbc\xa3\x0b\xac\xff\xf7\x15\x29\x0e\x20\xe8\x28\x35\xc1\xa2\x35\xf2\x20\x75\x47\xad\xe6\xea\x20\xe8\xc8\x8d\xea\x16\x16\x53\x26\xe4\xb5\x48\x01\x98\x51\x45\x9a\x30\x25\x8b\x92\x9f\x28\x07\x1c\x60\xe3\x14\x58\x64\x99\x92\x91\xed\x79\x43\x2a\x40\x2f\xa7\xb4\xbe\x9d\x63\xa7\x53\x53\xf6\x26\xb5\x08\xcc\xa9\x29\xab\xba\x9b\x8a\x8d\x5e\xce\xe8\xe7\x37\x8d\x37\x0a\x7a\xca\x77\x31\x77\x27\xe6\xf5\x7e\x04\xfd\xec\x9e\x71\xc2\x7f\x2f\xbd\xaf\x8c\xde\x24\x98\x3f\xb8\x32\xd1\xe5\xe5\x1d\xc6\x01\x56\x8a\x34\xe4\x5b\x70\xbf\x4d\x65\x9a\xf2\x47\xd0\x05\x2c\x86\xde\x25\x54\xa6\xd1\xae\xe4\x03\xbc\x79\xd2\xde\x72\xfd\x45\xb0\x92\x5a\x50\xff\xbc\xba\xfa\xfc\x7c\x3b\x4b\x0a\x53\x08\x41\xf0\xb8\xc4\x21\xf8\x2d\x42\x06\x38\xb8\x50\x8b\x4f\x79\x7c\x9f\x4f\x35\x72\x9c\xe4\x06\x2c\x86\xa3\x48\x0b\x8b\x49\xe2\x56\x9d\xb9\x5d\xb3\xcc\x6f\x7e\x24\xe9\x6f\x00\x00\x00\xff\xff\x98\x99\xb8\xc7\x85\x02\x00\x00"),
		},
		"/templates/rui_std.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rui_std.tmpl",
			modTime:          time.Date(2018, 12, 30, 9, 54, 50, 26099042, time.UTC),
			uncompressedSize: 794,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcf\x6a\xf3\x30\x10\xc4\xcf\xd1\x53\xec\x17\x3e\x1d\xad\xdc\x83\xe2\x53\xe8\xa9\xf4\x50\x87\x5e\xc3\xda\xda\x14\x81\x62\x1b\xfd\x09\x04\xe3\x77\x2f\x92\x65\xda\xb4\x49\x5b\x68\x7c\x32\xb3\xda\xf1\xcf\x33\x92\xff\x8a\x02\xaa\xdd\x16\x76\x74\xec\x0d\x7a\x72\x50\xfc\xed\x29\x19\xe3\x03\x28\x3a\xe8\x96\x60\x69\x83\xde\x3b\xaf\xf6\xb6\xeb\xfc\x72\xe4\x4c\x2a\x7d\x82\xc6\xa0\x73\x9b\xcb\x59\xc9\x16\x7c\x80\xff\xce\x9f\x0d\xc1\x7a\x03\xa2\x4a\x6f\x23\x4f\xba\xc5\xf6\x95\x40\x54\xa1\x7e\xea\x14\xb9\x24\x27\x3d\xe8\x79\xe7\x72\xc8\x07\xa0\x56\x8d\x7c\xc1\xe4\x4a\xe9\x53\xc9\x66\xe1\x2a\x9d\xf3\x36\x34\x89\xef\x06\xc3\x35\xec\xbc\x14\xad\xc5\x23\xd6\x64\x46\xbe\x96\xb5\x4d\x9f\xba\xc1\xf5\x3b\x9a\x23\xf6\x19\xc5\xe7\x52\xbe\x90\x82\x80\xe9\xc4\x37\xff\x64\x74\x43\x77\xf0\x41\x6b\xf1\x7c\x07\x1f\x45\x07\x0c\x66\xba\x04\xae\xc7\xf6\x73\x9c\xf3\xbc\xfc\x10\x27\xf0\x01\xc4\x0b\x9a\x10\x4b\x88\xd9\xca\x55\x5c\xfd\x21\xbf\xe7\xa0\x53\x71\x5b\xf4\x98\xc1\x63\x1f\xe2\x41\x93\x51\x6e\x2a\x35\xbb\xce\x5a\x3c\x9a\xfa\x79\x37\x7e\x0b\x00\x00\xff\xff\x8b\xc6\x43\x14\x1a\x03\x00\x00"),
		},
		"/templates/rui_table.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rui_table.tmpl",
			modTime:          time.Date(2018, 11, 26, 10, 8, 29, 126693315, time.UTC),
			uncompressedSize: 3452,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\xdf\x8f\x9c\x36\x10\xc7\x9f\xe1\xaf\x98\xba\xe2\xad\x2c\xdd\x1f\xf7\x52\x01\x52\xa4\xaa\x6a\xd5\x24\x8a\x7a\xa7\xbe\x9e\x0c\x0c\xac\x15\x83\xa9\x31\x77\x3d\xad\xf6\x7f\xaf\xc6\xfc\x58\xb8\x0d\xc9\x66\x57\xa9\xb2\x2f\xd8\xeb\x61\xfc\x99\xef\xcc\xd8\x78\x07\xc8\x30\x17\x15\x02\xd3\xad\x78\x34\x3c\x91\xf8\xa8\x95\x32\xec\xe8\xb9\x61\x26\x9e\x62\x37\xcc\x95\x2e\x41\x64\x11\xf3\x0e\xab\xf7\xbc\xc4\xa3\xc7\x62\x37\xfc\xc1\xf7\xc1\xe5\xa9\x11\xaa\x8a\x58\xd0\x0d\x1e\x6b\x5e\xe0\xaa\xde\xd7\x0c\x4a\x34\x7b\x95\x45\xac\x40\xc3\x00\x06\x3b\x0f\x0e\xb0\x7a\x63\x27\x70\x04\xef\x64\x66\x17\xde\xd9\x89\x5d\x70\x7d\x3f\x76\x1d\xef\x00\x9a\x57\x05\xc2\xea\xbe\x4d\xde\xab\x0c\x1b\x38\x7a\xae\xb3\xf8\xbf\x5d\x68\x05\x30\x1b\x07\x83\x15\x1c\x3d\xa7\xb3\xc7\x2a\x3b\x7a\x8e\xe3\x8e\x63\xd7\x09\x93\xd6\x18\x55\x81\x79\xa9\x31\x62\x4d\x9b\x94\xc2\x30\x48\x25\x6f\x9a\x88\x25\xa6\x82\xc4\x54\x7e\xad\x45\xc9\xf5\x0b\x8b\xc7\xe0\xc3\xa0\x7b\x2f\x76\x1c\xa7\x97\x81\x14\x8a\x18\x03\x7a\x2e\x86\x4a\x8b\x9f\x0d\x37\x0c\xc8\x84\x9e\x56\xf7\x11\xd4\xfd\x64\x92\x32\xcc\x79\x2b\x6d\x9e\x1c\x4a\xd4\x00\x4e\x3e\xfc\x42\xab\xb6\x06\xad\x9e\x59\xec\x02\x00\x38\xa1\xe4\x09\xca\xc1\x26\x55\xd2\x6f\x4a\x7f\x03\x34\xb0\x2f\xd8\x65\xcb\x48\x69\x86\xd5\x1f\x19\x50\x9a\x69\xf8\x96\x96\x28\x6c\x6b\x13\xbb\xce\x6c\xbb\xde\xd5\xfa\x67\x16\xbb\x4e\xb7\x95\xa8\xea\xd6\xc0\x8c\x27\x55\x95\xd1\x4a\xb2\xbe\x8c\x06\xff\x50\x4b\x9e\xe2\x5e\xc9\x0c\x75\xc4\x7e\xed\x22\xfa\x05\xc8\xe0\x4f\x51\x91\xc9\x4f\x76\xf2\xf0\x52\xe3\x69\x82\x65\x2d\xb9\x41\xca\x86\x75\xf2\xc4\x65\x8b\x9d\xdb\xbf\x69\x68\xc9\x09\xb3\xd3\x91\x90\x2e\x52\x34\x51\x4a\x7e\x33\x39\x6f\x54\xb3\xab\xd1\x74\x8f\xe9\xc7\x44\xfd\xcb\x2e\x12\xb7\xd7\xc5\xe8\x16\x19\x78\x07\x91\x9f\xf4\xb1\x8e\x30\xf3\x0e\x56\x8f\x98\xfa\xe2\xab\xe5\x6a\x8c\x6e\xd3\x6f\x57\x7f\x33\xc1\xe0\x02\xc5\x5e\x85\x1f\x0f\xe7\xc1\xfc\xa0\xe8\xce\x89\xa5\x83\x62\x76\x52\xf4\x3a\x50\x9b\x5f\x2c\xca\x6f\x42\xe2\x07\x6e\xf6\xdf\x77\x1d\xe5\x82\x82\xfe\x8a\x1a\x9a\xf7\xd6\x55\xe5\xf2\xc0\x8b\xb7\xa2\xf9\x9f\xea\xe5\x24\x8c\x33\x28\xf3\x2a\xb0\xde\x73\xc5\x9f\x12\xae\x67\x72\xcd\x95\x3c\x0f\xc0\x9a\xcc\x6c\xfa\x95\xdf\x91\x67\xa8\xfb\xf5\x65\x83\x07\x61\x24\xbe\x62\xb5\x02\x76\xaf\xcd\x2e\xa5\x6e\x32\xbd\x94\x48\x7d\xc7\x26\xc0\x5a\x73\xd8\x6b\xcc\x23\xf6\x23\x8b\x43\x51\x16\xd0\xe8\x94\x6e\xe3\xba\x0e\x44\x59\x04\xb5\x6c\x1b\xff\x79\x55\x57\x05\x03\x2e\x4d\xc4\x58\x1c\x06\x7c\xd8\x69\x72\x8f\x0d\x1e\x27\x24\x61\x2b\x7b\xc3\x53\x13\x0d\x25\xd0\x3b\x90\xc2\x5a\xc0\xe4\x37\x02\x11\xcf\x98\x62\x0e\x39\xf7\x13\xae\x1b\x7a\xca\x82\x28\x84\x15\xc0\xf6\x35\x9f\x79\x09\x03\xeb\xd6\x56\xc1\xd0\x8f\xc3\x96\x61\xd0\x43\x75\x9c\xf0\xf9\x32\xfc\x74\x1d\xde\x1b\xae\xff\xe2\x46\x54\xc5\xf7\x5d\x8a\xd0\xdd\x61\xb9\x40\x99\x35\x68\xce\xeb\x71\x12\x48\xa7\xc9\xb4\xc3\x35\xcf\x84\x1a\xba\xd9\x22\xf9\x77\x0c\x2a\x5e\xe2\xf4\x33\x6e\x68\xef\x3b\xba\x1f\x40\xe4\x80\xff\x0c\x49\xbe\x83\xa3\x07\xfd\x35\x01\x63\x16\x20\x88\xa7\x72\x40\x04\x2c\x6f\xe5\x29\xfa\x71\x27\x43\x45\x1e\xb1\x37\xcf\xd8\xa8\x12\xc1\x87\x3b\x68\x0c\xd7\x0d\xa5\x7e\x3c\xb1\xbe\x88\xbc\x5b\x46\xde\x9d\x23\xef\x6e\x40\xde\x8d\xc8\x1f\x34\x1a\xf3\x02\x85\x52\x19\xf8\xb0\xbb\x06\x7b\xbb\x8c\xbd\x3d\xc7\xde\xde\x80\xbd\x1d\xb1\xdf\xe1\x1e\x7c\xd8\x5e\x83\xbb\x59\xc6\xdd\x9c\xe3\x6e\x6e\xc0\xdd\x8c\xb8\xf4\x79\xc7\x21\xe1\xa4\xf1\xe6\x1a\xe8\xf5\x32\xf4\xfa\x1c\x7a\x7d\x03\xf4\x7a\x84\xbe\x6f\xd3\x8f\x0d\x24\xa2\x00\x23\x6c\x51\xaf\x2d\xf9\x0c\x3c\x0c\x86\x8e\xfd\xe2\x35\xf9\x5f\x00\x00\x00\xff\xff\x78\x34\xa6\xdf\x7c\x0d\x00\x00"),
		},
		"/templates/rui_tree.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rui_tree.tmpl",
			modTime:          time.Date(2018, 11, 26, 10, 8, 15, 63905048, time.UTC),
			uncompressedSize: 550,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x6a\xc3\x30\x10\x44\xcf\xd6\x57\x4c\x03\x3a\xb5\xb6\xef\x45\xd5\x27\xe4\x92\xd2\x4b\x08\x41\xb1\xd6\x20\x50\x6c\xa3\x95\x5a\x82\xf1\xbf\x17\xb9\x8e\x09\x6d\x7a\xc8\x9e\x3c\x3b\x66\xe6\xad\xd4\x53\x59\x22\x06\xa2\x63\xe8\xfb\x88\x52\x83\xa3\x5d\xbe\x1f\x19\xd1\xf6\xe1\xcb\x04\x8b\x48\xe7\xc1\x9b\x48\xfc\x82\x13\x35\x26\x31\xad\x2b\x04\xe2\xde\x7f\x12\x86\x40\x0d\x59\xea\x1a\x42\xeb\x3a\xcb\xb0\xd4\x9a\xe4\x23\x5a\x17\x38\x42\x3c\xd4\xfc\xff\x68\x21\xc7\x1c\xed\x3a\xc2\x26\x24\x77\x5c\xef\xdc\x4c\x32\x7b\x2b\xd8\xec\x5e\x0f\xdf\xa0\xc2\x8f\x4f\x9d\x9d\xa4\x10\x77\x63\x38\x86\xd4\xcc\x41\xca\x3b\x34\xde\x30\xbf\xfd\x71\xb5\x28\x14\x0f\xa6\xd3\xfb\xe7\x03\xe4\x88\x6a\x6b\xce\x84\x49\x42\xd5\xf3\x5a\x14\x2a\x79\x2d\x0a\x39\x22\x24\x87\x6a\x17\x2f\x9e\x50\xed\xd2\x69\xdb\x5b\x62\x4c\xb2\x10\x00\xa0\xea\xfc\x9b\xaa\xbd\xd3\x37\x5c\xf7\xb0\x96\x97\x5c\xb8\xd6\xfe\xdb\xee\x7d\x16\xef\x97\x21\x8b\x03\x5e\x67\xae\x0f\xe3\x53\xd6\x57\xae\x5f\x55\xdf\x01\x00\x00\xff\xff\x9e\x20\xee\xee\x26\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css/reflectui.css"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/rui_form.tmpl"].(os.FileInfo),
		fs["/templates/rui_internal.tmpl"].(os.FileInfo),
		fs["/templates/rui_std.tmpl"].(os.FileInfo),
		fs["/templates/rui_table.tmpl"].(os.FileInfo),
		fs["/templates/rui_tree.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
