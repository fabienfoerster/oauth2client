# Oauth2client


## Utilisation

```go
  conf := &oauth2.Config{
    ClientID:     "b695d5cdcb49",
    ClientSecret: "dfadc5cf323bbc5c588bba176960d5f6",
    RedirectURL:  "http:localhost:8080",
    Endpoint:     Endpoint,
  }

  client := oauth2client.NewClient(*conf)
  client.RetrieveCode()
  fmt.Println(client.Code)
  oauth2.RegisterBrokenAuthHeaderProvider(Endpoint.TokenURL)
  tok, err := client.RetrieveToken()
  if err != nil {
    fmt.Printf("Fuck, it went wrong : %s\n", err)
  }
  fmt.Printf("And the token is : %s\n", tok.AccessToken)
  ```
