package requests

import (
	"encoding/json"
	"net/http"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	if service.AuthorizeUser(r.Cookie("token")){
			response := dbaccess.GetAllProducts()
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
	}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")){
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if service.AuthorizeUser(token, err) {
			var product entities.Product
			json.NewDecoder(r.Body).Decode(&product)
			owner, _ := service.GetLoginAndEmailFromToken(token.Value)
			product.Owner = owner
			dbaccess.InsertProduct(product)
			return
	}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func PurchaseProduct(w http.ResponseWriter, r *http.Request){
	if service.AuthorizeUser(r.Cookie("token")){
			var productInfo entities.ProductInfo
			json.NewDecoder(r.Body).Decode(&productInfo)
			result := dbaccess.FindProductByName(productInfo.Name)
			var keyId entities.KeyIdJson
			keyId.KeyId = result.Id
			js, err := json.Marshal(keyId)
			if err!=nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
