package page

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type HandlerService struct{}

func (hs *HandlerService) Bootstrap(r *gin.Engine) {

	r.POST("/transaction", hs.CreateTransaction)
	r.GET("/transaction", hs.GetTransaction)
	r.DELETE("/transaction", hs.DeleteAllTranscation)
	r.POST("/location", hs.CreateLocation)
	r.POST("/location/:id", hs.UpdateLocation)

}

func (hs *HandlerService) CreateTransaction(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var req TransactionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	now := time.Now()
	res := now.Sub(req.CreatedAt)
	final := int(res.Seconds())

	if final > 300 {
		c.JSON(http.StatusNoContent, gin.H{"message": "time expired"})
		return
	}
	var transaction Transaction
	transaction.Amount = req.Amount
	transaction.CreatedAt = req.CreatedAt
	transaction.UserId = req.UserId
	if err := h.Table("transaction").Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction created successfully"})
}

func (hs *HandlerService) GetTransaction(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var transaction TransactionDetails
    where :="is_deleted=false and EXTRACT(EPOCH FROM (now() - created_at))<60"
	if err := h.Table("transaction").Select("sum(amount) as sum,max(amount) as max,min(amount) as min,avg(amount) as avg,count(amount) as count").Where(where).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, transaction)
}

// soft deleting the transaction
func (hs *HandlerService) DeleteAllTranscation(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var transaction Transaction
	transaction.IsDeleted = true
	if err := h.Table("transaction").Update(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "user location update successfully"})
}


func (hs *HandlerService) CreateLocation(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	var location LocationReq
	fmt.Println("location",location)
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    fmt.Println("location",location)
	var user UserDetails
	user.City = location.City
	if err := h.Table("user_details").Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}



func (hs *HandlerService) UpdateLocation(c *gin.Context) {
	h := c.MustGet("DB").(*gorm.DB)
	id := c.Param("id")
	var location LocationReq
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(time.Now().UTC())
	var user UserDetails
	user.City = location.City
	if err := h.Table("user_details").Where("id=?", id).Update(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}



