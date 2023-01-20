package repository

import (
	"context"
	"database/sql"

	"github.com/imniwa/go-rest-api-edu/model/entity"
)

type BaseCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) entity.Category
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
