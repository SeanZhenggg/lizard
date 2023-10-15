package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/source/model/bo"
	"lizard/source/utils/request"
	"net/http"
)

type ITrendSrv interface {
	GetTrends(ctx context.Context)
}

func ProviderITrendsSrv(logger logger.ILogger) ITrendSrv {
	return &trendSrv{
		logger: logger,
	}
}

type trendSrv struct {
	logger logger.ILogger
}

func (t *trendSrv) GetTrends(ctx context.Context) {
	client := request.NewClient(t.logger)

	response, err := client.HttpGet(http.MethodGet, map[string]string{
		"hl":  "zh-TW",
		"tz":  "-480",
		"geo": "TW",
		"ns":  "15",
	}, nil)
	if err != nil {
		return
	}

	trend := &bo.DailyTrends{}

	if err := json.Unmarshal(response, trend); err != nil {
		t.logger.Error(xerrors.Errorf("service GetTrends json unmarshal error: %w", err))
		return
	}

	t.logger.Info(fmt.Sprintf("response trends : %v", trend))
}
