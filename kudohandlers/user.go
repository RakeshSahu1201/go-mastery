package kudohandlers

import (
	"main/ent"
	"main/kudomodels"
	"main/kudostore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers handles GET /users — returns all users from the DB.
func GetUsers(c *gin.Context) {
	err := kudostore.WithTx(c.Request.Context(), func(client *ent.Client) error {
		users, err := kudomodels.GetAllUsers(c.Request.Context(), client)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, users)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
	}
}

// CreateUser handles POST /users — creates a new user.
func CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name"  binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := kudostore.WithTx(c.Request.Context(), func(client *ent.Client) error {
		user, err := kudomodels.CreateUser(c.Request.Context(), client, req.Name, req.Email)
		if err != nil {
			if ent.IsConstraintError(err) {
				c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
				return nil // Don't return error to WithTx so it doesn't log a 500
			}
			return err
		}
		c.JSON(http.StatusCreated, user)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
	}
}

// GetUserByID handles GET /users/:id — returns a single user by ID.
func GetUserByID(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err := kudostore.WithTx(c.Request.Context(), func(client *ent.Client) error {
		user, err := kudomodels.GetUserByID(c.Request.Context(), client, uri.ID)
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return nil
			}
			return err
		}
		c.JSON(http.StatusOK, user)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
	}
}
