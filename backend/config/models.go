package config

import (
	"time"

	"gorm.io/gorm"
)

type Response[T any] struct {
	StatusCode int    `json:"statusCode"`
	Data       T      `json:"data,omitempty"`
	Message    string `json:"message"`
	HasData    bool   `json:"hasData"`
}

type QueryHistory struct {
	gorm.Model
	ID           uint `gorm:"primary_key"`
	SearchTerm   string
	SearchResult string `gorm:"not null:false"`
	SearchTime   time.Time
}

func (history *QueryHistory) BeforeCreate(tx *gorm.DB) (err error) {
	var all_queries []QueryHistory
	err = DB.Find(&all_queries).Error
	if err == nil {
		if len(all_queries) >= 5 {
			for index, record := range all_queries {
				if index <= len(all_queries)-5 {
					DB.Unscoped().Delete(&record)
				}
			}
		}
	}
	return
}

// save new history as the latest of 5 maximum search histories
func (query *QueryHistory) SaveNew() error {
	return DB.Create(&query).Error

}

type MovieData struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []map[string]string
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}
