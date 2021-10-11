package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"tire/ent"
	"tire/ent/user"
	"tire/graph/model"

	"go.uber.org/zap"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (*model.User, error) {
	dbUser, err := r.dbClient.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		r.Error("create_user_in_database", zap.String("user_name", user.Name), zap.Error(err))
		return nil, err
	}
	return dbUser.ToGraph(), nil
}

func (r *mutationResolver) CreatePlant(ctx context.Context, plant model.NewPlant) (*model.Plant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.dbClient.User.Query().All(ctx)
	if err != nil {
		r.Error("serching_all_users_in_database", zap.Error(err))
		return nil, err
	}
	return ent.Users(users).ToGraph(), nil
}

func (r *queryResolver) UserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := r.dbClient.User.Query().
		Where(
			user.Name(name),
		).
		Only(ctx)
	if err != nil {
		r.Error("serching_user_in_database", zap.String("user_name", name), zap.Error(err))
		return nil, err
	}
	return user.ToGraph(), nil
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPlants(ctx context.Context) ([]*model.Plant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PlantByID(ctx context.Context, id string) (*model.Plant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
