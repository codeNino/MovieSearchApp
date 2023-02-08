package test

import (
	"PracticalTask/config"
	"PracticalTask/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200
func TestShowIndexPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Welcome to the Gin Backend for my Movie search application !",
		})
	})
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		_, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil

		return statusOK && pageOK
	})
}

func TestMovieHandler(t *testing.T) {
	config.ConnectDB(config.QueryHistory{})
	w := httptest.NewRecorder()
	r := getRouter(true)
	r.GET("/movies/search/:title", controller.MovieSearchHandler)
	searchQuery := "harry"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/movies/search/title=%v", searchQuery), nil)
	r.ServeHTTP(w, req)
	var respBody config.Response[config.MovieData]
	p, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)
	err = json.Unmarshal(p, &respBody)
	require.NoError(t, err)
	if w.Code == http.StatusOK {
		require.Equal(t, true, respBody.HasData)
		require.NotEmpty(t, respBody.Data)
	} else if w.Code == 404 {
		require.Equal(t, false, respBody.HasData)
		require.Empty(t, respBody.Data)
	} else if w.Code == 422 {
		require.Equal(t, "", searchQuery)
		require.Equal(t, false, respBody.HasData)
	} else {
		t.Fail()
	}
}
