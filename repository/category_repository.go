package repository

import (
	"Go-RESTfull-API/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cateogry domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, cateogry domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, cateogry domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
