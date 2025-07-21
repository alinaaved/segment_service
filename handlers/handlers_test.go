package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"segment_service/db"
	"segment_service/models"
	"segment_service/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Вспомогательная функция: создаёт in-memory SQLite базу и мигрирует схемы
func setupTestDB() {
	database, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.AutoMigrate(&models.User{}, &models.Segment{})
	db.DB = database
}

// Вспомогательная функция: создаёт тестовый маршрутизатор с подключёнными хендлерами
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

// Тест: создание сегмента через POST-запрос
func TestCreateSegment(t *testing.T) {
	setupTestDB()
	r := setupRouter()

	body := map[string]string{"name": "TEST_SEGMENT"}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/segments", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "TEST_SEGMENT")
}

// Тест: назначение сегмента конкретному пользователю
func TestAssignSegmentToUser(t *testing.T) {
	setupTestDB()
	r := setupRouter()
	database := db.GetDB()

	// Создаём пользователя и сегмент
	database.Create(&models.User{ID: 1})
	database.Create(&models.Segment{Name: "MAIL_TEST"})

	body := map[string]interface{}{
		"user_ids": []uint{1},
	}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/segments/MAIL_TEST/assign", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Тест: получение списка сегментов у пользователя
func TestGetUserSegments(t *testing.T) {
	setupTestDB()
	r := setupRouter()
	database := db.GetDB()

	// Создаём сегмент и пользователя с привязкой
	segment := models.Segment{Name: "S1"}
	database.Create(&segment)
	user := models.User{ID: 99, Segments: []models.Segment{segment}}
	database.Create(&user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/99/segments", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "S1")
}
