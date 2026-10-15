package db

import (
	"context"
	"fmt"
	"golink/internal/infra/env"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnectMongo(env *env.ENV) *mongo.Collection {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s/",
		env.MONGO_ROOT_USER,
		env.MONGO_ROOT_PASSWORD,
		env.MONGO_PORT,
	)

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌​ Erro ao ligar ao MongoDB.", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌​ Não foi possível alcançar o MongoDB.", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	collection := client.Database(env.MONGO_DB_NAME).Collection(env.MONGO_COLLECTION)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("❌ Erro ao criar índice único para e-mail: ", err)
	}

	return collection
}
