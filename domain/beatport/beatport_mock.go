package beatport

import "github.com/jeanmorais/beatport/domain"

// ServiceMock represents a mock of BeatPortService interface
type ServiceMock struct {
	GetGenresFn func() ([]domain.Genre, error)
	GetTop10Fn  func(genre string) ([]domain.Track, error)
}

// GetGenres indicates a call of GetGenres in the service layer
func (bsm *ServiceMock) GetGenres() ([]domain.Genre, error) {
	return bsm.GetGenresFn()
}

// GetTop10 indicates a call of GetTop10 in the service layer
func (bsm *ServiceMock) GetTop10(genre string) ([]domain.Track, error) {
	return bsm.GetTop10Fn(genre)
}

// ClientMock represents a mock of BeatPortClient interface
type ClientMock struct {
	GetGenresFn func() ([]domain.Genre, error)
	GetTop10Fn  func(genre *domain.Genre) ([]domain.Track, error)
}

// GetGenres indicates a call of GetGenres in the client layer
func (bcm *ClientMock) GetGenres() ([]domain.Genre, error) {
	return bcm.GetGenresFn()
}

// GetTop10 indicates a call of GetTop10 in the client layer
func (bcm *ClientMock) GetTop10(genre *domain.Genre) ([]domain.Track, error) {
	return bcm.GetTop10Fn(genre)
}
