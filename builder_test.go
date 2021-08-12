package errors_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MovieStoreGuy/errors"
)

func TestErrorBuilder(t *testing.T) {
	t.Parallel()

	var eb errors.Builder
	eb.WithError(nil)
	assert.Nil(t, eb.Build())

	eb.WithError(io.EOF)
	assert.Len(t, eb, 1)
	assert.ErrorIs(t, eb.Build(), io.EOF)

	eb.WithError(io.ErrClosedPipe)
	assert.Len(t, eb, 2)
	assert.ErrorIs(t, eb.Build(), errors.Combine(io.EOF, io.ErrClosedPipe))
}
