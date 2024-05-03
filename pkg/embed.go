package pkg

import (
	"fmt"

	cfg "github.com/MHNightCat/mhcat/config"
	it "github.com/MHNightCat/mhcat/locales"
	"github.com/bwmarrin/discordgo"
)

func ErrorEmbed(message string, err error, locale string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: it.I18n.Tr(locale, message),
		Description: fmt.Sprintf("```%s```", err.Error()),
		Color: ErrorColor,
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: cfg.ImageConfig.FooterIconUrl,
			Text: it.I18n.Tr(locale, "error.report_this_error"),
		},
	}
}

func Successful(message string, locale string, args string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: it.I18n.Tr(locale, message, args),
		Color: SuccessfulColor,
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: cfg.ImageConfig.FooterIconUrl,
			Text: it.I18n.Tr(locale, "info.bio"),
		},
	}
}