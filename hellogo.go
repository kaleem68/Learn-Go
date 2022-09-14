package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	conditional()
}
func conditional() {
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important")
	} else if iAge >= 65 {
		pl("Important")
	} else {
		pl("Not Important")
	}
}
