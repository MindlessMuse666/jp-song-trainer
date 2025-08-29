package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

type Material struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	Title       string          `json:"title" db:"title"`
	Description string          `json:"description" db:"description"`
	Difficulty  string          `json:"difficulty" db:"difficulty"`
	Status      string          `json:"status" db:"status"`
	CreatedBy   uuid.UUID       `json:"created_by" db:"created_by"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	Blocks      []MaterialBlock `json:"blocks,omitempty"`
}

type MaterialBlock struct {
	ID               uuid.UUID `json:"id" db:"id"`
	MaterialID       uuid.UUID `json:"material_id" db:"material_id"`
	OriginalText     string    `json:"original_text" db:"original_text"`
	TranslatedText   string    `json:"translated_text" db:"translated_text"`
	GrammarBreakdown string    `json:"grammar_breakdown" db:"grammar_breakdown"`
	BlockOrder       int       `json:"block_order" db:"block_order"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type PersonalDictionary struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Word        string    `json:"word" db:"word"`
	Translation string    `json:"translation" db:"translation"`
	Notes       string    `json:"notes" db:"notes"`
	AddedDate   time.Time `json:"added_date" db:"added_date"`
}

type Achievement struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Criteria    string    `json:"criteria" db:"criteria"`
}

type UserAchievement struct {
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	AchievementID uuid.UUID `json:"achievement_id" db:"achievement_id"`
	Progress      int       `json:"progress" db:"progress"`
	Achieved      bool      `json:"achieved" db:"achieved"`
	AchievedDate  time.Time `json:"achieved_date" db:"achieved_date"`
}

type ErrorReport struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	MaterialID  uuid.UUID `json:"material_id" db:"material_id"`
	Description string    `json:"description" db:"description"`
	IsFixed     bool      `json:"is_fixed" db:"is_fixed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
