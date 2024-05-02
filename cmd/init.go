package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml/v2"
	"github.com/kataras/i18n"
)

var I18n *i18n.I18n

type MHCATConfigType struct {
	DiscordToken string `toml:"discord_token"`
}

var MHCATConfig MHCATConfigType

func initBot() {
	fileData , err := os.ReadFile("config/config.toml")
	
	if err != nil {
		log.Error("Fail to load MHCAT config file,", err)
		return
	}

	err = toml.Unmarshal(fileData, &MHCATConfig)

	if err != nil {
		log.Error("Fail to unmarshal mhcat config file,", err)
		return
	}
}

func initLocales() {
	var err error
	I18n, err = i18n.New(i18n.Glob("./locales/*/*"), "en", "el-GR", "zh-CN")
	
	if err != nil {
		log.Error("Fail to load locales file", err)
		return
	}
	test := I18n.Tr("en", "slash_command.locales.options.language.name")
	log.Info(test)
}