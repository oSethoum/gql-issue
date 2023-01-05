package resolvers

import (
	"api/db"
	"api/ent"
	"api/graph/generated"
	"api/graph/models"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	Client              *ent.Client
	UserListenners      map[*chan *models.UserEvent]UserListenner
	UserListennersMutex sync.Mutex
	TodoListenners      map[*chan *models.TodoEvent]TodoListenner
	TodoListennersMutex sync.Mutex
}

var schema *graphql.ExecutableSchema

func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Client:              db.Client,
			UserListenners:      make(map[*chan *models.UserEvent]UserListenner),
			UserListennersMutex: sync.Mutex{},
			TodoListenners:      make(map[*chan *models.TodoEvent]TodoListenner),
			TodoListennersMutex: sync.Mutex{},
		}})
	}

	return *schema
}
