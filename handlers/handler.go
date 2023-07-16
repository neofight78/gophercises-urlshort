package handlers

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type Redirect struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}

	var redirects []Redirect
	err := yaml.Unmarshal(yml, &redirects)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, redirect := range redirects {
		pathsToUrls[redirect.Path] = redirect.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
