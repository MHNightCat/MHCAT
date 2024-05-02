package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)


type MHCATConfigType struct {
	DiscordToken string `toml:"discord_token"`
}

var MHCATConfig MHCATConfigType

func initBot() {
	fileData , err := os.ReadFile("config/config.toml")
	
	if err != nil {
		log.Error("Fail to load MHCAT config file,", err)
	}

	err = toml.Unmarshal(fileData, &MHCATConfig)

	if err != nil {
		log.Error("Fail to unmarshal mhcat config file,", err)
	}
}

func initLocales() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)

	bundle.MustLoadMessageFile("locales/en.json")
	
}