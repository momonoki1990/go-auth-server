# go-auth-server

Auth server by go
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

## Reference

[【Go×WAF】うほうほ！！gorilla を分かりやすくまとめてみた【mux/context/schema】Go,gorilla](https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a)
[Go 言語で理解する JWT 認証 実装ハンズオン](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
