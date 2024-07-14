package pkg

import (
	"errors"
	"fmt"
	"github.com/signintech/gopdf"
)

type Pdf struct {
	*gopdf.GoPdf
}

// 新建PDF对象
// important: 每次 SetXY 都会影响后续所有的位置

const (
	defaultFontSize = 12
)

func NewPdf() *Pdf {
	pdf := &Pdf{
		&gopdf.GoPdf{},
	}
	
	// 设置默认页面大小
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	
	// 设置默认字体
	err := pdf.AddTTFFont("roboto", "../../front/Roboto-Light.ttf")
	if err != nil {
		panic("添加TTF字体失败, err: " + err.Error())
	}
	
	err = pdf.SetFont("roboto", "", defaultFontSize)
	if err != nil {
		panic("设置字体失败, err: " + err.Error())
	}
	
	return pdf
}

// 设置TTF字体

type fontFamilyType string

const (
	fontFamilyRoboto      fontFamilyType = "roboto"
	fontFamilyZhiMangXing fontFamilyType = "zhimangxing"
	fontFamilyNoteSans    fontFamilyType = "noteSans"
)

func (pdf *Pdf) setFont(family fontFamilyType, style string, size int) error {
	switch family {
	case "roboto":
		if err := pdf.AddTTFFont("roboto", "../../front/Roboto-Light.ttf"); err == nil {
			return pdf.SetFont("roboto", style, size)
		} else {
			return err
		}
	
	case "zhimangxing":
		if err := pdf.AddTTFFont("zhimangxing", "../../front/ZhiMangXing-Regular.ttf"); err == nil {
			return pdf.SetFont("zhimangxing", style, size)
		} else {
			return err
		}
	
	case "noteSans":
		
		if err := pdf.AddTTFFont("noteSans", "../../front/NotoSansSC-VariableFont_wght.ttf"); err == nil {
			return pdf.SetFont("noteSans", style, size)
		} else {
			return err
		}
	
	default:
		return errors.New(fmt.Sprintf("未匹配字体: %s", family))
	}
}

// 添加一个头函数，如果存在，它将由AddPage（）自动调用 -- 一定需要在调用 AddPage 之前调用 AddHeader 函数

func (pdf *Pdf) addHeader(contents ...string) {
	pdf.AddHeader(func() {
		for _, content := range contents {
			pdf.CellWithOption(&gopdf.Rect{W: 150, H: 15}, content, gopdf.CellOption{
				Border: 0,
			})
		}
	})
	
}

// 添加一个页脚函数，如果存在，它将由AddPage（）自动调用 -- 一定需要在调用 AddPage 之前调用 AddFooter 函数
func (pdf *Pdf) addFooter(contents ...string) {
	pdf.AddFooter(func() {
		for _, content := range contents {
			pdf.Cell(&gopdf.Rect{W: 120, H: 15}, content)
		}
	})
}

// 添加行
type pdfLineType string

const pdfLineDashedType = "dashed"
const pdfLineDottedType = "dotted"

func (pdf *Pdf) addLine(lineType pdfLineType, x1 float64, y1 float64, x2 float64, y2 float64) {
	// 下划线
	// line type: "dashed" ,"dotted"
	pdf.SetLineWidth(1)
	
	switch lineType {
	case pdfLineDashedType:
		pdf.SetLineType("dashed")
	case pdfLineDottedType:
		pdf.SetLineType("dotted")
	default:
		pdf.SetLineType("dotted")
	}
	
	pdf.Line(x1, y1, x2, y2)
}

// 反转, 文字横排改成竖排

func (pdf *Pdf) rotate(x, y float64, text string) error {
	//pdf.SetXY(x, y)
	// angle 斜角
	pdf.Rotate(270.0, x, y)
	return pdf.Text(text)
}

// Pdf 中 Outline是大纲的意思 - macOS使用没发现什么效果 - 不建议使用
func (pdf *Pdf) addOutline(text string) {
	pdf.AddOutlineWithPosition(text)
}

type simpleMathType string

const (
	SimpleMathRand               simpleMathType = "rand"
	SimpleMathAddition           simpleMathType = "addition"
	SimpleMathAdditionX          simpleMathType = "additionX"
	SimpleMathAdditionY          simpleMathType = "additionY"
	SimpleMathSubtraction        simpleMathType = "subtraction"
	SimpleMathSubtractionX       simpleMathType = "subtractionX"
	SimpleMathSubtractionY       simpleMathType = "subtractionY"
	SimpleMathMultiplication     simpleMathType = "multiplication"
	SimpleMathDivision           simpleMathType = "division"
	SimpleMathDecimalAddition    simpleMathType = "decimalAddition"
	SimpleMathDecimalSubtraction simpleMathType = "decimalSubtraction"
)

// decimal: 当需要生成小数时，填入

func (pdf *Pdf) GenerateSimpleMathFile(max int, t simpleMathType, pageCount int, decimal int, ) error {
	
	const height = 25
	err := pdf.setFont(fontFamilyNoteSans, "", defaultFontSize)
	if err != nil {
		return err
	}
	
	for page := 1; page <= pageCount; page++ {
		// 添加 Header 内容
		pdf.addHeader("姓名: ______________", "开始时间: ______________", "结束时间: ______________", "分数: ______________")
		//pdf.addFooter("耗时:_________", "评分:___________")
		
		pdf.AddPage()
		//pdf.addLine(pdfLineDashedType, 10, 30, 1000, 30)
		
		var startX float64 = 10
		var startY float64 = 40
		
		for i := 1; i <= 31; i++ {
			pdf.SetXY(startX, startY)
			
			for ii := 1; ii <= 4; ii++ {
				var r string
				
				switch t {
				case SimpleMathAddition:
					r = generateSimpleMathIntAddition(max)
				case SimpleMathAdditionX:
					r = generateSimpleMathIntAdditionX(max)
				case SimpleMathAdditionY:
					r = generateSimpleMathIntAdditionY(max)
				case SimpleMathSubtraction:
					r = generateSimpleMathIntSubtraction(max)
				case SimpleMathSubtractionX:
					r = generateSimpleMathIntSubtractionX(max)
				case SimpleMathSubtractionY:
					r = generateSimpleMathIntSubtractionY(max)
				case SimpleMathMultiplication:
					r = generateSimpleMathMultiplication(max)
				case SimpleMathDivision:
					r = generateSimpleMathDivision(max)
				
				case SimpleMathDecimalAddition:
					r = generateSimpleMathDecimalAddition(decimal)
				
				case SimpleMathDecimalSubtraction:
					r = generateSimpleMathDecimalSubtraction(decimal)
				
				default:
					r = generateSimpleMathAdditionAndSubtraction(max)
				}
				
				err := pdf.CellWithOption(&gopdf.Rect{W: 150, H: 15}, r, gopdf.CellOption{
					Align:  gopdf.Left,
					Border: 1,
					Float:  2,
				})
				if err != nil {
					return err
				}
				
				//pdf.addLine(pdfLineDashedType, startX, startY, startX+gopdf.PageSizeA4.W-2*startX, startY)
			}
			
			startY += height
		}
	}
	
	return pdf.WritePdf("/Users/qianxing/Downloads/1.pdf")
}

func generatePdfPage() {

}
