package crm

type Question struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Statement string `gorm:"size:255"`
	Option1   string `gorm:"size:255"`
	Option2   string `gorm:"size:255"`
	Option3   string `gorm:"size:255"`
	Option4   string `gorm:"size:255"`
	Ans       int
}
