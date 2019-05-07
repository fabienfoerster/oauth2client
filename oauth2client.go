package oauth2client

import (
	"context"
	"net/http"

	"github.com/pkg/browser"

	"golang.org/x/oauth2"
)

type Oauth2Client struct {
	codeChan chan string // channel use to retrieve the code from the dummy server
	Conf     *oauth2.Config
	addr     string
	server   http.Server
}

func (o *Oauth2Client) handleCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	o.codeChan <- code
	w.Write([]byte(code))
}

func (o *Oauth2Client) Shutdown() {
	o.server.Shutdown(context.Background())
}

func NewClient(conf *oauth2.Config) *Oauth2Client {
	client := &Oauth2Client{
		codeChan: make(chan string),
		Conf:     conf,
		addr:     "http://localhost:3000",
	}
	client.Conf.RedirectURL = client.addr

	client.server = http.Server{}
	client.server.Addr = client.addr
	mux := http.NewServeMux()
	mux.HandleFunc("/", client.handleCode)
	client.server.Handler = mux

	return client
}

func (o *Oauth2Client) RetrieveCode() string {

	go o.server.ListenAndServe()
	url := o.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	code := <-o.codeChan
	return code
}
