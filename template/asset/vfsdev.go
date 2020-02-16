// +build dev

package asset

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("..")
