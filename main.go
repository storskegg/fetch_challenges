package main

import (
	"fmt"
)

func main() {
	fmt.Println("===================== [ Version ]======================")
	fmt.Println(VersionStringCompare("1.0.0", "1.0.1"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("2.0", "1.0.17"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("1.18.0", "1.18.0"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("1.18.0", "1.18"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("2.1", "2.1.0"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("1.18.1", "1.18"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println(VersionStringCompare("2.1", "2.1.1"))
	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------------------------------------------------")
	fmt.Println()
	fmt.Println("===================== [ Pyramid ]======================")
	fmt.Println(IsPyramidWord("banaNa"))
}
