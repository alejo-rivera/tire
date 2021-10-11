package ent

import (
	"testing"
	"tire/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestToGraph(t *testing.T) {
	expectedGraphUsers := []*model.User{
		{
			Name:  "alejo",
			Email: "alejorivera@protonmail.com",
			ID:    "0",
		},
		{
			Name:  "ErrorOnLine1024",
			Email: "test@neomontecito.com",
			ID:    "1",
		},
	}
	dbUsers := Users{
		{
			Name:  "alejo",
			Email: "alejorivera@protonmail.com",
			ID:    0,
		},
		{
			Name:  "ErrorOnLine1024",
			Email: "test@neomontecito.com",
			ID:    1,
		},
	}
	graphUsers := dbUsers.ToGraph()
	assert.Equal(t, graphUsers, expectedGraphUsers)
}
