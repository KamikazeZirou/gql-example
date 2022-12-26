package graph

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/kamikazezirou/gql-example/graph/model"
	"net/http"
)

type Loaders struct {
	UserById *dataloader.Loader[string, *model.User]
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: dataloader.NewBatchedLoader(func(ctx context.Context, userIds []string) []*dataloader.Result[*model.User] {
				fmt.Println("batch get users:", userIds)

				// ユーザIDのリストからユーザ情報を取得する
				// サンプル実装なので適当な値を返していますが、プロダクト実装では以下のようにしてください。
				//   - "SELECT * FROM users WHERE id IN (id1, id2, id3, ...)"のようなSQLでDBからユーザ情報を一括取得する
				//   - 他のサービスのBatch Read APIを呼ぶ
				// それでN+1問題を回避することができます。
				results := make([]*dataloader.Result[*model.User], len(userIds))
				for i, id := range userIds {
					results[i] = &dataloader.Result[*model.User]{
						Data:  &model.User{ID: id, Name: "user " + id},
						Error: nil,
					}
				}

				return results
			}),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

type contextKey int

var loadersKey contextKey

func ctxLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
