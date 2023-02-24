package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/evgeniy-dammer/emenu-api/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// TablePostgresql repository.
type TablePostgresql struct {
	db *sqlx.DB
}

// NewTablePostgresql is a constructor for TablePostgresql.
func NewTablePostgresql(db *sqlx.DB) *TablePostgresql {
	return &TablePostgresql{db: db}
}

// GetAll selects all tables from database.
func (r *TablePostgresql) GetAll(userID string, organizationID string) ([]model.Table, error) {
	var tables []model.Table

	query := fmt.Sprintf("SELECT id, name, organization_id FROM %s WHERE is_deleted = false AND organization_id = $1 ",
		tableTable)
	err := r.db.Select(&tables, query, organizationID)

	return tables, errors.Wrap(err, "tables select query error")
}

// GetOne select table by id from database.
func (r *TablePostgresql) GetOne(userID string, organizationID string, tableID string) (model.Table, error) {
	var table model.Table

	query := fmt.Sprintf(
		"SELECT id, name, organization_id FROM %s WHERE is_deleted = false AND organization_id = $1 AND id = $2 ",
		tableTable,
	)
	err := r.db.Get(&table, query, organizationID, tableID)

	return table, errors.Wrap(err, "table select query error")
}

// Create insert table into database.
func (r *TablePostgresql) Create(userID string, table model.Table) (string, error) {
	var tableID string

	query := fmt.Sprintf("INSERT INTO %s (name, organization_id, user_created) VALUES ($1, $2, $3) RETURNING id",
		tableTable)
	row := r.db.QueryRow(query, table.Name, table.OrganizationID, userID)
	err := row.Scan(&tableID)

	return tableID, errors.Wrap(err, "table create query error")
}

// Update updates table by id in database.
func (r *TablePostgresql) Update(userID string, input model.UpdateTableInput) error {
	setValues := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
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
		tableTable, setQuery, *input.OrganizationID, *input.ID)
	_, err := r.db.Exec(query, args...)

	return errors.Wrap(err, "table update query error")
}

// Delete deletes table by id from database.
func (r *TablePostgresql) Delete(userID string, organizationID string, tableID string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET is_deleted = true, deleted_at = $1, user_deleted = $2 "+
			"WHERE is_deleted = false AND id = $3 AND organization_id = $4",
		tableTable,
	)

	_, err := r.db.Exec(query, time.Now().Format("2006-01-02 15:04:05"), userID, tableID, organizationID)

	return errors.Wrap(err, "table delete query error")
}
