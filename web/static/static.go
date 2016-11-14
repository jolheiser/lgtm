package static

//go:generate go-bindata -pkg static -o bindata.go files/...
//go:generate go fmt bindata.go

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func FileSystem() http.FileSystem {
	fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "files"}
	return &binaryFileSystem{
		fs,
	}
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name[1:])
}
