package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

func main() {

	var i int = 42
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var r rune = 'Ю'

	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	var ptr uintptr = 0x1a2b3c4d

	var f32 float32 = 3.1415926535
	var f64 float64 = 3.14159265358979323846
	var c64 complex64 = 3 + 2i
	var c128 complex128 = cmplx.Sqrt(-5 + 12i)

	var s string = "Hello, world!!"
	var b byte = 255

	var isTrue bool = true

	fmt.Println("\nЦелочисленные типы со знаком:")
	fmt.Printf("int(%d байт): %d\n", unsafe.Sizeof(i), i)
	fmt.Printf("int8(%d байт): %d\n", unsafe.Sizeof(i8), i8)
	fmt.Printf("int16(%d байт): %d\n", unsafe.Sizeof(i16), i16)
	fmt.Printf("int32(%d байт): %d\n", unsafe.Sizeof(i32), i32)
	fmt.Printf("int64(%d байт): %d\n", unsafe.Sizeof(i64), i64)
	fmt.Printf("int32(%d байт): %d (rune: %U '%c')\n", unsafe.Sizeof(i32), i32, r, r)

	fmt.Println("\nЦелочисленные типы без знака:")
	fmt.Printf("uint(%d байт): %d\n", unsafe.Sizeof(ui), ui)
	fmt.Printf("uint8(%d байт): %d (byte: %d)\n", unsafe.Sizeof(ui8), ui8, b)
	fmt.Printf("uint16(%d байт): %d\n", unsafe.Sizeof(ui16), ui16)
	fmt.Printf("uint32(%d байт): %d\n", unsafe.Sizeof(ui32), ui32)
	fmt.Printf("uint64(%d байт): %d\n", unsafe.Sizeof(ui64), ui64)
	fmt.Printf("uintptr(%d байт): %#x\n", unsafe.Sizeof(ptr), ptr)

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Printf("float32(%d байт): %.7f (потеря точности)\n", unsafe.Sizeof(f32), f32)
	fmt.Printf("float64(%d байт): %.15f\n", unsafe.Sizeof(f64), f64)

	fmt.Println("\nКомплексные числа:")
	fmt.Printf("complex64(%d байт): %v\n", unsafe.Sizeof(c64), c64)
	fmt.Printf("complex128(%d байт): %v\n", unsafe.Sizeof(c128), c128)

	fmt.Println("\nСтроки:")
	fmt.Printf("string(%d байт): %s\n", unsafe.Sizeof(s), s)

	fmt.Println("\nЛогический тип:")
	fmt.Printf("bool(%d байт): %v", unsafe.Sizeof(isTrue))
}
