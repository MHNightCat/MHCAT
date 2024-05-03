package pkg

import (
	it "github.com/MHNightCat/mhcat/locales"
	"github.com/bwmarrin/discordgo"
)

func SlashCommandLocalizations(translate_id string) *map[discordgo.Locale]string {
	localizations := map[discordgo.Locale]string{
		discordgo.ChineseTW: it.I18n.Tr("zh-TW", translate_id),
		discordgo.ChineseCN: it.I18n.Tr("zh-CN", translate_id),
	}
	return &localizations
}

func SlashCommandOptionsLocalizations(translate_id string) map[discordgo.Locale]string {
	localizations := map[discordgo.Locale]string{
		discordgo.ChineseTW: it.I18n.Tr("zh-TW", translate_id),
		discordgo.ChineseCN: it.I18n.Tr("zh-CN", translate_id),
	}
	return localizations
}