package models

import (
	"errors"
	"fmt"
	"time"

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

	rows, err  := db.GetDB().QueryOne("SELECT id, email, password, name, updated_at, created_at FROM user WHERE email=LOWER(" + form.Email +") LIMIT 1")
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

	checkUser, err := getDb.QueryOne("SELECT count(id) FROM user WHERE email=LOWER('" +form.Email + "') LIMIT 1")
	fmt.Println("checkUser:", checkUser.NumRows())

	if err != nil {
		return user, err
	}

	if checkUser.NumRows() > 1 {
		return user, errors.New("User exists")
	}
    
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// simulate database/sql Prepare()
	
	// pattern := "INSERT INTO secret_agents(id, hero_name, abbrev) VALES (%d, '%s', '%3s')"
	pattern := "INSERT INTO user(email, password, name, updated_at, created_at) VALUES('%s', '%s', '%s', '%s', '%s')"
	
	statements := fmt.Sprintf(pattern,form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())
	res, err := getDb.WriteOne(statements)
	fmt.Println(" getDb.Exec:", res, err)

	if err == nil {
		pattern1 := "SELECT id, email, name, updated_at, created_at FROM user WHERE email=LOWER('%s') LIMIT 1"
		statements1 := fmt.Sprintf(pattern1,form.Email)

		rows ,err := getDb.QueryOne(statements1)
		fmt.Printf("query returned %d rows\n",rows.NumRows())
		if err == nil {
			for rows.Next() {
				err := rows.Scan(&user.ID, &user.Email,&user.Name,&user.UpdatedAt,&user.CreatedAt)
				if err == nil { return user, nil}
				fmt.Printf("this is row number %d\n",rows.RowNumber())
				fmt.Printf("there are %d rows \n",rows.NumRows())
			}
			return user, nil
		}
	}

	return user, errors.New("Not registered")
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	// err = db.GetDB().QueryOne("SELECT id, email, name FROM public.user WHERE id=" + userID)
	return user, err
}
