# AZLyrics Scrapper

A Scrapper for AZLyrics written in Go

## Getting Started

This repo is a scrapper for azlyrics that exposes an http server, with one endpoint /search which expects a get parameter for the "artist, song description" and then returns lyrics, 

### Prerequisites

What things you need to install the software and how to install them

```
Redis 3.0.6
Golang

```

### Installing

Clone or Download the repo

## Running

```
go to the cloned folder
go run main.go
```

### Test

/search?song=Avicii Somewhere in Stokholm

```
response 
```

## Built With

* [Golang](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Redis](https://maven.apache.org/) - To Cache frequent requests

## Contributing

Feel free to fork this repo.

## License

This project is licensed under the MIT License 

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
