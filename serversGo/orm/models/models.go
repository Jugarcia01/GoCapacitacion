package models

import (
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nombre      string
	Descripcion string
}
