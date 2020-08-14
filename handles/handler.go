package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong"
	"net/http"
)

func SetupRoutes(clnt, scrt, securityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, "", Index(tmpl), make(map[string]bool))).Methods(http.MethodGet)

	//r.HandleFunc("/blog", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.GetArticles(tmpl), "blog.articles.view")).Methods(http.MethodGet)
	//r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	//r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	//r.HandleFunc("/blog/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.ViewArticle(tmpl), "blog.articles.view")).Methods(http.MethodGet)

	return r
}
