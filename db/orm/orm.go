package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// TestOrm gorm
func TestOrm() {
	db, err := gorm.Open(mysql.Open("root:1234@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=UTC"))
	if err != nil {
		fmt.Printf("connection database error! %s", err)
	}

	err = db.AutoMigrate(User{})
	if err != nil {
		fmt.Printf("migrate table error! %s", err)
		return
	}

	birthday := time.Now()
	db.Table("user").Create(&User{
		Name:     "likeai",
		Age:      24,
		Birthday: &birthday,
	})

	var u []User
	//err = db.Table("user").Where("name like ?", "%like%").Find(&u).Error
	//if err != nil {
	//	fmt.Printf("select error! %s", err)
	//}
	//for i := range u {
	//	user := u[i]
	//	fmt.Printf("id %d name %s, age %d, birthday %s \t\n", user.ID, user.Name, user.Age, user.Birthday.Local().Format("2006-03-02 15:04:05"))
	//}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Table("user").Where("id", 7).Update("name", "dsada").Error
		err = tx.Table("user").Where("id", 4).Update("age", "dsafsafas").Error
		return err
	})
	if err != nil {
		fmt.Printf("tx error! %s", err)
	}
	//err = db.Table("user").Where("id", 2).Update("name", "chenaiquan").Error
	//if err != nil {
	//	fmt.Printf("update error! %s", err)
	//}
	//
	err = db.Table("user").Find(&u).Error
	if err != nil {
		fmt.Printf("select error! %s", err)
	}
	for i := range u {
		user := u[i]
		fmt.Printf("id %d name %s, age %d, birthday %s \t\n", user.ID, user.Name, user.Age, user.Birthday.Local().Format("2006-03-02 15:04:05"))
	}
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday *time.Time
}

//TableName 将 User 的表名设置为 `profiles`
//func (User) TableName() string {
//	return "user"
//}

func (u User) TableName() string {
	return "users"
}
