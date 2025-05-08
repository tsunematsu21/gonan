package gonan

import "fmt"

func Speek(speeker Character, message string) {
	fmt.Printf("(%s) %s\n", speeker.DisplayName(), message)
}
