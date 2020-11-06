package http_test

import (
	"fmt"
	"github.com/jeanmorais/beatport/domain"
	"github.com/jeanmorais/beatport/domain/beatport"
	internalHttp "github.com/jeanmorais/beatport/infrastructure/server/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	expectedGenres = []domain.Genre{{"Afro House", "afro-house-89", "https://www.beatport.com/genre/afro-house/89"}}
	expectedTracks = []domain.Track{
		{1, "Zvezdara", "Original Mix", []string{"Last95"}, "MIR MUSIC", "Progressive House", "https://www.beatport.com/track/zvezdara-original-mix/13418144", "1.29"}}
)

func TestGenres_Get(t *testing.T) {
	t.Run("should get genres successfully", func(t *testing.T) {

		service := &beatport.ServiceMock{GetGenresFn: func() ([]domain.Genre, error) {
			return expectedGenres, nil
		}}

		server := httptest.NewServer(internalHttp.NewHandler(service))
		defer server.Close()

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/genres", URL), nil)
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		responseBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)

		expectedBody := `
		[
			{
				"name": "Afro House",
				"key": "afro-house-89",
				"url": "https://www.beatport.com/genre/afro-house/89"
			}
		]`

		assert.JSONEq(t, expectedBody, string(responseBody))
	})

	t.Run("should return internal server error", func(t *testing.T) {

		service := &beatport.ServiceMock{GetGenresFn: func() ([]domain.Genre, error) {
			return nil, fmt.Errorf("error")
		}}

		server := httptest.NewServer(internalHttp.NewHandler(service))
		defer server.Close()

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/genres", URL), nil)
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

func TestTracksTop10_Get(t *testing.T) {
	t.Run("should get top10 by genre successfully", func(t *testing.T) {

		service := &beatport.ServiceMock{GetTop10Fn: func(genre string) ([]domain.Track, error) {
			return expectedTracks, nil
		}}

		server := httptest.NewServer(internalHttp.NewHandler(service))
		defer server.Close()

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/tracks/top10/progressive-house-15", URL), nil)
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		responseBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)

		expectedBody := `
			[
			   {
				  "artists": [
					 "Last95"
				  ],
				  "chartNumber":1,
				  "genre":"Progressive House",
				  "label":"MIR MUSIC",
				  "price":"1.29",
				  "remix":"Original Mix",
				  "title":"Zvezdara",
				  "url":"https://www.beatport.com/track/zvezdara-original-mix/13418144"
			   }
			]`

		assert.JSONEq(t, expectedBody, string(responseBody))
	})

	t.Run("should return internal server error", func(t *testing.T) {

		service := &beatport.ServiceMock{GetTop10Fn: func(genre string) ([]domain.Track, error) {
			return nil, fmt.Errorf("error")
		}}

		server := httptest.NewServer(internalHttp.NewHandler(service))
		defer server.Close()

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/tracks/top10/progressive-house-15", URL), nil)
		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
