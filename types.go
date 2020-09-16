package fleep

// FileType represents supported file type
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
