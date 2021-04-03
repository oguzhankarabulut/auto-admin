package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	mc *Client
}

func NewRepository(mc *Client) *repository {
	return &repository{mc: mc}
}

type Repository interface {
	All(coll string) ([]bson.M, error)
	Create(coll string, i interface{}) (interface{}, error)
}

func (r *repository) All(coll string) ([]bson.M, error) {
	dd, err := r.mc.All(coll, bson.M{})
	if err != nil {
		return nil, wrapError(errAll+coll, err)
	}
	return dd, err
}

func (r *repository) Create(coll string, i interface{}) (interface{}, error) {
	if err := r.mc.InsertOne(coll, i); err != nil {
		return nil, wrapError(errCreate+coll, err)
	}
	return i, nil
}
