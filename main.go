package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func response(w http.ResponseWriter, code int, src interface{}) {
	var body []byte

	switch s := src.(type) {
	case []byte:
		body = s
		fmt.Println("byte", s)

	case string:
		body = []byte(s)
	default:
		body, _ = json.Marshal(src)

	}

	w.WriteHeader(code)
	w.Write(body)

}

var mov = []string{"El diablo a todas horas ",
	"Guasón",
	"Liga de la justicia oscura: guerra Apokolips ",
	"Un lugar en silencio",
	"Annabelle 3: vuelve a casa",
	"Bad Boys para siempre",
	"It: capítulo dos",
	"Hellboy",
	"Godzilla 2: el rey de los monstruos",
	"Creed: corazón de campeón",
	"Creed II: defendiendo el legado",
	"La monja",
	"Mad Max: furia en el camino",
	"destiny 845",
	"Feliz dia de tu muerte",
}

func findAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := params["q"][0]
	var l []string

	for _, m := range mov {
		nameMovie := strings.Split(m, "")
		if query == nameMovie[0] {
			l = append(l, m)
		}
	}

	if 0 == len(l) {
		response(w, 200, "No hay nada")
		return
	}

	response(w, 200, l)
}

func main() {
	fmt.Println("Api running in :8080")

	r := mux.NewRouter()
	r.HandleFunc("/api/movies", findAll)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}

}
