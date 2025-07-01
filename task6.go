package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var ptr uintptr

	var f32 float32
	var f64 float64

	var c64 complex64
	var c128 complex128

	var s string
	var bl bool

	fmt.Println("Размеры типов данных в байтах:")
	fmt.Println(" ")
	fmt.Println("Целочисленные со знаком:")
	fmt.Printf("int:    %d байт\n", unsafe.Sizeof(i))
	fmt.Printf("int8:   %d байт\n", unsafe.Sizeof(i8))
	fmt.Printf("int16:  %d байт\n", unsafe.Sizeof(i16))
	fmt.Printf("int32:  %d байт (rune)\n", unsafe.Sizeof(i32))
	fmt.Printf("int64:  %d байт\n", unsafe.Sizeof(i64))

	fmt.Println("\nЦелочисленные без знака:")
	fmt.Printf("uint:    %d байт\n", unsafe.Sizeof(ui))
	fmt.Printf("uint8:   %d байт (byte)\n", unsafe.Sizeof(ui8))
	fmt.Printf("uint16:  %d байт\n", unsafe.Sizeof(ui16))
	fmt.Printf("uint32:  %d байт\n", unsafe.Sizeof(ui32))
	fmt.Printf("uint64:  %d байт\n", unsafe.Sizeof(ui64))
	fmt.Printf("uintptr: %d байт\n", unsafe.Sizeof(ptr))

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Printf("float32: %d байт\n", unsafe.Sizeof(f32))
	fmt.Printf("float64: %d байт\n", unsafe.Sizeof(f64))

	fmt.Println("\nКомплексные числа:")
	fmt.Printf("complex64:  %d байт\n", unsafe.Sizeof(c64))
	fmt.Printf("complex128: %d байт\n", unsafe.Sizeof(c128))

	fmt.Println("\nПрочие типы:")
	fmt.Printf("string: %d байт (базовый размер)\n", unsafe.Sizeof(s))
	fmt.Printf("bool:   %d байт\n", unsafe.Sizeof(bl))

	fmt.Println("\nПримечание: размер int/uint зависит от платформы (32 или 64 бита)")
}
