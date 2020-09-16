package fleep

type FileType int

const (
	_ FileType = iota
	RasterImage
	RawImage
	VectorImage
	ThreeDimensionalImage
	Audio
	Video
	Document
	Archive
	Executable
	Font
	System
	Database
)
