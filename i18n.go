package i18n

import (
	"fmt"
	"github.com/robfig/config"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const messageFilePattern = `^\w+\.[a-zA-Z]{2}$`

var locales = &translator{}

type translator struct {
	messages map[string]*config.Config
}

// loadMessages recursively read and cache all available messages from all message files on the given path.
func (c *translator) loadMessages(path string) {
	c.messages = make(map[string]*config.Config)
	if error := filepath.Walk(path, loadMessageFile); error != nil && !os.IsNotExist(error) {
		fmt.Println("Error reading messages files:", error)
	}
}

func (c *translator) translate(locale, message string, args []interface{}) string {
	language, region := parseLocale(locale)

	messageConfig, knownLanguage := c.messages[language]
	if !knownLanguage {
		fmt.Printf("Unsupported language for locale '%s' and message '%s' \n", locale, message)
		return message
	}

	value, error := messageConfig.String(region, message)

	if error != nil {
		fmt.Printf("Unknown message '%s' for locale '%s'  \n", message, locale)
		return message
	}

	if len(args) > 0 {
		value = fmt.Sprintf(value, args...)
	}

	return value
}

// listLanguages returns all currently loaded message languages.
func (c *translator) listLanguages() []string {
	languages := make([]string, len(c.messages))
	i := 0
	for language, _ := range c.messages {
		languages[i] = language
		i++
	}
	return languages
}

// Load a single message file
func loadMessageFile(path string, info os.FileInfo, osError error) error {
	if osError != nil {
		return osError
	}
	if info.IsDir() {
		return nil
	}
	if matched, _ := regexp.MatchString(messageFilePattern, info.Name()); matched {
		if config, error := config.ReadDefault(path); error != nil {
			return error
		} else {
			locale := parseLocaleFromFileName(info.Name())
			// If already parsed a message file for this locale, merge both
			if _, exists := locales.messages[locale]; exists {
				locales.messages[locale].Merge(config)
			} else {
				locales.messages[locale] = config
			}
		}
	} else {
		fmt.Printf("Ignoring file %s because it did not have a valid extension  \n", info.Name())
	}

	return nil
}

func parseLocaleFromFileName(file string) string {
	extension := filepath.Ext(file)[1:]
	return strings.ToLower(extension)
}

func parseLocale(locale string) (language, region string) {
	if strings.Contains(locale, "-") {
		languageAndRegion := strings.Split(locale, "-")
		return languageAndRegion[0], languageAndRegion[1]
	}
	return locale, ""
}

// LoadMessages loads messages from all message files on the given path.
func LoadMessages(path string) {
	locales.loadMessages(path)
}

// Translate translates content to target language.
func Translate(locale, message string, args ...interface{}) string {
	return locales.translate(locale, message, args)
}

// ListLanguages returns all currently loaded message languages.
func ListLanguages() []string {
	return locales.listLanguages()
}
