package service

import (
	"context"
	"encoding/json"
	encryptTool "github.com/SeanZhenggg/go-utils/encrypt"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/internal/config"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/model/po"
	"lizard/internal/mongo"
	"lizard/internal/repository"
	"lizard/internal/utils/request"
	"regexp"
)

type ITrendSrv interface {
	FetchTrends(ctx context.Context) ([]*po.Trend, error)
	UpsertTrends(ctx context.Context, data []*po.Trend) ([]*po.Trend, error)
}

func ProviderITrendsSrv(logger logger.ILogger, db mongo.IMongoCli, repo repository.ITrendRepository, config config.IConfigEnv) ITrendSrv {
	return &trendSrv{
		logger: logger,
		db:     db,
		repo:   repo,
		cfg:    config,
	}
}

type trendSrv struct {
	logger logger.ILogger
	repo   repository.ITrendRepository
	db     mongo.IMongoCli
	cfg    config.IConfigEnv
}

func (t *trendSrv) FetchTrends(ctx context.Context) ([]*po.Trend, error) {
	client := request.NewClient(t.logger)

	response, err := client.HttpGet(constant.GoogleDailyTrendUrl, map[string]string{
		"hl":  "zh-TW",
		"tz":  "-480",
		"geo": "TW",
		"ns":  "15",
	}, nil)
	if err != nil {
		return nil, xerrors.Errorf("trendsSrv FetchTrends client.HttpGet error: %w", err)
	}

	re, err := regexp.Compile(`{"default":{(.*?)}}`)
	if err != nil {
		return nil, xerrors.Errorf("trendsSrv FetchTrends regexp.Compile error: %w", err)
	}

	matched := re.FindString(string(response))

	trend := &bo.DailyTrends{}
	if err := json.Unmarshal([]byte(matched), trend); err != nil {
		return nil, xerrors.Errorf("trendsSrv FetchTrends json unmarshal error: %w", err)
	}

	if trend.Default == nil {
		return nil, xerrors.Errorf("trendsSrv fetch error: trend.Default is nil")
	}

	poTrends := make([]*po.Trend, 0, len(trend.Default.TrendingSearchesDays))

	for _, tr := range trend.Default.TrendingSearchesDays {
		for _, search := range tr.TrendingSearches {
			poTrend := &po.Trend{
				Date:             tr.Date,
				FormattedTraffic: search.FormattedTraffic,
				ShareUrl:         search.ShareUrl,
			}

			if search.Title != nil {
				poTrend.Title = search.Title.Query
				poTrend.TitleExploreLink = search.Title.ExploreLink
			}

			if search.Image != nil {
				poTrend.Image = search.Image.Source
				poTrend.ImageUrl = search.Image.ImageUrl
				poTrend.ImageNewsUrl = search.Image.NewsUrl
			}

			poTrends = append(poTrends, poTrend)
		}
	}

	return poTrends, nil
}

func (t *trendSrv) UpsertTrends(ctx context.Context, data []*po.Trend) ([]*po.Trend, error) {
	db := t.db.GetCollection(ctx, "trends")
	aiIdInfo, err := t.repo.GetMaxAiIDInfo(ctx, db)
	if err != nil {
		return nil, xerrors.Errorf("trendsSrv t.repo.GetMaxAiIDInfo error : %w", err)
	}

	var nextAiID int64
	if aiIdInfo == nil {
		nextAiID = constant.MIN_ID_VALUE
	} else {
		nextAiID = aiIdInfo.AiID + 1
	}

	for _, search := range data {
		search.UID = search.GenUID()
	}
	trendsInDb, err := t.repo.GetMatchedExistedTrends(ctx, db, data)
	if err != nil {
		return nil, xerrors.Errorf("trendsSrv t.repo.GetMatchedExistedTrends error : %w", err)
	}

	trendsInDbMap := make(map[string]*po.Trend)
	for _, v := range trendsInDb {
		trendsInDbMap[v.UID] = v
	}

	insertTrends := make([]*po.Trend, 0, len(data))
	updateTrends := make([]*po.Trend, 0, len(data))
	for _, search := range data {
		// for insert
		if trendsInDbMap[search.UID] == nil {
			search.AiID = nextAiID
			search.ShortUrl = encryptTool.GenShortUrlById(search.AiID)
			nextAiID += 1
			insertTrends = append(insertTrends, search)
			continue
		}

		// for update
		search.ID = trendsInDbMap[search.UID].ID
		updateTrends = append(updateTrends, search)
	}

	// for insert
	if len(insertTrends) > 0 {
		err = t.repo.BatchInsert(ctx, db, insertTrends)
		if err != nil {
			return nil, xerrors.Errorf("trendsSrv t.repo.BatchInsert error : %w", err)
		}
	}

	// for update
	if len(updateTrends) > 0 {
		err = t.repo.BatchUpdate(ctx, db, updateTrends, trendsInDbMap)
		if err != nil {
			return nil, xerrors.Errorf("trendsSrv t.repo.BatchUpdate error : %w", err)
		}
	}

	return insertTrends, nil
}
