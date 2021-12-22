package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

func CorsSetup(isPublic bool) *cors.Cors {
	options := cors.Options{
		AllowedOrigins: GetEnv().Cors.AllowedOrigins,
		AllowedMethods: GetEnv().Cors.AllowedMethods,
		AllowedHeaders: GetEnv().Cors.AllowedHeaders,
	}

	if isPublic {
		options.AllowedOrigins = []string{"*"}
	}

	return cors.New(options)
}

func SecureSetup() *secure.Secure {
	options := secure.Options{
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	}

	return secure.New(options)
}

func RouterConfig(r *chi.Mux) {
	res := func(message string) map[string]interface{} {
		return map[string]interface{}{
			"success": false,
			"data":    nil,
			"message": message,
		}
	}

	msg := NewMessage()

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		Render().JSON(w, http.StatusMethodNotAllowed, res(msg.MethodNotAllowed))
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		Render().JSON(w, http.StatusNotFound, res(msg.RouteNotFound))
	})

	fmt.Println("\nRegistered Routes")

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("\t%s\t%s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}
