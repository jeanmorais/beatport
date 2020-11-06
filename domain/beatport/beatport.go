package beatport

import "github.com/jeanmorais/beatport/domain"

type service struct {
	beatPortClient domain.BeatPortClient
}

// NewService create a new service
func NewService(beatPortClient domain.BeatPortClient) *service {
	return &service{
		beatPortClient: beatPortClient,
	}
}

// GetGenres get all genres
func (s *service) GetGenres() ([]domain.Genre, error) {
	return s.beatPortClient.GetGenres()
}

// GetGenres get the top 10 tracks by a genreKey
func (s *service) GetTop10(genreKey string) ([]domain.Track, error) {
	g := &domain.Genre{
		Key: genreKey,
	}

	if err := g.Validate(); err != nil {
		return nil, err
	}

	return s.beatPortClient.GetTop10(g)
}
