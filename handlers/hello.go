package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) HelloWorld(c *gin.Context) {
	// Implement your logic to retrieve all tables from the database
	// For example:
	rows, err := h.db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tables = append(tables, table)
	}

	// Return the tables as JSON
	c.JSON(http.StatusOK, tables)
}
