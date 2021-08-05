package gnar

import (
	"github.com/gnames/gnar/ent/archive"
)

type gnar struct {
	archive.Archive
}

func New(arch archive.Archive) GNar {
	return &gnar{Archive: arch}
}
