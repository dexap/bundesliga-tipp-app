package main

import (
	"context"
	"html/template"

	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var redisClient *redis.Client
var db *gorm.DB
var templates *template.Template

func init() {
	// Templates
	var err error
	templates, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// Redis-Verbindung
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// MariaDB-Verbindung
	db, err = gorm.Open(mysql.Open("root:rootpassword@tcp(localhost:3306)/tippapp?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migriere das Schema
	db.AutoMigrate(&User{})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Printf("HTTP Error: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	testConnections()

	e.GET("/", func(c echo.Context) error {
		log.Println("Handling root request")

		cookie, err := c.Cookie("session_token")
		if err != nil {
			log.Printf("No session token found: %v", err)
		}

		data := map[string]interface{}{"Content": "login"}

		if err == nil {
			// Überprüfe Token
			_, err = redisClient.Get(c.Request().Context(), cookie.Value).Result()
			if err != nil {
				log.Printf("Failed to get token from Redis: %v", err)
			} else {
				data["Content"] = "home"
			}
		}

		log.Printf("Attempting to execute template with data: %+v", data)

		err = templates.ExecuteTemplate(c.Response().Writer, "layout.html", data)
		if err != nil {
			log.Printf("Template execution error: %v", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		log.Println("Request handled successfully")
		return nil
	})

	e.POST("/login", handleLogin)

	e.Logger.Fatal(e.Start(":3000"))
}

func handleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generiere JWT Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return err
	}

	// Speichere Token in Redis
	err = redisClient.Set(c.Request().Context(), t, username, time.Hour).Err()
	if err != nil {
		return err
	}

	// Setze Cookie
	cookie := new(http.Cookie)
	cookie.Name = "session_token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour)
	c.SetCookie(cookie)

	return templates.ExecuteTemplate(c.Response().Writer, "layout.html", map[string]interface{}{"Content": "home"})
}

func testConnections() {
	// Test MariaDB connection
	var result int
	err := db.Raw("SELECT 1").Scan(&result).Error
	if err != nil {
		log.Fatalf("Failed to connect to MariaDB: %v", err)
	}
	log.Println("Successfully connected to MariaDB")

	// Test Redis connection
	ctx := context.Background()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis")
}
