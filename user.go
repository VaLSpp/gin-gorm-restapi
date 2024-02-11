package main

type Users struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string `json:"username" gorm:"type:varchar(50); unique_index"`
	PasswordHash string `json:"password_hash" gorm:"type:varchar(255); unique_index"`
	Email        string `json:"email" gorm:"type:varchar(255); unique_index"`
	RoleId       uint   `json:"role_id" gorm:"ref: > roles.id"`
}
