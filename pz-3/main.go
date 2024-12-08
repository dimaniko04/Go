package main

import (
	"fmt"
	"html/template"
	"net/http"
	"pz-3/database"
	"pz-3/routes"

	"github.com/gorilla/mux"
)

var templates *template.Template

func dict(values ...interface{}) map[string]interface{} {
	d := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		key, _ := values[i].(string)
		d[key] = values[i+1]
	}
	return d
}

func init() {
	funcMap := template.FuncMap{
		"dict": dict,
	}

	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*/*.go.tmpl"))
}

func main() {
	db, err := database.Db()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	routes.InitRoutes(r, templates)

	fmt.Println("Server is started on http://localhost:8080")
	panic(http.ListenAndServe(":8080", r))
}
