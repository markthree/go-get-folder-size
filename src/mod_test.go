package goGetFolderSize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileSize(t *testing.T) {
	is := assert.New(t)

	size, err := GetFolderSize("../fixture")

	is.Nil(err)

	is.Equal(size, int64(6))
}

func TestGetFolderSizeParallel(t *testing.T) {
	is := assert.New(t)

	size, err := GetFolderSizeParallel("../fixture")

	is.Nil(err)

	is.Equal(size, int64(6))
}
