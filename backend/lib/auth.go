package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/* Auth middleware */
func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		if bearerToken == "" {
			HttpError(w, http.StatusUnauthorized, "Authorization")
			return
		}

		token, err := VerifyToken(bearerToken)
		if err != nil {
			HttpError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claim := token.Claims.(jwt.MapClaims)
		userId := chi.URLParam(r, "userId")
		if userId != "" && userId != claim["user_id"].(string) {
			HttpError(w, http.StatusUnauthorized, "Authorization")
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func AdminMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		if bearerToken == "" {
			HttpError(w, http.StatusUnauthorized, "Authorization")
			return
		}

		token, err := VerifyToken(bearerToken)
		if err != nil {
			HttpError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claim := token.Claims.(jwt.MapClaims)
		if claim["role"].(string) != "admin" {
			HttpError(w, http.StatusUnauthorized, "Authorization")
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

/*JWT token */
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(userId string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":     1,
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"iat":     time.Now().Unix(),
			"role":    role,
		})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// TODO add expired token check

	return token, nil
}

/* Encrypt/decrypt */
func HashPasswordWithSalt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("[warning] unable to generate hash of password with length", len(password))
		return "", err
	}
	return string(bytes), nil
}

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
