package beatport_test

import (
	"fmt"
	"github.com/jeanmorais/beatport/domain"
	"github.com/jeanmorais/beatport/domain/beatport"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedGenres = []domain.Genre{{"Afro House", "afro-house-89", "https://www.beatport.com/genre/afro-house/89"}}

var expectedTracks = []domain.Track{
	{1, "Zvezdara", "Original Mix", []string{"Last95"}, "MIR MUSIC", "Progressive House", "https://www.beatport.com/track/zvezdara-original-mix/13418144", "1.29"}}

func TestService_GetGenres(t *testing.T) {
	t.Run("should return data successfully", func(t *testing.T) {
		clientMock := &beatport.ClientMock{
			GetGenresFn: func() ([]domain.Genre, error) {
				return expectedGenres, nil
			},
		}

		beatPortService := beatport.NewService(clientMock)
		genres, err := beatPortService.GetGenres()
		assert.NoError(t, err)
		assert.NotNil(t, genres)
		assert.Equal(t, expectedGenres, genres)
	})

	t.Run("should return error from get genres function", func(t *testing.T) {
		clientMock := &beatport.ClientMock{
			GetGenresFn: func() ([]domain.Genre, error) {
				return nil, fmt.Errorf("unexpected error")
			},
		}

		beatPortService := beatport.NewService(clientMock)
		genres, err := beatPortService.GetGenres()
		assert.Nil(t, genres)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unexpected error")
	})
}

func TestService_GetTop10(t *testing.T) {
	t.Run("should return data successfully", func(t *testing.T) {
		clientMock := &beatport.ClientMock{
			GetTop10Fn: func(genre *domain.Genre) ([]domain.Track, error) {
				return expectedTracks, nil
			},
		}

		beatPortService := beatport.NewService(clientMock)
		tracks, err := beatPortService.GetTop10("progressive-house-15")
		assert.NoError(t, err)
		assert.NotNil(t, tracks)
		assert.Equal(t, expectedTracks, tracks)
	})

	t.Run("should return invalid argument error from get top10 function", func(t *testing.T) {
		beatPortService := beatport.NewService(nil)
		tracks, err := beatPortService.GetTop10("progressive?15")
		assert.Nil(t, tracks)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Invalid argument: genreKey [progressive?15]")
	})

	t.Run("should return unexpected error from get top10 function", func(t *testing.T) {
		clientMock := &beatport.ClientMock{
			GetTop10Fn: func(genre *domain.Genre) ([]domain.Track, error) {
				return nil, fmt.Errorf("unexpected error")
			},
		}

		beatPortService := beatport.NewService(clientMock)
		tracks, err := beatPortService.GetTop10("progressive-house-15")
		assert.Nil(t, tracks)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unexpected error")
	})
}
