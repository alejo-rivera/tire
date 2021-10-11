package ent

import (
	"fmt"
	"tire/graph/model"
)

func (u User) ToGraph() *model.User {
	return &model.User{
		ID:    fmt.Sprintf("%d", u.ID),
		Name:  u.Name,
		Email: u.Email,
		// If the graph model ever includes plants
		// if plants, err := u.Edges.PlantsOrErr();err!=nil(
		// 	result.Plants
		// )
	}
}

func (u Users) ToGraph() []*model.User {
	users := make([]*model.User, len(u))
	for i, user := range u {
		users[i] = user.ToGraph()
	}
	return users
}
