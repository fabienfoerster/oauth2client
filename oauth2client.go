package oauth2client

import (
	"fmt"
	"net/http"

	"github.com/pkg/browser"

	"golang.org/x/oauth2"
)

type Oauth2Client struct {
	codeChan chan string // channel use to retrieve the code from the dummy server
	Conf     *oauth2.Config
	port     int
	Code     string
}

func (o *Oauth2Client) handleCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	o.codeChan <- code
	w.Write([]byte(code))
}

func (o *Oauth2Client) setPort(port int) {
	o.port = port
	o.Conf.RedirectURL = fmt.Sprintf("http://localhost:%d", port)
}

func NewClient(conf oauth2.Config) *Oauth2Client {
	client := &Oauth2Client{
		codeChan: make(chan string),
		Conf:     &conf,
		port:     3000,
	}
	client.Conf.RedirectURL = fmt.Sprintf("http://localhost:%d", client.port)
	return client
}

func (o *Oauth2Client) RetrieveCode() {
	http.HandleFunc("/", o.handleCode)
	port := fmt.Sprintf(":%d", o.port)
	go http.ListenAndServe(port, nil)
	url := o.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	o.Code = <-o.codeChan
}

// RetrieveToken is an exported function
func (o *Oauth2Client) RetrieveToken() (oauth2.Token, error) {
	tok, err := o.Conf.Exchange(oauth2.NoContext, o.Code)
	return *tok, err
}
