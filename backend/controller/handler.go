package controller

import (
	"PracticalTask/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func MovieSearchHandler(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	movieTitle, apiKey := strings.Split(c.Param("title"), "=")[1], config.OMDB_API_KEY

	if movieTitle == "" {
		c.JSON(422, config.Response[config.MovieData]{StatusCode: 422, HasData: false, Message: "query parameter cannot be empty"})
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		SaveLatestQuery(movieTitle)
		wg.Done()
	}()

	url := fmt.Sprintf("http://www.omdbapi.com/?t=%v&apikey=%v", movieTitle, apiKey)
	resp, err := http.Get(url)
	var respBody config.MovieData
	if err != nil {
		log.Println(err)
		c.JSON(500, config.Response[config.MovieData]{StatusCode: 500, HasData: false, Message: "oops! something went wrong"})
		wg.Wait()
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		c.JSON(500, config.Response[config.MovieData]{StatusCode: 500, HasData: false, Message: "oops! something went wrong"})
		wg.Wait()
		return
	}

	json.Unmarshal(body, &respBody)
	if respBody.Response == "False" {
		c.JSON(404, config.Response[config.MovieData]{StatusCode: 404, HasData: false, Message: "No movies found matching keyword"})
		wg.Wait()
		return
	}
	wg.Wait()
	c.JSON(200, config.Response[config.MovieData]{StatusCode: 200, HasData: true, Data: respBody, Message: "Success"})

}

func SaveLatestQuery(query string) {

	newHistory := config.QueryHistory{
		SearchTerm: query,
		SearchTime: time.Now(),
	}

	err := newHistory.SaveNew()

	if err != nil {
		log.Println(err)
	}

}
