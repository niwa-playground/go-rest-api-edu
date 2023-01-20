package repositories

import (
	"context"
	"database/sql"

	"github.com/imniwa/go-rest-api-edu/model/entity"
)

type BaseCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
