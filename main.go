package main

import (
	"fmt"
	"intl/translator"

	"golang.org/x/text/language"
)

func main() {
	translator := translator.GetTranslator()
	translator.Init(language.English)
	// translator.Init(language.Hebrew)

	fmt.Println(translator.Translate("HELLO", nil))
	fmt.Println(translator.Translate("GOODBYE", map[string]string{"Name": "Guy"}))
}
