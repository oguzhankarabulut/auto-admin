package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository struct {
	mc *Client
}

func NewRepository(mc *Client) *repository {
	return &repository{mc: mc}
}

type Repository interface {
	All(coll string) ([]bson.M, error)
	Create(coll string, i map[string]interface{}) (interface{}, error)
	Update(coll string, i map[string]interface{}) (interface{}, error)
	Delete(coll string, id string) error
	Single(coll string, id string) (interface{}, error)
	CollectionNames() ([]string, error)
	Count(coll string) (int64, error)
}

func (r *repository) All(coll string) ([]bson.M, error) {
	rr, err := r.mc.All(coll, bson.M{})
	if err != nil {
		return nil, wrapError(errAll+coll, err)
	}
	return rr, err
}

func (r *repository) Single(coll string, id string) (interface{}, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	q := bson.M{keyId: objId}
	m := make(map[string]interface{})
	if err := r.mc.FindOne(coll, q, m); err != nil {
		return nil, wrapError(errSingle+coll, err)
	}
	return m, nil
}

func (r *repository) Create(coll string, i map[string]interface{}) (interface{}, error) {
	if len(i) == 0 {
		return nil, wrapError(errCreateObjectNil, nil)
	}
	if err := r.mc.InsertOne(coll, i); err != nil {
		return nil, wrapError(errCreate+coll, err)
	}
	return i, nil
}

func (r *repository) Update(coll string, i map[string]interface{}) (interface{}, error) {
	objId, _ := primitive.ObjectIDFromHex(fmt.Sprintf("%v", i[keyId]))
	q := bson.M{keyId: objId}
	delete(i, keyId)
	if err := r.mc.Update(coll, q, i); err != nil {
		return nil, wrapError(errUpdate+coll, err)
	}
	return i, nil
}

func (r *repository) Delete(coll string, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	q := bson.M{keyId: objId}
	if err := r.mc.Delete(coll, q); err != nil {
		return wrapError(errDelete+coll, err)
	}

	return nil
}

func (r *repository) CollectionNames() ([]string, error) {
	cc, err := r.mc.CollectionNames()
	if err != nil {
		return nil, wrapError(errCollectionNames, err)
	}

	return cc, nil
}

func (r *repository) Count(coll string) (int64, error) {
	c, err := r.mc.Count(coll)
	if err != nil {
		return 0, wrapError(errCollectionCount, err)
	}

	return c, nil
}
