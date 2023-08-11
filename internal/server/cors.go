package server

import (
	"net/http"
	"strings"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}

		if strings.HasPrefix(r.URL.Path, "/preview") {
			pathSplitted := strings.Split(r.URL.Path, "/")
			if len(pathSplitted) <= 3 {
				next.ServeHTTP(w, r)
			} else {
				workspaceName := pathSplitted[2]
				pageName := pathSplitted[3]
				fs := http.FileServer(http.Dir("/data/repository/" + workspaceName + "/" + pageName + "/build/"))
				x := http.StripPrefix("/preview/"+workspaceName+"/"+pageName+"/", fs)
				x.ServeHTTP(w, r)
			}

		} else {
			next.ServeHTTP(w, r)
		}

	})
}
