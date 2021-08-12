package gnzip

import (
	"github.com/gnames/gnar/ent/archive"
)

type gnzip struct {
	path, ext string
}

func New(path, ext string) (archive.Archive, error) {
	return gnzip{path: path, ext: ext}, nil
}
