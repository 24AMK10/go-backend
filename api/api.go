package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"backend.com/example/go-backend/json_structs"
	"backend.com/example/go-backend/mongo_ops"
)


type ApiRes struct {
	Message string `json:"message"`
	// Resolver string ``
}
func StartMongo()  {
	fmt.Println("conncting to Mongo", time.Now())

	_ = mongo_ops.ConnectMongo()

	fmt.Println("mongo Connected ", time.Now())
}

func HandleSignUp(w http.ResponseWriter, req *http.Request)  {
	if req.Method == http.MethodPost {
		decoder := json.NewDecoder(req.Body)
		signup := json_structs.SignupDetails{}
		err := decoder.Decode(&signup)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Request Body Invalid", http.StatusBadRequest)
			return
		}
		result := mongo_ops.SignupMethod(signup)

		if result == true{
			jsonData := ApiRes{"success"}
			w.Header().Set("content-type", "application/json")
			js, _ := json.Marshal(jsonData)
			w.Write(js)
			fmt.Println(time.Now())
		}else{
			jsonData := ApiRes{"UserExists"}
			w.Header().Set("content-type", "application/json")
			js, _ := json.Marshal(jsonData)
			w.Write(js)
			fmt.Println(time.Now())
		}
		
	}else{
		http.Error(w, "Invalid Method for the Route", http.StatusNotFound)
		return
	}

}