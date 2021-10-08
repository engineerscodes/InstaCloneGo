package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

   if err != nil {
    log.Fatal(err)
  }
  ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
   if err != nil {
    log.Fatal(err)
  }

   defer client.Disconnect(ctx)

   err = client.Ping(ctx,readpref.Primary())

   if err != nil {
    log.Fatal(err)
  }

  database,err :=client.ListDatabaseNames(ctx,bson.M{})
  if err != nil {
	   log.Fatal(err)
  }
   
  fmt.Println(database)

   handleRequest()
  

}