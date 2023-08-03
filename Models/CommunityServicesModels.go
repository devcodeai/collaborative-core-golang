package Models

type Community struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;type: int not null"`
	Name        string `gorm:"type: varchar(255) not null"`
	Description string `gorm:"type: varchar(255) not null"`
	Members     string `gorm:"type: varchar(255) not null"`
}
