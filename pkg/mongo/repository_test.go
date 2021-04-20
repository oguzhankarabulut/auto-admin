package mongo

import (
	"testing"
)

func TestAllWithValidDocument(t *testing.T) {
	mc := NewClient("mongodb://root:example@localhost:27017")
	mc.SetDB("auto-admin")
	r := NewRepository(mc)
	d, _ := r.CollectionNames()
	_, err := r.All(d[0])
	if err == nil {
		t.Logf("success")
	}
}

func TestAllWithInvalidDocument(t *testing.T) {
	mc := NewClient("mongodb://root:example@localhost:27017")
	mc.SetDB("auto-admin")
	r := NewRepository(mc)
	_, err := r.All("test")
	if err != nil {
		t.Logf("success")
	}
}
