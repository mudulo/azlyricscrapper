# AZLyrics Scrapper

A Scrapper for AZLyrics written in Go

## Getting Started

This repo is a scrapper for azlyrics that exposes an http server, with one endpoint /search which expects a get parameter for the "artist, song description" and then returns lyrics, 

## Built With

* [Golang](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Redis](https://maven.apache.org/) - To Cache frequent requests

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
        "body": "\n\u003c!-- Usage of azlyrics.com content by any third-party lyrics provider is prohibited by our licensing agreement. Sorry about that. --\u003e\nThat\u0026#39;s me right there on the corner, listening to Wu in my Walkman\u003cbr/\u003e\nNeon lights hit the water, reflecting the city I\u0026#39;m lost in\u003cbr/\u003e\nThat\u0026#39;s me right there on the corner, I one day would be leaving\u003cbr/\u003e\nFor a dream that I didn\u0026#39;t have, that I\u0026#39;d one day would believe in\u003cbr/\u003e\nStrange how the same place I ran from\u0026#39;s the same place I think of whenever the chance comes\u003cbr/\u003e\nIt\u0026#39;s inevitable cause wherever I go\u003cbr/\u003e\n\u003cbr/\u003e\nI hear echoes of a thousand screams\u003cbr/\u003e\nAs I lay me down to sleep\u003cbr/\u003e\nThere\u0026#39;s a black hole deep inside of me\u003cbr/\u003e\nReminding me, that I\u0026#39;ve lost my backbone\u003cbr/\u003e\nSomewhere in Stockholm\u003cbr/\u003e\nI lost my backbone, somewhere in Stockholm\u003cbr/\u003e\n\u003cbr/\u003e\nI\u0026#39;m from a place where we never, openly show our emotions\u003cbr/\u003e\nWe drown our sorrows in bottomless bottles and leave them to float in the ocean\u003cbr/\u003e\nI\u0026#39;m from a place where we never, separate people from people\u003cbr/\u003e\nSome generalize, but in general I still believe that we are treated as equals\u003cbr/\u003e\nMy father, my mother, my sister, my brother, my friends and my family\u0026#39;s there\u003cbr/\u003e\nMy hope and my money, my innocence in a sense, almost lost everything here\u003cbr/\u003e\nRight where I was founded, is right where I\u0026#39;ll be found dead\u003cbr/\u003e\nThese streets are my backbone, until I get back home\u003cbr/\u003e\n\u003cbr/\u003e\nI hear echoes of a thousand screams\u003cbr/\u003e\nAs I lay me down to sleep\u003cbr/\u003e\nThere\u0026#39;s a black hole deep inside of me\u003cbr/\u003e\nReminding me, that I\u0026#39;ve lost my backbone\u003cbr/\u003e\nSomewhere in Stockholm\u003cbr/\u003e\nI lost my backbone, somewhere in Stockholm\u003cbr/\u003e\n\u003cbr/\u003e\nI\u0026#39;m not alone, I am the fire that burns not in the city, but out in the burbs\u003cbr/\u003e\nA river that\u0026#39;s dying of thirst, I am a reverend lying in church\u003cbr/\u003e\nA crack in the pattern, a miracle waiting to happen\u003cbr/\u003e\nA promise that never was kept, one of those moments you\u0026#39;ll never forget\u003cbr/\u003e\nI am that feeling inside the one we all know but can\u0026#39;t really describe\u003cbr/\u003e\nI am the blood spill, but I\u0026#39;m in love still\u003cbr/\u003e\n\u003cbr/\u003e\nSomewhere in Stockholm, but I\u0026#39;m not alone\u003cbr/\u003e\nDon\u0026#39;t have to get by on my own, I\u0026#39;m finally home\u003cbr/\u003e\nHemma I Stockholm\u003cbr/\u003e\nDär jag hör hemma\u003cbr/\u003e\n\u003cbr/\u003e\nI hear echoes of a thousand screams\u003cbr/\u003e\nAs I lay me down to sleep\u003cbr/\u003e\nThere\u0026#39;s a black hole deep inside of me\u003cbr/\u003e\nReminding me, that I\u0026#39;ve lost my backbone\u003cbr/\u003e\nSomewhere in Stockholm\u003cbr/\u003e\nI lost my backbone, somewhere in Stockholm\n",
        "copyright": "Avicii lyrics are property and copyright of their owners. \u0026#34;Somewhere In Stockholm\u0026#34; lyrics provided for educational purposes and personal use only"
    }
}


```
## Contributing

Feel free to fork this repo.

## License

This project is licensed under the MIT License 


