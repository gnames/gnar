package gnar

import "github.com/gnames/gnar/ent/archive"

type GNar interface {
	archive.Unpacker
	archive.Packer
}
