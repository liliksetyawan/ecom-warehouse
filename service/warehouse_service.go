package service

import (
	"database/sql"
	"ecom-warehouse/dao"
	"ecom-warehouse/dto/in"
	"ecom-warehouse/dto/out"
	"ecom-warehouse/repository"
	"ecom-warehouse/server"
	"fmt"
	"log"
)

func CreateWarehouse(req *in.WarehouseDTOIn) (*out.WarehouseDTOOut, error) {
	if req.ShopID == 0 || req.ProductID == 0 || req.Code == "" || req.Name == "" {
		return nil, fmt.Errorf("shop_id, product_id, code, and name are required")
	}

	warehouse := &repository.WarehouseModel{
		ShopID:    sql.NullInt64{Int64: int64(req.ShopID), Valid: true},
		ProductID: sql.NullInt64{Int64: int64(req.ProductID), Valid: true},
		Code:      sql.NullString{String: req.Code, Valid: true},
		Name:      sql.NullString{String: req.Name, Valid: true},
		Stock:     sql.NullInt64{Int64: int64(req.Stock), Valid: true},
		Location:  sql.NullString{String: req.Location, Valid: req.Location != ""},
		Status:    sql.NullString{String: req.Status, Valid: req.Status != ""},
	}

	db := server.DBConn
	id, err := dao.CreateWarehouse(db, warehouse)
	if err != nil {
		return nil, fmt.Errorf("error saving warehouse: %v", err)
	}

	return &out.WarehouseDTOOut{
		ID:        int(id),
		ShopID:    req.ShopID,
		ProductID: req.ProductID,
		Code:      req.Code,
		Name:      req.Name,
		Stock:     req.Stock,
		Location:  req.Location,
		Status:    req.Status,
	}, nil
}

func GetWarehouseByID(id int64) (*out.WarehouseDTOOut, error) {
	db := server.DBConn
	warehouseModel, err := dao.GetWarehouseByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("warehouse not found")
		}
		return nil, err
	}

	return &out.WarehouseDTOOut{
		ID:        int(warehouseModel.ID.Int64),
		ShopID:    int(warehouseModel.ShopID.Int64),
		ProductID: int(warehouseModel.ProductID.Int64),
		Code:      warehouseModel.Code.String,
		Name:      warehouseModel.Name.String,
		Stock:     int(warehouseModel.Stock.Int64),
		Location:  warehouseModel.Location.String,
		Status:    warehouseModel.Status.String,
		CreatedAt: warehouseModel.CreatedAt.String,
		UpdatedAt: warehouseModel.UpdatedAt.String,
	}, nil
}

func GetWarehouses(req *in.GetListDTO) ([]*out.WarehouseDTOOut, error) {
	db := server.DBConn

	pagination := in.Pagination{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
	}

	warehouses, err := dao.GetListWarehouses(db, pagination)
	if err != nil {
		log.Println("Error fetching warehouses:", err)
		return nil, err
	}

	var warehouseDTOs []*out.WarehouseDTOOut
	for _, w := range warehouses {
		dto := &out.WarehouseDTOOut{
			ID:        int(w.ID.Int64),
			ShopID:    int(w.ShopID.Int64),
			ProductID: int(w.ProductID.Int64),
			Code:      w.Code.String,
			Name:      w.Name.String,
			Stock:     int(w.Stock.Int64),
			Location:  w.Location.String,
			Status:    w.Status.String,
			CreatedAt: w.CreatedAt.String,
			UpdatedAt: w.UpdatedAt.String,
		}
		warehouseDTOs = append(warehouseDTOs, dto)
	}

	return warehouseDTOs, nil
}
