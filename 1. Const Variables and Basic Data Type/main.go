package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ==========================================
	// CONSTANTS
	// ==========================================
	// Constants are declared with the `const` keyword.
	// Their values cannot be changed or reassigned later.
	const pi = 3.14159
	fmt.Println("const pi =", pi)
	// Will show compiler error if you try to modify it:
	// pi = "Another Value"
	
	// Constants can optionally have explicit types:
	const myTypedConst string = "Constant String"

	// ==========================================
	// BASIC DATA TYPES: INTEGERS
	// ==========================================
	
	// 'int' size is architecture dependent (32-bit on 32-bit OS, 64-bit on 64-bit OS)
	var intNum int = 1
	fmt.Println("\n--- Integers ---")
	fmt.Println("int:", intNum)
	
	// Explicit size integers (signed): int8, int16, int32, int64
	var intNum16 int16 = 32767
	// Will show compiler error if a value exceeds the maximum for its type:
	// var intNum16 int16 = 32767 + 1
	fmt.Println("int16:", intNum16)

	var intNum32 int32 = 2147483647
	fmt.Println("int32:", intNum32)

	var intNum64 int64 = 9223372036854775807
	fmt.Println("int64:", intNum64)
	
	// Unsigned integers (only positive values): uint8, uint16, uint32, uint64, uint
	var uintNum uint = 1
	fmt.Println("uint:", uintNum)

	// 'byte' is an alias for uint8, typically used for raw binary data.
	var myByte byte = 255
	fmt.Println("byte (alias for uint8):", myByte)

	// ==========================================
	// BASIC DATA TYPES: FLOATING-POINT NUMBERS
	// ==========================================
	fmt.Println("\n--- Floating-Point Numbers ---")
	
	// float32 offers smaller precision (typically 7 decimal digits)
	var floatNum32 float32 = 12345.9
	fmt.Println("float32:", floatNum32)

	// float64 offers larger precision (typically 15 decimal digits). Default for Go floats.
	var floatNum float64 = 12345.9
	fmt.Println("float64:", floatNum)

	// ==========================================
	// BASIC DATA TYPES: COMPLEX NUMBERS
	// ==========================================
	fmt.Println("\n--- Complex Numbers ---")
	
	// complex64 (float32 real/imaginary parts) & complex128 (float64 real/imaginary parts)
	var compNum complex128 = complex(5, 7) // 5 + 7i
	fmt.Println("complex128:", compNum)
	fmt.Printf("Real: %v, Imaginary: %v\n", real(compNum), imag(compNum))

	// ==========================================
	// TYPE CONVERSION
	// ==========================================
	fmt.Println("\n--- Type Conversion ---")
	
	// Error when doing mathematical operations with mixed types
	// var result = intNum + floatNum32 // Mismatched types 'int' and 'float32'

	// But works with explicit type conversion. Note: converting float to int truncates the decimal.
	var result = intNum + int(floatNum32)
	fmt.Println("Result of int + int(float32):", result)
	
	// ==========================================
	// BASIC DATA TYPES: STRINGS
	// ==========================================
	fmt.Println("\n--- Strings ---")
	
	// Double quotes for standard single-line strings
	var singleLine string = "Single Line"
	
	// Backticks (`) for raw multi-line strings (preserves newlines & ignores escapes)
	var multiLine string = `Multi
	Line`
	
	// Concatenating strings
	var concat = singleLine + " | " + multiLine
	fmt.Println("Concatenated:", concat)

	// len() returns the number of BYTES in the string (not necessarily the number of characters)
	fmt.Println("Byte length of concat:", len(concat))

	// To count actual characters (runes), especially for non-ASCII characters, use utf8 package
	fmt.Println("Character (rune) count:", utf8.RuneCountInString(concat))

	// ==========================================
	// BASIC DATA TYPES: RUNE (CHARACTERS)
	// ==========================================
	fmt.Println("\n--- Rune ---")
	
	// 'rune' is an alias for int32. It represents a single Unicode code point.
	// Declared using single quotes.
	var myRune rune = 'a'
	// It prints the Unicode int value of 'a', which is 97
	fmt.Println("rune value for 'a':", myRune)
	// Formatting to print the actual character
	fmt.Printf("Formatted rune char: %c\n", myRune)

	// ==========================================
	// BASIC DATA TYPES: BOOLEAN
	// ==========================================
	fmt.Println("\n--- Boolean ---")
	
	// Can only be true or false
	var myBool bool = true
	fmt.Println("bool:", myBool)

	// ==========================================
	// DEFAULT ZERO VALUES
	// ==========================================
	fmt.Println("\n--- Default Zero Values ---")
	
	// Variables declared without initial value are assigned the default "zero value":
	// 0 for numbers, false for booleans, "" for strings.
	var defaultInt int 
	fmt.Println("Default int:", defaultInt) // Output: 0
	
	var defaultString string
	fmt.Printf("Default string: %q\n", defaultString) // Output: ""

	var defaultBool bool
	fmt.Println("Default bool:", defaultBool) // Output: false

	// ==========================================
	// VARIABLE DECLARATION / ASSIGNMENT SHORTCUTS
	// ==========================================
	fmt.Println("\n--- Variable Assignments ---")
	
	// := is shorthand for variable declaration & assigning without explicitly defining type.
	// Used inside a function ONLY. The type is inferred based on the value on the right.
	myVar := "text"
	fmt.Println("Shorthand inferred string:", myVar)

	// Multiple assignments
	var var1, var2 = 1, "text"
	fmt.Println("Multiple assignment:", var1, var2)
}