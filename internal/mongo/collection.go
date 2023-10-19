package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
}

type Collection struct {
	*mongo.Collection
}
