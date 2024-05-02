package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/MHNightCat/mhcat/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/servusdei2018/shards/v2"
)

func Start() {
	initBot()
	initLocales()
	var err error

	Mgr, err := shards.New("Bot " + MHCATConfig.DiscordToken)
	if err != nil {
		log.Error("Error creating manager,", err)
		return
	}

	Mgr.AddHandler(handler.MessageCreate)
	Mgr.AddHandler(handler.OnConnect)

	Mgr.RegisterIntent(discordgo.IntentsGuildMessages)

	log.Info("Starting shard manager...")

	err = Mgr.Start()
	if err != nil {
		log.Error("Error starting manager,", err)
		return
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	log.Info("Stopping shard manager...")
	Mgr.Shutdown()
	log.Info("Shard manager stopped. Bot is shut down.")
}