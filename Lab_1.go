package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func applyXOR(input string, key byte) string {

	data := []byte(input)
	result := make([]byte, len(data))

	for i := range data {
		result[i] = data[i] ^ key
	}
	return string(result)
}

func main() {

	// Lab 1: Basic Arithmetic

	fmt.Println("--- Lab 1 ---")

	val := 4
	val += 5
	val -= 3
	val *= 2
	val /= 3
	val %= 2
	fmt.Println("Final result for Lab 1:", val)

	// Lab 2: Logical Operators

	fmt.Println("\n--- Lab 2 ---")

	y := 1

	fmt.Println(y < 5 && y < 3)

	fmt.Println(y < 5 || y < 0)

	fmt.Println(!(y > 5 && y > 3))

	// Lab 3: Bitwise Operations
	fmt.Println("\n--- Lab 3 ---")

	p, q := 9, 10
	notP, notQ := ^4, ^5

	fmt.Printf("p      = %b\n", p)
	fmt.Printf("q      = %b\n", q)
	fmt.Printf("notP   = %b\n", notP)
	fmt.Printf("notQ   = %b\n", notQ)

	fmt.Printf("p & q  = %b\n", p&q)
	fmt.Printf("p | q  = %b\n", p|q)
	fmt.Printf("p ^ q  = %b\n", p^q)

	fmt.Printf("NOT p (using notP) = %b\n", notP)
	fmt.Printf("NOT q (using notQ) = %b\n", notQ)

	fmt.Printf("p << 2 = %b\n", p<<2)
	fmt.Printf("q >> 2 = %b\n", q>>2)

	// Lab 4: Mini Calculator
	fmt.Println("\n--- Lab 4 ---")

	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Addition")
		fmt.Println("2) Subtraction")
		fmt.Println("3) Multiplication")
		fmt.Println("4) Division")
		fmt.Println("5) Modulo")
		fmt.Println("6) Exit")

		var choice int
		fmt.Print("Select an option: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting calculator...")
			break
		}

		var num1, num2 int
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		switch choice {
		case 1:
			fmt.Println("Result:", num1+num2)
		case 2:
			fmt.Println("Result:", num1-num2)
		case 3:
			fmt.Println("Result:", num1*num2)
		case 4:
			if num2 != 0 {
				fmt.Println("Result:", num1/num2)
			} else {
				fmt.Println("Error: Cannot divide by zero.")
			}
		case 5:
			if num2 != 0 {
				fmt.Println("Result:", num1%num2)
			} else {
				fmt.Println("Error: Cannot modulo by zero.")
			}
		default:
			fmt.Println("Invalid selection. Try again.")
		}
	}

	// Lab 5: Encoding (Binary, Hex, Base64)
	fmt.Println("\n--- Lab 5 ---")
	fmt.Println("===== Data Encodings =====")

	var userInput string
	fmt.Print("Enter a single word to encode: ")
	fmt.Scan(&userInput)

	byteData := []byte(userInput)

	fmt.Println("Original text: ", userInput)
	fmt.Printf("Binary format: %b\n", byteData)
	fmt.Println("Hex format:    ", hex.EncodeToString(byteData))
	fmt.Println("Base64 format: ", base64.StdEncoding.EncodeToString(byteData))

	// Lab 6: XOR Encryption
	fmt.Println("\n--- Lab 6 ---")

	var secretMsg string
	fmt.Print("Enter a secret message: ")
	fmt.Scan(&secretMsg)

	encryptionKey := byte('Z')

	encrypted := applyXOR(secretMsg, encryptionKey)
	decrypted := applyXOR(encrypted, encryptionKey)

	fmt.Println("Original text:  ", secretMsg)
	fmt.Printf("Encrypted (Hex): %x\n", encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}
