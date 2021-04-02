package entity

// membuat struct untuk dijadikan penyimpanan data di dalam slice
type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) Catatan() string {
	if p.Price > 200000000 {
		return "harga mobil mewah"
	} else {
		return "harga masih standar"
	}
}
