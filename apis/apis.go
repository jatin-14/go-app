package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jatin-14/go-app/models"
	uuid "github.com/satori/go.uuid"
)

type ErrorMessage struct {
	message string
}

func GetMovies (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	db := DBConnection()
	rows, err := db.Query("SELECT * FROM movies ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
	}
	defer rows.Close()
	defer db.Close()
	
	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Rating)
        if err != nil {
            panic(err.Error())
		}
		movies = append(movies, movie)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func GetMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	movieId := mux.Vars(r)["id"]
	db := DBConnection()
	row := db.QueryRow("SELECT * FROM movies WHERE id = ?", movieId)

	defer db.Close()
	
	var movie models.Movie
	err := row.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Rating)
    if err != nil {
		//Respond when the movie with this Id is not found
		if err == sql.ErrNoRows{
			// message := ErrorMessage{message: "Record not found."}
			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(message)

			http.Error(w, "Not found", 200)
			return
		}else {
			panic(err.Error())
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)

}

func CreateMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	movie.Id = uuid.NewV4().String()

	db := DBConnection()
	defer db.Close()
	_, err := db.Exec("INSERT INTO movies(id, name, year, rating) VALUES( ?, ?, ?, ? )", 
		movie.Id, movie.Name, movie.Year, movie.Rating,
	)
    if err != nil {
        panic(err.Error())
	}
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	movieId := mux.Vars(r)["id"]

	db := DBConnection()
	defer db.Close()
	result , err := db.Exec("DELETE FROM movies WHERE id = ?", movieId)
    if err != nil {
        panic(err.Error())
	}
	rows, err := result.RowsAffected()
    if err != nil {
        panic(err.Error())
	} else {
		log.Println(rows)
		if rows == int64(1) {
			// w.WriteHeader(http.StatusOK)
			http.Error(w, "1 Record deleted", 200)
		}else {
			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(message)
			http.Error(w, "Record not found.", 200)
		}
	}
}

func UpdateMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	movieId := mux.Vars(r)["id"]
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	log.Println(movieId)
	db := DBConnection()
	result, err := db.Exec("UPDATE movies SET name = ?, year = ?, rating = ? WHERE id = ?", movie.Name, movie.Year, movie.Rating, movieId)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()
	rows, err := result.RowsAffected()
    if err != nil {
        panic(err.Error())
	} else {
		log.Println(rows)
		if rows == int64(1) {
			w.WriteHeader(http.StatusOK)
			http.Error(w, "1 Record updated", 200)
		}else {
			w.WriteHeader(http.StatusOK)
			http.Error(w, "Record not found.", 200)
		}
	}
}