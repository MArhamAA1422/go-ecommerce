package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please give GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Please give POST request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server is running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product {
		ID: 1,
		Title: "Orange",
		Description: "Orange is red. I Love orange.",
		Price: 100,
		ImgUrl: "https://thfvnext.bing.com/th/id/OIP.hPuzy_Bzmw0VA7Vgjkwk4wHaFV?w=224&h=180&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	}

	prd2 := Product {
		ID: 2,
		Title: "Apple",
		Description: "Apple is green. I Love apple.",
		Price: 90,
		ImgUrl: "https://thfvnext.bing.com/th/id/OIP.n_CQTHRbeNZD_bP_zVa0NwHaH3?w=176&h=187&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	}

	prd3 := Product {
		ID: 3,
		Title: "Banana",
		Description: "Banana is yellow. I Love banana.",
		Price: 50,
		ImgUrl: "https://thfvnext.bing.com/th/id/OIP.Ji1rlIbzuRqNEc9nBHEH3QHaGu?w=217&h=197&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	}

	// prd4 := Product {
	// 	ID: 4,
	// 	Title: "Mango",
	// 	Description: "Mango is beautiful. I Love mango.",
	// 	Price: 150,
	// 	ImgUrl: "https://thfvnext.bing.com/th/id/OIP.GqqWLmKBZrwrWWFRPIFFEQHaHa?w=199&h=199&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	// }

	// prd5 := Product {
	// 	ID: 5,
	// 	Title: "Grapes",
	// 	Description: "Grapes is a fruit. I Love grapes.",
	// 	Price: 130,
	// 	ImgUrl: "https://thfvnext.bing.com/th/id/OIP.kAZRHXusfaaVeBxDTJC8wQHaLI?w=115&h=180&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	// }

	// prd6 := Product {
		// ID: 6,
		// Title: "Guava",
		// Description: "Guava is green. I Love guava.",
		// Price: 60,
		// ImgUrl: "https://thfvnext.bing.com/th/id/OIP.veuJftghwFhQP487poLEtgHaEy?w=288&h=187&c=7&r=0&o=7&cb=thfvnext&dpr=1.3&pid=1.7&rm=3",
	// }

	productList = append(productList, prd1, prd2, prd3)
}