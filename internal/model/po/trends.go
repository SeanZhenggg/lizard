package po

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Trend struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	AiID             int64              `bson:"ai_id"`
	UID              string             `bson:"uid"`
	Title            string             `bson:"title"`
	TitleExploreLink string             `bson:"title_explore"`
	FormattedTraffic string             `bson:"formatted_traffic"`
	Image            string             `bson:"image"`
	ImageUrl         string             `bson:"image_url"`
	ImageNewsUrl     string             `bson:"image_news_url"`
	Date             string             `bson:"date"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

type TrendAiIDInfo struct {
	AiID int64 `bson:"ai_id"`
}

func (t *Trend) GenUID() string {
	return fmt.Sprintf("%s_%s_%s", t.Date, t.Title, t.TitleExploreLink)
}
