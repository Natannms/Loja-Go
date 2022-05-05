package models

import (
	"goweb/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAll() []Produto {
	db := db.ConectaComBancoDeDados()

	query, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func GetById(id string) []Produto {
	db := db.ConectaComBancoDeDados()

	query, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		return []Produto{}
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}

func Create(p Produto) {
	db := db.ConectaComBancoDeDados()

	_, err := db.Exec("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)", p.Nome, p.Descricao, p.Preco, p.Quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func Update(id string, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	query, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConectaComBancoDeDados()

	delete, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}
