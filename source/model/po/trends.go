package po

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Trend struct {
	ID               primitive.ObjectID `bson:"_id"`
	TitleQuery       string             `bson:"title_query"`
	TitleExploreLink string             `bson:"title_exploreLink"`
	FormattedTraffic string             `bson:"formattedTraffic"`
	ImageNewsUrl     string             `bson:"image_newsUrl"`
	ImageSource      string             `bson:"image_source"`
	ImageUrl         string             `bson:"image_Url"`
	Date             time.Time          `bson:"date"`
	CreatedAt        string             `bson:"created_at"`
	UpdatedAt        string             `bson:"updated_at"`
}
