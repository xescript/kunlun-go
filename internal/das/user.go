package das

import (
	"context"
	"fmt"

	"kunlun/ent"
	"kunlun/ent/namespace"
	"kunlun/ent/user"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client}
}

// FindByNamespace return nil if user not found
func (s *UserService) FindByNamespace(path string) *ent.User {
	u, err := s.client.User.Query().
		Where(user.HasNamespaceWith(namespace.PathEqualFold(path))).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil
	}
	if err != nil {
		panic(fmt.Errorf("failed to query user: %v", err))
	}
	return u
}
