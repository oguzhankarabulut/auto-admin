package service

import (
	"auto-admin/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const (
	collections = "Collections"
)

type WebService interface {
	Dashboard() (*dashboardResponse, error)
	Table(coll string) (*tableResponse, error)
	Detail(coll string, id string) (*detailResponse, error)
}

type webService struct {
	repository      mongo.Repository
	collectionNames []string
}

func NewWebService(r mongo.Repository, cn []string) *webService {
	return &webService{repository: r, collectionNames: cn}
}

type dashboardResponse struct {
	Collections []string
	Metric      map[string]int64
	Title       string
}

type tableResponse struct {
	CollectionName string
	Collections    []string
	Documents      []bson.M
}

type detailResponse struct {
	CollectionName string
	Collections    []string
	Detail         interface{}
}

func (s webService) Dashboard() (*dashboardResponse, error) {
	cn := s.collectionNames

	m := make(map[string]int64)
	if len(cn) != 0 {
		for i := 0; i < len(cn); i++ {
			c, _ := s.repository.Count(cn[i])
			m[cn[i]] = c
		}
	}

	return &dashboardResponse{Collections: cn, Metric: m, Title: collections}, nil
}

func (s webService) Table(coll string) (*tableResponse, error) {
	cc, err := s.repository.All(coll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &tableResponse{CollectionName: coll, Collections: s.collectionNames, Documents: cc}, nil
}

func (s webService) Detail(coll string, id string) (*detailResponse, error) {
	d, err := s.repository.Single(coll, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &detailResponse{CollectionName: coll, Collections: s.collectionNames, Detail: d}, nil
}
