package service

import (
	"auto-admin/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type ApiService interface {
	All(coll string) ([]bson.M, error)
	Single(coll string, id string) (interface{}, error)
	Create(coll string, m map[string]interface{}) (interface{}, error)
	Update(coll string, m map[string]interface{}) (interface{}, error)
	Delete(coll string, id string) error
}

type apiService struct {
	repository mongo.Repository
}

func NewApiService(r mongo.Repository) *apiService {
	return &apiService{repository: r}
}

func (s apiService) All(coll string) ([]bson.M, error) {
	cc, err := s.repository.All(coll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return cc, nil
}

func (s apiService) Single(coll string, id string) (interface{}, error) {
	c, err := s.repository.Single(coll, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return c, nil
}

func (s apiService) Create(coll string, m map[string]interface{}) (interface{}, error) {
	i, err := s.repository.Create(coll, m)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return i, nil
}

func (s apiService) Update(coll string, m map[string]interface{}) (interface{}, error) {
	i, err := s.repository.Update(coll, m)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return i, nil
}

func (s apiService) Delete(coll string, id string) error {
	err := s.repository.Delete(coll, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
