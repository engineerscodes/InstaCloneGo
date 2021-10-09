package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
var Id autoInc
type NewUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Time  time.Time `json:"Time"`
}



type DbUser struct{
	UserID          int      `json:"UserID" bson:"UserID"`
	Name   string        `json:"Name" bson:"Name"`
	Email string         `json:"Email" bson:"Email"`
	Password string      `json:"Password" bson:"Password"`
	Time  time.Time      `json:"Time" bson:"Time"`

}

type Post struct{
	PostID          int      `json:"PostID" bson:"PostID"`
	Email string         `json:"Email" bson:"Email"`
	Caption string        `json:"caption" bson:"caption"`
	Time  time.Time      `json:"Time" bson:"Time"`
	ImageURl string      `json:"ImageURl" bson:"ImageURl"`
}

type Invalid struct{
	Message string  `json:"data"`
}



func index() {
	http.HandleFunc("/",HttpReqHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

var(
	users=regexp.MustCompile(`^\/users[\/]*$`)
	getuser=regexp.MustCompile(`^\/users\/([0-9]+)$`)
	CreatePost=regexp.MustCompile(`^\/posts[\/]*$`)
	GetPost=regexp.MustCompile(`^\/posts\/([0-9]+)$`)
	GetUserPost=regexp.MustCompile(`^\/posts\/users\/([0-9]+)$`) ///posts/users/<Id here>' user=id od the user which is unqiue

)

func HttpReqHandler(w http.ResponseWriter, r *http.Request) {
	 
    switch {
		 
	case users.MatchString(r.URL.Path):
		CreateUser(w,r)
		return
	case getuser.MatchString(r.URL.Path):
	    GetUser(w,r)
		return
	case CreatePost.MatchString(r.URL.Path):
		PostData(w,r)
		return
	case GetPost.MatchString(r.URL.Path):
	GetPostId(w,r)
	return
	case GetUserPost.MatchString(r.URL.Path):
	GetPostOFUser(w,r)
	return
	}

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "text/html")
		t, _ := template.ParseFiles("Templates/User.html")
		t.Execute(w, nil)
	} else {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	
	col :=client.Database("users").Collection("usersinstaclone")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(col)
		//r.ParseForm()
		r.ParseMultipartForm(10<<20)
		w.Header().Set("content-type", "application/json")
		//fmt.Println("Creating User ")
		var user NewUser = NewUser{Name: r.FormValue("name"), Email: r.FormValue("email"), Time: time.Now()}
		CheckUser:=DbUser{}
		err2:=col.FindOne(ctx, bson.M{"Email":r.FormValue("email")}).Decode(&CheckUser)
		if err2 !=nil{   //if email is not already reg
	    var dbdata DbUser=DbUser{UserID:Id.ID(),Name :r.FormValue("name"),Email: r.FormValue("email"),Password: r.FormValue("password") ,Time: time.Now()}
		res,Inserterr:=col.InsertOne(ctx,dbdata)
		if Inserterr != nil {
			fmt.Println("ERROR IN INSERT TO DB")
		}
		fmt.Println(reflect.TypeOf( res))
		json.NewEncoder(w).Encode(user)
		}else{
        data:="Email already Registered ! "+r.FormValue("email")
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
		}
		
	}
	
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := getuser.FindStringSubmatch(r.URL.Path)[1]
	IntUserID, _ := strconv.ParseInt(userID, 0, 64)
    fmt.Println(IntUserID)
	w.Header().Set("content-type", "application/json")
  if r.Method == "GET" {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	res:=DbUser{}
	collection := client.Database("users").Collection("usersinstaclone")
	databases, _ := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)
	collection.FindOne(ctx, bson.M{"UserID":IntUserID}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	if res.Email !="" ||  res.Name !="" ||  res.Password !="" {
	json.NewEncoder(w).Encode(res)
    }else{
	
		data:="USER ID : "+userID+" NOT FOUND"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
	}else{
		
		data:="Method Not Allowed"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
}

func PostData(w http.ResponseWriter, r *http.Request){
	if r.Method =="GET"{
		data:="Method not Allowed"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
	if r.Method == "POST" {
	w.Header().Set("content-type", "application/json")
      r.ParseMultipartForm(10<<20)
	  imageurl:=r.FormValue("ImageURl")
	  UEmail:=r.FormValue("email")
	  Caption:=r.FormValue("Caption")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	res:=DbUser{}
	UserCollection:=client.Database("users").Collection("usersinstaclone")
	err2:=UserCollection.FindOne(ctx, bson.M{"Email":UEmail}).Decode(&res)
	if err2 ==nil{ //there is a mail match
	PostCollection:=client.Database("users").Collection("Posts")
	UserPost :=Post{PostID :Id.ID(),Email:UEmail,ImageURl: imageurl,Caption: Caption,Time: time.Now()}
	response,Inserterr:=PostCollection.InsertOne(ctx,UserPost)
	if Inserterr != nil {
		log.Fatal(Inserterr)
	}
	fmt.Print(response)
	json.NewEncoder(w).Encode(UserPost)
	}else{  // mail is not reg
		data:="EMAIL NOT FOUND !"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
	}
 
}
func GetPostId(w http.ResponseWriter, r *http.Request) {
	PostID := GetPost.FindStringSubmatch(r.URL.Path)[1]
	IntPostID, _ := strconv.ParseInt(PostID, 0, 64)
    fmt.Println(IntPostID)
	w.Header().Set("content-type", "application/json")
	 if r.Method == "GET" { 
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	res:=Post{}
	collection := client.Database("users").Collection("Posts")
	collection.FindOne(ctx, bson.M{"PostID":IntPostID}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	if res.Email !="" ||  res.Caption !="" ||  res.ImageURl !="" {
	json.NewEncoder(w).Encode(res)
    }else{
	
		data:="POST ID : "+PostID+" NOT FOUND"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
	
      
	}else{
		data:="Method Not Allowed"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}

}


func GetPostOFUser(w http.ResponseWriter, r *http.Request){
    UserID:=GetUserPost.FindStringSubmatch(r.URL.Path)[1]
	IntUserID, _ := strconv.ParseInt(UserID, 0, 64)
	w.Header().Set("content-type", "application/json")
	if r.Method == "GET" {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	res:=DbUser{}
	collection := client.Database("users").Collection("usersinstaclone")
	collection.FindOne(ctx, bson.M{"UserID":IntUserID}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	if res.Email !="" ||  res.Name !="" ||  res.Password !="" {
        
		PostCol:=client.Database("users").Collection("Posts")
		//Postres:=Post{}
		cursor,_:=PostCol.Find(ctx, bson.M{"Email":res.Email})
		  for cursor.Next(ctx) {
			var Postres bson.M
			if err3 := cursor.Decode(&Postres); err3 != nil {
				log.Fatal(err3)
			}
          json.NewEncoder(w).Encode(Postres)
		
		}


    }else{
	
		data:="USER ID : "+UserID+" NOT FOUND"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
	}else{
		data:="Method Not Allowed"
		invalidmess:=Invalid{Message:data }
		json.NewEncoder(w).Encode(invalidmess)
	}
}




func main() {
	index()
}