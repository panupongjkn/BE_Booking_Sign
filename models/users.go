package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Password     string    `json:"password" db:"password"`
	Fname        string    `json:"fname" db:"fname"`
	Lname        string    `json:"lname" db:"lname"`
	Organization string    `json:"organization" db:"organization"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) CreateModel(data map[string]interface{}) bool {
	if data["fname"] == nil {
		return false
	}
	if data["lname"] == nil {
		return false
	}
	if data["organization"] == nil {
		return false
	}
	u.Fname = data["fname"].(string)
	u.Lname = data["lname"].(string)
	u.Organization = data["organization"].(string)
	return true
}