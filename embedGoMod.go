package main

import (
	"embed"
)

//go:embed go/models go.mod
var embeddedModelsDir embed.FS
