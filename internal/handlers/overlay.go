package handlers

import (
	"errors"
	"io/fs"
	"os"
)

type overlayFS struct {
	override fs.FS
	fallback fs.FS
}

func (o overlayFS) Open(name string) (fs.File, error) {
	f, err := o.override.Open(name)
	if err == nil {
		return f, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return o.fallback.Open(name)
	}
	return nil, err
}

func NewOverlayFS(overridePath string, fallback fs.FS) fs.FS {
	return overlayFS{
		override: os.DirFS(overridePath),
		fallback: fallback,
	}
}
