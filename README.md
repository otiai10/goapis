# goapis

Yet another Web API Client library for Go.

# [Google Custom Search](https://github.com/otiai10/goapis/tree/main/google)

```go
client := google.Client{
    APIKey:               yourAPIKey,
    CustomSearchEngineID: yourSearchEngineID,
}
query := url.Values{"q": "golang"}

res, err := client.CustomSearch(query)
```
