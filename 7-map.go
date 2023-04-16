package main

import (
	"fmt"
)

func main() {
	StudentAge := make(map[string]int)

	StudentAge["Me"] = 25
	StudentAge["John"] = 30
	StudentAge["Joe"] = 22
	StudentAge["Harry"] = 28

	fmt.Println(StudentAge)
	for name, value := range StudentAge {
		fmt.Printf("Name: %s, Age: %d\n", name, value)
	}

	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
