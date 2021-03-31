package mongo

import (
	"fmt"
)

const (
	errSaveSerialization = "mongo: client insert one serialization"
	errSaveInsertOne     = "mongo: client insert one"
	errFindAll           = "mongo: client find all"
	errUpdateById        = "mongo: client update by id can not find document"
	errIterateAll        = "mongo: client iterate all"
	errCount             = "mongo: client count"
	errInsertMany        = "mongo: client insert many"
	errCollectionNames   = "mongo: client collection names"
)

func wrapError(localErr string, err error) error {
	return fmt.Errorf("%v:%v", localErr, err)
}
