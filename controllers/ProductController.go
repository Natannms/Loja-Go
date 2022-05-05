package controllers

import (
	"goweb/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func UpdatePge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		produtos := models.GetById(id)
		temp.ExecuteTemplate(w, "Update", produtos)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAll()

	temp.ExecuteTemplate(w, "Produtos", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("name")
		descricao := r.FormValue("description")
		preco := r.FormValue("price")
		quantidade := r.FormValue("quantity")

		isAuthorized := true

		convPreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			isAuthorized = false
			log.Println("erro ao converter preço para int64", err)
		}

		convQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			isAuthorized = false
			log.Println("erro ao converter quantidade para int", err)
		}

		if isAuthorized {
			produto := models.Produto{Nome: nome, Descricao: descricao, Preco: convPreco, Quantidade: convQuantidade}
			models.Create(produto)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("name")
		descricao := r.FormValue("description")
		preco := r.FormValue("price")
		quantidade := r.FormValue("quantity")

		isAuthorized := true

		convPreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			isAuthorized = false
			log.Println("erro ao converter preço para int64", err)
		}

		convQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			isAuthorized = false
			log.Println("erro ao converter quantidade para int", err)
		}

		if isAuthorized {
			models.Update(id, nome, descricao, convPreco, convQuantidade)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		models.Delete(id)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
