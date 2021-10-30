package main

import (

        "flag"

        "google.golang.org/api/youtube/v3"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "net/url"
        "os"
        "os/user"
        "path/filepath"

        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
)

// This variable indicates whether the script should launch a web server to
// initiate the authorization flow or just display the URL in the terminal
// window. Note the following instructions based on this setting:
// * launchWebServer = true
//   1. Use OAuth2 credentials for a web application
//   2. Define authorized redirect URIs for the credential in the Google APIs
//      Console and set the RedirectURL property on the config object to one
//      of those redirect URIs. For example:
//      config.RedirectURL = "http://localhost:8090"
//   3. In the startWebServer function below, update the URL in this line
//      to match the redirect URI you selected:
//         listener, err := net.Listen("tcp", "localhost:8090")
//      The redirect URI identifies the URI to which the user is sent after
//      completing the authorization flow. The listener then captures the
//      authorization code in the URL and passes it back to this script.
// * launchWebServer = false
//   1. Use OAuth2 credentials for an installed application. (When choosing
//      the application type for the OAuth2 client ID, select "Other".)
//   2. Set the redirect URI to "urn:ietf:wg:oauth:2.0:oob", like this:
//      config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
//   3. When running the script, complete the auth flow. Then copy the
//      authorization code from the browser and enter it on the command line.
const launchWebServer = false

const missingClientSecretsMessage = `
Please configure OAuth 2.0
To make this sample run, you need to populate the client_secrets.json file
found at:
   %v
with information from the {{ Google Cloud Console }}
{{ https://cloud.google.com/console }}
For more information about the client_secrets.json file format, please visit:
https://developers.google.com/api-client-library/python/guide/aaa_client_secrets
`

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(scope string) *http.Client {
        ctx := context.Background()
        
        b, err := ioutil.ReadFile(*clientSecret)
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }
        
        // If modifying the scope, delete your previously saved credentials
        // at ~/.credentials/youtube-go.json
        config, err := google.ConfigFromJSON(b, scope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        
        // Use a redirect URI like this for a web app. The redirect URI must be a
        // valid one for your OAuth2 credentials.
        config.RedirectURL = "http://localhost:8090"
        // Use the following redirect URI if launchWebServer=false in oauth2.go
        // config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
        
        cacheFile, err := tokenCacheFile()
        if err != nil {
                log.Fatalf("Unable to get path to cached credential file. %v", err)
        }
        tok, err := tokenFromFile(cacheFile)
        if err != nil {
                authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
                if launchWebServer {
                        fmt.Println("Trying to get token from web")
//                        tok, err = getTokenFromWeb(config, authURL)
                } else {
                        fmt.Println("Trying to get token from prompt")
                        tok, err = getTokenFromPrompt(config, authURL)
                }
                if err == nil {
                        saveToken(cacheFile, tok)
                }
        }
        return config.Client(ctx, tok)
}

// Exchange the authorization code for an access token
func exchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
        tok, err := config.Exchange(oauth2.NoContext, code)
        if err != nil {
                log.Fatalf("Unable to retrieve token %v", err)
        }
        return tok, nil
}

// getTokenFromPrompt uses Config to request a Token and prompts the user
// to enter the token on the command line. It returns the retrieved Token.
func getTokenFromPrompt(config *oauth2.Config, authURL string) (*oauth2.Token, error) {
        var code string
        fmt.Printf("Go to the following link in your browser. After completing " +
                "the authorization flow, enter the authorization code on the command " +
                "line: \n%v\n", authURL)

        if _, err := fmt.Scan(&code); err != nil {
                log.Fatalf("Unable to read authorization code %v", err)
        }
        fmt.Println(authURL)
        return exchangeToken(config, code)
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
        usr, err := user.Current()
        if err != nil {
                return "", err
        }
        tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
        os.MkdirAll(tokenCacheDir, 0700)
        return filepath.Join(tokenCacheDir,
                url.QueryEscape("youtube-go.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        t := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(t)
        defer f.Close()
        return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
        fmt.Println("trying to save token")
        fmt.Printf("Saving credential file to: %s\n", file)
        f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.



func handleError(err error, message string) {
        if message == "" {
                message = "Error making API call"
        }
        if err != nil {
                log.Fatalf(message+": %v", err.Error())
        }
}
var (
	method = flag.String("method", "list", "The API method to execute. (List is the only method that this sample currently supports.")

	channelId              = flag.String("channelId", "", "Retrieve playlists for this channel. Value is a YouTube channel ID.")
	hl                     = flag.String("hl", "", "Retrieve localized resource metadata for the specified application language.")
	maxResults             = flag.Int64("maxResults", 50, "The maximum number of playlist resources to include in the API response.")
	mine                   = flag.Bool("mine", true, "List playlists for authenticated user's channel. Default: false.")
	onBehalfOfContentOwner = flag.String("onBehalfOfContentOwner", "", "Indicates that the request's auth credentials identify a user authorized to act on behalf of the specified content owner.")
	pageToken              = flag.String("pageToken", "", "Token that identifies a specific page in the result set that should be returned.")
	part                   = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	playlistId             = flag.String("playlistId", "", "Retrieve information about this playlist.")
	clientSecret           = flag.String("clientSecret", "client_secret.json", "")
)

func playlistsList(service *youtube.Service, part []string, channelId string, hl string, maxResults int64, mine bool, onBehalfOfContentOwner string, pageToken string, playlistId string) *youtube.PlaylistListResponse {
	call := service.Playlists.List(part)
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	if hl != "" {
		call = call.Hl(hl)
	}
	call = call.MaxResults(maxResults)
	if mine != false {
		call = call.Mine(true)
	}
	if onBehalfOfContentOwner != "" {
		call = call.OnBehalfOfContentOwner(onBehalfOfContentOwner)
	}
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if playlistId != "" {
		call = call.Id(playlistId)
	}
	response, err := call.Do()
	handleError(err, "")
	return response
}

func main() {
	flag.Parse()

	println("Note if the code in the address bar has %2F, change it to /")
	if *channelId == "" && *mine == false && *playlistId == "" {
		log.Fatalf("You must either set a value for the channelId or playlistId flag or set the mine flag to 'true'.")
	}
	client := getClient(youtube.YoutubeReadonlyScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	response := playlistsList(service, []string{"snippet", "contentDetails"}, *channelId, *hl, *maxResults, *mine, *onBehalfOfContentOwner, *pageToken, *playlistId)

	for _, playlist := range response.Items {
		playlistId := playlist.Id
		playlistTitle := playlist.Snippet.Title

		// Print the playlist ID and title for the playlist resource.
		fmt.Println("https://www.youtube.com/playlist?list=" + playlistId, " :: ", playlistTitle)
	}
}
