package controller

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "strconv"
    "github.com/go-redis/redis"
    "github.com/fatih/structs"
    "github.com/PuerkitoBio/goquery"
    "golyrics/model"
    "time"
    "jaytaylor.com/html2text"
)

var client = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: os.Getenv("REDIS_PASSWORD"), 
    DB:       getEnvAsInt("REDIS_DB", 0), 
})


func SearchLyrics(query string) interface{} {
    url := fmt.Sprintf("https://search.azlyrics.com/search.php?q=%s", query)
    resp, err := http.Get(url)
    if err != nil{
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        response := model.ErrorResponse{500, "An error occurred"}
        return response;
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        response := model.ErrorResponse{500, "An error occurred"}
        return response;
    }

    url, found := doc.Find("td.text-left a").Attr("href")
    if(found) {
    	return GetLyrics(url)
    } else {
    	fmt.Println("Could not find lyrics");
        response := model.ErrorResponse{500, "An error occurred"}
        return response;
    }
   
}

func GetLyrics(url string) interface{} {

    vals, err := client.HMGet(url, "Body", "Artist", "Title", "Copyright").Result()

    fmt.Println(vals[1])

    if err != nil {

        response := model.ErrorResponse{500, "An unknown Server Error occured."}
        return response;
            
    } else {

         if vals[0] != nil {
            
            lyrics := model.Lyric{vals[2].(string), vals[1].(string), vals[0].(string), vals[3].(string)}
        
            ttl, _ := time.ParseDuration("24h")
            client.Expire(url, ttl) 
            
            response := model.LyricsResponse{200, lyrics};
            //fmt.Println(response)
            return response;
        } else {

            resp, err := http.Get(url)
            if err != nil{
                log.Fatal(err)
            }
            defer resp.Body.Close()

            if resp.StatusCode != 200{
                log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
            }

            doc, err := goquery.NewDocumentFromReader(resp.Body)
            if err != nil {
                response := model.ErrorResponse{500, "An error occurred"}
                return response;
            }

            header := doc.Find("title").Text()
            meta := strings.Split(header, "-")
            artist := meta[0];
            title := strings.TrimSpace(strings.Split(meta[1], "Lyrics")[0])

            feat := doc.Find(".feat").Text()
            artist_feat := fmt.Sprintf("%s%s", artist,feat)

            copyright, err := doc.Find(".footer-wrap").Find("small").Html();
            if err != nil {
                copyright = "AZLyrics"
            } else {
                copyright = strings.TrimSpace(strings.Split(copyright, "<br/>")[0])
            }
             
            html, err := doc.Find(".ringtone ~ div").Html()
            if err != nil {
                response := model.ErrorResponse{404, "Could not find lyrics"}
                return response;
            } else {
                text, err := html2text.FromString(html, html2text.Options{PrettyTables: false})
                
                if err != nil {

                }  else {

                } 

                lyrics := model.Lyric{title, artist_feat, text, copyright}
                client.HMSet(url, structs.Map(lyrics));
                ttl, _ := time.ParseDuration("24h")
                client.Expire(url, ttl)
                response := model.LyricsResponse{200, lyrics};

                return response;
            }
        }

    }

}

func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
    return value
    }

    return defaultVal
}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
    return value
    }

    return defaultVal
}