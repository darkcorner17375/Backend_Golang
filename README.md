# Golang heroku deploy
This is a simple restful API project that golang + mysql deployed to heroku. You can access it with this url:
```
https://golang-gaintime.herokuapp.com/
```

# Documentation
## Health check
check server is alive:
<b>GET</b>
```
https://golang-gaintime.herokuapp.com/health/
```
Response (Status: 200)
```
{
message: "Successful Go-Lives!!!!!!"
}
```
<img src="/result/health.png" />

## Register
Registering a new user

<b>POST</b>
```
https://golang-gaintime.herokuapp.com/api/auth/register
```

Request Body
```
{
    "name": "User1",
    "email": "test@gmail.com",
    "password": "123456"
}
```

Response success (Status: 201)
<img src="/result/register.png" />

Response error (Status : 400)
<img src="/result/register-400.png" />

## Login
Authenticate by email & password

<b>POST</b>
```
https://golang-gaintime.herokuapp.com/api/auth/login
```

Request body
```
{
    "email": "test@gmail.com",
    "password": "123456"
}
```
Response Success (Status: 200)
<img src="/result/login.png" />

Response error, wrong credential (Status: 401)
<img src="/result/login-401.png" />

## Get Profile
Get user info from logged

<b>GET</b>
```
https://golang-gaintime.herokuapp.com/api/user/profile
```

Headers
```
Authorization: yourToken
```
<img src="/result/header.png" />

Response success (status: 200)
<img src="/result/profile.png" />

## Update profile
Update user data who logged

<b>PUT</b>
```
https://golang-gaintime.herokuapp.com/api/useer/profile
```

Headers
```
Authorization: yourToken
```

Request Body
<img src="/result/update-body.png" />

Response success (Status: 200)
<img src="/result/update-profile.png" />


<b>=============================================</b>


## All Todos (Based on user who logged)
Only shows todo item by user

<b>GET</b>
```
https://golang-gaintime.herokuapp.com/api/todos
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
<img src="/result/todo-all.png" />

Response success (Status: 401)
<img src="/result/todo-noToken.png" />


## Create Todo item
Create a todo item with owner

<b>POST</b>
```
https://golang-gaintime.herokuapp.com/api/todos
```

Headers
```
Authorization: yourToken
```

Request body
<img src="/result/todo-insert.png" />

Response success (Status: 201)
<img src="/result/todo-insert201.png" />

## Find Todo by ID
Find todo by id

<b>GET</b>
```
https://golang-gaintime.herokuapp.com/api/todos/{id}
```

Headers
```
Authorization: yourToken
```


Response success (Status: 200)
<img src="/result/todo-id200.png" />

Response error, wrong id (Status: 404)
request body
```
{id} = 1, but no "1" data
```
<img src="/result/todo-id404.png" />

## Update Todo by ID
Update todo info from owner by todo id
<b>PUT</b>
```
https://golang-gaintime.herokuapp.com/api/todos/{id}
```

origin data
<img src="/result/todo-origin-data.png" />

Request body
<img src="/result/todo-update-new.png" />

Response success (Status: 200)
<img src="/result/todo-update200.png" />


## Delete Todo
only delete onw todo by id

<b>DELETE</b>
```
https://golang-gaintime.herokuapp.com/api/todos/{id}
```

origin data
<img src="/result/user data.png" />

Response success (Status: 200)
delete "id = 5" item
<img src="/result/todo-delete.png" />

Check result
<img src="/result/delete-result.png" />

<b>=============================================</b>

This is Restful API demo
Thanks!!
