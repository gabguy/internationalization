package translator

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var translator = Translator{}

type Translator struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func (t *Translator) Init(langTag language.Tag) {
	(*t).bundle = i18n.NewBundle(langTag)
	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	lang := langTag.String()
	t.bundle.LoadMessageFile(fmt.Sprintf("i18n/%s.json", lang))
	t.localizer = i18n.NewLocalizer(t.bundle, lang)
}

func (t Translator) Translate(key string, data map[string]string) string {
	return t.localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: key},
		TemplateData:   data,
	})
}

func GetTranslator() *Translator {
	return &translator
}
