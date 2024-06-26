package product

import (
	"humoshop/pkg/db"
	"log"
)

// Запрос к базе данных о продукте
func GetProductDB(product_id ProductDatabaseStruct) (p ProductDatabaseStruct, err error) {
	row := db.DB.QueryRow("SELECT * FROM product WHERE product_id=$1", product_id.Product_id)

	err = row.Scan(&p.Product_id, &p.Name, &p.Price, &p.In_stock, &p.Category_id)
	if err != nil {
		log.Println("Ошибка row")
	}
	return
}

// Добавление нового продукта В базу данных
func CreateProductDB(p ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("INSERT INTO product(name, price, in_stock) VALUES ($1, $2, $3)", p.Name, p.Price, p.In_stock, p.Category_id)
	return
}

// Обновление информации о продукте
func UpdateProductDB(p ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("UPDATE product SET name=$1, price=$2, in_stock=$3 WHERE product_id=$4", p.Name, p.Price, p.In_stock, p.Product_id, p.Category_id)
	return	
}

// Удаление продукта из Базы Данных
func DeleteProductDB(p ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("DELETE FROM product WHERE product_id=$1", p.Product_id)
	return
}
