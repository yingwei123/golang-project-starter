package server

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Router struct {
	ResourcesPath string
	ServerURL     string
	MongoDBClient mongoDBClient
}

type mongoDBClient interface {
}

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	template := rt.NewTemplateHandlerFactory(filepath.Join(rt.ResourcesPath, "templates"))

	router.Handle("/", template.DefaultHandler("default.gohtml"))

	n := negroni.New()
	n.Use(negroni.NewLogger())

	faviconMiddleware := negroni.NewStatic(http.Dir(filepath.Join(rt.ResourcesPath, "favicon.ico")))
	faviconMiddleware.Prefix = "/favicon.ico"
	n.Use(faviconMiddleware)

	publicMiddleware := negroni.NewStatic(http.Dir(filepath.Join(rt.ResourcesPath, "public")))
	publicMiddleware.Prefix = "/public"

	n.Use(publicMiddleware)

	n.UseHandler(router)
	n.ServeHTTP(w, r)
}
