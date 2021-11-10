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
Docker Test: "OK!!!!!!!",
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