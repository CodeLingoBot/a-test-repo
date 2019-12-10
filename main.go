package main

import "fmt"

func main() {
	doThisOrThat(false)
}

func doThisOrThat(flag bool) {
	if flag {
		fmt.Println("Hey")
	} else {
		fmt.Println("There")
	}
}
