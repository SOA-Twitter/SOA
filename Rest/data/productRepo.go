package data

import (
	"errors"
	"strconv"
	"time"
)

//In this file we localize all the operations on our database (currently in-memory database with productList)
//When we decide to migrate this to a genuine solution we would have to make changes just to this file

//Return all the products
func GetAll() Products {
	return productList
}

//Return only active products
func GetProducts() Products {
	list := Products{}

	for _, prod := range productList {
		if len(prod.DeletedOn) == 0 {
			list = append(list, prod)
		}
	}

	return list
}

func Addproduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = ""
	productList = append(productList, p)
}

func PutProduct(p *Product, id int) error {
	for _, currentProd := range GetProducts() {
		if currentProd.ID == id && len(currentProd.DeletedOn) == 0 {
			currentProd.Name = p.Name
			currentProd.Description = p.Description
			currentProd.Price = p.Price
			currentProd.SKU = p.SKU
			currentProd.UpdatedOn = time.Now().UTC().String()

			return nil
		}
	}

	return errors.New("Item with id " + strconv.FormatInt(int64(id), 10) + " not found")
}

func DeleteProduct(id int) error {
	for _, currentProd := range GetProducts() {
		if currentProd.ID == id && len(currentProd.DeletedOn) == 0 {
			currentProd.DeletedOn = time.Now().UTC().String()
			currentProd.UpdatedOn = time.Now().UTC().String()
			return nil
		}
	}
	return errors.New("Item with id " + strconv.FormatInt(int64(id), 10) + " not found")
}

func getNextID() int {
	max := 0

	for _, currentProd := range GetProducts() {
		if currentProd.ID > max {
			max = currentProd.ID
		}
	}

	return max + 1
}

//Our initial database
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
