package cfg

type BotConfigType struct {
	DiscordToken         string `toml:"discord_token"`
	MongodbConnectString string `toml:"mongodb_connect_string"`
	MongodbDatabaseName  string `toml:"mongodb_database_name"`
}

type ImageConfigType struct {
	FooterIconUrl string `toml:"footer_icon_url"`
}

var BotConfig BotConfigType
var ImageConfig ImageConfigType
