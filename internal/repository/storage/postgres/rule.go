package postgres

import (
	"fmt"
	"strings"

	"github.com/evgeniy-dammer/emenu-api/internal/domain/rule"
	"github.com/evgeniy-dammer/emenu-api/pkg/context"
	"github.com/pkg/errors"
)

// RuleGetAll selects all rules from database.
func (r *Repository) RuleGetAll(ctx context.Context, userID string) ([]rule.Rule, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var rules []rule.Rule

	query := fmt.Sprintf("SELECT id, ptype, v0, v1, v2, v3, v4, v5 FROM %s ", ruleTable)
	err := r.database.Select(&rules, query)

	return rules, errors.Wrap(err, "rules select query error")
}

// RuleGetOne select rule by id from database.
func (r *Repository) RuleGetOne(ctx context.Context, userID string, ruleID string) (rule.Rule, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var rle rule.Rule

	query := fmt.Sprintf("SELECT id, ptype, v0, v1, v2, v3, v4, v5 FROM %s WHERE id = $1 ", ruleTable)
	err := r.database.Get(&rle, query, ruleID)

	return rle, errors.Wrap(err, "rule select query error")
}

// RuleCreate insert rule into database.
func (r *Repository) RuleCreate(ctx context.Context, userID string, input rule.CreateRuleInput) (string, error) {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var ruleID string

	query := fmt.Sprintf("INSERT INTO %s (ptype, v0, v1, v2, v3, v4, v5) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		ruleTable)

	row := r.database.QueryRow(query, input.Ptype, input.V0, input.V1, input.V2, input.V3, input.V4, input.V5)
	err := row.Scan(&ruleID)

	return ruleID, errors.Wrap(err, "rule create query error")
}

// RuleUpdate updates rule by id in database.
func (r *Repository) RuleUpdate(ctx context.Context, userID string, input rule.UpdateRuleInput) error {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	setValues := make([]string, 0, 8)
	args := make([]interface{}, 0, 8)
	argID := 1

	if input.Ptype != nil {
		setValues = append(setValues, fmt.Sprintf("ptype=$%d", argID))
		args = append(args, *input.Ptype)
		argID++
	}

	if input.V0 != nil {
		setValues = append(setValues, fmt.Sprintf("v0=$%d", argID))
		args = append(args, *input.V0)
		argID++
	}

	if input.V1 != nil {
		setValues = append(setValues, fmt.Sprintf("v1=$%d", argID))
		args = append(args, *input.V1)
		argID++
	}

	if input.V2 != nil {
		setValues = append(setValues, fmt.Sprintf("v2=$%d", argID))
		args = append(args, *input.V2)
		argID++
	}

	if input.V3 != nil {
		setValues = append(setValues, fmt.Sprintf("v3=$%d", argID))
		args = append(args, *input.V3)
		argID++
	}

	if input.V4 != nil {
		setValues = append(setValues, fmt.Sprintf("v4=$%d", argID))
		args = append(args, *input.V4)
		argID++
	}

	if input.V5 != nil {
		setValues = append(setValues, fmt.Sprintf("v5=$%d", argID))
		args = append(args, *input.V5)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s'", ruleTable, setQuery, *input.ID)
	_, err := r.database.Exec(query, args...)

	return errors.Wrap(err, "rule update query error")
}

// RuleDelete deletes rule by id from database.
func (r *Repository) RuleDelete(ctx context.Context, userID string, ruleID string) error {
	ctx = ctx.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ruleTable)
	_, err := r.database.Exec(query, ruleID)

	return errors.Wrap(err, "rule delete query error")
}
