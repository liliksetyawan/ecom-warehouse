package in

type WarehouseDTOIn struct {
	ID        int    `json:"id"`
	ShopID    int    `json:"shop_id"`
	ProductID int    `json:"product_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Stock     int    `json:"stock"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
