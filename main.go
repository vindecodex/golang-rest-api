package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type sidebar struct {
	ID      int    `json:"ID"`
	Element string `json:"Element"`
}

var aside = []sidebar{
	{
		ID:      2,
		Element: "<div class='left-menu indeiFlower'><ul><li><a href='#'>Pokemon 1</a></li><li><a href='#'>Pokemon 2</a></li><li><a href='#'>Pokemon 3</a></li><li><a href='#'>Pokemon 4</a></li><li><a href='#'>Pokemon 5</a></li></div>",
	},
	{
		ID:      3,
		Element: "<div class='left-menu indeiFlower'><ul><li><a href='#'>Pokemon 2</a></li></ul></div>",
	},
}

func createSidebar(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(aside)
	fmt.Println("----------New Request----------")
}

func main() {
	fmt.Println("***API IS RUNNING***")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create-sidebar", createSidebar)
	http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
}
