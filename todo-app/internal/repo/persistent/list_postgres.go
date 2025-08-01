package persistent

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/transport/todo_list"
	"github.com/London57/todo-app/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ListRepo struct {
	*postgres.Postgres
}

func NewListRepo(pg *postgres.Postgres) *ListRepo {
	return &ListRepo{pg}
}

func (r *ListRepo) Create(ctx context.Context, userID uuid.UUID, list todo_list.TodoListRequest) (int, error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return 0, fmt.Errorf("ListRepo - Create - tx.BeginTx: %W", err)
	}
	stmt, args, err := r.Builder.
		Insert("todo_lists").
		Columns("title", "description").
		Values(list.Title, list.Description).
		Suffix("returning id").
		ToSql()

	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("ListRepo - Create - r.Builder: %w", err)
	}

	var id int
	err = tx.QueryRow(ctx, stmt, args...).Scan(&id)
	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("ListRepo - Create - r.Pool.QueryRow: %w", err)
	}

	stmt, args, err = r.Builder.
		Insert("users_lists").
		Values(userID, id).
		ToSql()

	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("ListRepo - Create - r.Builder: %w", err)
	}

	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		tx.Rollback(ctx)
		return 0, fmt.Errorf("ListRepo - Create - r.Pool.Exec: %w", err)
	}

	return id, tx.Commit(ctx)
}

func (r *ListRepo) GetAll(ctx context.Context, userID uuid.UUID) ([]domain.TodoList, error) {
	stmt, args, err := r.Builder.
		Select("tl.*").
		From("todo_lists tl").
		InnerJoin("users_lists ul on tl.id = ul.list_id").
		Where(squirrel.Eq{"ul.user_id": userID}).
		ToSql()
	
	if err != nil {
		return nil, fmt.Errorf("ListRepo - GetAll - r.Builder: %w", err)
	}

	var todoLists []domain.TodoList
	err = r.Pool.QueryRow(ctx, stmt, args...).Scan(&todoLists)
	if err != nil {
		return nil, fmt.Errorf("ListRepo - GetAll - r.Pool.QueryRow: %w", err)
	}

	return todoLists, nil
}

func (r *ListRepo) GetById(ctx context.Context, userID, listID int) (domain.TodoList, error) {
	stmt, args, err := r.Builder.
		Select("tl.*").
		From("todo_lists tl").
		InnerJoin("users_lists ul on tl.id = ul.list_id").
		Where(squirrel.Eq{"ul.user_id": userID, "ul.list_id": listID}).
		ToSql()
	if err != nil {
		return domain.TodoList{}, err
	}

	var todolist domain.TodoList
	err = r.Pool.QueryRow(ctx, stmt, args...).Scan(&todolist)
	if err != nil {
		return domain.TodoList{}, fmt.Errorf("ListRepo - GetById - r.Pool.QueryRow: %w", err)
	}
	return todolist, nil
}

func (r *ListRepo) Delete(ctx context.Context, userID uuid.UUID, listID int) error {
	stmt, args, err := r.Builder.Delete("todo_lists tl").
		Suffix("Using users_lists ul").
		Where("tl.id = ul.list_id").
		Where(squirrel.Eq{"ul.list_id": listID, "ul.user_id": userID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("ListRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("ListRepo - Delete - r.Pool.Exec: %w", err)
	}
	return nil
}

func (r *ListRepo) Update(ctx context.Context, userID uuid.UUID, listID int, input todo_list.UpdateListRequest) error {
	updateBuilder := r.Builder.Update("todo_lists tl")

	if input.Title != nil {
		updateBuilder = updateBuilder.Set("title", input.Title)
	}
	if input.Description != nil {
		updateBuilder = updateBuilder.Set("description", input.Description)
	}

	stmt, args, err := updateBuilder.From("users_lists ul").
		Where(squirrel.Eq{"ul.user_id": userID, "ul.list_id": listID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("ListRepo - Update - r.Builder: %w", err)
	}
	
	_, err = r.Pool.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("ListRepo - Update - r.Pool.Exec: %w", err)
	}
	return nil
}