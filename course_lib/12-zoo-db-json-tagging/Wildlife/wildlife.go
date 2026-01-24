package Wildlife

// Remember: While the name of the files in a package can be anything, the package name itself (line above)
//  should be the same as this containing folder name.
import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Zoo struct {
	Title         string   `json:"title"`
	Basic_Animals []Animal `json:"animals"`
	Fish_Types    []Fish   `json:"fishes"`
}

func (z *Zoo) New(title string) {
	z.Title = title
}

func (z *Zoo) Add_Animal(name, species, skin_type, ecosystem, diet string) error {
	var animal Animal
	err := animal.New(name, species, skin_type, ecosystem, diet)
	if err != nil {
		return err
	}
	z.Basic_Animals = append(z.Basic_Animals, animal)
	return nil
}

func (z *Zoo) Add_Fish(name, species, skin_type, ecosystem, diet string, expected_depth int) {
	var fish Fish
	fish.New_Fish(name, species, skin_type, ecosystem, diet, expected_depth)
	z.Fish_Types = append(z.Fish_Types, fish)
}

func (z Zoo) Save() {
	fileName := strings.ReplaceAll(z.Title, " ", "_")

	json, err := json.MarshalIndent(z, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	os.WriteFile(fmt.Sprintf("%v.json", fileName), json, 0644)
}

type Animal struct {
	Name      string `json:"name"`
	Species   string `json:"species"`
	Skin_type string `json:"skin_type"`
	Ecosystem string `json:"ecosystem"`
	Diet      string `json:"diet"`
}

// Public method: accessible from outside the Wildlife package
func (a *Animal) New(name, species, skin_type, ecosystem, diet string) error {
	if name == "" {
		return errors.New("The Name of the animal must be prvided.")
	}

	(*a).Name = name
	a.Species = species
	a.Skin_type = skin_type
	a.Ecosystem = ecosystem
	a.Diet = diet
	return nil
}

func (a Animal) Print_Profile() {
	fmt.Println("|----------------|")
	fmt.Printf("| Name     : %s\n", a.Name)
	fmt.Printf("| Species  : %s\n", a.Species)
	fmt.Printf("| Skin Type: %s\n", a.Skin_type)
	fmt.Printf("| Ecosystem: %s\n", a.Ecosystem)
	fmt.Printf("| Diet     : %s\n", a.Diet)
}

type Fish struct {
	Animal             // Embedding the Animal struct
	Expected_Depth int `json:"expected_depth"`
}

func (f *Fish) New_Fish(name, species, skin_type, ecosystem, diet string, expected_depth int) {
	f.New(name, species, skin_type, ecosystem, diet)
	f.Expected_Depth = expected_depth
}

func (f *Fish) Print_Fish_Profile() {
	f.Print_Profile()
	fmt.Printf("| Expected Depth: %d meters\n", f.Expected_Depth)
}
