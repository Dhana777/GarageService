package page

import (
	"go_test/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type HandlerService struct{}

func (hs *HandlerService) Bootstrap(r *gin.Engine) {


	r.POST("/signup", hs.SignUp)
	r.GET("/login", hs.Login)
	r.POST("/car", hs.CreateCarDetails)
	r.GET("/car", hs.GetCarDetails)
	r.GET("/car/:id", hs.GetCarDetailsById)


}

func (hs *HandlerService) Login(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var user1 User
	var user2 User
	if err := c.ShouldBindJSON(&user1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.Table("user").Where("user_name=? and password=?", user1.UserName, user1.Password).Find(&user2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	token := middleware.GenerateToken(user2.Id)
	var res http.ResponseWriter
	SetSession("setsession", res)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (hs *HandlerService) SignUp(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var user1 UserReq
	var user2 User
	if err := c.ShouldBindJSON(&user1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user2.UserName = user1.Firstname + user1.Lastname
	user2.Phonenumber = user1.Phonenumber
	user2.Password = user1.Password
	if err := h.Table("user").Create(&user2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	}

	c.JSON(http.StatusOK, gin.H{"res": "user created succesfully"})
}

func (hs *HandlerService) GetCarDetails(c *gin.Context) {
	var cardetails []Car
	h := c.MustGet("DB").(*gorm.DB)
	//	var r *http.Request
	// middleware.TokenValid(r)
	if err := h.Table("car").Find(&cardetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, cardetails)

}

func (hs *HandlerService) GetCarDetailsById(c *gin.Context) {
	var cardetails []Car
	h := c.MustGet("DB").(*gorm.DB)
	if err := h.Table("car").Where("id=?", c.Param("id")).Find(&cardetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, cardetails)

}

func (hs *HandlerService) CreateCarDetails(c *gin.Context) {
	var cardetails Car

	if err := c.ShouldBindJSON(&cardetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	h := c.MustGet("DB").(*gorm.DB)
	if err := h.Table("car").Create(&cardetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

}
func (hs *HandlerService) Logout(c *gin.Context) {
	var r http.ResponseWriter
	ClearSession(r)
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetSession(request *http.Request) (yourName string) {
	if cookie, err := request.Cookie("your-name"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("name", cookie.Value, &cookieValue); err == nil {
			yourName = cookieValue["name"]
		}
	}
	return yourName
}

func SetSession(yourName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": yourName,
	}
	if encoded, err := cookieHandler.Encode("your-name", value); err == nil {
		cookie := &http.Cookie{
			Name:   "name",
			Value:  encoded,
			Path:   "/",
			MaxAge: 3600,
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "name",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
