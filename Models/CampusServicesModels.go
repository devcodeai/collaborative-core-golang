package Models

type Campus struct {
	ID             uint    `gorm:"primaryKey;autoIncrement;type: int not null"`
	UniversityName string  `gorm:"type: varchar(255) not null"`
	Location       string  `gorm:"type: varchar(255) not null"`
	Website        string  `gorm:"type: varchar(255) not null"`
	Majors         []Major `gorm:"constraint:OnDelete:CASCADE"` // Relationship with 'Majors' table
}

// specify table name as 'campuses' for the 'Campus' model
func (Campus) TableName() string {
	return "campuses"
}

type Major struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type: int not null"`
	Name     string `gorm:"type: varchar(255) not null"`
	CampusID uint   `gorm:"type: int not null"`
}
