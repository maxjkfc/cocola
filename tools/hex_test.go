package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_switchHex(t *testing.T) {

	assert.Equal(t, switchHex(10), 'A')

}
