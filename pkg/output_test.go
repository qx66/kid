package pkg

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestGenerateSimpleMathFile(t *testing.T) {
	logger, err := zap.NewProduction()
	assert.NoError(t, err, "new logger failed")
	
	pdf := NewPdf(logger)
	err = pdf.GenerateSimpleMathFile(1000, SimpleMathRand, 30, 2)
	
	assert.NoError(t, err, "生成PDF文件失败")
}

func TestPdf_GenerateCopyTextFile(t *testing.T) {
	logger, err := zap.NewProduction()
	assert.NoError(t, err, "new logger failed")
	
	pdf := NewPdf(logger)
	
	texts := []string{"▶️", "@", "！", "O"}
	err = pdf.GenerateCopyTextFile(1, texts)
	
	assert.NoError(t, err, "生成PDF文件失败")
}

func TestPdf_GenerateCopyTextImageFile(t *testing.T) {
	logger, err := zap.NewProduction()
	assert.NoError(t, err, "new logger failed")
	
	pdf := NewPdf(logger)
	
	words := []Word{
		Word{
			Header: "/Users/qianxing/Downloads/chinese/word/难/header.png",
			Stroke: "/Users/qianxing/Downloads/chinese/word/难/stroke.png",
			Word:   "难",
		},
		Word{
			Header: "/Users/qianxing/Downloads/chinese/word/好/header.png",
			Stroke: "/Users/qianxing/Downloads/chinese/word/好/stroke.png",
			Word:   "好",
		},
	}
	
	err = pdf.GenerateHanYuWordFile(1, words)
	assert.NoError(t, err, "生成PDF文件失败")
}
