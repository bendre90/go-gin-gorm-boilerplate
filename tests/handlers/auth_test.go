package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/daystram/go-gin-gorm-boilerplate/constants"
	"github.com/daystram/go-gin-gorm-boilerplate/utils"
)

func TestAuthOnly_Unauthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(constants.IsAuthenticatedKey, false)

	utils.AuthOnly(c)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestAuthOnly_Authenticated(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(constants.IsAuthenticatedKey, true)

	utils.AuthOnly(c)

	if c.IsAborted() {
		t.Errorf("middleware aborted despite authenticated context")
	}
}
