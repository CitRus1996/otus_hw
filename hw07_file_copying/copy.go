package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fi, err := os.Lstat(fromPath)
	if err != nil {
		return err
	}
	size := fi.Size()

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if limit <= 0 {
		limit = size
	}

	sourceFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}

	dataReader := io.NewSectionReader(sourceFile, offset, limit+1)

	targetFile, err := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(targetFile, dataReader)
	if err != nil {
		return err
	}
	return nil
}
