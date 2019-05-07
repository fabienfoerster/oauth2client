# Oauth2client


## Utilisation

```go
conf := &oauth2.Config{
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    Scopes:       []string{"SCOPE1", "SCOPE2"},
    Endpoint: oauth2.Endpoint{
        AuthURL:  "https://provider.com/o/oauth2/auth",
        TokenURL: "https://provider.com/o/oauth2/token",
    },
}

client := oauth2client.NewClient(conf)
code := client.RetrieveCode()

tok, err := conf.Exchange(oauth2.NoContext, code)
if err != nil {
    log.Fatal(err)
}
  ```
