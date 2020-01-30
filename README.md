# AZLyrics Scrapper

A Scrapper for AZLyrics written in Go

## Getting Started

This repo is a scrapper for azlyrics that exposes an http endpoint , with one endpoint /search which expects a get parameter for the "artist, song description (no specific order or delimiters since this is a search query)" and then returns lyrics, 

## Built With

* [Golang](https://golang.org/) - The programming alnguage used
* [Redis](https://redis.io/) - To Cache frequent requests for max 24 hours

### Installing

Clone or Download the repo

## Running

```
go to the cloned folder
go run main.go
```

### Example

**/search?song=Avicii Somewhere in Stokholm**

```
response 

json
{
    "code": 200,
    "lyrics": {
        "title": "Somewhere In Stockholm",
        "artist": "Avicii",
        "body": "That's me right there on the corner, listening to Wu in my Walkman\nNeon lights hit the water, reflecting the city I'm lost in\nThat's me right there on the corner, I one day would be leaving\nFor a dream that I didn't have, that I'd one day would believe in\nStrange how the same place I ran from's the same place I think of whenever the chance comes\nIt's inevitable cause wherever I go\n\nI hear echoes of a thousand screams\nAs I lay me down to sleep\nThere's a black hole deep inside of me\nReminding me, that I've lost my backbone\nSomewhere in Stockholm\nI lost my backbone, somewhere in Stockholm\n\nI'm from a place where we never, openly show our emotions\nWe drown our sorrows in bottomless bottles and leave them to float in the ocean\nI'm from a place where we never, separate people from people\nSome generalize, but in general I still believe that we are treated as equals\nMy father, my mother, my sister, my brother, my friends and my family's there\nMy hope and my money, my innocence in a sense, almost lost everything here\nRight where I was founded, is right where I'll be found dead\nThese streets are my backbone, until I get back home\n\nI hear echoes of a thousand screams\nAs I lay me down to sleep\nThere's a black hole deep inside of me\nReminding me, that I've lost my backbone\nSomewhere in Stockholm\nI lost my backbone, somewhere in Stockholm\n\nI'm not alone, I am the fire that burns not in the city, but out in the burbs\nA river that's dying of thirst, I am a reverend lying in church\nA crack in the pattern, a miracle waiting to happen\nA promise that never was kept, one of those moments you'll never forget\nI am that feeling inside the one we all know but can't really describe\nI am the blood spill, but I'm in love still\n\nSomewhere in Stockholm, but I'm not alone\nDon't have to get by on my own, I'm finally home\nHemma I Stockholm\nDär jag hör hemma\n\nI hear echoes of a thousand screams\nAs I lay me down to sleep\nThere's a black hole deep inside of me\nReminding me, that I've lost my backbone\nSomewhere in Stockholm\nI lost my backbone, somewhere in Stockholm","copyright":"Avicii lyrics are property and copyright of their owners. \u0026#34;Somewhere In Stockholm\u0026#34; lyrics provided for educational purposes and personal use only.",
        "copyright": "Avicii lyrics are property and copyright of their owners. \u0026#34;Somewhere In Stockholm\u0026#34; lyrics provided for educational purposes and personal use only"
    }
}


```
## Contributing

Feel free to fork this repo.

## License

This project is licensed under the MIT License 


