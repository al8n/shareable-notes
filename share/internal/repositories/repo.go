package repositories

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/al8n/shareable-notes/share/config"
	"github.com/al8n/shareable-notes/share/internal/utils"
	"github.com/al8n/shareable-notes/share/model"
	stdopentracing "github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const mongoOPName = "MongoDB"

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
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, mongoOPName)
	defer span.Finish()

	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	now := time.Now().Unix()
	note = &model.Note{
		Name:          name,
		Content:       content,
		Deactivated:   false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	rst, err = collection.InsertOne(spanCtx, note)
	span.LogKV("operation",  "share note", "db.insertOne", name)
	if err != nil {
		utils.SetTracerSpanError(span, err)
		return "", "", err
	}

	shareID = rst.InsertedID.(primitive.ObjectID).Hex()
	url = cfg.Address + "/share/v1/note/" + base64.URLEncoding.EncodeToString([]byte(shareID))
	return
}

func (repo Repo) PrivateNote(ctx context.Context, id string) (err error)  {
	var (
		cfg = config.GetConfig()
		collection *mongo.Collection
		oid primitive.ObjectID
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, mongoOPName)
	defer span.Finish()

	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	span.LogKV("operation",  "private note", "db.updateOne", id)

	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.SetTracerSpanError(span, err)
		return err
	}

	_, err = collection.UpdateOne(spanCtx, bson.M{"_id": oid}, bson.D{
		{
			"$set",
			bson.D{
				{Key: "deactivated", Value: true},
				{Key: "deactivated_at", Value: time.Now().Unix()},
			},
		},
	})

	if err != nil {
		utils.SetTracerSpanError(span, err)
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
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, mongoOPName)
	defer span.Finish()

	span.LogKV("operation",  "get note", "db.findOne", id)

	collection = repo.MongoDB.Database(cfg.Mongo.DB).Collection(cfg.Mongo.Collection)

	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.SetTracerSpanError(span, err)
		return "", "", err
	}

	err = collection.FindOne(spanCtx,
		bson.D{
			{
				Key: "_id",
				Value: oid,
			},
		},
	).Decode(&note)
	if err != nil {
		utils.SetTracerSpanError(span, err)
		return "", "", err
	}

	if note.Deactivated {
		return "", "", errors.New("note cannot be found")
	}

	name = note.Name
	content = note.Content
	return
}


