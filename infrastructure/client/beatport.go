package client

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jeanmorais/beatport/domain"
	httpClient "github.com/jeanmorais/beatport/pkg/http_client"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"strings"
)

type beatPortClient struct {
	httpClient  httpClient.HTTPClient
	beatPortURL string
}

// NewBeatPortClient Create a new client
func NewBeatPortClient(client httpClient.HTTPClient, beatPortURL string) *beatPortClient {
	return &beatPortClient{
		httpClient:  client,
		beatPortURL: beatPortURL,
	}
}

// GetGenres Get all genres
func (bc *beatPortClient) GetGenres() ([]domain.Genre, error) {

	req, err := http.NewRequest(http.MethodGet, bc.beatPortURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := bc.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("unexpected error getting data from beatport [%s]: %v", bc.beatPortURL, err)
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Printf("(GetGenres) error closing response body: %v", err)
		}
	}()

	if res.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected error to extract html data from beatport [%s]: %s", bc.beatPortURL, err)
		}

		regexp, _ := regexp.Compile(`\w+[^/genre/].*`)
		genres := []domain.Genre{}
		doc.Find(".head-drop .genre-drop-list__item").
			Each(func(i int, s *goquery.Selection) {
				name := s.Find("a").Text()
				href, _ := s.Find("a").Attr("href")
				key := strings.Replace(regexp.FindAllString(href, -1)[0], "/", "-", -1)
				genre := domain.Genre{
					Name: name,
					Key:  key,
					URL:  fmt.Sprintf("%s%s", bc.beatPortURL, href),
				}
				genres = append(genres, genre)
			})
		return genres, nil
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("genres not found")
	}

	return nil, fmt.Errorf("unexpected status code [%d] from beatport [%s]", res.StatusCode, bc.beatPortURL)
}

// GetTop10 get the top 10 tracks by genre
func (bc *beatPortClient) GetTop10(genre *domain.Genre) ([]domain.Track, error) {

	ID, name := genre.ParseKey()
	URL := fmt.Sprintf("%s/genre/%s/%s", bc.beatPortURL, name, ID)

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	res, err := bc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Printf("(GetTop10) error to close response body: %v", err)
		}
	}()

	if res.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		tracks := []domain.Track{}
		doc.Find(".top-ten-track").EachWithBreak(func(i int, s *goquery.Selection) bool {
			number, _ := strconv.Atoi(s.Find(".top-ten-track-num").Text())
			title := s.Find(".top-ten-track-primary-title").Text()
			remix := s.Find(".top-ten-track-remixed").Text()
			artists := []string{}
			s.Find(".top-ten-track-artists a").Each(func(i int, s *goquery.Selection) {
				artists = append(artists, s.Text())
			})
			label := s.Find(".top-ten-track-label a").Text()
			genre, _ := s.Attr("data-ec-d3")
			price, _ := s.Attr("data-ec-price")
			href, _ := s.Find("a").Attr("href")
			trackURL := fmt.Sprintf("%s%s", bc.beatPortURL, href)
			if i == 10 {
				return false
			}

			track := domain.Track{
				ChartNumber: number,
				Title:       title,
				Remix:       remix,
				Artists:     artists,
				Label:       label,
				Genre:       genre,
				URL:         trackURL,
				Price:       price,
			}

			tracks = append(tracks, track)
			return true
		})
		return tracks, nil
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("top10 tracks not found [%s]", URL)
	}

	return nil, fmt.Errorf("unexpected status code [%d] from beatport [%s]", res.StatusCode, URL)

}
