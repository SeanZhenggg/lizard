package service

import (
	"context"
	"encoding/json"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/source/model/bo"
	"lizard/source/model/po"
	"lizard/source/mongo"
	"lizard/source/repository"
	"lizard/source/utils/request"
	"regexp"
)

type ITrendSrv interface {
	FetchTrends(ctx context.Context) error
}

func ProviderITrendsSrv(logger logger.ILogger, db mongo.IMongoCli, repo repository.ITrendRepository) ITrendSrv {
	return &trendSrv{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}

type trendSrv struct {
	logger logger.ILogger
	repo   repository.ITrendRepository
	db     mongo.IMongoCli
}

func (t *trendSrv) FetchTrends(ctx context.Context) error {
	client := request.NewClient(t.logger)

	response, err := client.HttpGet("https://trends.google.com/trends/api/dailytrends", map[string]string{
		"hl":  "zh-TW",
		"tz":  "-480",
		"geo": "TW",
		"ns":  "15",
	}, nil)
	if err != nil {
		return xerrors.Errorf("trendsSrv FetchTrends client.HttpGet error: %w", err)
	}

	re, err := regexp.Compile(`{"default":{(.*?)}}`)
	if err != nil {
		return xerrors.Errorf("trendsSrv FetchTrends regexp.Compile error: %w", err)
	}

	matched := re.FindString(string(response))

	trend := &bo.DailyTrends{}
	if err := json.Unmarshal([]byte(matched), trend); err != nil {
		return xerrors.Errorf("trendsSrv FetchTrends json unmarshal error: %w", err)
	}

	if trend.Default == nil {
		return xerrors.Errorf("trendsSrv fetch error: trend.Default is nil")
	}

	var testFlag bool
	poTrends := make([]*po.Trend, 0, len(trend.Default.TrendingSearchesDays))
	for _, tr := range trend.Default.TrendingSearchesDays {
		poTrend := &po.Trend{}
		poTrend.Date = tr.Date

		for i, search := range tr.TrendingSearches {
			if i > 0 {
				testFlag = true
				break
			}

			if search.Title != nil {
				poTrend.Title = search.Title.Query
				poTrend.TitleExploreLink = search.Title.ExploreLink
			}

			poTrend.FormattedTraffic = search.FormattedTraffic
			if search.Image != nil {
				poTrend.Image = search.Image.Source
				poTrend.ImageUrl = search.Image.ImageUrl
				poTrend.ImageNewsUrl = search.Image.NewsUrl
			}

			poTrends = append(poTrends, poTrend)
		}

		if testFlag {
			break
		}
	}

	db := t.db.GetCollection(ctx, "trends")
	err = t.repo.BatchInsert(ctx, db, poTrends)
	if err != nil {
		return xerrors.Errorf("trendsSrv FetchTrends t.repo.BatchInsert error: %w", err)
	}

	return nil
}
