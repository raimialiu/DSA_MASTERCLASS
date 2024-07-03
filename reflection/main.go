package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	Name, Category string
	Price          float64
}
type Customer struct {
	Name, City string
}

func printDetails(values ...Product) {
	for _, elem := range values {
		fmt.Printf("Product: Name: %v, Category: %v, Price: %v",
			elem.Name, elem.Category, elem.Price)
	}
}

func printDetailsFunc(value ...interface{}) {
	for _, v := range value {
		elemType := reflect.TypeOf(v)
		elemValue := reflect.ValueOf(v)

		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				field := elemType.Field(i)
				fieldValue := elemValue.Field(i)
				fmt.Printf("Field: %v, FieldValue: %v, ", field.Name, fieldValue)
			}
		}
	}
}

func main() {

	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}

	printDetailsFunc(product)
}
