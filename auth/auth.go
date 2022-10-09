package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenCookieName string = "user"

// JWT claim struct
type User struct {
	Admin bool
	Sub string
    Name string
	Iat int64
	Exp int64
	jwt.StandardClaims
}

// Generate token
func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &User{
		Admin: true,
		Sub: "54546557354",
		Name: "taro",
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString, err
}

// Verify JWT token in cookie (middleware)
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sample Middleware has been called")

		// Get token from cookie
		tokenCookie, err := r.Cookie(TokenCookieName)
		// including no cookie error ('http: named cookie not present')
		if err != nil {
			log.Println("Error at cookie parse:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
			return
		}
		log.Println(tokenCookie.Value)

		// Verify token string
		token, err := jwt.ParseWithClaims(tokenCookie.Value, &User{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			log.Println("Error at token string parse:", err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		// Decode jwt claims to struct
		claims := token.Claims.(*User)
		fmt.Printf("%+v", claims)

		next.ServeHTTP(w, r)
	}) 
}