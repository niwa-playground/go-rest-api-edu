package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/imniwa/go-rest-api-edu/helper"
	"github.com/imniwa/go-rest-api-edu/model/entity"
)

type CategoryRepository struct {
}

func NewCategoryRepository() BaseCategoryRepository {
	return &CategoryRepository{}
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
	SQL := "UPDATE category SET name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}
func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := entity.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}
func (repository *CategoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
