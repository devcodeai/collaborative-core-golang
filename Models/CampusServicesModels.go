package Models

type Campus struct {
	ID             uint    `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	UniversityName string  `gorm:"type: varchar(255) not null" json:"university_name,omitempty"`
	Location       string  `gorm:"type: varchar(255) not null" json:"location,omitempty"`
	Website        string  `gorm:"type: varchar(255) not null" json:"website,omitempty"`
	Majors         []Major `gorm:"constraint:OnDelete:CASCADE" json:"-"` // Relationship with 'Majors' table
}

// specify table name as 'campuses' for the 'Campus' model
func (Campus) TableName() string {
	return "campuses"
}

type Major struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name     string `gorm:"type: varchar(255) not null" json:"name,omitempty"`
	CampusID uint   `gorm:"type: int not null" json:"campus_id,omitempty"`
}
