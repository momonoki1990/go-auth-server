# go-auth-server

Auth server by go, returning jwt token in cookie if access to /token.
Use:

- gorilla/mux (for simple router)
- golang-jwt/jwt

## Get started

```
$ vim .env
# Generate by '$ openssl rand -hex 32'
JWT_SECRET_KEY={random value}

$ go build
$ ./go-auth-server
```

## Send request to server

```
# Access to public endpoint
$ curl http://localhost:8080/public
This is public handler

# Access to private endpoint without token
$ curl http://localhost:8080/private
Internal Server Error

# Get token
$ curl -X POST http://localhost:8080/token -d "email=sample@sample.com&password=mypassword" -H "application/x-www-form-urlencoded" -c cookie.txt
Succesfully token returned

# Access to private endpoint with token
$ curl http://localhost:8080/private  -b cookie.txt
Authorization succeeded (successfully accessed to private handler)

# Modify token string in cookie.txt
# Access to private endpoint with invalid token
Invalid token
```

## Reference

[【Go×WAF】うほうほ！！gorilla を分かりやすくまとめてみた【mux/context/schema】Go,gorilla](https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a)
[Go 言語で理解する JWT 認証 実装ハンズオン](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
