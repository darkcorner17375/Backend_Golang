# Golang heroku deploy
This is a simple restful API project that golang + mysql deployed to heroku. You can access it with this url:
```
https://golang-gaintime.herokuapp.com/
```

#Documentation
## health check
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