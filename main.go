package fleep

import (
	"bytes"
	"errors"
)

// GetInfo returns file info
func GetInfo(file []byte) (Info, error) {

	var info Info

	sequence, err := getSequenceFromFile(file)

	if err != nil {
		return info, err
	}

	info, err = getInfo(sequence)

	if err != nil {
		return info, err
	}

	return info, nil
}

func getSequenceFromFile(file []byte) ([]byte, error) {
	sequenceLength := 128
	if len(file) < sequenceLength {
		return nil, errors.New("not enough bytes")
	}
	sequence := file[:sequenceLength]
	return sequence, nil
}

func getInfo(sequence []byte) (Info, error) {

	var info Info

	items := findMatchesInCollection(sequence)

	if len(items) == 0 {
		return info, errors.New("unknown file format")
	}

	info = constructInfo(items)

	return info, nil

}

func findMatchesInCollection(sequence []byte) []collectionItem {
	var selectedItems []collectionItem
	for _, item := range collection {
		if itemMatches(item, sequence) {
			selectedItems = append(selectedItems, item)
		}
	}

	return selectedItems
}

func itemMatches(item collectionItem, sequence []byte) bool {
	for _, signature := range item.Signature {
		if signatureMatches(signature, item.Offset, sequence) {
			return true
		}
	}
	return false
}

func signatureMatches(signature []byte, offset int, sequence []byte) bool {
	return bytes.Equal(sequence[offset:offset+len(signature)], signature)
}

func constructInfo(items []collectionItem) Info {
	var info Info
	for _, item := range items {
		info.Type = append(info.Type, item.Type)
		info.Extension = append(info.Extension, item.Extension)
		info.Mime = append(info.Mime, item.Mime)
	}
	return info
}
