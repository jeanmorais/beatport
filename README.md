# beatport

![Build Status](https://github.com/jeanmorais/beatport/workflows/ci/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/jeanmorais/beatport)](https://goreportcard.com/report/github.com/jeanmorais/beatport)
[![codecov](https://codecov.io/gh/jeanmorais/beatport/branch/main/graph/badge.svg)](https://codecov.io/gh/jeanmorais/beatport)
> A REST API to follow the most hyped tracks on the dance floors :notes: :dancers:

This project has no link with [beatport.com](https://beatport.com). It's just an open source project to facilitate the search for tracks through an API.

## Getting Started 

### Requirements

- [Golang](http://golang.org/) (14.0 or higher)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](http://docker.com)

### Setting up and running locally
```bash
# Install dependencies
make install

# Run server
make run
```

### Running the tests
```
make test
```

## Build process

### Build

```bash
make build
```

### Create a Docker image

```bash
make image
```

### Running the app using Docker

```bash
make run-docker
```

## Sample
### Genres

Get all available genres on Beatport:

`GET /genres `
```
[
    {
        "name": "Afro House",
        "key": "afro-house-89",
        "url": "https://www.beatport.com/genre/afro-house/89"
    },
    {
        "name": "Bass House",
        "key": "bass-house-91",
        "url": "https://www.beatport.com/genre/bass-house/91"
    },
...
]
```

### Tracks
#### TOP 10
Get the top 10 tracks by a genre key:

`GET /tracks/top10/:genreKey`

```
[
    {
        "chartNumber": 1,
        "title": "Mona Ki Ngi Xica",
        "remix": "Pablo Fierro Remix",
        "artists": [
            "Pablo Fierro",
            "Bonga"
        ],
        "label": "MoBlack Records",
        "genre": "Afro House",
        "url": "https://www.beatport.com/track/mona-ki-ngi-xica-pablo-fierro-remix/14279156",
        "price": "1.29"
    },
    {
        "chartNumber": 2,
        "title": "Feeling Good",
        "remix": "Original Mix",
        "artists": [
            "Javi Colors",
            "Dr. Alfred"
        ],
        "label": "SP Recordings",
        "genre": "Afro House",
        "url": "https://www.beatport.com/track/feeling-good-original-mix/13781081",
        "price": "1.29"
    },
...
]

```