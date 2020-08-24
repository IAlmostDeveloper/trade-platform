package requests

import (
	"encoding/json"
	"net/http"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			response := dbaccess.GetAllProducts()
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
	}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil{
		if service.CheckToken(token.Value) {
			id := service.GetIdFromPath(r.URL.Path)
			response := dbaccess.FindProductById(id)
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
	}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			var product entities.Product
			json.NewDecoder(r.Body).Decode(&product)
			owner := service.GetLoginFromToken(token.Value)
			product.Owner = owner
			dbaccess.InsertProduct(product)
			return
		}
	}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
