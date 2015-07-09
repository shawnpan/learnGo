package testbase

import (
	"fmt"
	"math"
)

func TestMath() {
	fmt.Println(math.Abs(11))
	fmt.Println(math.Abs(-12))
	fmt.Println(math.Abs(-12.34))
	fmt.Println(math.Dim(15, 17))
	fmt.Println(math.Dim(17, 15))

	fmt.Println(math.Max(15, 17))
	fmt.Println(math.Min(15.6, 15.7))

	// 取余
	fmt.Println(math.Mod(10, 4))

	// 返回值1，整数部分，返回值2，小数部分
	fmt.Println(math.Modf(15.6))

	// 幂运算
	fmt.Println(math.Pow(2, 3))

	fmt.Println(math.Sqrt(9))

	fmt.Println("MaxInt8 = ", math.MaxInt8)
	fmt.Println("MinInt8 = ", math.MinInt8)
	fmt.Println("MaxInt16 = ", math.MaxInt16)
	fmt.Println("MinInt16 = ", math.MinInt16)
	fmt.Println("MaxInt32 = ", math.MaxInt32)
	fmt.Println("MinInt32 = ", math.MinInt32)
	fmt.Println("MaxInt64 = ", math.MaxInt64)
	fmt.Println("MinInt64 = ", math.MinInt64)

	fmt.Println("MaxUint8 = ", math.MaxUint8)
	fmt.Println("MaxUint16 = ", math.MaxUint16)
	fmt.Println("MaxUint32 = ", math.MaxUint32)
	// fmt.Println("MaxUint64 = ", math.MaxUint64)

	fmt.Println("MaxInt16 = ", math.MaxInt16)
	fmt.Println("MinInt16 = ", math.MinInt16)

}
