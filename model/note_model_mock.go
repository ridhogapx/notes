package model

import (

	"github.com/stretchr/testify/mock"
)

type MockNoteModel struct {
  mock.Mock
}

func (m *MockNoteModel) CreateNote(data *Notes) error {
  arg := m.Called(data)

  return arg.Error(0)
}

func (m *MockNoteModel) FindNote(id string) (*Notes, error) {
  arg := m.Called(id)
  
  res := arg.Get(0).(Notes)
  return &res, arg.Error(0)
}  
