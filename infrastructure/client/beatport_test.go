package client_test

import (
	"fmt"
	"github.com/jeanmorais/beatport/domain"
	"github.com/jeanmorais/beatport/infrastructure/client"
	httpClient "github.com/jeanmorais/beatport/pkg/http_client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

var expectedGenres = []domain.Genre{
	{
		"Afro House", "afro-house-89", "https://www.beatport.com/genre/afro-house/89",
	},
	{
		"Bass House", "bass-house-91", "https://www.beatport.com/genre/bass-house/91",
	},
	{
		"Big Room", "big-room-79", "https://www.beatport.com/genre/big-room/79",
	},
	{
		"Breaks", "breaks-9", "https://www.beatport.com/genre/breaks/9",
	},
	{
		"DJ Tools", "dj-tools-16", "https://www.beatport.com/genre/dj-tools/16",
	},
	{
		"Dance", "dance-39", "https://www.beatport.com/genre/dance/39",
	},
	{
		"Deep House", "deep-house-12", "https://www.beatport.com/genre/deep-house/12",
	},
	{
		"Drum & Bass", "drum-and-bass-1", "https://www.beatport.com/genre/drum-and-bass/1",
	},
	{
		"Dubstep", "dubstep-18", "https://www.beatport.com/genre/dubstep/18",
	},
	{
		"Electro House", "electro-house-17", "https://www.beatport.com/genre/electro-house/17",
	},
	{
		"Electronica / Downtempo", "electronica-downtempo-3", "https://www.beatport.com/genre/electronica-downtempo/3",
	},
	{
		"Funky / Groove / Jackin' House", "funky-groove-jackin-house-81", "https://www.beatport.com/genre/funky-groove-jackin-house/81",
	},
	{
		"Future House", "future-house-65", "https://www.beatport.com/genre/future-house/65",
	},
	{
		"Garage / Bassline / Grime", "garage-bassline-grime-86", "https://www.beatport.com/genre/garage-bassline-grime/86",
	},
	{
		"Hard Dance / Hardcore", "hard-dance-hardcore-8", "https://www.beatport.com/genre/hard-dance-hardcore/8",
	},
	{
		"Hip-Hop / R&B", "hip-hop-r-and-b-38", "https://www.beatport.com/genre/hip-hop-r-and-b/38",
	},
	{
		"House", "house-5", "https://www.beatport.com/genre/house/5",
	},
	{
		"Indie Dance", "indie-dance-37", "https://www.beatport.com/genre/indie-dance/37",
	},
	{
		"Leftfield Bass", "leftfield-bass-85", "https://www.beatport.com/genre/leftfield-bass/85",
	},
	{
		"Leftfield House & Techno", "leftfield-house-and-techno-80", "https://www.beatport.com/genre/leftfield-house-and-techno/80",
	},
	{
		"Melodic House & Techno", "melodic-house-and-techno-90", "https://www.beatport.com/genre/melodic-house-and-techno/90",
	},
	{
		"Minimal / Deep Tech", "minimal-deep-tech-14", "https://www.beatport.com/genre/minimal-deep-tech/14",
	},
	{
		"Nu Disco / Disco", "nu-disco-disco-50", "https://www.beatport.com/genre/nu-disco-disco/50",
	},
	{
		"Progressive House", "progressive-house-15", "https://www.beatport.com/genre/progressive-house/15",
	},
	{
		"Psy-Trance", "psy-trance-13", "https://www.beatport.com/genre/psy-trance/13",
	},
	{
		"Reggae / Dancehall / Dub", "reggae-dancehall-dub-41", "https://www.beatport.com/genre/reggae-dancehall-dub/41",
	},
	{
		"Tech House", "tech-house-11", "https://www.beatport.com/genre/tech-house/11",
	},
	{
		"Techno (Peak Time / Driving / Hard)", "techno-peak-time-driving-hard-6", "https://www.beatport.com/genre/techno-peak-time-driving-hard/6",
	},
	{
		"Techno (Raw / Deep / Hypnotic)", "techno-raw-deep-hypnotic-92", "https://www.beatport.com/genre/techno-raw-deep-hypnotic/92",
	},
	{
		"Trance", "trance-7", "https://www.beatport.com/genre/trance/7",
	},
	{
		"Trap / Future Bass", "trap-future-bass-87", "https://www.beatport.com/genre/trap-future-bass/87",
	},
}

var top10Response = []domain.Track{
	{1, "Zvezdara", "Original Mix", []string{"Last95"}, "MIR MUSIC", "Progressive House", "https://www.beatport.com/track/zvezdara-original-mix/13418144", "1.29"},
	{2, "Lost In You", "Extended Mix", []string{"Marsh"}, "Anjunadeep", "Progressive House", "https://www.beatport.com/track/lost-in-you-extended-mix/13857948", "1.29"},
	{3, "Discopolis 2.0", "MEDUZA Extended Remix", []string{"Meduza", "Lifelike", "Kris Menace"}, "Armada Music", "Progressive House", "https://www.beatport.com/track/discopolis-2-0-meduza-extended-remix/13616384", "1.29"},
	{4, "Inside Me", "Extended Mix", []string{"The Dualz"}, "Anjunadeep", "Progressive House", "https://www.beatport.com/track/inside-me-extended-mix/13770067", "1.29"},
	{5, "No Time To Wait", "Extended Mix", []string{"Jerome Isma-Ae", "Milkwish"}, "Anjunabeats", "Progressive House", "https://www.beatport.com/track/no-time-to-wait-extended-mix/13843498", "1.29"},
	{6, "Wherever You Are feat. Margret", "Tinlicker Remix", []string{"Tinlicker", "Nils Hoffmann", "MARGRET"}, "Poesie Musik", "Progressive House", "https://www.beatport.com/track/wherever-you-are-feat-margret-tinlicker-remix/13782984", "1.29"},
	{7, "The Great Escape", "Original Mix", []string{"Volen Sentir"}, "Lost & Found", "Progressive House", "https://www.beatport.com/track/the-great-escape-original-mix/13074964", "1.29"},
	{8, "Zooz", "Original Mix", []string{"Stereo Underground"}, "Lost & Found", "Progressive House", "https://www.beatport.com/track/zooz-original-mix/13776028", "1.29"},
	{9, "Apollo", "Extended Mix", []string{"Melody Stranger", "Sean & Dee"}, "UV", "Progressive House", "https://www.beatport.com/track/apollo-extended-mix/13592753", "1.29"},
	{10, "Mars", "Extended Mix", []string{"Shadow Child"}, "Armada Electronic Elements", "Progressive House", "https://www.beatport.com/track/mars-extended-mix/13859752", "1.29"},
}

const (
	beatPortURL = "https://www.beatport.com"
)

func TestBeatPortClient_GetGenres(t *testing.T) {
	t.Run("should get genres successfully", func(t *testing.T) {

		body, err := os.Open("../../testdata/index.html")
		assert.NoError(t, err)

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 200,
			ResponseBody:   body,
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genres, err := beatPortClient.GetGenres()
		assert.NoError(t, err)
		assert.Equal(t, len(genres), 31)
		assert.Equal(t, genres, expectedGenres)

	})

	t.Run("should return not found error http 404", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 404,
			ResponseBody:   ioutil.NopCloser(strings.NewReader("not found")),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genres, err := beatPortClient.GetGenres()
		assert.Error(t, err)
		assert.Nil(t, genres)
		assert.Equal(t, http.StatusNotFound, httpClientMock.ResponseStatus)
	})

	t.Run("should return unexpected http status code", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 500,
			ResponseBody:   ioutil.NopCloser(strings.NewReader("error")),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genres, err := beatPortClient.GetGenres()
		assert.Error(t, err)
		assert.Nil(t, genres)
		assert.Equal(t, http.StatusInternalServerError, httpClientMock.ResponseStatus)
	})

	t.Run("should return http request error", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			Error: fmt.Errorf("http request error"),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genres, err := beatPortClient.GetGenres()
		assert.Error(t, err)
		assert.Nil(t, genres)
	})
}

func TestBeatPortClient_GetTop10(t *testing.T) {
	t.Run("should get top10 tracks successfully", func(t *testing.T) {

		body, err := os.Open("../../testdata/progressive.html")
		assert.NoError(t, err)

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 200,
			ResponseBody:   body,
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genre := &domain.Genre{Key: "progressive-house-15"}
		tracks, err := beatPortClient.GetTop10(genre)

		assert.NoError(t, err)
		assert.Equal(t, len(tracks), 10)
		assert.Equal(t, tracks, top10Response)

	})

	t.Run("should return not found error http 404", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 404,
			ResponseBody:   ioutil.NopCloser(strings.NewReader("not found")),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genre := &domain.Genre{Key: "progressive-10"}

		tracks, err := beatPortClient.GetTop10(genre)
		assert.Error(t, err)
		assert.Nil(t, tracks)
		assert.Equal(t, http.StatusNotFound, httpClientMock.ResponseStatus)
	})

	t.Run("should return unexpected http status code", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			ResponseStatus: 500,
			ResponseBody:   ioutil.NopCloser(strings.NewReader("error")),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genre := &domain.Genre{Key: "progressive-10"}
		tracks, err := beatPortClient.GetTop10(genre)
		assert.Error(t, err)
		assert.Nil(t, tracks)
		assert.Equal(t, http.StatusInternalServerError, httpClientMock.ResponseStatus)
	})

	t.Run("should return http request error", func(t *testing.T) {

		httpClientMock := &httpClient.ClientMock{
			Error: fmt.Errorf("http request error"),
		}

		beatPortClient := client.NewBeatPortClient(httpClientMock, beatPortURL)
		assert.NotNil(t, beatPortClient)

		genre := &domain.Genre{Key: "progressive-10"}
		tracks, err := beatPortClient.GetTop10(genre)
		assert.Error(t, err)
		assert.Nil(t, tracks)
	})
}
