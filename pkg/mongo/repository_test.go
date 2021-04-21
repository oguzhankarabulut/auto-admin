package mongo

import (
	"testing"
)

func BaseRepositoryTest() *repository {
	mc := NewClient("mongodb://root:example@localhost:27017")
	mc.SetDB("auto-admin")
	return NewRepository(mc)
}

func TestAllWithValidDocument(t *testing.T) {
	r := BaseRepositoryTest()
	d, _ := r.CollectionNames()
	_, err := r.All(d[0])
	if err == nil {
		t.Logf("success")
	}
}

func TestAllWithInvalidDocument(t *testing.T) {
	r := BaseRepositoryTest()
	_, err := r.All("test")
	if err != nil {
		t.Logf("success")
	}
}

func TestAllWithNilDocument(t *testing.T) {
	r := BaseRepositoryTest()
	_, err := r.All("")
	if err != nil {
		t.Logf("success")
	}
}
