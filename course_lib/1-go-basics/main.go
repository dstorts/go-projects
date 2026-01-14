package main

//above is the reserved package name, required to run a build and make an exe, tells go where to start execution

import "fmt" //fmt is a package from Go std lib.

//in order to apply a 'go build' cmd on this app, you need to make this package/collection of packages a module
//in order to initiate this folder as a module you need to run the 'go mod init' command but must include a path the module
//you can use a dummy url and path, just know that the service piece of the url will become the name of the module
//now you can run 'go build'
//on windows that would make an executable, able to be run without Go installed
//in terminal you can call that exe like './<NAME_OF_EXE>' and it will run the same as if you had performed the 'go run' cmd

func main() {
	//while multiple files can be grouped under the 'main' reserved package, the main() must be a singleton in a given package
	//you only NEED a main() if you are going to be providing this pacakge as an exe
	fmt.Print("Hello World.")
}
