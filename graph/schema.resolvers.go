package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/graph-gophers/dataloader"
	"github.com/kamikazezirou/gql-example/graph/generated"
	"github.com/kamikazezirou/gql-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("todo:%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.User, error) {
	// サンプル実装なので固定値を返しているだけですが、
	// プロダクトコードでは、ユーザを認証して、そのユーザの情報を返すようにしてください。
	return &model.User{
		ID:   "user:1",
		Name: "user1",
	}, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	s := strings.Split(id, ":")
	t := s[0]

	switch t {
	case "todo":
		for _, todo := range r.todos {
			if todo.ID == id {
				return todo, nil
			}
		}
		return nil, errors.New("not found")
	default:
		return nil, fmt.Errorf("unknwon type:%s", t)
	}
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	thunk := ctxLoaders(ctx).UserById.Load(ctx, dataloader.StringKey(obj.UserID))
	item, err := thunk()
	if err != nil {
		return nil, err
	} else {
		return item.(*model.User), nil
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
