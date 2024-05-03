package cmd

import (
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/MHNightCat/mhcat/config"
	"github.com/MHNightCat/mhcat/handler"
	it "github.com/MHNightCat/mhcat/locales"
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/servusdei2018/shards/v2"
)


func Start() {
	var err error

	initConfig()
	initImageConfig()
	it.InitLocales()
	initCommand()
	connectToMongodb()

	Mgr, err := shards.New("Bot " + cfg.BotConfig.DiscordToken)
	if err != nil {
		log.Error("Error creating manager,", err)
		return
	}

	Mgr.AddHandler(handler.MessageCreate)
	Mgr.AddHandler(handler.OnConnect)
	Mgr.AddHandler(handler.OnSlashCommand)
	Mgr.RegisterIntent(discordgo.IntentsGuildMessages)

	log.Info("Starting shard manager...")

	err = Mgr.Start()
	if err != nil {
		log.Error("Error starting manager,", err)
		return
	}

	for _, v := range slashCommandsList {
		err := Mgr.ApplicationCommandCreate("", v)
		if err != nil {
			log.Error("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	log.Info("Stopping shard manager...")
	Mgr.Shutdown()
	log.Info("Shard manager stopped. Bot is shut down.")
}
