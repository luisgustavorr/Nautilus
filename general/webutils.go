package General

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var languageUsed = language.BrazilianPortuguese

func urlToTitle(url string) string {
	strings.ReplaceAll(url, "_", " ")
	caser := cases.Title(languageUsed)
	return caser.String(url)
}
func CreateBindInfos(url string) fiber.Map {
	return fiber.Map{
		"Title": "Nautilus ° " + urlToTitle(url),
	}
}
