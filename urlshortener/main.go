package main

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//	yaml := `
	//- path: /urlshort
	//	url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//	url: https://github.com/gophercises/urlshort/tree/final
	//`

	//yamlHandler, err := urlshort.YAML

}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandleFunc {
	return func(w http.ResonseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
