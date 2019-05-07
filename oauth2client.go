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
}

func (o *Oauth2Client) handleCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	o.codeChan <- code
	w.Write([]byte(code))
}

func NewClient(conf *oauth2.Config) *Oauth2Client {
	client := &Oauth2Client{
		codeChan: make(chan string),
		Conf:     conf,
		addr:     "0.0.0.0:3000",
	}
	client.Conf.RedirectURL = client.addr
	return client
}

func (o *Oauth2Client) RetrieveCode() string {
	httpServer := http.Server{}
	httpServer.Addr = o.addr
	mux := http.NewServeMux()
	mux.HandleFunc("/", o.handleCode)
	httpServer.Handler = mux
	go httpServer.ListenAndServe()
	url := o.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	code := <-o.codeChan
	httpServer.Shutdown(context.Background())
	return code
}
