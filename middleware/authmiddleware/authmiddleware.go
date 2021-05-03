package authmiddleware

import (
	"context"
	"encoding/json"
	"net/http"
)

type contextKey string

var (
	contextKeyAuthtoken = contextKey("user")
)

func authHandler(ctx context.Context, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		user, err := getuser(authToken)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAuthtoken, user)
		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getuser(authToken string) (string, error) {
	return "sourav", nil
}

func adminHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	user := ctx.Value(contextKeyAuthtoken)
	json.NewEncoder(w).Encode(user)
}

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(adminHandler)
	// commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	// http.Handle("/admin", commonHandlers.Append(authHandler).ThenFunc(adminHandler))
	// http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	// http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	// http.ListenAndServe(":8080", nil)
}
