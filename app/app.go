package main

import(
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w, "The Restful API working ...")
}

func main(){
	router := httprouter.New()
	router.GET("/", indexHandler)

	env := os.Getenv("APP_ENV")
	if(env == "production"){
		log.Println("Rest Server in Production mode")
	}else{
		log.Println("Rest Server is non Prod environment")
	}

	http.ListenAndServe(":8080", router)
}