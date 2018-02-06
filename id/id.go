package id

import (
	"github.com/teris-io/shortid"
)

func NextId() string {
	id, err := shortid.Generate()
	if err != nil {
		panic("id create error")
	}
	return id
}
