package repo

import (
	"fmt"
	"github.com/JYama321/gin-crud-template/entity"
	"github.com/jinzhu/gorm"
	"os"
	"testing"
)

func connectTestDatabase() (db *gorm.DB, err error) {
	db,err = gorm.Open("mysql", "root:db_root_password@tcp(db:3306)/gin_template_test?charset=utf8&parseTime=True&loc=Local")
	return
}

func connectTest() *gorm.DB {
	db, err := connectTestDatabase()
	if err != nil {
		panic(err)
	}
	return db
}

func newTestRepository() *UserRepository{
	db := connectTest()
	return &UserRepository{db}
}



func setup() {
	r := newTestRepository()
	r.db.DropTableIfExists("users")

	r.db.CreateTable(&entity.User{})
	testPassword := generatePasswordDigest("testPassword")
	user := entity.User{UserName: "test1", Email: "test1@gmail.com", PasswordDigest: testPassword}
	r.db.Create(&user)
	user = entity.User{UserName: "test2", Email: "test2@gmail.com", PasswordDigest: testPassword}
	r.db.Create(&user)
}

func clearDb(){
	r := newTestRepository()
	r.db.DropTableIfExists("users")
}

func TestMain(m *testing.M) {
	fmt.Println("Before testing")
	setup()
	code := m.Run()
	println("After Testing")
	clearDb()
	os.Exit(code)
}

func TestCreateUserSuccess(t *testing.T) {
	r := newTestRepository()
	var users []entity.User
	beforeCount := 0
	afterCount := 0
	r.db.Find(&users).Count(&beforeCount)
	println(beforeCount)
	r.CreateUser("createTest","createtest@gmail.com","createTEstPassword")
	r.db.Find(&users).Count(&afterCount)
	println(afterCount)
	if (beforeCount+1) != afterCount {
		t.Fail()
	}
}