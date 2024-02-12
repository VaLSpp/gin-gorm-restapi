package models

type Users struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string `json:"username" gorm:"type:varchar(50); unique_index"`
	PasswordHash string `json:"password_hash" gorm:"not null; type:varchar(255); unique_index default:null"`
	Email        string `json:"email" gorm:"not null; type:varchar(255); unique_index default:null"`
	RoleId       uint   `json:"role_id" gorm:"unique_index"`
}
