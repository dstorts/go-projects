package main

import (
	"fmt"

	"gozoo.com/simple-structs/Wildlife"
	// Note: The import of a local package must start with the module name as defined in go.mod file.
	//       Then you add the name of the package folder (which must also match the package name, case sensitive).
)

func main() {
	var dog Wildlife.Animal

	dog.New("Dog", "Canis Lupus Familiaris", "Fur", "Domestic", "Omnivore")

	dog.Print_Profile()
	fmt.Println(dog.Name)

	var goldfish Wildlife.Fish
	goldfish.New_Fish("Goldfish", "Carassius auratus", "Scales", "Freshwater", "Omnivore", 15)
	goldfish.Print_Fish_Profile()
}
