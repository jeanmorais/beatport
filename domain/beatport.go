package domain

import (
	"fmt"
	"regexp"
	"strings"
)

// Genre represents a musical genre
type Genre struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	URL  string `json:"url"`
}

// Validate validates Genre fields
func (g *Genre) Validate() error {

	matched, _ := regexp.MatchString(`^[a-z-]+\-\d{2}$`, g.Key)
	if !matched {
		return fmt.Errorf("Invalid argument: genreKey [%s]", g.Key)
	}
	return nil
}

// ParseKey parses the genre's key
func (g *Genre) ParseKey() (ID, name string) {
	arr := strings.Split(g.Key, "-")
	ID = arr[len(arr)-1]
	name = strings.Join(arr[:len(arr)-1], "-")
	return
}

// Track represents a track
type Track struct {
	ChartNumber int      `json:"chartNumber"`
	Title       string   `json:"title"`
	Remix       string   `json:"remix"`
	Artists     []string `json:"artists"`
	Label       string   `json:"label"`
	Genre       string   `json:"genre"`
	URL         string   `json:"url"`
	Price       string   `json:"price"`
}

// BeatPortService interface
type BeatPortService interface {
	GetTop10(genre string) ([]Track, error)
	GetGenres() ([]Genre, error)
}

// BeatPortClient interface
type BeatPortClient interface {
	GetTop10(genre *Genre) ([]Track, error)
	GetGenres() ([]Genre, error)
}
