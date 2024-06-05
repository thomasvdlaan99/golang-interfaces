package learning

import (
	"fmt"
	"strings"
)

// we use a stringBuilder std package to preallocate strings to run the application faster
func Strb() {

	var strSlice = []string{"r", "é", "s"}
	var stringBuilder strings.Builder

	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}
	var catStr = stringBuilder.String()
	fmt.Printf("\n%v", catStr)

}
