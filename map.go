package main

import "fmt"

func main() {
	capitalCity := map[string]string{
		"Indonesia": "Jakarta",
		"France":    "Paris",
		"Singapore": "Singapore",
	}

	fmt.Println(capitalCity)
	fmt.Println(capitalCity["Indonesia"])
	fmt.Println(capitalCity["Malaysia"]) // Will print the zero value of string

	if value, ok := capitalCity["Malaysia"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value does not exist")
	}

	capitalCity["Malaysia"] = "Kuala Lumpur"
	fmt.Println(capitalCity["Malaysia"])

	for key, value := range capitalCity {
		fmt.Printf("Capital city of %v is %v\n", key, value)
	}

	delete(capitalCity, "Singapore")
	fmt.Println(capitalCity)
}
