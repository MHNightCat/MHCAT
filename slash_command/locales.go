package slashcommand

import (
	"github.com/MHNightCat/mhcat/db"
	it "github.com/MHNightCat/mhcat/locales"
	"github.com/MHNightCat/mhcat/model"
	"github.com/MHNightCat/mhcat/pkg"
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

var LocalesCommand = discordgo.ApplicationCommand{}

func InitLocalesCommand() {
	LocalesCommand = discordgo.ApplicationCommand{
		Name:                     it.I18n.Tr("en", "slashcmd.locales.name"),
		Description:              it.I18n.Tr("en", "slashcmd.locales.description"),
		NameLocalizations:        pkg.SlashCommandLocalizations("slashcmd.locales.name"),
		DescriptionLocalizations: pkg.SlashCommandLocalizations("slashcmd.locales.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     it.I18n.Tr("en", "slashcmd.locales.options.language.name"),
				Description:              it.I18n.Tr("en", "slashcmd.locales.options.language.description"),
				NameLocalizations:        pkg.SlashCommandOptionsLocalizations("slashcmd.locales.options.language.name"),
				DescriptionLocalizations: pkg.SlashCommandOptionsLocalizations("slashcmd.locales.options.language.description"),
				Required:                 true,
				Type:                     discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "English",
						Value: "en",
					},
					{
						Name:  "繁體中文",
						Value: "zh-TW",
					},
					{
						Name:  "简体中文",
						Value: "zh-CN",
					},
				},
			},
		},
	}
}

func LocalesCommandRun(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	insertData := model.Lanuage{
		GuildId: i.GuildID,
		Language: optionMap["language"].Value.(string),
	}

	udateData, err := db.FileOneAndUpdateGuildLanguageSetting(insertData)
	if err != nil {
		log.Error("Error update language data to database, data:", insertData, err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					pkg.ErrorEmbed("slashcmd.locales.error.error_database_operate", nil, udateData.Language),
				},
			},
		})
		return
	}
	
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				pkg.Successful("slashcmd.locales.successful.successful_update_data", udateData.Language, optionMap["language"].Value.(string)),
			},
		},
	})
}
