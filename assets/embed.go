package assets

import "embed"

//go:embed data/*.json
var FS embed.FS
