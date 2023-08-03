package Models

type Company struct {
	ID       uint      `gorm:"primaryKey;autoIncrement;type: int not null"`
	Name     string    `gorm:"type: varchar(255) not null"`
	Address  string    `gorm:"type: varchar(255) not null"`
	Email    string    `gorm:"type: varchar(255) not null"`
	Phone    string    `gorm:"type: varchar(255) not null"`
	Products []Product `gorm:"constraint:OnDelete:CASCADE"` // Relationship with 'Products' table
}

type Product struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;type: int not null"`
	Name      string `gorm:"type: varchar(255) not null"`
	CompanyID uint   `gorm:"type: int not null"`
}
