package wrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AutoWrapSline(t *testing.T) {

	d := "1.建議依照之前所說的那幾項修改.\n2.建議根據第一項所說的那些在延伸修改\n3.建議這幾次的修改都要確實達成並且提供驗證之項目"

	s := AutoWrapSline(d, 20)

	assert.NotEmpty(t, s)
}
