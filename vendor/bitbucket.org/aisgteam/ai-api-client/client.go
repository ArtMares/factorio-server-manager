package aiapic

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/url"

    "bitbucket.org/aisgteam/ai-struct"
)

type Client struct {
    Name        string
    Version     aistruct.Version
    BaseUrl     *url.URL
    SubPath     string
    UserAgent   string
    Auth        AuthInterface
    httpClient  *http.Client
}

func (api *Client) New(u string) error {
    if u == "" {
        return api.Err("New", "Empty URL")
    }
    ur, err := url.Parse(u)
    if err != nil {
        return api.Err("New", err.Error())
    }
    api.BaseUrl = ur
    api.SubPath = ur.Path

    return nil
}

func (api *Client) FullName() string {
    return fmt.Sprintf("%s (%s)", api.Name, api.Version)
}

func (api *Client) PrepareUri(uri string) *url.URL {
    rel := &url.URL{}
    if api.SubPath != "" {
        rel.Path = api.SubPath + uri
    } else {
        rel.Path = uri
    }
    rel = api.BaseUrl.ResolveReference(rel)

    return rel
}

func (api *Client) Prepare(method string, u *url.URL, body io.Reader) (*http.Request, error) {
    req, err := http.NewRequest(method, u.String(), body)
    if err != nil {
        return nil, api.Err("Prepare", err.Error())
    }
    if api.UserAgent == "" {
        api.UserAgent = fmt.Sprintf("Golang Ai API Client (%s)", VERSION)
    }
    api.SetHeader(req, "User-Agent", api.UserAgent)
    if api.Auth != nil {
        api.Auth.Use(req)
    }

    return req, nil
}

func (api *Client) Exec(r *http.Request) (*http.Response, error) {
    if api.httpClient == nil {
        api.httpClient = &http.Client{}
    }
    res, err := api.httpClient.Do(r)
    if err != nil {
        return nil, api.Err("Exec", err.Error())
    }

    return res, nil
}

func (api *Client) SetAuth(auth AuthInterface) {
    api.Auth = auth
}

func (api *Client) SetHeaders(r *http.Request, headers map[string]string) {
    for key, value := range headers {
        api.SetHeader(r, key, value)
    }
}

func (api *Client) SetHeader(r *http.Request, key, value string) {
    r.Header.Set(key, value)
}

func (api *Client) Err(method, value string) error {
    name := api.Name
    if name == "" {
        name = "Ai API Client"
    }
    return errors.New(fmt.Sprintf("%s.%s(): %s", api.Name, method, value))
}

