// +build dev

package data

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("base")
