package controllers

import "lizard/source/mongo"

type ITrendCtrl interface {
}

func ProviderITrendsCtrl(db mongo.IMongoCli) ITrendCtrl {
	return &trendCtrl{
		db: db,
	}
}

type trendCtrl struct {
	db mongo.IMongoCli
}
