package controllers

import (
	"curso3/database"
	"curso3/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	
	var ps []models.Personalidade
	rows, err := database.DB.Query("select id, nome, historia from personalidades")
	if err != nil {
		fmt.Println("Erro na query")
	}
	for rows.Next() {
		p := models.Personalidade{}
		err := rows.Scan(&p.Id, &p.Nome, &p.Historia)
		if err != nil {
			log.Panic()

		}
		ps = append(ps, p)
	}
	json.NewEncoder(w).Encode(ps)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	row := database.DB.QueryRow("select id, nome, historia from personalidades where id = $1::int", id)

	err := row.Scan(&p.Id, &p.Nome, &p.Historia)
	if err != nil {
		log.Panic()
	}
	json.NewEncoder(w).Encode(p)
}

func Criar(w http.ResponseWriter, r *http.Request) {

	var data models.NovaPersonalidade

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// vars := mux.Vars(r)
	// nome := vars["nome"]
	// historia := vars["historia"]
	// var personalidade models.Personalidade
	// json.NewDecoder(r.Body).Decode(&personalidade)

	_, err = database.DB.Exec("insert into personalidades (nome, historia) values ($1, $2)", data.Nome, data.Historia)
	if err != nil {
		http.Error(w, "Erro ao inserir a personalidade", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, `{"status":"Sucesso","mensagem": "inserido com sucesso"}`)
	fmt.Println("Personalidade " + data.Nome + " inserida com sucesso")

}

func Deletar(w http.ResponseWriter, r *http.Request) {
	var data models.NovaPersonalidade
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("delete from personalidades where nome = $1", data.Nome)
	if err != nil {
		http.Error(w, "Erro ao deletar a personalidade", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, `{"status":"Sucesso","mensagem": "Deletado com sucesso"}`)
	fmt.Println("Personalidade " + data.Nome + " deletada com sucesso")
}

func Editar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	if data.Historia == "" {
		_, err = database.DB.Exec("update personalidades set nome = $1 where id = $2", data.Nome, id)
		if err != nil {
			http.Error(w, "Erro ao atualizar a personalidade: ", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
	} else if data.Nome == "" {
		_, err = database.DB.Exec("update personalidades set historia = $1 where id = $2", data.Historia, id)
		if err != nil {
			http.Error(w, "Erro ao atualizar a personalidade", http.StatusBadRequest)
			return
		}
	} else {
		_, err = database.DB.Exec("update personalidades set nome = $1, historia = $2 where id = $3", data.Nome, data.Historia, id)
		if err != nil {
			http.Error(w, "Erro ao atualizar a personalidade", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
	}
	fmt.Fprintf(w, `{"status":"Sucesso","mensagem": "Editado com sucesso"}`)
	fmt.Println("Personalidade atualizada com sucesso")

}
