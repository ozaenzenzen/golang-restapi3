package entities

type ProductResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Product
}
