package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "os"
)

func main() {
    // Get the user's TikTok ID.
    userID := os.Args[1]

    // Create the URL to the TikTok API.
    url := fmt.Sprintf("https://www.tiktok.com/api/user/videos/by/%s/", userID)

    // Make the request to the TikTok API.
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read the response body.
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Println(string(body))

    // Unmarshal the response body into a JSON object.
    var videos []struct {
        URL string `json:"url"`
    }
    err = json.Unmarshal(body, &videos)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the videos.
    for _, video := range videos {
        fmt.Println(video.URL)
    }
}

