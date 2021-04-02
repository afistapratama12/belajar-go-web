package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/afistapratama12/belajar-go-web/entity"
)

//Function home route
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	// log.Println("ini adalah path yang dimasukkan : ", r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "learning golang web",
		"status":  200,
		"author":  "afista pratama",
		"content": "ini adalah home page golang web dengan menggunakan native golang",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

}

func HandelListProduct(w http.ResponseWriter, r *http.Request) {
	// dataEntity := entity.Product{ID: 1, Name: "Mobilio", Price: 200_000_000, Stock: 5}

	// template view action
	dataEntity := []entity.Product{
		{ID: 1, Name: "Mobilio1", Price: 200_000_000, Stock: 5},
		{ID: 2, Name: "Mobilio2", Price: 250_000_000, Stock: 10},
		{ID: 3, Name: "Mobilio3", Price: 300_000_000, Stock: 3},
	}

	tmpl, err := template.ParseFiles(path.Join("views", "list.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, dataEntity)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}
}

//Function to render router Author
func HandlerAuthor(w http.ResponseWriter, r *http.Request) {
	data := "Afista Pratama"

	// using write to render data if user get path
	w.Write([]byte(data))
}

//Function to render router Product with query
func HandlerProduct(w http.ResponseWriter, r *http.Request) {

	// get query /produts?id=5
	id := r.URL.Query().Get("id")

	// convert id from string to number
	idNumber, err := strconv.Atoi(id)

	log.Println(id, "ini adalah id")
	log.Println(idNumber, "ini adalah id number")

	// validate data query
	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"queryId": &idNumber,
		"author":  "afista pratama",
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error inernal server", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

	// product := "Number product " + id
	// w.Write([]byte(product))
	// fmt.Fprintf(w, "Product page : %d", idNumber)
}

func HandleTables(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "table.html"))

	if err != nil {

		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {

		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "form.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error internal server", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET, aku kembalian datanya"))
	case "POST":
		w.Write([]byte("ini adalah POST, data diterima"))
	default:
		http.Error(w, "error bukan GET dan POST", http.StatusBadRequest)
	}
}

func HandleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// mendapat data dari method POST form
		errParse := r.ParseForm()

		if errParse != nil {
			log.Println(errParse.Error())
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")
		check := r.Form.Get("checkbox")

		data := map[string]interface{}{
			"name":  name,
			"pesan": pesan,
			"check": check,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "eror internal server", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "ini bukan post", http.StatusBadRequest)
}
