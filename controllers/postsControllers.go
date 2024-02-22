package controllers

import (
	"github.com/gabriel-leone/go-jwt/initializers"
	"github.com/gabriel-leone/go-jwt/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create post!",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch posts!",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch post!",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	c.Bind(&body)

	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Failed to fetch post!",
		})
		return
	}

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update post!",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Failed to fetch post!",
		})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete post!",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted!",
	})
}