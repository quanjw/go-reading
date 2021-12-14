package initDB

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestInit(t *testing.T) {

	var err error
	t.Log(Db)
	err = Db.Ping()
	t.Log(err)
	assert.Equal(t, nil, err)

}
