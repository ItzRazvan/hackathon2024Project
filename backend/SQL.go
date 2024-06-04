package backend

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:razvan@tcp(127.0.0.1:3306)/school"

//create the User

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   string `json:"access"`
}

func comparePasswordAndEmail(email string, password string) (bool, uint) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return false, 0
	}

	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return false, 0
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false, 0
	}

	var id uint
	db.Raw("SELECT id FROM users WHERE email = ?", email).Scan(&id)

	return true, id

}

func createUser(name string, email string, password string, access string) (uint, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return 0, err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: hash,
		Access:   access,
	}

	db.Create(&user)

	var id uint
	db.Raw("SELECT id FROM users WHERE email = ?", email).Scan(&id)

	return id, nil
}

type UserAbsenceMonth struct {
	Id        uint `json:"id"`
	January   int  `json:"january"`
	February  int  `json:"february"`
	March     int  `json:"march"`
	April     int  `json:"april"`
	May       int  `json:"may"`
	June      int  `json:"june"`
	July      int  `json:"july"`
	August    int  `json:"august"`
	September int  `json:"september"`
	October   int  `json:"october"`
	November  int  `json:"november"`
	December  int  `json:"december"`
}

type UserAbsenceYear struct {
	Id    uint `json:"id"`
	Y2023 int  `json:"2023"`
	Y2024 int  `json:"2024"`
	Y2025 int  `json:"2025"`
	Y2026 int  `json:"2026"`
}

func (UserAbsenceMonth) TableName() string {
	return "user_absence_month"
}

func (UserAbsenceYear) TableName() string {
	return "user_absence_year"
}

func createUserAbsenceMonth(id uint) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	userAbsenceMonth := UserAbsenceMonth{
		Id:        id,
		January:   0,
		February:  0,
		March:     0,
		April:     0,
		May:       0,
		June:      0,
		July:      0,
		August:    0,
		September: 0,
		October:   0,
		November:  0,
		December:  0,
	}

	db.Create(&userAbsenceMonth)

	return nil
}

func createUserAbsenceYear(id uint) error {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	userAbsenceYear := UserAbsenceYear{
		Id:    id,
		Y2023: 0,
		Y2024: 0,
		Y2025: 0,
		Y2026: 0,
	}

	db.Create(&userAbsenceYear)

	return nil
}

func isAdmin(c echo.Context) bool {
	id := getId(c)

	if id == 0 {
		fmt.Println("Nu am gasit userul in baza de date")
		return false
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return false
	}

	var user User
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return false
	}

	return user.Access == "admin"

}
