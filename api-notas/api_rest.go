package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura de uma nota
type Nota struct {
	ID       int    `json:"id"`
	Titulo   string `json:"titulo"`
	Conteudo string `json:"conteudo"`
}

// Slice para armazenar as notas (simulando um banco de dados)
var notas = []Nota{
	{ID: 1, Titulo: "Primeira Nota", Conteudo: "Essa é minha primeira nota!"},
}

// Criar uma nova nota
func criarNota(c *gin.Context) {
	var novaNota Nota
	if err := c.BindJSON(&novaNota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novaNota.ID = len(notas) + 1
	notas = append(notas, novaNota)
	c.JSON(http.StatusCreated, novaNota)
}

// Listar todas as notas
func listarNotas(c *gin.Context) {
	c.JSON(http.StatusOK, notas)
}

// Buscar uma nota por ID
func buscarNota(c *gin.Context) {
	id := c.Param("id")

	for _, nota := range notas {
		if id == string(rune(nota.ID)) {
			c.JSON(http.StatusOK, nota)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"erro": "Nota não encontrada"})
}

// Atualizar uma nota existente
func atualizarNota(c *gin.Context) {
	var notaAtualizada Nota
	if err := c.BindJSON(&notaAtualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	id := c.Param("id")

	for i, nota := range notas {
		if id == string(rune(nota.ID)) {
			notas[i] = notaAtualizada
			notas[i].ID = nota.ID
			c.JSON(http.StatusOK, notas[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"erro": "Nota não encontrada"})
}

// Deletar uma nota
func deletarNota(c *gin.Context) {
	id := c.Param("id")

	for i, nota := range notas {
		if id == string(rune(nota.ID)) {
			notas = append(notas[:i], notas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "Nota deletada"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"erro": "Nota não encontrada"})
}

// Configurar as rotas da API
func main() {
	r := gin.Default()

	r.POST("/notas", criarNota)
	r.GET("/notas", listarNotas)
	r.GET("/notas/:id", buscarNota)
	r.PUT("/notas/:id", atualizarNota)
	r.DELETE("/notas/:id", deletarNota)

	r.Run(":8080") // Inicia o servidor na porta 8080
}
