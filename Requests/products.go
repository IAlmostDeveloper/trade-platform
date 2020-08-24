package requests

import (
	"encoding/json"
	"net/http"
	dbaccess "trade-platform/DBAccess"
	service "trade-platform/Service"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	response := dbaccess.GetAllProducts()
	//fmt.Println(r.Cookie("token"))
	//service.WriteToken("1234")
	token, err := r.Cookie("token")
	if token != nil{
		if !service.CheckToken(token.String()){
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := service.GetIdFromPath(r.URL.Path)
	response := dbaccess.FindProductById(id)
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil{
		var result = service.GetLoginFromToken(token.Value)
		w.Write([]byte(result))
	}
}
