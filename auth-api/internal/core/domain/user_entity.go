package domain

type StudentEntity struct{
 	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name      string
    Email     string `gorm:"uniqueIndex"`
    Username  string `gorm:"uniqueIndex"`
    Password  string
    CPF       string `gorm:"uniqueIndex"`
    Phone     string
    Address   string
    CreatedAt string `gorm:"autoCreateTime"`
    UpdatedAt *string
}

type CustomerEntity struct{
	 ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name      string
    Email     string `gorm:"uniqueIndex"`
    Username  string `gorm:"uniqueIndex"`
    Password  string
    CNPJ      string
    CPF       string `gorm:"uniqueIndex"`
    Phone     string
    Address   string
    CreatedAt string `gorm:"autoCreateTime"`
    UpdatedAt *string
}
