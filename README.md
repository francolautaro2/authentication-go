# authentication-go
I am learning user authentication with jwt token in Golang.

## Project structure
This project uses jwt to authenticate users.<br>
* In pkg/server are all functios to create a http server and listen.<br>
* In pkg/token are all functions to make and validate jwt tokens.<br>
* In pkg/handlers are all routes as a logIn and SignUp.<br>

## How does this project work
We create a post request to log in to the server: <br>
![LogIn](https://github.com/francolautaro2/authentication-go/assets/69493845/664075c5-c451-4345-98f8-287b9e6fbe6c)

We get home page from server: <br>
![home](https://github.com/francolautaro2/authentication-go/assets/69493845/05f361b2-1762-43d5-806d-04933c4b160c)
