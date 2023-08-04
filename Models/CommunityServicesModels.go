package Models

type Community struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name        string `gorm:"type: varchar(255) not null" json:"name"`
	Description string `gorm:"type: varchar(255) not null" json:"description"`
	Members     string `gorm:"type: varchar(255) not null" json:"members"`
}
