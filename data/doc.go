package data

import (
	// The blank import is to make govendor happy.
	_ "github.com/shurcooL/vfsgen"
)

//go:generate go run -tags=dev assets_generate.go
