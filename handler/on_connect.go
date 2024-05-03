package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

func OnConnect(s *discordgo.Session, evt *discordgo.Connect) {
	log.Info(fmt.Sprintf("Shard #%v connected.", s.ShardID))
}
