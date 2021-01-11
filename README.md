# JWT_REST_GIN_GORM_MySQL

Web service CRUD using Golang with GIN for create REST api, MySQL as database, Viper as environment variable, JWT for secure service and Cookies to store token.

**Prerequisites**

1. [Go](https://golang.org/)
2. [Gin](github.com/gin-gonic/gin)
3. [Mysql](https://www.mysql.com/downloads/)
4. [Viper](https://github.com/spf13/viper)
5. [BCrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
6. [JWT](https://github.com/dgrijalva/jwt-go)
7. [UUID](https://github.com/segmentio/ksuid)

## Getting Started
1. Firstly, we need to get Gin, MySQL, Viper, jwt and ksuid for UUID library dependencies for install it
```
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/segmentio/ksuid
```

2. Import dump.sql to your MySQL and configure your credential in folder resource.

3. To run application,open cmd in your project directory and type
```
go run main.go
```