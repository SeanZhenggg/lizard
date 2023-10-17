package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/source/model/bo"
	"lizard/source/utils/request"
	"regexp"
)

type ITrendSrv interface {
	GetTrends(ctx context.Context) (*bo.DailyTrends, error)
}

func ProviderITrendsSrv(logger logger.ILogger) ITrendSrv {
	return &trendSrv{
		logger: logger,
	}
}

type trendSrv struct {
	logger logger.ILogger
}

func (t *trendSrv) GetTrends(ctx context.Context) (*bo.DailyTrends, error) {
	client := request.NewClient(t.logger)

	response, err := client.HttpGet("https://trends.google.com/trends/api/dailytrends", map[string]string{
		"hl":  "zh-TW",
		"tz":  "-480",
		"geo": "TW",
		"ns":  "15",
	}, nil)
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(`{"default":{(.*?)}}`)
	if err != nil {
		return nil, err
	}

	matched := re.FindString(string(response))

	trend := &bo.DailyTrends{}
	if err := json.Unmarshal([]byte(matched), trend); err != nil {
		t.logger.Error(xerrors.Errorf("service GetTrends json unmarshal error: %w", err))
		return nil, err
	}

	t.logger.Info(fmt.Sprintf("response trends : %v", trend))

	return trend, nil
}
