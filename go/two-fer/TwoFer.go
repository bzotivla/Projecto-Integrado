package twofer

import "fmt"

 
const baseOutput string = "One for %v, one for me."

 
func ShareWith(input string) string {
	value := input
	if len(input) == 0 {
		value = "you"
	}
	return fmt.Sprintf(baseOutput, value)
}