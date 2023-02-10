package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	is := assert.New(t)

	size, err := Parallel("../fixture")

	is.Nil(err)

	is.Equal(size, int64(6))
}
