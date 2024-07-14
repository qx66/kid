package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSimpleMathFile(t *testing.T) {
	pdf := NewPdf()
	err := pdf.GenerateSimpleMathFile(100, SimpleMathSubtraction, 1, 2)
	
	assert.NoError(t, err, "生成PDF文件失败")
}
