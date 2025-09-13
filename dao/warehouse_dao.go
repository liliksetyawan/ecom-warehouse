package dao

import (
	"database/sql"
	"ecom-warehouse/dto/in"
	"ecom-warehouse/repository"
	"fmt"
	"log"
	"strings"
)

func CreateWarehouse(db *sql.DB, warehouse *repository.WarehouseModel) (int64, error) {
	query := `
		INSERT INTO warehouses (shop_id, product_id, code, name, stock, location, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id int64
	err := db.QueryRow(
		query,
		warehouse.ShopID,
		warehouse.ProductID,
		warehouse.Code,
		warehouse.Name,
		warehouse.Stock,
		warehouse.Location,
		warehouse.Status,
	).Scan(&id)

	if err != nil {
		log.Println("Error CreateWarehouse:", err)
		return 0, err
	}

	return id, nil
}

func GetWarehouseByID(db *sql.DB, id int64) (*repository.WarehouseModel, error) {
	query := `
		SELECT id, shop_id, product_id, code, name, stock, location, status, created_at, updated_at
		FROM warehouses
		WHERE id = $1
	`

	row := db.QueryRow(query, id)

	var warehouse repository.WarehouseModel
	err := row.Scan(
		&warehouse.ID,
		&warehouse.ShopID,
		&warehouse.ProductID,
		&warehouse.Code,
		&warehouse.Name,
		&warehouse.Stock,
		&warehouse.Location,
		&warehouse.Status,
		&warehouse.CreatedAt,
		&warehouse.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("warehouse not found")
		}
		return nil, err
	}

	return &warehouse, nil
}

func GetListWarehouses(db *sql.DB, pagination in.Pagination) ([]repository.WarehouseModel, error) {
	baseQuery := `
		SELECT id, shop_id, product_id, code, name, stock, location, status, created_at, updated_at
		FROM warehouses
	`
	var conditions []string
	var args []interface{}
	argID := 1

	if pagination.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR code ILIKE $%d)", argID, argID+1))
		args = append(args, "%"+pagination.Search+"%", "%"+pagination.Search+"%")
		argID += 2
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	baseQuery += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, pagination.Limit, (pagination.Offset-1)*pagination.Limit)

	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		log.Println("Error GetListWarehouses:", err)
		return nil, err
	}
	defer rows.Close()

	var warehouses []repository.WarehouseModel
	for rows.Next() {
		var warehouse repository.WarehouseModel
		err := rows.Scan(
			&warehouse.ID,
			&warehouse.ShopID,
			&warehouse.ProductID,
			&warehouse.Code,
			&warehouse.Name,
			&warehouse.Stock,
			&warehouse.Location,
			&warehouse.Status,
			&warehouse.CreatedAt,
			&warehouse.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}
