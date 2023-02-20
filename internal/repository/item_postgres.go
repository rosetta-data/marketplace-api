package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/evgeniy-dammer/emenu-api/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// ItemPostgresql repository.
type ItemPostgresql struct {
	db *sqlx.DB
}

// NewItemPostgresql is a constructor for ItemPostgresql.
func NewItemPostgresql(db *sqlx.DB) *ItemPostgresql {
	return &ItemPostgresql{db: db}
}

// GetAll selects all items from database.
func (r *ItemPostgresql) GetAll(userID string, organizationID string) ([]model.Item, error) {
	var items []model.Item

	query := fmt.Sprintf(
		"SELECT id, name, price, category_id, organization_id FROM %s "+
			"WHERE is_deleted = false AND organization_id = $1 ",
		itemTable,
	)

	err := r.db.Select(&items, query, organizationID)

	return items, errors.Wrap(err, "items select query error")
}

// GetOne select item by id from database.
func (r *ItemPostgresql) GetOne(userID string, organizationID string, itemID string) (model.Item, error) {
	var item model.Item

	query := fmt.Sprintf(
		"SELECT id, name, price, category_id, organization_id FROM %s "+
			"WHERE is_deleted = false AND organization_id = $1 AND id = $2 ",
		itemTable,
	)
	err := r.db.Get(&item, query, organizationID, itemID)

	return item, errors.Wrap(err, "item select query error")
}

// Create insert item into database.
func (r *ItemPostgresql) Create(userID string, item model.Item) (string, error) {
	var itemID string

	query := fmt.Sprintf(
		"INSERT INTO %s (name, price, category_id, organization_id, user_created) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		itemTable,
	)

	row := r.db.QueryRow(query, item.Name, item.Price, item.CategoryID, item.OrganizationID, userID)
	err := row.Scan(&itemID)

	return itemID, errors.Wrap(err, "item create query error")
}

// Update updates item by id in database.
func (r *ItemPostgresql) Update(userID string, input model.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	if input.CategoryID != nil {
		setValues = append(setValues, fmt.Sprintf("category_id=$%d", argID))
		args = append(args, *input.CategoryID)
		argID++
	}

	if input.OrganizationID != nil {
		setValues = append(setValues, fmt.Sprintf("organization_id=$%d", argID))
		args = append(args, *input.OrganizationID)
		argID++
	}

	setValues = append(setValues, fmt.Sprintf("user_updated=$%d", argID))
	args = append(args, userID)
	argID++

	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argID))
	args = append(args, time.Now().Format("2006-01-02 15:04:05"))

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE is_deleted = false AND organization_id = '%s' AND id = '%s'",
		itemTable, setQuery, *input.OrganizationID, *input.ID)

	_, err := r.db.Exec(query, args...)

	return errors.Wrap(err, "item update query error")
}

// Delete deletes item by id from database.
func (r *ItemPostgresql) Delete(userID string, organizationID string, itemID string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET is_deleted = true, deleted_at = $1, user_deleted = $2 "+
			"WHERE is_deleted = false AND id = $3 AND organization_id = $4",
		itemTable,
	)

	_, err := r.db.Exec(query, time.Now().Format("2006-01-02 15:04:05"), userID, itemID, organizationID)

	return errors.Wrap(err, "item delete query error")
}
