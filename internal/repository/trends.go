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
	GetMaxAiIDInfo(ctx context.Context, db *mongo.Collection) (*po.TrendMaxAiID, error)
	GetMatchedExistedTrends(ctx context.Context, db *mongo.Collection, data []*po.Trend) ([]*po.Trend, error)
	BatchInsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error
	BatchUpdate(ctx context.Context, db *mongo.Collection, data []*po.Trend, inDbMap map[string]*po.Trend) error
	GetTrendByUrl(ctx context.Context, db *mongo.Collection, cond *po.TrendUrlCond) (*po.Trend, error)
}

func ProvideTrendRepository() ITrendRepository {
	return &trendRepo{}
}

type trendRepo struct {
}

func (repo *trendRepo) GetMaxAiIDInfo(ctx context.Context, db *mongo.Collection) (*po.TrendMaxAiID, error) {
	var res *po.TrendMaxAiID
	if err := db.
		FindOne(
			ctx,
			bson.M{},
			options.FindOne().SetSort(bson.M{"ai_id": -1}).SetProjection(bson.M{"ai_id": 1})).
		Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, xerrors.Errorf("trendRepo GetMaxAiIDInfo FindOne.Decode error : %w", err)
	}

	return res, nil
}

func (repo *trendRepo) GetMatchedExistedTrends(ctx context.Context, db *mongo.Collection, data []*po.Trend) ([]*po.Trend, error) {
	currentUIds := make([]string, 0, len(data))
	var trendsInDb []*po.Trend

	for _, trend := range data {
		currentUIds = append(currentUIds, trend.UID)
	}

	result, err := db.Find(ctx, bson.M{"uid": bson.M{"$in": currentUIds}})
	defer func() { result.Close(ctx) }()

	if err != nil {
		return nil, xerrors.Errorf("trendRepo GetMatchedExistedTrends db.Find error : %w", err)
	}

	err = result.All(ctx, &trendsInDb)
	if err != nil {
		return nil, xerrors.Errorf("trendRepo GetMatchedExistedTrends db.Find.All error : %w", err)
	}

	return trendsInDb, nil
}

func (repo *trendRepo) BatchInsert(ctx context.Context, db *mongo.Collection, data []*po.Trend) error {
	writeModels := make([]mongo.WriteModel, 0, len(data))

	for _, trend := range data {
		trend.ID = primitive.NewObjectID()
		trend.CreatedAt = time.Now()
		trend.UpdatedAt = time.Now()
		insertModel := mongo.NewInsertOneModel().SetDocument(trend)
		writeModels = append(writeModels, insertModel)
	}

	_, err := db.BulkWrite(ctx, writeModels, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return xerrors.Errorf("trendRepo BatchInsert db.BulkWrite error : %w", err)
	}

	return nil
}

func (repo *trendRepo) BatchUpdate(ctx context.Context, db *mongo.Collection, data []*po.Trend, inDbMap map[string]*po.Trend) error {
	writeModels := make([]mongo.WriteModel, 0, len(data))

	for _, trend := range data {
		toUpdate := trend.ToUpdate(inDbMap[trend.UID])
		if len(toUpdate) > 0 {
			toUpdate["updated_at"] = time.Now()
			updateModel := mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": trend.ID}).
				SetUpdate(
					bson.M{"$set": toUpdate})
			writeModels = append(writeModels, updateModel)
		}
	}

	if len(writeModels) > 0 {
		_, err := db.BulkWrite(ctx, writeModels, options.BulkWrite().SetOrdered(false))
		if err != nil {
			return xerrors.Errorf("trendRepo BatchUpdate db.BulkWrite error : %w", err)
		}
	}

	return nil
}

func (repo *trendRepo) GetTrendByUrl(ctx context.Context, db *mongo.Collection, cond *po.TrendUrlCond) (*po.Trend, error) {
	var res *po.Trend
	if err := db.FindOne(ctx, cond).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, xerrors.Errorf("trendRepo GetTrendByUrl FindOne error : %w", err)
	}

	return res, nil
}
