package models

import (
	"errors"
	"fmt"
	// "time"

	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/forms"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        int    `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//UserModel ...
type UserModel struct{}

//Signin ...
func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {

	rows, err  := db.GetDB().QueryOne("SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER(" + form.Email +") LIMIT 1")
	if err != nil {
		return user, err
	}
	fmt.Printf("query returned %d rows\n",rows.NumRows())

	err = rows.Scan(user.Password)

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, errors.New("Invalid password")
	}

	return user, nil
}

//Signup ...
func (m UserModel) Signup(form forms.SignupForm) (user User, err error) {
	getDb := db.GetDB()

	checkUser, err := getDb.QueryOne("SELECT count(id) FROM public.user WHERE email=LOWER(" +form.Email + ") LIMIT 1")
	fmt.Println("checkUser:", checkUser, err)

	if err != nil {
		return user, err
	}

	if checkUser.NumRows() > 0 {
		return user, errors.New("User exists")
	}
    
	// bytePassword := []byte(form.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := getDb.WriteOne("INSERT INTO public.user(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())
	// fmt.Println(" getDb.Exec:", res, err)

	// if res != nil && err == nil {
	// 	err = getDb.QueryOne("SELECT id, email, name, updated_at, created_at FROM public.user WHERE email=LOWER(" + form.Email + ") LIMIT 1")
	// 	if err == nil {
	// 		return user, nil
	// 	}
	// }

	return user, errors.New("Not registered")
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	// err = db.GetDB().QueryOne("SELECT id, email, name FROM public.user WHERE id=" + userID)
	return user, err
}
