## HTTP
- protocol
    - rules
- `GET, POST, PUT, PATCH, DELETE`
    - world's 99% applications were built by these 5
- status code
    - 200: OK
    - 201: resource create in server
    - 400: dead request
    - 404: not found
    - 500: internal server error


## Handler function
- r: request from frontend to backend with info
- w: response from backend

```go
func getProducts(w http.RequestWriter, r *http.Request) {
    if r.Method != http.MethodGet {  // also, we can write "GET"
        http.Error(w, "Plz give me GET request", 400);
        return;
    }

    encoder := json.NewEncoder(w)
	encoder.Encode(productList)  // productList: a slice of Product struct
}

mux.HandleFunc("/products", getProducts);
```