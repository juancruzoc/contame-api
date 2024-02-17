package controllers

import (
	"example/contame/initializers"
	"example/contame/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler to create a new category
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userID = c.GetUint("userid")
	category.UserID = userID

	// Save the category to the database
	if err := initializers.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// Handler to update an existing category
func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := initializers.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var userID = c.GetUint("userid")
	if category.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the category in the database
	if err := initializers.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, category)
}

// Handler to delete an existing category
func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := initializers.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var userID = c.GetUint("userid")
	if category.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Delete the category from the database
	if err := initializers.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
