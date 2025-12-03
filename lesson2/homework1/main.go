package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) Totalvalue(store []Product) float64 {
	total := p.Price * float64(p.Stock)
	return total
}
func (p *Product) Sell(amount int) (success bool, message string) {
	str := "售卖" + strconv.Itoa(amount) + "本:"
	if amount > p.Stock {
		return false, str + "库存不足"
	}
	p.Stock -= amount
	return true, str + "销售成功,剩余：" + strconv.Itoa(p.Stock) + "本"
}
func (p *Product) Restock(amount int) {
	p.Stock += amount
}
func main() {
	store := (Product{
		Name:  "Go语言编程书",
		Price: 89.5,
		Stock: 10,
	})
	_, message := store.Sell(5)
	fmt.Println(message)
	store.Restock(20)
	_, message = store.Sell(30)
	fmt.Println(message)
	fmt.Printf("商品信息：\n商品：%s，单价：￥%g，库存：%d件\n库存总价值：%.2f\n", store.Name, store.Price, store.Stock, store.Totalvalue([]Product{store}))
}