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
	} else {
		t.Errorf("All with valid document failed")
	}
}

func TestAllWithInvalidDocument(t *testing.T) {
	r := BaseRepositoryTest()
	cc, err := r.All("test")
	if cc == nil {
		t.Logf("success")
	} else {
		t.Errorf("all with invalid document failed %v, %v", cc, err)
	}
}

func TestAllWithNilDocument(t *testing.T) {
	r := BaseRepositoryTest()
	_, err := r.All("")
	if err != nil {
		t.Logf("success")
	} else {
		t.Errorf("all with nil document failed")
	}
}

func TestCreateWithValidValue(t *testing.T) {
	r := BaseRepositoryTest()
	m := map[string]interface{}{
		"day":         5,
		"bool_value":  true,
		"title":       "test title",
		"description": "test description",
	}
	_, err := r.Create("post", m)
	if err == nil {
		t.Logf("success")
	} else {
		t.Errorf("create valid value failed")
	}
}

func TestCreateWithNilValue(t *testing.T) {
	r := BaseRepositoryTest()
	m := make(map[string]interface{})
	_, err := r.Create("post", m)
	if err != nil {
		t.Logf("success")
	} else {
		t.Errorf("create wit nil value failed")
	}
}
