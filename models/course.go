package models

import (
	"gorm.io/gorm"
)

type Course struct {
	Cod                 string `json:"cdsCod" gorm:"primaryKey"`
	Nome                string `json:"nomeCds"`
	NomeEng             string `json:"nomeCdsEng"`
	TipoCorso           string `json:"tipoCorsoDes"`
	Durata              int    `json:"durataAnni"`
	ClaMiurCod          string `json:"claMiurCod"`
	ClaMiurDes          string `json:"claMiurDes"`
	Sedi                string `json:"sedi"`
	CoordinatoreCognome string `json:"cognome"`
	CoordinatoreNome    string `json:"nome"`
	CoordinatoreEmail   string `json:"email"`
	AaAttId             int    `json:"aaAttId"`
	LinguaIta           string `json:"linguaIta"`
	LinguaEng           string `json:"linguaEng"`
	LastSuaYear         int    `json:"lastSuaYear"`
	StatoAttivazione    string `json:"statoAttivazione"`
	Attivo              bool   `json:"attivo"`
	Esaurimento         bool   `json:"esaurimento"`
	InBreve             string `json:"inBreve"`

	DipartimentoCod string     `json:"dipCod" gorm:"size:10"`
	Dipartimento    Department `gorm:"foreignKey:DipartimentoCod;references:Cod" json:"-"`

	Teachings []Teaching `gorm:"foreignKey:CodCorso;references:Cod" json:"-"`
}

type CoursesResponse struct {
	TotalItems  int      `json:"totalItems"`
	TotalPages  int      `json:"totalPages"`
	CurrentPage int      `json:"currentPage"`
	Items       []Course `json:"items"`
}

type CourseDescriptionResponse struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	ValidityYear int    `json:"validityYear"`
}

func (Course) TableName() string {
	return "courses"
}

func (c *Course) AfterFind(tx *gorm.DB) (err error) {
	if c.Dipartimento == (Department{}) {
		dipartimento := Department{}
		if err := tx.Where("cod = ?", c.DipartimentoCod).First(&dipartimento).Error; err == nil {
			c.Dipartimento = dipartimento
		}
	}

	return
}
