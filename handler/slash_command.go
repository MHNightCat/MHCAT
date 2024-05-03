package handler

import (
	slashcommand "github.com/MHNightCat/mhcat/slash_command"
	"github.com/bwmarrin/discordgo"
)

func OnSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"set-language": slashcommand.LocalesCommandRun,
}