package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/go-jet/jet/v2/mysql"
	_ "github.com/go-sql-driver/mysql"
	. "gojet/gen/membership/table"

	"gojet/gen/membership/model"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbName   = "membership"
)

func main() {
	r := gin.Default()

	// Connect to database
	var connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		stmt := SELECT(
			Users.AllColumns,
			Roles.Name.AS("role"),
			Groups.Name.AS("groups.name"),
		)

		stmt.
			FROM(Users.
				INNER_JOIN(Roles, Users.RoleID.EQ(Roles.ID)).
				LEFT_JOIN(UserGroups, Users.ID.EQ(UserGroups.UserID)).
				LEFT_JOIN(Groups, UserGroups.GroupID.EQ(Groups.ID)),
			)

		type groups struct {
			Name *string
		}

		var users []struct {
			model.Users
			Role   *string
			Groups []groups
		}

		err := stmt.Query(db, &users)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"users": users,
		})

	})

	r.GET("/users-raw", func(c *gin.Context) {
		stmt := "SELECT * FROM users"

		var users []model.Users

		rows, err := db.Query(stmt)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var user model.Users
			err = rows.Scan(
				&user.ID,
				&user.CreatedAt,
				&user.UpdatedAt,
				&user.DeletedAt,
				&user.Name,
				&user.Username,
				&user.Email,
				&user.Password,
				&user.IsSuperAdmin,
				&user.ActivatedAt,
				&user.LanguageID,
				&user.OrganizationID,
				&user.RoleID,
				&user.StatusID)
			if err != nil {
				panic(err)
			}
			users = append(users, user)
		}

		c.JSON(200, gin.H{
			"users": users,
		})
	})

	type users struct {
		ID        string
		Name      *string
		CreatedAt *time.Time
	}

	r.GET("/users-raw-2", func(c *gin.Context) {
		stmt := RawStatement(`SELECT users.id AS "users.id" , users.created_at AS "users.created_at", users.name AS "users.name" FROM users`)

		var userList []users

		err := stmt.Query(db, &userList)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"users": userList,
		})

	})

	r.Run()
}
