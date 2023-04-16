package main

import "fmt"

func main() {
	var age int = 18

	if age < 18 {
		fmt.Println("not old enough")
	} else if age == 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("older generate also can vote")
	}

	switch age {
	case 16:
		fmt.Println("prepare collage")
	case 18:
		fmt.Println("dont run after a girl")
	case 20:
		fmt.Println("get yourself a job")
	}
}
