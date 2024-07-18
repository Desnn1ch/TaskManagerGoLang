package controllers

import (
	"TaskManagerGoLang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type TaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTask(c *gin.Context, db *sqlx.DB) {
	var taskReq TaskRequest
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	query := `INSERT INTO tasks (title, status, user_id) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, taskReq.Title, taskReq.Status, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func GetTasks(c *gin.Context, db *sqlx.DB) {
	userID, _ := c.Get("userID")

	var tasks []models.Task
	query := `SELECT id, title, status FROM tasks WHERE user_id = $1`
	err := db.Select(&tasks, query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context, db *sqlx.DB) {
	var taskReq TaskRequest
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	userID, _ := c.Get("userID")

	query := `UPDATE tasks SET title = $1, status = $2 WHERE id = $3 AND user_id = $4`
	_, err = db.Exec(query, taskReq.Title, taskReq.Status, taskID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context, db *sqlx.DB) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	userID, _ := c.Get("userID")

	query := `DELETE FROM tasks WHERE id = $1 AND user_id = $2`
	_, err = db.Exec(query, taskID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
