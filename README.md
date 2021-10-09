# InstaCloneGo
Instagram Backend API Using GO 

# Setup

```go 
 go run server.go
```
# API END POINTS 
1. CREATE USERS :http://localhost:8001/users
2. GET A USER USING ID : http://localhost:8001/users/:id
3. CREATE A POST  : http://localhost:8001/posts
4. GET A POST USING ID :http://localhost:8001/posts/:id
5. LIST ALL POSTS OF A USER http://localhost:8001/posts/users/:id

# CREATE USERS
1. GET 
![image](https://user-images.githubusercontent.com/68312849/136656301-061c2b53-3fef-4097-8c64-19f551f57dbc.png)
2. POST
![image](https://user-images.githubusercontent.com/68312849/136656570-12509640-89b6-4196-863a-146956d98066.png)
![image](https://user-images.githubusercontent.com/68312849/136656586-eb464c6c-a7c3-4ff9-939b-618859e4dba6.png)
3. Email must be unique
![image](https://user-images.githubusercontent.com/68312849/136656634-78db628d-b6f6-4e20-9652-c7ee70b5f17d.png)

 # GET A USER USING ID
1. GET
![image](https://user-images.githubusercontent.com/68312849/136656736-12096711-1555-4da1-a68c-f4cfa943d5a5.png)
2. User Not Found
![image](https://user-images.githubusercontent.com/68312849/136656763-d8c3f252-6715-493e-9d5a-34179e873828.png)
3. POST 
![image](https://user-images.githubusercontent.com/68312849/136656775-22d1c7e2-58b6-4297-b4c6-73c445ab0d03.png)
# CREATE A POST
1. POST 
![image](https://user-images.githubusercontent.com/68312849/136656920-ba02a189-6986-4b77-af22-2fa9e567419e.png)
![image](https://user-images.githubusercontent.com/68312849/136657037-958f3962-2887-4fbb-a62e-b488b31fd145.png)
2. using unauthorized email
![image](https://user-images.githubusercontent.com/68312849/136657091-818a0f54-ca7b-4f2b-a3ca-dfe5acc8404e.png)

#  GET A POST USING ID 
1. GET 
![image](https://user-images.githubusercontent.com/68312849/136657180-b9fda864-7542-45e5-9587-3c2386db5ee2.png)
2. Invalid Post Id
![image](https://user-images.githubusercontent.com/68312849/136657196-3974579b-6551-47d9-9a57-53d5b7f52dde.png)
3. POST 
![image](https://user-images.githubusercontent.com/68312849/136657208-2517c862-cc43-415a-92e8-d351c8de8a51.png)

# LIST ALL POSTS OF A USER
1. GET
![image](https://user-images.githubusercontent.com/68312849/136657239-6535ad57-7258-424b-b7b9-f887db9730d7.png)
2. INVALID USER ID
![image](https://user-images.githubusercontent.com/68312849/136657250-a00b0f5d-236e-41ba-b22d-e5f967630edd.png)
3. POST
![image](https://user-images.githubusercontent.com/68312849/136657255-bb122c5f-7db8-4107-a967-f69ee634153d.png)



