package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name string

	Email string              'gorm:"uniqueIndex"'

	Screenings []Screening    'gorm:foreignKey:UserID'
}

type Patient struct {
	gorm.Model

	Name string

	Email string           'gorm:"uniqueIndex"'

	Screenings []Screening  'gorm:foreignKey:PatientID'
}

type Symptom struct {
	gorm.Model

	Name string

	Screenings []Screening  'gorm:foreignKey:SymptomID'
}
type Covid struct {
	gorm.Model

	Name string

	Screenings []Screening  'gorm:foreignKey:CovidID'
}
type Screening struct {
	gorm.Model

	UserID	*uint
	User   =  User
	
	PatientID *uint
	Patient   = Patient

	SymptomID  *uint
	Symptom	   = Symptom

	CovidID		*uint
	Covid		= Covid
}
