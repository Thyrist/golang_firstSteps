package main
import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello, World!")

	var hello_world_string string = "Hello, \nWorld!" + " concatenated"	//single line string
	var hello_world_multiline_string string = `Hello,
World!`	//multiline string
	var intnum uint16 = 1337	//possible int types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
	var floatnum float32 = 3.14	//possible float types: float32, float64

	//automatic type setting
	var auto_intnum = 1337
	auto_intnum_2 := 1337	//short variable declaration

	//multiple variable declaration
	var a, b, c int = 1, 2, 3
	var d, e, f = 4, 5, 6
	g, h, i := 7, 8, 9

	//constants; set once and cannot be changed
	const pi float64 = 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229
	

}