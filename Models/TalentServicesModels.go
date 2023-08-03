package Models

type Talent struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;type: int not null"`
	Name   string `gorm:"type: varchar(255) not null"`
	Email  string `gorm:"type: varchar(255) not null"`
	Skills string `gorm:"type: varchar(255) not null"`
}
