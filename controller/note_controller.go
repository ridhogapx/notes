package controller

import (
	"log"
	"net/http"
	"notes/model"
	"notes/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NoteController struct {
	Repos repository.NoteRepository
}

func (controller *NoteController) AddNote(ctx *gin.Context) {
	var payload model.NoteRequest

	s := sessions.Default(ctx)
	ctx.ShouldBind(&payload)

	id := uuid.New().String()

	err := controller.Repos.CreateNote(&model.Note{
		ID:     id,
		Title:  payload.Title,
		Body:   payload.Body,
		Author: s.Get("user_email").(string),
	})

	if err != nil {
		log.Println("Can't create note:", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Catatan berhasil ditambahkan",
	})
}
