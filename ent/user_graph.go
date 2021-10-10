package ent

import (
	"fmt"
	"tire/graph/model"
)

func (u User) ToGraph() *model.User {
	return &model.User{
		ID:           fmt.Sprintf("%d", u.ID),
		Name:         u.Name,
		MaxHealth:    u.MaxHealth,
		CurentHealth: u.CurrentHealth,
	}
	// if the graph model ever includes items
	// if items, err := u.Edges.ItemsOrErr();err!=nil(
	// 	result.Items
	// )
}

func (u Users) ToGraph() []*model.User {
	users := make([]*model.User, len(u))
	for i, user := range u {
		users[i] = user.ToGraph()
	}
	return users
}
