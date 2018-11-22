package repo

import (
	"fmt"
	"github.com/JYama321/gin-crud-template/entity"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	r := NewRepository()
	testPassword := generatePasswordDigest("testPassword")
	user := entity.User{UserName: "test1", Email: "test1@gmail.com", PasswordDigest: testPassword}
	r.db.Create(&user)
	user = entity.User{UserName: "test2", Email: "test2@gmail.com", PasswordDigest: testPassword}
	r.db.Create(&user)

	code := m.Run()

	println("After Testing")

	os.Exit(code)
}

func TestCreateUserSuccess(t *testing.T) {
	println("Before Testing")
	r := NewRepository()
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