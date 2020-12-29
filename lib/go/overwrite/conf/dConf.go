// Code generated by archivist. DO NOT EDIT.

package conf

import (
	"time"

	"github.com/kingsgroupos/archivist/lib/go/archivist"
	"github.com/kingsgroupos/misc/wtime"
	"github.com/pkg/errors"
)

var (
	_ = time.After
	_ = errors.New
	_ = archivist.NewArchivist
	_ = wtime.ParseDuration
)

// easyjson:json
type DConf struct {
	D1 int64 `json:"D1" bson:"D1"`
	D2 int64 `json:"D2" bson:"D2"`
	D3 int64 `json:"D3" bson:"D3"`
}

func (this *DConf) bindRefs(c *Collection) error {
	if this == nil {
		return nil
	}

	var ok bool
	_ = ok

	return nil
}