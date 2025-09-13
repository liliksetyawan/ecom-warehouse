package repository

import "database/sql"

type WarehouseModel struct {
	ID        sql.NullInt64
	ShopID    sql.NullInt64
	ProductID sql.NullInt64
	Code      sql.NullString
	Name      sql.NullString
	Stock     sql.NullInt64
	Location  sql.NullString
	Status    sql.NullString
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}
