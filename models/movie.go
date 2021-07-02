package models

type Movie struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Year int16 `json:"year"`
    Rating float32 `json:"rating"`
}