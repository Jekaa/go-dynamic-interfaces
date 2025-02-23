package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ================================
	// Базовый пример с type switch
	// ================================
	example1 := "Hello World"
	printAnyType(example1)
	printAnyType(42)
	printAnyType(true)

	// ================================
	// Работа с структурами и interface{}
	// ================================
	box := &SmartBox{Content: "Secret Message"}
	box.DisplayContent()

	box.Content = 3.1415
	box.DisplayContent()

	// ================================
	// Обработка кастомных ошибок
	// ================================
	err := processRequest("valid")
	fmt.Println("\nProcessing valid request:", err)

	err = processRequest("invalid")
	fmt.Println("Processing invalid request:", err)

	// ================================
	// Расширенное использование с рефлексией
	// ================================
	analyzeType("Go is awesome!")
	analyzeType(map[string]int{"one": 1})
}

// 1. Базовый пример с type switch
func printAnyType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s (Length: %d)\n", v, len(v))
	case int:
		fmt.Printf("Integer: %d (Double: %d)\n", v, v*2)
	case bool:
		fmt.Printf("Boolean: %t (Inverted: %t)\n", v, !v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// 2. Пример со структурой, содержащей interface{}
type SmartBox struct {
	Content interface{}
}

func (b *SmartBox) DisplayContent() {
	fmt.Printf("\nSmartBox contains: ")
	switch content := b.Content.(type) {
	case string:
		fmt.Printf("String -> %q\n", content)
	case int:
		fmt.Printf("Integer -> %d (Hex: 0x%x)\n", content, content)
	case float64:
		fmt.Printf("Float -> %.2f (Scientific: %e)\n", content, content)
	default:
		fmt.Printf("Unsupported type: %T\n", content)
	}
}

// 3. Пример обработки кастомных ошибок
type NetworkError struct{ Message string }
type ValidationError struct{ Field string }

func (e NetworkError) Error() string    { return "Network Error: " + e.Message }
func (e ValidationError) Error() string { return "Validation Error in field: " + e.Field }

func processRequest(data string) error {
	if data == "invalid" {
		return ValidationError{Field: "email"}
	}
	return NetworkError{Message: "connection timeout"}
}

func handleError(err error) {
	switch e := err.(type) {
	case NetworkError:
		fmt.Printf("Handling Network Error: %s\n", e.Message)
	case ValidationError:
		fmt.Printf("Handling Validation Error: %s\n", e.Field)
	default:
		fmt.Printf("Unknown error type: %T\n", e)
	}
}

// 4. Пример с рефлексией для расширенного анализа
func analyzeType(data interface{}) {
	fmt.Println("\nType analysis:")
	fmt.Println("Type:", reflect.TypeOf(data))
	fmt.Println("Value:", reflect.ValueOf(data))
	fmt.Println("Kind:", reflect.TypeOf(data).Kind())
}
