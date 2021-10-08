package main 

import (
        "fmt"
        "log"
        "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Welcome!\n")
}

// domain/users
func CreateUser(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Creating New User!\n")
}

// domain/users/<id here>
func GetUser(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "GET user by ID\n")
}


// domain/posts

func UserPost(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Post Users Data\n")
}

// domain/posts/<id here>
func GetPost(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "GET POST \n")
}

//domain/posts/users/<Id here> 
// apply pagination 

func AllUserPost(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "All User Post Using Pagination \n")
}


func handleRequest(){
	http.HandleFunc("/",Index)
	http.HandleFunc("/users",CreateUser)
	http.HandleFunc("/users/:id",GetUser)
	http.HandleFunc("/posts",UserPost)
	log.Fatal(http.ListenAndServe(":8001",nil))
}


func main(){
   handleRequest()
}