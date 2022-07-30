package entity

type Prodcut struct {
	ID int
	Name string
	Price int
	Stock int
}

func (p Prodcut) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stock terbtas"
	}
	return status
}