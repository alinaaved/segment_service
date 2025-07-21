package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"segment_service/db"
	"segment_service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSegment(c *gin.Context) {
	type Req struct {
		Name string `json:"name" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()
	s := models.Segment{Name: req.Name}
	if err := database.Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, s)
}

func DeleteSegment(c *gin.Context) {
	name := c.Param("name")
	database := db.GetDB()
	if err := database.Where("name = ?", name).Delete(&models.Segment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func AssignSegment(c *gin.Context) {
	name := c.Param("name")
	var body struct {
		UserIDs []uint `json:"user_ids"`
		Percent int    `json:"percent"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()
	var segment models.Segment
	if err := database.Where("name = ?", name).First(&segment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "segment not found"})
		return
	}

	if body.Percent > 0 {
		var users []models.User
		database.Find(&users)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
		count := int(float64(len(users)) * float64(body.Percent) / 100.0)
		selected := users[:count]
		for _, u := range selected {
			database.Model(&u).Association("Segments").Append(&segment)
		}
	} else if len(body.UserIDs) > 0 {
		for _, uid := range body.UserIDs {
			var user models.User
			if err := database.FirstOrCreate(&user, uid).Error; err == nil {
				database.Model(&user).Association("Segments").Append(&segment)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user_ids or percent provided"})
		return
	}

	c.Status(http.StatusOK)
}

func GetUserSegments(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	database := db.GetDB()
	var user models.User
	if err := database.Preload("Segments").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	names := []string{}
	for _, s := range user.Segments {
		names = append(names, s.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  user.ID,
		"segments": strings.Join(names, ", "),
	})
}
