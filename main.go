package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != http.MethodGet { //r.Method != "GET"
		http.Error(w, "Please give me GET request", 400)
		return
	}
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != http.MethodPost {
		http.Error(w, "Please give me POST request", 400)
		return
	}
	//r.Body => description, title, price, imageUrl => Product struct
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server is running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "A fresh orange from the orange tree",
		Price:       10.99,
		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "A fresh apple from the apple tree",
		Price:       13.99,
		ImgUrl:      "https://5.imimg.com/data5/AK/RA/MY-68428614/apple.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "A fresh banana from the banana tree",
		Price:       1.99,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/banana-cavendish_0.png?itok=xIgYOIE_-9FKLRtCr",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "A fresh grapes from the grapes tree",
		Price:       2.99,
		ImgUrl:      "https://farmfreshfundraising.com/wp-content/uploads/2017/10/271156-grapes.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "A fresh mango from the mango tree",
		Price:       3.99,
		ImgUrl:      "https://png.pngtree.com/png-vector/20240125/ourmid/pngtree-sweet-mango-fruit-png-png-image_11495826.png",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "A fresh pineapple from the pineapple tree",
		Price:       4.99,
		ImgUrl:      "https://www.healthxchange.sg/adobe/dynamicmedia/deliver/dm-aid--c06c2aed-90cf-4360-a423-7f053b2a44d9/pineapple-health-benefits-and-ways-to-enjoy.jpg?preferwebp=true",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
