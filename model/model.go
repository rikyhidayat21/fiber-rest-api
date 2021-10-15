package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model																		// Add some metadata fields to the table
	ID 					uuid.UUID	`gorm:"type:uuid"`			// Explicitly specify the type
	Title 			string
	SubTitle		string
	Text				string

}