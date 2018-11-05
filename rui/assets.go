// +build dev

package rui

import (
	"net/http"
)

var Assets http.FileSystem = http.Dir("assets")
