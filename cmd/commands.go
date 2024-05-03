package cmd

import (
	slashcommand "github.com/MHNightCat/mhcat/slash_command"
	"github.com/bwmarrin/discordgo"
)

var slashCommandsList = []*discordgo.ApplicationCommand{
	&slashcommand.LocalesCommand,
}