# gql-issue

## Reproduction steps:

- run commands

```console
  go generate ./ent/generate
```

```console
  go get github.com/99designs/gqlgen
```

```console
  go run github.com/99designs/gqlgen
```

- Create User

```gql
mutation CreateUser {
  createUser(input: { name: "Coder" }) {
    id
    name
  }
}
```

- Subscribe to todo

```gql
subscription WatchTodo {
  todo {
    event
    todo {
      id
      text
      done
      user {
        id
        name
      }
    }
  }
}
```

- Create Todo multiple times

```gql
mutation CreateTodo($userId: ID!) {
  createTodo(input: { text: "Do something", userID: $userId }) {
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

## Issues:

- when the edge Todo -> User is Required `CreateTodo` and `WatchTodo` return an error

```json
{
  "errors": [
    {
      "message": "ent: user not found",
      "path": ["createTodo", "user"]
    }
  ],
  "data": null
}
```

- when the edge Todo -> User is not Required `CreateTodo` returns null for user field.

```json
{
  "data": {
    "createTodo": {
      "id": "1",
      "text": "Do something",
      "done": false,
      "user": null
    }
  }
}
```
