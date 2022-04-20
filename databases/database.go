package databases

import (
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jmoiron/sqlx"
)

type Condition struct {
	Column     string
	Expression string
	Value      string
}

type Join struct {
	TableName  string
	ForeignKey string
	LocalKey   string
}

type Manager struct {
	DB *sqlx.DB
}

func (db *Manager) IsExists(table string, conditions []Condition, joins []Join) bool {
	return db.GetCount(table, conditions, joins) > 0
}

func (db *Manager) GetCount(table string, conditions []Condition, joins []Join) int {
	var result int

	query := sq.Select("COUNT(*)").From(table)

	if len(joins) > 0 {
		for _, join := range joins {
			query = query.Join(join.TableName + " ON " + join.LocalKey + " = " + join.ForeignKey).
				Where(join.TableName + ".deleted_at IS NULL")
		}
	}

	if len(conditions) > 0 {
		for _, condition := range conditions {
			query = query.Where(condition.Column+" "+condition.Expression+" ?", condition.Value)
		}
	}

	query = query.Where(table + ".deleted_at IS NULL")
	sql, args, err := query.ToSql()

	if err != nil {
		return 0
	}

	err = db.DB.Get(&result, sql, args...)

	if err != nil {
		return 0
	}

	return result
}
