package Wildlife

// Remember: While the name of the files in a package can be anything, the package name itself (line above)
//  should be the same as this containing folder name.
import (
	"errors"
	"fmt"
)

/*
------Public vs Private Attributes in Go Structs------
In Go, the visibility of struct attributes (fields) is determined by the case of the first letter of the attribute name.
So, in the example below, the Name attribute is public (exported) because it starts with an uppercase letter,
while the species, skin_type, ecosystem, and diet attributes are private (unexported) because they start with lowercase letters.
This means that code which uses this package can create an Animal struct, directly access the Name attribute,
but cannot directly access the species, skin_type, ecosystem, or diet attributes. The function 'get_ecosystem' is also private,
because its name starts with a lowercase letter. It can only be called from within the scope of this package.
*/

type Animal struct {
	Name      string
	species   string
	skin_type string
	ecosystem string
	diet      string
}

// Public method: accessible from outside the Wildlife package
func (a *Animal) New(name, species, skin_type, ecosystem, diet string) error {
	if name == "" {
		return errors.New("The Name of the animal must be prvided.")
	}

	(*a).Name = name    //true syntax to dereference the name attribute of the current struct pointer
	a.species = species //Go allows you to skip the dereference operator on struct attributes like this
	a.skin_type = skin_type
	a.ecosystem = ecosystem
	a.diet = diet
	return nil
}

// Private method: only accessible within the Wildlife package
func (a *Animal) get_ecosystem() string {
	return a.ecosystem
}

func (a Animal) Print_Profile() {
	fmt.Println("|----------------|")
	fmt.Printf("| Name     : %s\n", a.Name)
	fmt.Printf("| Species  : %s\n", a.species)
	fmt.Printf("| Skin Type: %s\n", a.skin_type)
	fmt.Printf("| Ecosystem: %s\n", a.ecosystem)
	fmt.Printf("| Diet     : %s\n", a.diet)
	fmt.Println("|----------------|")
}
