package i18n

import (
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
    "gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

func Init() {
    bundle = i18n.NewBundle(language.English)
    bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

    // Load translation files
    bundle.MustLoadMessageFile("locales/en.yaml")
    bundle.MustLoadMessageFile("locales/vi.yaml")
}

func Translate(lang, messageID string, templateData map[string]interface{}) string {
    localizer := i18n.NewLocalizer(bundle, lang)
    message, _ := localizer.Localize(&i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: templateData,
    })
    return message
}
