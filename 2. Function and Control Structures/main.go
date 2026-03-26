package main

import (
	"errors"
	"fmt"
)

func main() {
	// ==========================================
	// FUNCTIONS
	// ==========================================
	// Functions are declared using the 'func' keyword.
	// main() is the entry point of the executable program.
	fmt.Println("--- Function Calls ---")
	printMessage("Hello World from printMessage function!")
}

// ==========================================
// FUNCTION WITH PARAMETERS
// ==========================================
// printMessage takes one parameter of type string.
func printMessage(printValue string) {
	fmt.Println("Message:", printValue)

	fmt.Println("\n--- Multiple Return Values & Error Handling ---")
	// Calling a function that returns multiple values.
	// In Go, functions can return more than one value, commonly used for returning a result and an error.
	var result, remainder, err = intDivision(10, 3)
	
	// ==========================================
	// CONTROL STRUCTURES: IF / ELSE IF / ELSE
	// ==========================================
	// The basic if statement. Notice that conditions don't need parentheses ().
	if err != nil {
		fmt.Println("Error:", err)
	} else if remainder == 0 {
		fmt.Println("Result:", result, "(No remainder)")
	} else {
		fmt.Println("Result:", result)
		fmt.Println("Remainder:", remainder)
	}

	// ==========================================
	// LOGICAL OPERATORS: AND (&&), OR (||)
	// ==========================================
	fmt.Println("\n--- AND / OR Logic Conditions ---")
	
	age := 25
	hasID := true

	// AND (&&): Both conditions must be true
	if age >= 18 && hasID {
		fmt.Println("Access Granted: Over 18 and has ID.")
	} else {
		fmt.Println("Access Denied.")
	}

	isWeekend := true
	isHoliday := false

	// OR (||): At least one condition must be true
	if isWeekend || isHoliday {
		fmt.Println("You don't have to work today!")
	} else {
		fmt.Println("It's a workday.")
	}

	// ==========================================
	// CONTROL STRUCTURES: SWITCH
	// ==========================================
	fmt.Println("\n--- Switch Statements ---")
	
	dayOfWeek := "Tuesday"

	// Basic switch statement. In Go, cases break automatically; no need for a 'break' keyword.
	switch dayOfWeek {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Tuesday", "Wednesday", "Thursday": // Multiple condition matching in one case
		fmt.Println("Midweek days.")
	case "Friday":
		fmt.Println("Almost the weekend!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default: // The catch-all case if nothing matches
		fmt.Println("Invalid day.")
	}

	// Switch without an expression (acts like a clean if/else chain)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// ==========================================
// FUNCTION WITH MULTIPLE RETURN VALUES
// ==========================================
// intDivision takes two integers and returns three values: int, int, error.
// Specifying exactly what is returned inside the parentheses: (int, int, error).
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	
	// Check for division by zero to prevent a runtime panic
	if denominator == 0 {
		// errors.New creates a simple text-based error
		err = errors.New("cannot divide by zero")
		// We return 0 for the ints and the error
		return 0, 0, err
	}
	
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	
	// Returning the three variables matching the (int, int, error) signature
	return result, remainder, err
}
