package main

import (
	"encoding/xml"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var SECRET = []byte("secret")

type User struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

func login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Username == "admin" && u.Password == "admin123!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "John Snow"
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response
		t, err := token.SignedString(SECRET)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{
		"alive": true,
	})
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	role := claims["role"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"-["+role+"]!")
}

type Burp struct {
	XMLName     xml.Name `xml:"issues"`
	ExportTime  string   `json:"exportTime" xml:"exportTime,attr"`
	BurpVersion string   `json:"burpVersion" xml:"burpVersion,attr"`
	Issue       []struct {
		//XMLName         xml.Name `xml:"issue"`
		SerialNumber     int    `json:"serialNumber" xml:"serialNumber"`
		Type             int    `json:"type" xml:"type"`
		Name             string `json:"name" xml:"name"`
		Host             string `json:"host" xml:"host"`
		Path             string `json:"path" xml:"path"`
		Location         string `json:"location" xml:"location"`
		Severity         string `json:"severity" xml:"severity"`
		Confidence       string `json:"confidence" xml:"confidence"`
		Request          string `json:"request" xml:"requestresponse>request"`
		Response         string `json:"response" xml:"requestresponse>response"`
		ResponseRedirect bool   `json:"responseRedirect" xml:"request>responseRedirect"`
	} `json:"issue" xml:"issue"`
}

func burpParse(c echo.Context) error {
	p := new(Burp)
	if err := c.Bind(p); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)
	e.GET("/health-check", HealthCheckHandler)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT(SECRET))
	r.GET("", restricted)
	r.POST("", burpParse)

	e.Logger.Fatal(e.Start(":8086"))
}
