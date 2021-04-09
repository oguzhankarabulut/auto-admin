package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	set          = "$set"
	keyTimestamp = "timestamp"
	keyId        = "_id"
)

type Client struct {
	mc *mongo.Client
	db string
}

func NewClient(conn string) *Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return &Client{
		mc: client,
	}
}

func (c *Client) SetDB(db string) {
	c.db = db
}

func (c *Client) FindOne(coll string, q bson.M, v interface{}) error {
	col := c.mc.Database(c.db).Collection(coll)
	if err := col.FindOne(context.Background(), q).Decode(v); err != nil {
		return err
	}
	return nil
}

func (c *Client) InsertOne(coll string, v interface{}) error {
	col := c.mc.Database(c.db).Collection(coll)
	f, err := bson.Marshal(v)
	if err != nil {
		return wrapError(errSaveSerialization, err)
	}
	if _, err = col.InsertOne(context.Background(), f); err != nil {
		return wrapError(errSaveInsertOne, err)
	}
	return nil
}

func (c *Client) Update(coll string, f bson.M, v interface{}) error {
	col := c.mc.Database(c.db).Collection(coll)
	u := bson.D{{set, v}}
	var updatedDocument bson.M
	err := col.FindOneAndUpdate(context.Background(), f, u).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return wrapError(errUpdateById, err)
		}
	}
	return nil
}

func (c *Client) All(coll string, f bson.M) ([]bson.M, error) {
	col := c.mc.Database(c.db).Collection(coll)
	opt := options.Find()
	opt.SetSort(bson.D{{keyTimestamp, -1}})
	cursor, err := col.Find(context.Background(), f, opt)
	if err != nil {
		return nil, wrapError(errFindAll, err)
	}
	var list []bson.M
	if err = cursor.All(context.Background(), &list); err != nil {
		return nil, wrapError(errIterateAll, err)
	}
	return list, nil
}

func (c *Client) Pagination(coll string, page int64, limit int64, q bson.M, s bson.D) ([]bson.M, error) {
	col := c.mc.Database(c.db).Collection(coll)
	opt := options.Find()
	skip := (page - 1) * limit
	if page == 1 {
		skip = 0
	}
	opt.SetSort(s)
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	cursor, err := col.Find(context.Background(), q, opt)
	if err != nil {
		return nil, wrapError(errFindAll, err)
	}
	var list []bson.M
	if err = cursor.All(context.Background(), &list); err != nil {
		return nil, wrapError(errIterateAll, err)
	}
	return list, nil
}

func (c *Client) Count(coll string) (int64, error) {
	col := c.mc.Database(c.db).Collection(coll)
	count, err := col.EstimatedDocumentCount(context.Background())
	if err != nil {
		return 0, wrapError(errCount, err)
	}
	return count, nil
}

func (c *Client) InsertBulk(coll string, doc []interface{}) error {
	col := c.mc.Database(c.db).Collection(coll)
	_, err := col.InsertMany(context.Background(), doc)
	if err != nil {
		return wrapError(errInsertMany, err)
	}
	return nil
}

func (c *Client) CollectionNames() ([]string, error) {
	cc, err := c.mc.Database(c.db).ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		return nil, wrapError(errCollectionNames, err)
	}

	return cc, nil
}

func (c *Client) Delete(coll string, q bson.M) error {
	col := c.mc.Database(c.db).Collection(coll)
	_, err := col.DeleteOne(context.Background(), q)
	if err != nil {
		return wrapError(errDeleteOne, err)
	}

	return nil
}
