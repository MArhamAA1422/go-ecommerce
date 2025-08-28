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

    encoder := json.NewEncoder(w)  // json is a package
	encoder.Encode(productList)  // productList: a slice of Product struct
}

mux.HandleFunc("/products", getProducts);
```

### `[]` = list, `{}` = object

## JSON encoder customization
- json tag
- we want to pass "id" not "ID"
- json package can only access prop starting with uppercase letter
```go
type Product struct {
	ID int `json:"id"`
}
```

## Response Header
- CORS error handle
```go
w.Header().Set("Access-Control-Allow-Origin", "*")
```
- Content type (by default: text)
```go
w.Header().Set("Content-Type", "application/json")
```