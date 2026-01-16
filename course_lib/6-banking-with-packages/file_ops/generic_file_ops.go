package file_ops

import (
	"errors"
	"fmt"
	"os"
)

func Write_Float_To_File(value float64, fileName string) {
	content := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(content), 0644)
	// above: []byte(content) turns the string into a byte slice, which is what WriteFile needs
	// 0644 is the file permission setting, owner can read/write, group can read, others can read
}

func Get_Float_From_File(fileName string) (float64, error) {
	contents, err := os.ReadFile(fileName)
	// above: '_' can be used in place of 'err' to ignore the error value returned but acknowledge that there is one
	var contents_float float64
	if err != nil {
		//here if failed to read file
		return 1.0, errors.New("Failed to parse float from file.")
	} else {
		fmt.Sscan(string(contents), &contents_float)
	}
	return contents_float, nil
}
