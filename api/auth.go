package api

import (
	"net/http"

	"github.com/spf13/viper"
)

func (a *API) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if viper.GetBool("fail2rest.auth_enabled") {
			token := r.Header.Get("X-Auth-Token")

			if token == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if token != viper.GetString("fail2rest.secret") {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
