package po

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Trend struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	AiID             int64              `bson:"ai_id"`
	UID              string             `bson:"uid"`
	ShortUrl         string             `bson:"short_url"`
	Title            string             `bson:"title"`
	TitleExploreLink string             `bson:"title_explore"`
	FormattedTraffic string             `bson:"formatted_traffic"`
	Image            string             `bson:"image"`
	ImageUrl         string             `bson:"image_url"`
	ImageNewsUrl     string             `bson:"image_news_url"`
	Date             string             `bson:"date"`
	ShareUrl         string             `bson:"share_url"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

type TrendMaxAiID struct {
	AiID int64 `bson:"ai_id"`
}

type TrendUrlCond struct {
	ShortUrl string `bson:"short_url"`
}

func (t *Trend) GenUID() string {
	return fmt.Sprintf("%s_%s_%s", t.Date, t.Title, t.ShareUrl)
}

func (t *Trend) ToUpdate(inDb *Trend) bson.M {
	toUpdate := bson.M{}

	if t.Title != inDb.Title {
		toUpdate["title"] = t.Title
	}
	if t.TitleExploreLink != inDb.TitleExploreLink {
		toUpdate["title_explore"] = t.TitleExploreLink
	}
	if t.FormattedTraffic != inDb.FormattedTraffic {
		toUpdate["formatted_traffic"] = t.FormattedTraffic
	}
	if t.Image != inDb.Image {
		toUpdate["image"] = t.Image
	}
	if t.ImageUrl != inDb.ImageUrl {
		toUpdate["image_url"] = t.ImageUrl
	}
	if t.ImageNewsUrl != inDb.ImageNewsUrl {
		toUpdate["image_news_url"] = t.ImageNewsUrl
	}
	if t.Date != inDb.Date {
		toUpdate["date"] = t.Date
	}
	if t.ShareUrl != inDb.ShareUrl {
		toUpdate["share_url"] = t.ShareUrl
	}

	return toUpdate
}
