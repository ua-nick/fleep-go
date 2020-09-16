package fleep

type Info struct {
	Type      []FileType
	Extension []string
	Mime      []string
}

func (info *Info) TypeMatches(value FileType) bool {
	for item := range info.Type {
		if info.Type[item] == value {
			return true
		}
	}
	return false
}

func (info *Info) ExtensionMatches(value string) bool {
	for item := range info.Extension {
		if info.Extension[item] == value {
			return true
		}
	}
	return false
}

func (info *Info) MimeMatches(value string) bool {
	for item := range info.Mime {
		if info.Mime[item] == value {
			return true
		}
	}
	return false
}

func (info *Info) IsRasterImage() bool {
	return info.TypeMatches(RasterImage)
}

func (info *Info) IsRawImage() bool {
	return info.TypeMatches(RawImage)
}

func (info *Info) IsVectorImage() bool {
	return info.TypeMatches(VectorImage)
}

func (info *Info) Is3DImage() bool {
	return info.TypeMatches(ThreeDimensionalImage)
}

func (info *Info) IsAudio() bool {
	return info.TypeMatches(Audio)
}

func (info *Info) IsVideo() bool {
	return info.TypeMatches(Video)
}

func (info *Info) IsDocument() bool {
	return info.TypeMatches(Document)
}

func (info *Info) IsArchive() bool {
	return info.TypeMatches(Archive)
}

func (info *Info) IsExecutable() bool {
	return info.TypeMatches(Executable)
}

func (info *Info) IsFont() bool {
	return info.TypeMatches(Font)
}

func (info *Info) IsSystem() bool {
	return info.TypeMatches(System)
}

func (info *Info) IsDatabase() bool {
	return info.TypeMatches(Database)
}
