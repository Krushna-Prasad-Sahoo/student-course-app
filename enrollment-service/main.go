package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)


var db *sql.DB

type Enrollment struct {
    ID        int    `json:"id"`
    StudentID string `json:"student_id"`
    CourseID  int    `json:"course_id"`
}

func main() {
    // MySQL env config
    dbUser := os.Getenv("MYSQL_USER")
    dbPass := os.Getenv("MYSQL_PASSWORD")
    dbName := os.Getenv("MYSQL_DATABASE")
    dbHost := os.Getenv("MYSQL_HOST")
    
    // reads MySQL credentials from env
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

    // opens DB connections
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("DB connection error:", err)
    }

    // Retry ping with backoff
    maxRetries := 10
    for i := 0; i < maxRetries; i++ {
        err = db.Ping()
        if err == nil {
            break
        }
        log.Printf("Waiting for MySQL to be ready... attempt %d/%d\n", i+1, maxRetries)
        time.Sleep(2 * time.Second)
    }
    if err != nil {
        log.Fatal("DB ping error after retries:", err)
    }

    // ensures the table schema is created with required details  
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS enrollments (
        id INT AUTO_INCREMENT PRIMARY KEY,
        student_id VARCHAR(255),
        course_id INT
    )`)
    if err != nil {
        log.Fatal("Failed to create table:", err)
    }

    router := gin.Default()

    // health check
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Enrollment Service is running")
    })
    

    // Fetch enrollments for a specific student 
    router.GET("/enrollments/student/:student_id", getEnrollmentsByStudent)

    // enroll a new student with a course
    router.POST("/enroll", handleEnroll)

    // get all enrollments
    router.GET("/enrollments", getAllEnrollments)

    router.Run(":8000")
}

// function definition for enrolling a new student
func handleEnroll(c *gin.Context) {
    var e Enrollment
    if err := c.ShouldBindJSON(&e); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }

    // Validate student
    studentURL := fmt.Sprintf("http://student-service:3000/students/%s", e.StudentID)
    if !validateService(studentURL) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
        return
    }

    // Validate course
    courseURL := fmt.Sprintf("http://course-service:5000/courses/%d", e.CourseID)
    if !validateService(courseURL) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    // Store enrollment
    res, err := db.Exec("INSERT INTO enrollments (student_id, course_id) VALUES (?, ?)", e.StudentID, e.CourseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll", "details": err.Error()})
        return
    }

    id, _ := res.LastInsertId()
    e.ID = int(id)
    c.JSON(http.StatusCreated, e)
}


// Handler: fetch all enrollments
func getAllEnrollments(c *gin.Context) {
    rows, err := db.Query("SELECT id, student_id, course_id FROM enrollments")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments", "details": err.Error()})
        return
    }
    defer rows.Close()

    var enrollments []Enrollment
    for rows.Next() {
        var e Enrollment
        if err := rows.Scan(&e.ID, &e.StudentID, &e.CourseID); err == nil {
            enrollments = append(enrollments, e)
        }
    }

    c.JSON(http.StatusOK, enrollments)
}


// function definition for fetching enrollments by student id
func getEnrollmentsByStudent(c *gin.Context) {
    studentID := c.Param("student_id")
    rows, err := db.Query("SELECT id, student_id, course_id FROM enrollments WHERE student_id = ?", studentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
        return
    }
    defer rows.Close()

    var enrollments []Enrollment
    for rows.Next() {
        var e Enrollment
        if err := rows.Scan(&e.ID, &e.StudentID, &e.CourseID); err == nil {
            enrollments = append(enrollments, e)
        }
    }
    c.JSON(http.StatusOK, enrollments)
}


// validate if a service URL returns 200 and JSON 
func validateService(url string) bool {
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != 200 {
        return false
    }

    body, _ := ioutil.ReadAll(resp.Body)
    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        return false
    }

    return true
}

