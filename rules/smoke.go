package main

import "fmt"

func main() {
		// ruleid: tqa-smoke-test
		fmt.Println("this should be flagged")

		// ok: tqa-smoke-test
		fmt.Printf("this should not be flagged\n")
}