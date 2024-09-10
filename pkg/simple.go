package pkg

import (
	"fmt"
	"math/rand"
	"strconv"
)

func generateSimpleMathAdditionAndSubtraction(max int) string {
	f := rand.Intn(5)
	switch f {
	case 0:
		return generateSimpleMathIntAddition(max)
	case 1:
		return generateSimpleMathIntAdditionX(max)
	case 2:
		return generateSimpleMathIntAdditionY(max)
	case 3:
		return generateSimpleMathIntSubtraction(max)
	case 4:
		return generateSimpleMathIntSubtractionX(max)
	case 5:
		return generateSimpleMathIntSubtractionY(max)
	default:
		// 不会进入该default，仅为了保险
		return generateSimpleMathIntAddition(max)
	}
}

// 整数加法
func generateSimpleMathIntAddition(max int) string {
	x := rand.Intn(max + 1)
	y := rand.Intn(max + 1 - x)
	return fmt.Sprintf("%d + %d = (     )\n", x, y)
}

func generateSimpleMathDecimalAddition(decimal int) string {
	x := generateSimpleMathDecimal(decimal)
	y := generateSimpleMathDecimal(decimal)
	return fmt.Sprintf("%.*f + %.*f = (     )\n", decimal, x, decimal, y)
}

func generateSimpleMathIntAdditionX(max int) string {
	sum := rand.Intn(max + 1)
	y := rand.Intn(sum + 1)
	return fmt.Sprintf("(     ) + %d = %d\n", y, sum)
}

func generateSimpleMathIntAdditionY(max int) string {
	sum := rand.Intn(max + 1)
	x := rand.Intn(sum + 1)
	return fmt.Sprintf("%d + (     ) = %d\n", x, sum)
}

// 整数减法
func generateSimpleMathIntSubtraction(max int) string {
	x := rand.Intn(max + 1)
	y := rand.Intn(x + 1)
	return fmt.Sprintf("%d - %d = (     )\n", x, y)
}

func generateSimpleMathDecimalSubtraction(decimal int) string {
	x := generateSimpleMathDecimal(decimal)
	y := generateSimpleMathDecimal(decimal)
	
	if x > y {
		return fmt.Sprintf("%.*f - %.*f = (     )\n", decimal, x, decimal, y)
	} else if y > x {
		return fmt.Sprintf("%.*f - %.*f = (     )\n", decimal, y, decimal, x)
	} else {
		return fmt.Sprintf("%.*f - %.*f = (     )\n", decimal, x, decimal, y)
	}
}

func generateSimpleMathIntSubtractionY(max int) string {
	first := rand.Intn(max + 1)
	diff := rand.Intn(first + 1)
	
	return fmt.Sprintf("%d - (     ) = %d\n", first, diff)
}

func generateSimpleMathIntSubtractionX(max int) string {
	diff := rand.Intn(max + 1)
	second := rand.Intn(max - diff + 1)
	return fmt.Sprintf("(     ) - %d = %d\n", second, diff)
}

// 乘法
func generateSimpleMathMultiplication(max int) string {
	x := rand.Intn(max + 1)
	y := rand.Intn(max + 1)
	return fmt.Sprintf("%d x %d = (     )\n", x, y)
}

// 除法
func generateSimpleMathDivision(max int) string {
	x := rand.Intn(max + 1)
	y := rand.Intn(max+1) + 1 // 可能存在为0的情况
	s := x * y
	return fmt.Sprintf("%d ÷ %d = (     )\n", s, x)
}

//
func generateSimpleMathDecimal(decimal int) float64 {
	f := rand.Float64()
	var r float64
	var err error
	
	switch decimal {
	case 1:
		r, err = strconv.ParseFloat(fmt.Sprintf("%.1f", f*100), 64)
	case 2:
		r, err = strconv.ParseFloat(fmt.Sprintf("%.2f", f*100), 64)
	case 3:
		r, err = strconv.ParseFloat(fmt.Sprintf("%.3f", f*100), 64)
	case 4:
		r, err = strconv.ParseFloat(fmt.Sprintf("%.4f", f*100), 64)
	default:
		r, err = strconv.ParseFloat(fmt.Sprintf("%.2f", f*100), 64)
	}
	
	if err != nil {
		return f * 100
	}
	return r
}
