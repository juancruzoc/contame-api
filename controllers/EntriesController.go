package controllers

import (
	"example/contame/initializers"
	"example/contame/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler to create a new entry
func CreateEntry(c *gin.Context) {
	var entry models.Entry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userID = c.GetUint("userid")
	entry.UserID = userID

	// Save the entry to the database
	if err := initializers.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

// Handler to update an existing entry
func UpdateEntry(c *gin.Context) {
	var entry models.Entry
	id := c.Param("id")

	if err := initializers.DB.First(&entry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	var userID = c.GetUint("userid")
	if entry.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the entry in the database
	if err := initializers.DB.Save(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

// Handler to delete an existing entry
func DeleteEntry(c *gin.Context) {
	var entry models.Entry
	id := c.Param("id")

	if err := initializers.DB.First(&entry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	var userID = c.GetUint("userid")
	if entry.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	// Delete the entry from the database
	if err := initializers.DB.Delete(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Handler to get all entries
func GetAllEntries(c *gin.Context) {
	var entries []models.Entry
	var userID = c.GetUint("userid")

	if err := initializers.DB.Where("user_id = ?", userID).Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}
