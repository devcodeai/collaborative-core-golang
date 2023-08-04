package Models

type Community struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name        string `gorm:"type: varchar(255) not null" json:"name,omitempty"`
	Description string `gorm:"type: varchar(255) not null" json:"description,omitempty"`
	Members     string `gorm:"type: varchar(255) not null" json:"members,omitempty"`
}
