package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"net/url"

	// "os"
	 "fmt"
	"github.com/Mac-5/controllers"
	"github.com/Mac-5/routes"
	"github.com/Mac-5/services"
	"github.com/gin-gonic/gin"

	github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func createDBConnection()(*sql.DB, error){
	
	connectionString := "postgres://admin:HLf8mi6V44obIvk5lHzuE7ay6QjT23fb@dpg-ck0oohb6fquc738ue2cg-a.frankfurt-postgres.render.com/admin_pn1t"

    // Parse the URL
    u, err := url.Parse(connectionString)
    if err != nil {
        panic("Failed to parse connection string: " + err.Error())
    }

    // Extract components from the URL
    dbUser := u.User.Username()
    dbPassword, _ := u.User.Password()
    dbHost := u.Host
    dbName := u.Path[1:] // Remove the leading slash

    // Build the connection string
    dbConnection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=require",
        dbUser, dbPassword, dbName, dbHost)

	db, err := sql.Open("postgres",dbConnection)
	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := createDBConnection()
	if  err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULl,
		email VARCHAR(150)
	);
`

// Execute the SQL statement to create the table
_, err = db.Exec(createTableSQL)
if err != nil {
	panic("Failed to create table: " + err.Error())
}


	ctx := context.TODO()
	router := gin.Default()
	personController := controllers.NewPersonController(services.NewPersonSeviceImpl(db, ctx))
	
	router.GET("/api/status",  func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	routes.SetupRoutes(router, personController)

	router.Run(":8080")

}