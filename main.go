package main 

import (
        "fmt"
        "log"
        "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Welcome!\n")
}


func handleRequest(){

	http.HandleFunc("/",Index)
	log.Fatal(http.ListenAndServe(":8001",nil))


}


func main(){
   handleRequest()
}