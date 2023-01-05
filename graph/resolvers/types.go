package resolvers

import (
	"api/graph/models"
)

type UserListenner struct {
	ID      *int
	Events  []models.Event
	Channel chan *models.UserEvent
}
type TodoListenner struct {
	ID      *int
	Events  []models.Event
	Channel chan *models.TodoEvent
}
