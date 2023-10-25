package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
	"lizard/internal/model/po"
	"time"
)

type ITrendRepository interface {
	GetAiIDInfo(ctx context.Context, db *mongo.Collection) *po.TrendAiIDInfo
	BatchUpsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error
}

func ProvideTrendRepository() ITrendRepository {
	return &trendRepo{}
}

type trendRepo struct {
}

func (repo *trendRepo) GetAiIDInfo(ctx context.Context, db *mongo.Collection) *po.TrendAiIDInfo {
	var res *po.TrendAiIDInfo
	if err := db.
		FindOne(
			ctx,
			bson.M{},
			options.FindOne().SetSort(bson.M{"ai_id": -1}).SetProjection(bson.M{"ai_id": 1})).
		Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res
		}
	}

	return res
}

func (repo *trendRepo) BatchUpsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error {
	writeModels := make([]mongo.WriteModel, 0, len(data))

	for _, trend := range data {
		var trendInDb po.Trend

		uid := trend.GenUID()
		err := db.FindOne(ctx, bson.M{"uid": uid}).Decode(&trendInDb)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				return xerrors.Errorf("trendRepo BatchUpsert db.FindOne error : %w", err)
			}

			// insert
			trend.ID = primitive.NewObjectID()
			trend.UID = uid
			trend.CreatedAt = time.Now()
			trend.UpdatedAt = time.Now()
			insertModel := mongo.NewInsertOneModel().SetDocument(trend)
			writeModels = append(writeModels, insertModel)
			continue
		}

		// update
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": trendInDb.ID}).
			SetUpdate(
				bson.M{"$set": bson.M{
					"title":             trend.Title,
					"title_explore":     trend.TitleExploreLink,
					"formatted_traffic": trend.FormattedTraffic,
					"image":             trend.Image,
					"image_url":         trend.ImageUrl,
					"image_news_url":    trend.ImageNewsUrl,
					"updated_at":        time.Now(),
				}})
		writeModels = append(writeModels, updateModel)
	}

	_, err := db.BulkWrite(ctx, writeModels, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return xerrors.Errorf("trendRepo BatchUpsert db.BulkWrite error : %w", err)
	}

	return nil
}
