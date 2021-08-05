package archive

type Unpacker interface {
	Unpack(destDir string) error
}

type Packer interface {
	Pack(srcDir string) error
}

type Archive interface {
	Unpacker
	Packer
}
