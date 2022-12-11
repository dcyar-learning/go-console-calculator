package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println(">>>>> CALCULADORA <<<<<")
		fmt.Println("Escribe una de las opciones para continuar:")
		fmt.Println("1: suma")
		fmt.Println("2: resta")
		fmt.Println("3: multiplicación")
		fmt.Println("4: division")
		fmt.Println("5: salir")
		fmt.Println("")
		fmt.Print("-> ")
		scanner.Scan()
		optionMenu := scanner.Text()
		fmt.Println("")

		if optionMenu == "salir" || optionMenu == "5" {
			fmt.Println("Adios")
			break
		}

		switch {
		case optionMenu == "suma" || optionMenu == "1":
			fmt.Println("OPERACIÓN DE SUMA")
			val1, val2 := getValues(*scanner)

			printResult("suma", val1, val2, val1+val2)
		case optionMenu == "resta" || optionMenu == "2":
			fmt.Println("OPERACIÓN DE SUMA")
			val1, val2 := getValues(*scanner)

			printResult("resta", val1, val2, val1-val2)
		case optionMenu == "multiplicación" || optionMenu == "multiplicacion" || optionMenu == "3":
			fmt.Println("OPERACIÓN DE MULTIPLICACIÓN")
			val1, val2 := getValues(*scanner)

			printResult("multiplicación", val1, val2, val1*val2)
		case optionMenu == "división" || optionMenu == "division" || optionMenu == "4":
			fmt.Println("OPERACIÓN DE DIVISIÓN")
			val1, val2 := getValues(*scanner)

			printResult("división", val1, val2, val1/val2)
		default:
			fmt.Println("!! Opción no válida, inténtalo de nuevo")
		}
	}
}

func getValues(scanner bufio.Scanner) (int, int) {
	val1 := inputNumber("Ingresa el primer valor: ", scanner)
	val2 := inputNumber("Ingresa el segundo valor: ", scanner)

	return val1, val2
}

func inputNumber(text string, scanner bufio.Scanner) int {
	for {
		fmt.Print(text)
		scanner.Scan()
		inputText := scanner.Text()

		val, err := strconv.Atoi(inputText)

		if err == nil && val > 0 {
			return val
		}

		fmt.Printf("!! El valor \"%s\" no es aceptado, tienes que ingresar un número entero", inputText)
		fmt.Println("")
	}
}

func printResult(operator string, val1, val2 int, result int) {
	fmt.Println("")
	fmt.Printf(">>> La %s de %d / %d es: %d", operator, val1, val2, result)
	fmt.Println("")
}
