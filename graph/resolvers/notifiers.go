package resolvers

import (
	"api/ent"
	"api/graph/models"
)

func EventIn(event models.Event, events []models.Event) bool {
	for i := range events {
		if events[i] == event {
			return true
		}
	}
	return false
}

func NotifyUserListenners(r *mutationResolver, event models.Event, entity *ent.User) {
	r.UserListennersMutex.Lock()
	userEvent := &models.UserEvent{
		Event: event,
	}
	if event == models.EventDelete {
		userEvent.User = entity
	} else {
		userEvent.User = entity.Unwrap()
	}
	for key := range r.UserListenners {
		notify := (r.UserListenners[key].ID == nil || *r.UserListenners[key].ID == entity.ID) &&
			(r.UserListenners[key].Events == nil || EventIn(event, r.UserListenners[key].Events))
		if notify {
			r.UserListenners[key].Channel <- userEvent
		}
	}
	r.UserListennersMutex.Unlock()
}

func NotifyTodoListenners(r *mutationResolver, event models.Event, entity *ent.Todo) {
	r.TodoListennersMutex.Lock()

	todoEvent := &models.TodoEvent{
		Event: event,
	}

	if event == models.EventDelete {
		todoEvent.Todo = entity
	} else {
		todoEvent.Todo = entity.Unwrap()
	}

	for key := range r.TodoListenners {
		notify := (r.TodoListenners[key].ID == nil || *r.TodoListenners[key].ID == entity.ID) &&
			(r.TodoListenners[key].Events == nil || EventIn(event, r.TodoListenners[key].Events))
		if notify {
			r.TodoListenners[key].Channel <- todoEvent
		}
	}

	r.TodoListennersMutex.Unlock()
}
