package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Product struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price,omitempty"`
	InStock bool    `json:"in_stock,omitempty"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type Person2 struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Product2 struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Person3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name:  "said",
		Age:   17,
		Email: "said@said.com",
	}

	fmt.Println("---Ex1---")
	jsonPerson := marshalPerson(person)
	fmt.Println(jsonPerson)

	fmt.Println("---Ex2---")
	p := unmarshalPerson(jsonPerson)
	fmt.Println(p)

	fmt.Println("---Ex3---")
	data := map[string]int{
		"apples":  5,
		"oranges": 10,
	}
	jsonMap := marshalMap(data)
	fmt.Println(jsonMap)

	fmt.Println("---Ex4---")
	dataMap := unmarshalToMap(jsonMap)
	fmt.Println(dataMap)

	fmt.Println("---Ex5---")
	product := Product{
		Name: "Cheese",
	}
	jsonProduct := marshalOmittingEmpty(product)
	fmt.Println(jsonProduct)

	fmt.Println("---Ex6---")
	person2 := Person2{
		Name:  "said",
		Age:   17,
		Email: "said@said.com",
		Address: Address{
			Street: "m. knolov",
			City:   "dushanbe",
		},
	}
	jsonPerson2 := marshalEmbeddingStructs(person2)
	fmt.Println(jsonPerson2)
	person2 = unmarshalEmbeddingStructs(jsonPerson2)
	fmt.Println(person2)

	fmt.Println("---Ex7---")

	data2 := map[string]interface{}{
		"name":  "John",
		"age":   30,
		"email": "john@example.com",
	}
	jsonData2 := marshalMapInterface(data2)
	fmt.Println(jsonData2)

	fmt.Println("---Ex8---")
	f, err := os.OpenFile("person.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	err = encoder.Encode(person)

	person = decodeJSONFile("person.json")
	fmt.Println(person)

	fmt.Println("---Ex9---")
	product2 := Product2{
		Name:  "beef",
		Price: 100,
	}
	encodeToFile(product2, "output.json")

	fmt.Println("---Ex10---")
	jsonString := `{"name":"said","age":17,"email":"said@said.com"}`
	jsonStruct := unmarshalIgnoringNil(jsonString)
	fmt.Println(jsonStruct)
}
