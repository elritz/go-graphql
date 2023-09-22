//go:generate go run github.com/99designs/gqlgen -v

package graphql

import "github.com/elritz/meetmeup/domain"

type Resolver struct {
	Domain *domain.Domain
}
