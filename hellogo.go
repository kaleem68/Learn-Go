package main

import (
	"fmt"
	"strings"
)

var out = fmt.Println

func main() {
	testStrings()
}
func testStrings() {
	//replace string
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	out(sV2)

	//string length
	out(len(sV1))

	//contains
	out(sV1, " contains ", "word ", strings.Contains(sV1, "word"))

	//first matching index
	out(strings.Index(sV1, "ord"))

	sV3 := "word of worn"
	//replace all 'o' with '0'
	// if last instead of -1 we had 2 that means replace first 2 occurrence
	sV4 := strings.Replace(sV3, "o", "0", -1)
	out(sV4)

	//remove whitespace from beg and end of the string
	sV3 = "\t Some words"
	sV3 = strings.TrimSpace(sV3)
	out(sV3)

	//split based on "-"
	sV5 := "a-b-c-d"
	out(strings.Split(sV5, "-"))

	//lower and upercase
	out(strings.ToLower("LOts of Love"))
	out(strings.ToUpper("lots of Love"))

	//prefix and suffix
	out(strings.HasPrefix("tree", "tr"))
	out(strings.HasSuffix("tree", "ee"))

}
