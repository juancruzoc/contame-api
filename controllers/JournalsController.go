package controllers

import (
	"example/contame/initializers"
	"example/contame/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler to create a new entry
func CreateJournal(c *gin.Context) {
	var journal models.Journal
	if err := c.ShouldBindJSON(&journal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userID = c.GetUint("userid")
	journal.UserID = userID

	// Save the entry to the database
	if err := initializers.DB.Create(&journal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, journal)
}

// Handler to update an existing entry
func UpdateJournal(c *gin.Context) {
	var journal models.Journal
	id := c.Param("id")

	if err := initializers.DB.First(&journal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	var userID = c.GetUint("userid")
	if journal.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	if err := c.ShouldBindJSON(&journal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the entry in the database
	if err := initializers.DB.Save(&journal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, journal)
}

// Handler to delete an existing entry
func DeleteJournal(c *gin.Context) {
	var journal models.Journal
	id := c.Param("id")

	if err := initializers.DB.First(&journal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	var userID = c.GetUint("userid")
	if journal.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	// Delete the entry from the database
	if err := initializers.DB.Delete(&journal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Handler to get all entries
func GetAllJournals(c *gin.Context) {
	var entries []models.Journal
	var userID = c.GetUint("userid")

	if err := initializers.DB.Where("user_id = ?", userID).Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}
