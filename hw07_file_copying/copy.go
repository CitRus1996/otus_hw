package main

import (
	"errors"
	"io"
	"math"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const (
	bufsize = 1 << 10
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fi, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	size := fi.Size()

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if limit < 0 {
		limit = size
	}

	if limit == 0 || limit+offset > fi.Size() {
		limit = fi.Size() - offset
	}

	sourceFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	dataReader := io.NewSectionReader(sourceFile, offset, limit+1)

	targetFile, err := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	bar := pb.Full.Start64(limit)
	bar.SetMaxWidth(100)
	bar.Set(pb.Bytes, true)

	for i := offset; i < offset+limit; i += bufsize {
		copylen := int64(math.Min(bufsize, float64(limit)))
		c, err := io.CopyN(targetFile, dataReader, copylen)
		bar.Add64(c)
		if (err != nil && !errors.Is(err, io.EOF)) || c > limit {
			return err
		}
	}
	bar.Finish()

	return nil
}