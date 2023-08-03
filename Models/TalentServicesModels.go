package Models

type Talent struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name   string `gorm:"type: varchar(255) not null" json:"name"`
	Email  string `gorm:"type: varchar(255) not null" json:"email"`
	Skills string `gorm:"type: varchar(255) not null" json:"skills"`
}
