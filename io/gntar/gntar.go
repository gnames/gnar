package gntar

import (
	"github.com/gnames/gnar/ent/archive"
)

type gntar struct {
	path, ext string
}

func New(path, ext string) (archive.Archive, error) {
	return gntar{path: path, ext: ext}, nil
}
