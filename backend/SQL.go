package backend

import (
	"fmt"
	"strconv"
	"time"

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

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	var id uint
	db.Raw("SELECT id FROM users WHERE email = ?", email).Scan(&id)

	return id, nil
}

func getNameFromId(id uint) (string, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err
	}

	var name string
	db.Raw("SELECT name FROM users WHERE id = ?", id).Scan(&name)

	return name, nil

}

type UserAbsenceMonth struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	January   int    `json:"january"`
	February  int    `json:"february"`
	March     int    `json:"march"`
	April     int    `json:"april"`
	May       int    `json:"may"`
	June      int    `json:"june"`
	July      int    `json:"july"`
	August    int    `json:"august"`
	September int    `json:"september"`
	October   int    `json:"october"`
	November  int    `json:"november"`
	December  int    `json:"december"`
}

type UserAbsenceYear struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Y2023 int    `json:"2023"`
	Y2024 int    `json:"2024"`
	Y2025 int    `json:"2025"`
	Y2026 int    `json:"2026"`
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

	name, err := getNameFromId(id)

	if err != nil {
		return err
	}

	userAbsenceMonth := UserAbsenceMonth{
		Id:        id,
		Name:      name,
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

	name, err := getNameFromId(id)

	if err != nil {
		return err
	}

	userAbsenceYear := UserAbsenceYear{
		Id:    id,
		Name:  name,
		Y2023: 0,
		Y2024: 0,
		Y2025: 0,
		Y2026: 0,
	}

	db.Create(&userAbsenceYear)

	return nil
}

func deleteUser(id uint) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Exec("DELETE FROM users WHERE id = ?", id)

	return nil
}

func isAdmin(c echo.Context) bool {
	id := getId(c)

	if id == 0 {
		fmt.Println("Nu am gasit userul in baza de date")
		return false
	}

	user := getUserById(id)

	return user.Access == "admin"

}

func getUserById(id uint) User {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return User{}
	}

	var user User
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return User{}
	}

	return user
}

func getAbsenteNow(id uint, month string, year string) (int, int) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return 0, 0
	}

	var absenteLuna int
	var absenteAn int

	db.Raw("SELECT "+month+" FROM user_absence_month WHERE id = ?", id).Scan(&absenteLuna)
	db.Raw("SELECT "+year+" FROM user_absence_year WHERE id = ?", id).Scan(&absenteAn)

	return absenteLuna, absenteAn
}

func addAbsentaToDB(id uint, month string, year string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	actualYear := year[1:]
	actualYearInt, err := strconv.Atoi(actualYear)
	if err != nil {
		return err
	}

	if actualYearInt == time.Now().Year() {
		db.Exec("UPDATE user_absence_month SET "+month+" = "+month+" + 1 WHERE id = ?", id)
	}
	db.Exec("UPDATE user_absence_year SET "+year+" = "+year+" + 1 WHERE id = ?", id)

	return nil
}

func addAbsenteToAllAbsences(id uint, day int, hour int, minute string, month string, year int) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	absenta := absenta{
		Id:     id,
		Year:   year,
		Month:  month,
		Day:    day,
		Hour:   hour,
		Minute: minute,
	}

	db.Table("absente").Create(&absenta)

	return nil
}

func getAllAbsenteMonth() []UserAbsenceMonth {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	var absenteLuna []UserAbsenceMonth

	db.Find(&absenteLuna)

	return absenteLuna
}

func getAllAbsenteYear() []UserAbsenceYear {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	var absenteAn []UserAbsenceYear

	db.Find(&absenteAn)

	return absenteAn
}

func getAbsenteMonthById(id uint) UserAbsenceMonth {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return UserAbsenceMonth{}
	}

	var absenteLuna UserAbsenceMonth

	db.Where("id = ?", id).First(&absenteLuna)

	return absenteLuna
}

func getAllAbsencesForUser(id uint) []absenta {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	var absences []absenta

	db.Table("absente").Where("id = ?", id).Find(&absences)

	return absences
}
