package main

import (
	"gozoo.com/simple-structs/Wildlife"
	// Note: The import of a local package must start with the module name as defined in go.mod file.
	//       Then you add the name of the package folder (which must also match the package name, case sensitive).
)

func main() {
	var my_zoo Wildlife.Zoo

	my_zoo.New("The Stupendous Go Zoo")
	my_zoo.Add_Animal("Lion", "Panthera leo", "Fur", "Savannah", "Carnivore")
	my_zoo.Add_Animal("Kiwi", "Apteryx", "Feathers", "Forest", "Omnivore")
	my_zoo.Add_Fish("Clownfish", "Amphiprioninae", "Scales", "Coral Reef", "Omnivore", 12)
	my_zoo.Add_Fish("Goliath Grouper", "Epinephelus", "Scales", "Coral Reef", "Carnivore", 30)
	my_zoo.Save()
}
