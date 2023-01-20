package repository

import (
	"context"
	"database/sql"

	"github.com/imniwa/go-rest-api-edu/helper"
	"github.com/imniwa/go-rest-api-edu/model/entity"
)

type CategoryRepository struct {
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "INSERT INTO category(name) values (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	panic("err")
}
func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	panic("err")
}
func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) entity.Category {
	panic("err")
}
func (repository *CategoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	panic("err")
}
