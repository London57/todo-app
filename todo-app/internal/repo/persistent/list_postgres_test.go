package persistent

import (
	"context"
	"testing"

	"github.com/London57/todo-app/internal/transport/todo_list"
	"github.com/London57/todo-app/pkg/postgres"
	mock_postgres "github.com/London57/todo-app/pkg/postgres/mocks"
	"github.com/Masterminds/squirrel"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestListRepo_create(t *testing.T) {
	testcase := []struct {
		testname string
		setup func(*mock_postgres.MockPool, *mock_postgres.MockPgxTx, *mock_postgres.MockPgxRow)
		withError bool
	} {
		{
			"OK",
			func(pool *mock_postgres.MockPool, tx *mock_postgres.MockPgxTx, row *mock_postgres.MockPgxRow) {
				// var tx_return pgx.Tx = postgres.NewPgxTxAdapter(tx)
				tx_adapter := postgres.NewPgxTxAdapter(tx)
				pool.EXPECT().BeginTx(
					gomock.Any(),
 					pgx.TxOptions{IsoLevel: pgx.Serializable},
				).Return(tx_adapter, nil)
				tx_adapter.GetTx().EXPECT().QueryRow(
					gomock.Any(),
					"INSERT INTO todo_lists (title,description) VALUES ($1,$2) returning id",
					"title", "description",
				).Return(row)
				row.EXPECT().Scan(
					gomock.Any(),	
				).SetArg(0, 1).Return(nil)
				tx_adapter.GetTx().EXPECT().Exec(
					gomock.Any(),
					"INSERT INTO users_lists VALUES ($1,$2)",
					"f489ac09-8039-4a59-b6d6-46c6dbee4a1a", 1,
				).Return(pgconn.CommandTag{}, nil)
				tx_adapter.GetTx().EXPECT().Commit(
					gomock.Any(),
				).Return(nil)
			},
			false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.testname, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pool := mock_postgres.NewMockPool(c)
			tx := mock_postgres.NewMockPgxTx(c)
			row := mock_postgres.NewMockPgxRow(c)

			p := &postgres.Postgres{
				Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
				Pool: pool,
			}

			r := NewListRepo(p)

			tc.setup(pool, tx, row)
			ctx := context.Background()
			user_id := uuid.MustParse("f489ac09-8039-4a59-b6d6-46c6dbee4a1a")
			listID, err := r.Create(ctx, user_id, todo_list.TodoListRequest{
				Title: "title",
				Description: "description",
			})

			if tc.withError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, 1, listID)
		})
	}

}