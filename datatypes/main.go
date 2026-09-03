package main

import (
	"fmt"
	"time"
)

func main() {
	//Boolean
	var boolean_var bool = true
	fmt.Println(boolean_var)

	//Numeric types
	var int_value int = 58
	fmt.Println(int_value)

	//int8
	var int8_value int8 = 127
	fmt.Println(int8_value)
	//int16
	//int32
	//int64
	//uint
	//uint8
	//uint16
	//uint32
	//unint64
	//uintptr
	//float32
	//float64
	//complex64
	//complex128
	//string
	/*
			 Type     Size (bits)                    Range                                                    Notes

		 `int8`   8                              128 to 127                                              1 byte
		 `int16`  16                             32,768 to 32,767                                        2 bytes
		 `int32`  32                             2,147,483,648 to 2,147,483,647                          4 bytes, same as `rune`
		 `int64`  64                             9,223,372,036,854,775,808 to 9,223,372,036,854,775,807  8 bytes
		 `int`    platform dependent (32 or 64)  same as `int32` or `int64` depending on system           usually 32bit on 32bit systems, 64bit on 64bit systems

		 Type       Size (bits)                    Range                                           Notes

		 `uint8`    8                              0 to 255                                        alias `byte`
		 `uint16`   16                             0 to 65,535                                     2 bytes
		 `uint32`   32                             0 to 4,294,967,295                              4 bytes
		 `uint64`   64                             0 to 18,446,744,073,709,551,615                 8 bytes
		 `uint`     platform dependent (32 or 64)  same as `uint32` or `uint64`                    depends on system
		 `uintptr`  platform dependent             unsigned integer large enough to store pointer  used for memory addresses


		 Type       Size (bits)  Approx Range                Precision

		 `float32`  32           ±1.18×10⁻³⁸ to ±3.4×10³⁸    ~6–7 decimal digits
		 `float64`  64           ±2.23×10⁻³⁰⁸ to ±1.8×10³⁰⁸  ~15–16 decimal digits

		 Type          Components                       Size (bits)  Notes

		 `complex64`   `float32` real + `float32` imag  64           2×32bit floats
		 `complex128`  `float64` real + `float64` imag  128          2×64bit floats

	*/
	//Aggregate / Composite data types
	//array
	array_data := []int{1, 45, 4, 64, 564}
	fmt.Println(array_data)

	//slice
	slice_Data := []int{5, 6, 4, 6, 5, 4, 4}
	fmt.Println(slice_Data)

	//struct
	type Structure_value struct {
		Value int
	}

	//pointer
	pointervalue := 52
	var pointer *int = &pointervalue
	fmt.Println(pointer)

	//map
	mapdata := map[string]int{"one": 1, "Two": 2}
	fmt.Println(mapdata)

	//function
	//func Hello(){}

	//interface
	go func() {
		type interface_methods interface {
			Method1()
		}
	}()

	//channel
	go func() {
		fmt.Println("Executing the channel")
		channel := make(chan int, 1)
		channel <- 1
		fmt.Println("channel data :", <-channel)
	}()

	time.Sleep(2 * time.Second)

	//Special / Advanced Types

	//nil
	/*
		nil
			nil represents the zero value for:
			Pointers
			Slices
			Maps
			Channels
			Interfaces
			Functions
			It means “no value” or “not initialized.”
	*/
	var p *int     // pointer
	fmt.Println(p) // Output: <nil>

	var s []int    // slice
	fmt.Println(s) // Output: []

	var m map[string]int
	fmt.Println(m) // Output: map[] (nil map)

	var ch chan int
	fmt.Println(ch) // Output: <nil>

	var f func()
	fmt.Println(f) // Output: <nil>

	//Type Alias
	type ExistingType int
	type NewName = ExistingType

	//Custom types
	type NewType ExistingType

	//composite Literals

	//unsafe/Lowlevel types
}

/*
//Default values
| Category     | Type / Example                                 | Zero Value         |
| ------------ | ---------------------------------------------- | ------------------ |
| Boolean      | `bool`                                         | false              |
| Integer      | `int, int8, int16, int32, int64`               | 0                  |
| Unsigned Int | `uint, uint8, uint16, uint32, uint64, uintptr` | 0                  |
| Float        | `float32, float64`                             | 0                  |
| Complex      | `complex64, complex128`                        | 0+0i               |
| String       | `string`                                       | ""                 |
| Array        | `[N]Type`                                      | zero of element    |
| Slice        | `[]Type`                                       | nil                |
| Struct       | `struct`                                       | zero of each field |
| Pointer      | `*Type`                                        | nil                |
| Map          | `map[Key]Value`                                | nil                |
| Function     | `func()`                                       | nil                |
| Interface    | `interface{}`                                  | nil                |
| Channel      | `chan Type`                                    | nil                |

*/
