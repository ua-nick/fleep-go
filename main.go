package fleep

import (
	"bytes"
	"errors"
)

const sequenceLength = 128

type Info struct {
	Type      []string
	Extension []string
	Mime      []string
}

func (info *Info) TypeMatches(value string) bool {
	matches := false
	for item := range info.Type {
		if info.Type[item] == value {
			matches = true
			break
		}
	}
	return matches
}

func (info *Info) ExtensionMatches(value string) bool {
	matches := false
	for item := range info.Extension {
		if info.Extension[item] == value {
			matches = true
			break
		}
	}
	return matches
}

func (info *Info) MimeMatches(value string) bool {
	matches := false
	for item := range info.Mime {
		if info.Mime[item] == value {
			matches = true
			break
		}
	}
	return matches
}

func GetInfo(file []byte) (Info, error) {
	info := Info{}
	if len(file) < 128 {
		return info, errors.New("not enough bytes")
	}
	sequence := file[:sequenceLength]
	for _, item := range collectionData {
		for _, signature := range item.Signature {
			if bytes.Equal(sequence[item.Offset:item.Offset+len(signature)], signature) {
				info.Type = append(info.Type, item.Type)
				info.Extension = append(info.Extension, item.Extension)
				info.Mime = append(info.Mime, item.Mime)
			}
		}
	}
	if len(info.Type) == 0 {
		return info, errors.New("unknown file format")
	}
	return info, nil
}
