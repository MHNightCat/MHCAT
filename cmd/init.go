package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"
	"time"

	cfg "github.com/MHNightCat/mhcat/config"
	"github.com/MHNightCat/mhcat/db"
	slashcommand "github.com/MHNightCat/mhcat/slash_command"
	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func initConfig() {
	fileData, err := os.ReadFile("config/config.toml")

	if err != nil {
		log.Error("Fail to load MHCAT config file,", err)
		return
	}

	err = toml.Unmarshal(fileData, &cfg.BotConfig)

	if err != nil {
		log.Error("Fail to unmarshal mhcat config file,", err)
		return
	}
	
}

func initImageConfig() {
	fileData, err := os.ReadFile("config/imageConfig.toml")

	if err != nil {
		log.Error("Fail to load MHCAT iamge config file,", err)
		return
	}

	err = toml.Unmarshal(fileData, &cfg.ImageConfig)

	if err != nil {
		log.Error("Fail to unmarshal mhcat iamge config file,", err)
		return
	}
	
}

func initCommand() {
	slashcommand.InitLocalesCommand()
}

func connectToMongodb() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.BotConfig.MongodbConnectString))

	if err != nil {
		log.Error("Error connect to mongodb", err)
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Fail to connect to mongodb", err)
		return
	}
	db.Database = client.Database(cfg.BotConfig.MongodbDatabaseName)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collections, _ := db.Database.ListCollectionNames(ctx, bson.D{})

	// if database doesn't exist than create one
	if len(collections) == 0 {
		log.Info(fmt.Sprintf("Database '%s' does not exist. Creating...", db.Database.Name()))

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := db.Database.Collection("temp").InsertOne(ctx, bson.M{"check": true})
		if err != nil {
			log.Fatal(err)
		}
		log.Info(fmt.Sprintf("Database '%s' created successfully.", db.Database.Name()))
	} else if slices.Contains(collections, "temp") && len(collections) > 1 {

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db.Database.Collection("temp").Drop(ctx)
	}

	db.DatabaseCollectionSet()

	log.Info("Successful connect to mongodb")
}
