package memo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type wrapped struct {
	msg string
	err error
}

func (e wrapped) Error() string { return e.msg } // Errorはメッセージ表示
func (e wrapped) Unwrap() error { return e.err } // 深堀り

func TestUnwrap(t *testing.T) {
	err1 := errors.New("1")
	err2 := wrapped{"wrap 2", err1}

	assert.Equal(t, nil, errors.Unwrap(wrapped{"wrapped", nil}))
	assert.Equal(t, err1, errors.Unwrap(wrapped{"wrapped", err1}))
	assert.Equal(t, err2, errors.Unwrap(wrapped{"wrapped", err2}))

	assert.Equal(t, "wrap 2", err2.Error())
	assert.Equal(t, "wrap 2", errors.Unwrap(wrapped{"wrapped", err2}).Error())
}
