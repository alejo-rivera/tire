package ent

import (
	"testing"
	"tire/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestToGraph(t *testing.T) {
	expectedGraphUsers := []*model.User{
		{
			Name:          "alejo",
			CurrentHealth: 10,
			MaxHealth:     10,
			ID:            "0",
		},
		{
			Name:          "ErrorOnLine1024",
			CurrentHealth: 50,
			MaxHealth:     50,
			ID:            "1",
		},
	}
	dbUsers := Users{
		{
			Name:          "alejo",
			CurrentHealth: 10,
			MaxHealth:     10,
			ID:            0,
		},
		{
			Name:          "ErrorOnLine1024",
			CurrentHealth: 50,
			MaxHealth:     50,
			ID:            1,
		},
	}
	graphUsers := dbUsers.ToGraph()
	assert.Equal(t, graphUsers, expectedGraphUsers)
}
