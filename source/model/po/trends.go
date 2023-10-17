package po

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Trend struct {
	ID               primitive.ObjectID `bson:"_id"`
	Title            string             `bson:"title"`
	TitleExploreLink string             `bson:"title_explore"`
	FormattedTraffic string             `bson:"formatted_traffic"`
	Image            string             `bson:"image"`
	ImageUrl         string             `bson:"image_Url"`
	ImageNewsUrl     string             `bson:"image_newsUrl"`
	Date             string             `bson:"date"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

func (t *Trend) GenID() string {
	return fmt.Sprintf("%s_%s_%s", t.Title, t.TitleExploreLink, t.Date)
}
