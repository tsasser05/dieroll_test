package main

import (
	"fmt"

	"github.com/tsasser05/dieroll"
)

func main() {
	dieroll.Init()
	//dierollsvc.Example()
	var numRolls int = 16
	var myRolls = make([]int, numRolls)
	//fmt.Println("Initialized slice:  ", myRolls)
	myRolls = dieroll.Rolldice(10, numRolls)
	fmt.Println("The results of the dice rolled:  ", myRolls)
}
