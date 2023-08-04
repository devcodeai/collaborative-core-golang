package Models

type Company struct {
	ID       uint      `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name     string    `gorm:"type: varchar(255) not null" json:"name,omitempty"`
	Address  string    `gorm:"type: varchar(255) not null" json:"address,omitempty"`
	Email    string    `gorm:"type: varchar(255) not null" json:"email,omitempty"`
	Phone    string    `gorm:"type: varchar(255) not null" json:"phone,omitempty"`
	Products []Product `gorm:"constraint:OnDelete:CASCADE" json:"-"` // Relationship with 'Products' table
}

type Product struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;type: int not null" json:"id"`
	Name      string `gorm:"type: varchar(255) not null" json:"name,omitempty"`
	CompanyID uint   `gorm:"type: int not null" json:"company_id,omitempty"`
}
