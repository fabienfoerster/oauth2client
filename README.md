# Oauth2client


## Utilisation

```go
  conf := &oauth2.Config{
    ClientID:     "xxx",
    ClientSecret: "xxx",
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
