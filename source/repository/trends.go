package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
	"lizard/source/model/po"
	"time"
)

type ITrendRepository interface {
	BatchInsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error
}

func ProvideTrendRepository() ITrendRepository {
	return &trendRepo{}
}

type trendRepo struct {
}

func (repo *trendRepo) BatchInsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error {
	writeModels := make([]mongo.WriteModel, 0, len(data))

	for _, trend := range data {
		var trendInDb *po.Trend
		id := trend.GenID()
		err := db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(trendInDb)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				return xerrors.Errorf("trendRepo BatchInsert db.FindOne error : %w", err)
			}

			trend.ID = primitive.NewObjectID()
			trend.CreatedAt = time.Now()
			trend.UpdatedAt = time.Now()
			writeModel := mongo.NewInsertOneModel().SetDocument(trend)
			writeModels = append(writeModels, writeModel)
			continue
		}

		trend.UpdatedAt = time.Now()
		writeModel := mongo.NewUpdateOneModel().SetFilter(bson.M{"_id": trendInDb.ID}).SetUpdate(bson.M{"$set": trend})
		writeModels = append(writeModels, writeModel)
	}

	_, err := db.BulkWrite(context.TODO(), writeModels, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return xerrors.Errorf("trendRepo BatchInsert db.BulkWrite error : %w", err)
	}

	return nil
}
