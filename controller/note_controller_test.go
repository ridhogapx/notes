package controller

import (
	"notes-v1/model"
	"testing"
	"time"
)

func GetSingleNote(data *model.Notes) *model.Notes {
  return data
}

func TestNoteController(t *testing.T) {
  // Create instance of mock 
  m := new(model.MockNoteModel)
  
  // Setup expectation
  m.On("FindNote", "id-note-foo-bar").Return(&model.Notes{
    ID: "id-note-foo-bar",
    Title: "Foo",
    Body: "Bar",
    IsDone: false,
    CreatedAt: time.Now(),
  }, nil)

  // Todo: 
  // Add stub 
}