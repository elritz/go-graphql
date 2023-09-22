package graphql

import (
	"context"

	"github.com/elritz/meetmeup/models"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.Domain.MeetupsRepo.GetMeetupsForUser(obj)
}
