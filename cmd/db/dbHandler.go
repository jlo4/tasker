package dbHandler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Task struct {
	ID            int64
	CompletedDate int64
	Created       int64
	Description   string
	DueDate       int64
	Title         string
}

func Connect() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = sql.Open("sqlite3", "tasker.db")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DropTaskTable() {
	query := `DROP TABLE IF EXISTS task`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error dropping task table:", err.Error())
	}

	fmt.Println("Task table drpped")
}

func CreateTaskTable() {
	query := `CREATE TABLE task (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, ` +
		`created INTEGER NOT NULL DEFAULT (strftime('%s', 'now')), due_date INTEGER DEFAULT NULL, completed_date INTEGER DEFAULT NULL, title TEXT NOT NULL, description TEXT NOT NULL)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating table:", err.Error())
	}

	fmt.Println("Database connected and tabled created")
}

func InsertTask(task Task) (int64, error) {
	result, err := db.Exec("INSERT INTO task (title, description) VALUES (?, ?)", task.Title, task.Description)
	if err != nil {
		return 0, fmt.Errorf("insertTask: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertTask: %v", err)
	}
	return id, nil
}

func TaskByID(id int64, orderBy string) (Task, error) {
	var task Task
	if orderBy == "" {
		orderBy = "due_date"
	}
	fmt.Printf("orderby - %s", orderBy)
	row := db.QueryRow("SELECT * FROM task WHERE id = ? ORDER BY ?", id, orderBy)
	if err := row.Scan(&task.ID, &task.CompletedDate, &task.Created, &task.Description, &task.DueDate, &task.Title); err != nil {
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("task with ID [%d] could not be found", id)
		}
		return task, fmt.Errorf("calling TaskById %d: %v", id, err)
	}
	return task, nil
}
