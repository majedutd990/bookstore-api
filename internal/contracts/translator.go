package contracts

type (
	Translator interface {
		Translate(lang Lang, key string) string
		TranslateEn(key string) string
	}
	Lang string
)

const (
	FA = "fa"
	EN = "en"
)
