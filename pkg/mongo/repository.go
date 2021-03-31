package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	mc *Client
}

func NewRepository(mc *Client) *repository {
	return &repository{mc: mc}
}

type Repository interface {
	GetAll(coll string) ([]bson.M, error)
}

func (r *repository) GetAll(coll string) ([]bson.M, error) {
	dd, err := r.mc.All(coll, bson.M{})
	fmt.Println(coll)
	return dd, err
}