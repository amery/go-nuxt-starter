package client

//go:generate go run go.sancus.dev/web/cmd/embed -n embedded -o files.go dist

import (
	"log"

	"go.sancus.dev/web"
	"go.sancus.dev/web/fs"
)

func Middleware() web.MiddlewareHandlerFunc {
	fsys, err := fs.Sub(&embedded, "dist")
	if err != nil {
		log.Panic(err)
	}

	overlay := fs.FS{
		FS:    fsys,
		Index: "index.html",
	}

	return overlay.Middleware
}
