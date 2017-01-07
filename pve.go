package pve

import (
	"github.com/dghubble/sling"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"log"
	"strings"
	"time"
)

type PveClient struct {
	sling      *sling.Sling
	username   string
	realm      string
	password   string
	httpClient *pveHttpClient
}

// nomp dont send the api response with content type
// we fix this: set content type to json
type pveHttpClient struct {
	client    *http.Client
	debug     bool
	ticket    string
	token     string
	timestamp time.Time
	useragent string
}

func (d pveHttpClient) Do(req *http.Request) (*http.Response, error) {
	if d.debug {
		d.dumpRequest(req)
	}
	if d.useragent != "" {
		req.Header.Set("User-Agent", d.useragent)
	}
	if d.ticket != "" {
		req.AddCookie(&http.Cookie{Name: "PVEAuthCookie", Value: d.ticket})
		if req.Method != "GET" {
			req.Header.Set("CSRFPreventionToken", d.token)
		}
	}
	client := func() (*http.Client) {
		if d.client != nil {
			return d.client
		} else {
			return http.DefaultClient
		}
	}()
	if client.Transport != nil {
		if transport, ok := client.Transport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	} else {
		if transport, ok := http.DefaultTransport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	}
	resp, err := client.Do(req)
	if d.debug {
		d.dumpResponse(resp)
	}
	if err == nil {
		contenttype := resp.Header.Get("Content-Type");
		if len(contenttype) == 0 || strings.HasPrefix(contenttype, "text/html") {
			resp.Header.Set("Content-Type", "application/json")
		}
	}
	return resp, err
}

func (d pveHttpClient) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (d pveHttpClient) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func NewPveClient(client *http.Client, BaseURL string, UserAgent string, Username string, Password string, Realm string) *PveClient {
	httpClient := &pveHttpClient{client:client, useragent:UserAgent}
	return &PveClient{
		httpClient: httpClient,
		username:Username,
		password:Password,
		realm:Realm,
		sling: sling.New().Doer(httpClient).Base(BaseURL).Path("api2/json/"),
	}
}

func (client PveClient) SetDebug(debug bool) {
	client.httpClient.debug = debug
}

