package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/evgeniy-dammer/emenu-api/internal/domain/table"
	"github.com/evgeniy-dammer/emenu-api/pkg/context"
	"github.com/pkg/errors"
)

// TableGetAll selects all tables from database.
func (r *Repository) TableGetAll(ctx context.Context, userID string, organizationID string) ([]table.Table, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var tables []table.Table

	query := fmt.Sprintf("SELECT id, name, organization_id FROM %s WHERE is_deleted = false AND organization_id = $1 ",
		tableTable)
	err := r.database.Select(&tables, query, organizationID)

	return tables, errors.Wrap(err, "tables select query error")
}

// TableGetOne select table by id from database.
func (r *Repository) TableGetOne(ctx context.Context, userID string, organizationID string, tableID string) (table.Table, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var tble table.Table

	query := fmt.Sprintf(
		"SELECT id, name, organization_id FROM %s WHERE is_deleted = false AND organization_id = $1 AND id = $2 ",
		tableTable,
	)
	err := r.database.Get(&tble, query, organizationID, tableID)

	return tble, errors.Wrap(err, "table select query error")
}

// TableCreate insert table into database.
func (r *Repository) TableCreate(ctx context.Context, userID string, input table.CreateTableInput) (string, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var tableID string

	query := fmt.Sprintf("INSERT INTO %s (name, organization_id, user_created) VALUES ($1, $2, $3) RETURNING id",
		tableTable)
	row := r.database.QueryRow(query, input.Name, input.OrganizationID, userID)
	err := row.Scan(&tableID)

	return tableID, errors.Wrap(err, "table create query error")
}

// TableUpdate updates table by id in database.
func (r *Repository) TableUpdate(ctx context.Context, userID string, input table.UpdateTableInput) error {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

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
	_, err := r.database.Exec(query, args...)

	return errors.Wrap(err, "table update query error")
}

// TableDelete deletes table by id from database.
func (r *Repository) TableDelete(ctx context.Context, userID string, organizationID string, tableID string) error {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	query := fmt.Sprintf(
		"UPDATE %s SET is_deleted = true, deleted_at = $1, user_deleted = $2 "+
			"WHERE is_deleted = false AND id = $3 AND organization_id = $4",
		tableTable,
	)

	_, err := r.database.Exec(query, time.Now().Format("2006-01-02 15:04:05"), userID, tableID, organizationID)

	return errors.Wrap(err, "table delete query error")
}
