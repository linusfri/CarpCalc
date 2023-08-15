package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/linusfri/calc-api/helper"
	authModel "github.com/linusfri/calc-api/models/auth"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile("./logs/logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

		if err != nil {
			helper.HandleErr(err)
			return
		}

		defer f.Close()

		f.WriteString(fmt.Sprintf(
			"Got a %s request on url %s \n",
			r.Method,
			r.URL,
		))

		next.ServeHTTP(w, r)
	})
}

func AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("jwt")

		if err != nil {
			w.WriteHeader((http.StatusUnauthorized))
			return
		}

		if !authModel.VerifyJWT(authCookie.Value) || authCookie.Value == "" {
			w.WriteHeader((http.StatusUnauthorized))
			return
		}

		next.ServeHTTP(w, r)
	})
}
