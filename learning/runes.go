package learning

import "fmt"

func Runes() {
	myStringRunes := []rune("Résumé")
	myString := "Résumé"
	fmt.Println(myStringRunes)
	fmt.Println(myString)

	for i, v := range myString {
		fmt.Println(i, v)
	}
}
