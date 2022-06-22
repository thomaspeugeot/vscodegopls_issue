package main

import (
	"embed"
)

//go:embed go/models
var embeddedModelsDir embed.FS
