package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"vidannotate/constant"
	"vidannotate/model"
	"vidannotate/sqlite"

	"github.com/gorilla/mux"
)

var db *sqlite.DB

// Create a new video
func createVideo(w http.ResponseWriter, r *http.Request) {
	var video model.Video
	json.NewDecoder(r.Body).Decode(&video)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	_, err := db.Create(constant.InsertIntoVideosTable,
		video.Title, video.URL, video.Duration, video.CreatedAt, video.UpdatedAt)
	if err != nil {
		log.Fatalf("error inserting video: %v", err)
	}

	json.NewEncoder(w).Encode(video)
}

// Create a new annotation
func createAnnotation(w http.ResponseWriter, r *http.Request) {
	var annotation model.Annotation
	json.NewDecoder(r.Body).Decode(&annotation)

	// Check if annotation start and end time are within video duration
	video := getVideo(annotation.VideoID)
	if annotation.StartTime < 0 || annotation.EndTime > video.Duration {
		json.NewEncoder(w).Encode(Error{Message: "Annotation time out of bounds"})
		return
	}

	annotation.CreatedAt = time.Now()
	annotation.UpdatedAt = time.Now()

	res, err := db.Create("INSERT INTO annotations (video_id, start_time, end_time, type, notes, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		annotation.VideoID, annotation.StartTime, annotation.EndTime, annotation.Type, annotation.Notes, annotation.CreatedAt.Format("2006-01-02 15:04:05"), annotation.UpdatedAt.Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatalf("error inserting annotation: %v", err)
	}
	json.NewEncoder(w).Encode(res)
}

// Get all annotations for a specific video
func getAnnotations(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: log error
	}
	rows, err := db.Query("SELECT * FROM annotations WHERE video_id = ?", videoID)
	if err != nil {
		// TODO: log error
	}
	defer rows.Close()

	annotations := []model.Annotation{}
	for rows.Next() {
		var a model.Annotation
		var createdAt, updatedAt string
		if err := rows.Scan(&a.ID, &a.VideoID, &a.StartTime, &a.EndTime, &a.Type, &a.Notes, &createdAt, &updatedAt); err != nil {
			// TODO: log error
		}
		a.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			// TODO: log error
		}
		a.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
		if err != nil {
			// TODO: log error
		}
		annotations = append(annotations, a)
	}
	json.NewEncoder(w).Encode(annotations)
}
func updateAnnotation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	intID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: log error
	}

	var updatedAnnotation model.Annotation
	json.NewDecoder(r.Body).Decode(&updatedAnnotation)

	// Check if updated annotation start and end time are within video duration
	video := getVideo(updatedAnnotation.VideoID)
	if updatedAnnotation.StartTime < 0 || updatedAnnotation.EndTime > video.Duration {
		json.NewEncoder(w).Encode(Error{Message: "Annotation time out of bounds"})
		return
	}

	updatedAnnotation.ID = intID
	updatedAnnotation.CreatedAt = time.Now()
	updatedAnnotation.UpdatedAt = time.Now()

	// Update the annotation in the database
	result, err := db.Exec("UPDATE annotations SET video_id = ?, start_time = ?, end_time = ?, type = ?, notes = ?, updated_at = ? WHERE id = ?",
		updatedAnnotation.VideoID, updatedAnnotation.StartTime, updatedAnnotation.EndTime, updatedAnnotation.Type, updatedAnnotation.Notes, updatedAnnotation.UpdatedAt, intID)
	if err != nil {
		// TODO: log error
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// TODO: log error
	}

	if rowsAffected == 0 {
		json.NewEncoder(w).Encode(Error{Message: "Annotation not found"})
		return
	}

	json.NewEncoder(w).Encode(updatedAnnotation)
}

func deleteAnnotation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	intID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: log error
	}

	stmt, err := db.Prepare("DELETE FROM annotations WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(intID)
	if err != nil {
		json.NewEncoder(w).Encode(Error{Message: "Annotation not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "Annotation deleted successfully"})
}

// Delete a video and all related annotations
func deleteVideo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	intID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: log error
	}

	// Delete all related annotations
	_, err = db.Exec("DELETE FROM annotations WHERE video_id = $1", intID)
	if err != nil {
		log.Fatal(err)
	}

	// Delete video
	_, err = db.Exec("DELETE FROM videos WHERE id = $1", intID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "Video and related annotations deleted"})
}
// Get video by ID
func getVideo(id int) model.Video {
	var video model.Video
	row := db.QueryRow("SELECT id, title, description, duration, created_at, updated_at FROM videos WHERE id = ?", id)
	err := row.Scan(&video.ID, &video.Title, &video.URL, &video.Duration, &video.CreatedAt, &video.UpdatedAt)
	if err != nil {
		// TODO: log error
	}
	return video
}


// Main function
func main() {
	// Initialise connection
	var err error
	db, err = sqlite.New("file:production.db")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Initialise tables
	// Table 1: Videos
	_, err = db.Create(constant.CreateVideosTable)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	_, err = db.Create(constant.CreateAnnotationTable)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	// Initialize router
	router := mux.NewRouter()
	// Add routes
	router.HandleFunc("/videos", createVideo).Methods("POST")
	router.HandleFunc("/videos/{id}", deleteVideo).Methods("DELETE")
	router.HandleFunc("/annotations", createAnnotation).Methods("POST")
	router.HandleFunc("/annotations/{id}", updateAnnotation).Methods("PUT")
	router.HandleFunc("/annotations/{id}", deleteAnnotation).Methods("DELETE")
	router.HandleFunc("/videos/{id}/annotations", getAnnotations).Methods("GET")
	// Add middleware for API key or JWT token validation
	router.Use(validateAPIKeyMiddleware)
	router.Use(validateJWTMiddleware)

	// Start server
	http.ListenAndServe(":8000", router)
}

// Error struct
type Error struct {
	Message string `json:"message"`
}
