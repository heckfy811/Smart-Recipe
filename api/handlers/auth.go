package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"smart-recipe/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type FullUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Регистрация
func SignUp(c *gin.Context) {
	var fullUser FullUser

	//Создаем JSON из данных и проверяем все ли на месте
	ok := c.BindJSON(&fullUser)
	if ok != nil {
		c.JSON(400, ok.Error())
		fmt.Println("SignUp:", ok)
		return
	}

	email := fullUser.Email
	password := fullUser.Password

	//Шифруем паролич
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		fmt.Println("SignUp:", err)
		c.JSON(500, gin.H{"message": "could not register user"})
		return
	}

	user := database.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	fmt.Println(user)
	//Заносим в бд
	err = database.Database.Users.Add(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "could not register user"})
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	var user database.User
	err := database.Database.Users.Db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", data["email"]).
		Scan(&user.Id, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		fmt.Println(err)
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "could not retrieve user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "incorrect password"})
		return
	}

	token, err := GenerateToken(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.JSON(500, gin.H{"message": "could not login"})
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 730),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(200, cookie)
}
