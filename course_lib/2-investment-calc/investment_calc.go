package main

//above is the reserved package name, required to run a build and make an exe, tells go where to start execution

//fmt is a package from Go std lib.
import (
	"fmt"
	"math"
)

//in order to apply a 'go build' cmd on this app, you need to make this package/collection of packages a module
//in order to initiate this folder as a module you need to run the 'go mod init' command but must include a path the module
//you can use a dummy url and path, just know that the service piece of the url will become the name of the module
//ex: go mod init example.com/investment-calc
//now you can run 'go build'
//on windows that would make an executable, able to be run without Go installed
//in terminal you can call that exe like './<NAME_OF_EXE>' and it will run the same as if you had performed the 'go run' cmd
//you can then run 'go run .' on a terminal in the dir of your initialized module, instead of 'go run <file_name>.go'

func main() {
	const inflationRate = 2.5 //always need an initial value

	var investmentAmount float64 //can declare wihtout value IFF you use the 'var' keyword declaration
	expectedReturnRate := 5.5    //the ':=' assignment operator is a way to declare a var without the var keyword IFF you will be using Go's infered var type
	var years float64 = 10

	//investmentAmount, years, expectedReturnRate := 10000.0, 10.0, 5.5

	fmt.Print("Please enter inital investment amount:")
	//'&' is the pointer char
	fmt.Scan(&investmentAmount)
	fmt.Print("Please enter num. years of investment:")
	fmt.Scan(&years)
	fmt.Print("Please enter expected return rate:")
	fmt.Scan(&expectedReturnRate)

	//var futureValue = float64(investmentAmount) * math.Pow(1+(expectedReturnRate/100), float64(years))
	var futureValue = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+(inflationRate/100), years)

	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	formattedV := fmt.Sprintf(
		`Future Value: %.1f
Future Value (adjusted for inflation): %.1f`,
		futureValue,
		futureRealValue)

	fmt.Print(formattedV)
}
