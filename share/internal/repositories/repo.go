package repositories

import (
	"context"
	"encoding/base64"
	"github.com/ALiuGuanyan/margin/share/config"
	"github.com/ALiuGuanyan/margin/share/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


type Repo struct {
	MongoDB *mongo.Client
}

func NewRepo() (repo *Repo, err error ) {

	var (
		client *mongo.Client
		opt *options.ClientOptions
	)

	opt, err = config.GetConfig().Mongo.Standardize()
	if err != nil {
		return nil, err
	}

	if client, err = mongo.Connect(context.TODO(), opt); err != nil {
		return
	}

	return &Repo{
		MongoDB: client,
	}, nil
}

func (repo Repo) ShareNote(ctx context.Context, name, content string) (url, shareID string, err error)  {
	var (
		cfg = config.GetConfig()
		rst *mongo.InsertOneResult
		collection *mongo.Collection
		note *model.Note
	)

	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	now := time.Now().Unix()
	note = &model.Note{
		Name:          name,
		Content:       content,
		Deactivated:   false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	rst, err = collection.InsertOne(ctx, note)
	if err != nil {
		return "", "", err
	}

	shareID = rst.InsertedID.(primitive.ObjectID).Hex()
	url = cfg.Address + "/share/note/" + base64.URLEncoding.EncodeToString([]byte(shareID))
	return
}

func (repo Repo) PrivateNote(ctx context.Context, id string) (err error)  {
	var (
		cfg = config.GetConfig()
		collection *mongo.Collection
		oid primitive.ObjectID
	)
	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.D{
		{
			"$set",
			bson.D{
				{Key: "deactivated", Value: true},
				{Key: "deactivated_at", Value: time.Now().Unix()},
			},
		},
	})

	if err != nil {
		return err
	}

	return
}


func (repo Repo) GetNote(ctx context.Context, id string) (name, content string, err error)  {
	var (
		cfg = config.GetConfig()
		collection *mongo.Collection
		note model.Note
		oid primitive.ObjectID
	)

	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", "", err
	}

	err = collection.FindOne(ctx,
		bson.D{
			{
				Key: "_id",
				Value: oid,
			},
		},
	).Decode(&note)
	if err != nil {
		return "", "", err
	}

	name = note.Name
	content = note.Content
	return
}


