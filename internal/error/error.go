package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	WantErr  bool
	Contains string
}

func AssertError(t *testing.T, expect Result, err error) {
	if expect.WantErr {
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), expect.Contains)
		}
		return
	}

	assert.Nil(t, err)
}
