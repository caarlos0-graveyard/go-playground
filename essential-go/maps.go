package main

import "fmt"

func main() {
	age := make(map[string]int)

	age["carlos"] = 24
	age["jordie"] = 21
	age["josh"] = 27
	fmt.Println(age)

	fmt.Println("carlos' age: ", age["carlos"])
	delete(age, "carlos")
	fmt.Println(age)

	m := map[string]int{
		"carlos": 24,
		"blah":   12,
	}

	fmt.Println(m)

	for n, a := range m {
		fmt.Printf("%v is %v years old\n", n, a)
	}

	ma := map[string]int{
		"jeremy": 24,
		"jordie": 22,
		"josh":   27,
	}

	fmt.Println(extractKeys(ma))
}

func extractKeys(m map[string]int) []string {
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}
