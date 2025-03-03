package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/michaljemala/pqerror"

	"github.com/singl3focus/stats-project/collector/internal/domain"
)

var (
	ErrDataAlreadyExists = errors.New("data already exists")
)

func (r *Database) AddCall(ctx context.Context, userID, serviceID int32) error {
    query := `INSERT INTO stats (user_id, service_id, count)
	VALUES ($1, $2, 1)
	ON CONFLICT (user_id, service_id) 
	DO UPDATE SET count = stats.count + 1`
    
    _, err := r.db.ExecContext(ctx, query, userID, serviceID)

    return err
}

// не использую query builder, т.к. делаю на скорую руку :) для ТЗ, думаю, хватит

func (r *Database) Calls(ctx context.Context, filter domain.CallsFilter) ([]domain.Call, error) {
	query := `SELECT user_id, service_id, count FROM stats`
	
	// Немного быстрого говнокода (т.к параметры валидируются в сервисе, то исключаем SQL-иньекции)
	if (filter.UserID != nil && *filter.UserID > 0) && (filter.ServiceID != nil && *filter.ServiceID > 0){
		queryPart := fmt.Sprintf(`WHERE user_id = %d AND service_id = %d`, *filter.UserID, *filter.ServiceID)
		query = query + " " + queryPart
	} else if (filter.UserID != nil && *filter.UserID > 0) {
		queryPart := fmt.Sprintf(`WHERE user_id = %d`, *filter.UserID)
		query = query + " " + queryPart
	} else if (filter.ServiceID != nil && *filter.ServiceID > 0) {
		queryPart := fmt.Sprintf(`WHERE service_id = %d`, *filter.ServiceID)
		query = query + " " + queryPart
	}
	
	if (filter.Page != nil && *filter.Page > 0) && (filter.PerPage != nil && *filter.PerPage > 0) {
		queryPart := fmt.Sprintf(`LIMIT %d OFFSET %d`, *filter.PerPage, *filter.PerPage * *filter.Page)
		query = query + " " + queryPart
	}
	
	queryPart := fmt.Sprintf(`ORDER BY count %s`, filter.Sort)
	query = query + " " + queryPart

    rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	calls := make([]domain.Call, 0, 1) 
	for rows.Next() {
		var call domain.Call
		if err := rows.StructScan(&call); err != nil {
			return nil, err
		}

		calls = append(calls, call)
	}

    return calls, err
}

func (r *Database) AddService(ctx context.Context, name, description string) (int32, error) {
    query := `INSERT INTO services (name, description) VALUES ($1, $2) RETURNING id`
	
	var id int32
    err := r.db.QueryRowContext(ctx, query, name, description).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == pqerror.UniqueViolation {
			return 0, ErrDataAlreadyExists
		}

		return 0, err
	}

    return id, err
}