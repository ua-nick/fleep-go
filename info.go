package fleep

// Info represents file info
type Info struct {
	Type      []FileType
	Extension []string
	Mime      []string
}

// TypeMatches checks if the file matches provided file type
func (info *Info) TypeMatches(value FileType) bool {
	for item := range info.Type {
		if info.Type[item] == value {
			return true
		}
	}
	return false
}

// ExtensionMatches checks if the file matches provided extension
func (info *Info) ExtensionMatches(value string) bool {
	for item := range info.Extension {
		if info.Extension[item] == value {
			return true
		}
	}
	return false
}

// MimeMatches checks if the file matches provided MIME type
func (info *Info) MimeMatches(value string) bool {
	for item := range info.Mime {
		if info.Mime[item] == value {
			return true
		}
	}
	return false
}

// IsRasterImage checks if the file is a raster image
func (info *Info) IsRasterImage() bool {
	return info.TypeMatches(RasterImage)
}

// IsRawImage checks if the file is a raw image
func (info *Info) IsRawImage() bool {
	return info.TypeMatches(RawImage)
}

// IsVectorImage checks if the file is a vector image
func (info *Info) IsVectorImage() bool {
	return info.TypeMatches(VectorImage)
}

// Is3DImage checks if the file is a 3D image
func (info *Info) Is3DImage() bool {
	return info.TypeMatches(ThreeDimensionalImage)
}

// IsAudio checks if the file is an audio
func (info *Info) IsAudio() bool {
	return info.TypeMatches(Audio)
}

// IsVideo checks if the file is a video
func (info *Info) IsVideo() bool {
	return info.TypeMatches(Video)
}

// IsDocument checks if the file is a document
func (info *Info) IsDocument() bool {
	return info.TypeMatches(Document)
}

// IsArchive checks if the file is an archive
func (info *Info) IsArchive() bool {
	return info.TypeMatches(Archive)
}

// IsExecutable checks if the file is an executable
func (info *Info) IsExecutable() bool {
	return info.TypeMatches(Executable)
}

// IsFont checks if the file is a font
func (info *Info) IsFont() bool {
	return info.TypeMatches(Font)
}

// IsSystem checks if the file is a system
func (info *Info) IsSystem() bool {
	return info.TypeMatches(System)
}

// IsDatabase checks if the file is a database
func (info *Info) IsDatabase() bool {
	return info.TypeMatches(Database)
}
