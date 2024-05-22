package data

import (
	"fmt"
)

// Product defines the structure for an API product
type Product struct {
	// the ID for the product
	//
	// required : false
	// min : 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required : true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required : false
	// max length: 10000
	Description string `json:"description"`

	// the description for this product
	//
	// required : true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// The SKU for the product
	//
	// required : true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// Products defines a slice of Product
// Products: Collection of Product
type Products []*Product

// GetProducts() : function
// Gets the collection of products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the database
// If a product is not found, returns ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// UpdateProduct replaces a product in the database with the given product p
// If a product with the given p.ID does not exist in the database,
//
//	return ProductNotFound Error
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product to the database
	productList[i] = &p

	return nil
}

// AddProduct adds a product to the database
func AddProduct(p Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])
	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
// HELPER Function
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Ice Americano",
		Description: "Espresso with ice and water",
		Price:       3.99,
		SKU:         "ice-amr-abc",
	},

	&Product{
		ID:          2,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc142",
	},
}
