package aiapic

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/url"

    "bitbucket.org/aisgteam/ai-struct"
    "github.com/afex/hystrix-go/hystrix"
)

type HystrixClient struct {
    Name        string
    Version     aistruct.Version
    BaseUrl     *url.URL
    SubPath     string
    UserAgent   string
    Auth        AuthInterface
    httpClient  *http.Client
}

func (api *HystrixClient) New(u string) error {
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

func (api *HystrixClient) FullName() string {
    return fmt.Sprintf("%s (%s)", api.Name, api.Version)
}

func (api *HystrixClient) PrepareUri(uri string) *url.URL {
    rel := &url.URL{}
    if api.SubPath != "" {
        rel.Path = api.SubPath + uri
    } else {
        rel.Path = uri
    }
    rel = api.BaseUrl.ResolveReference(rel)

    return rel
}

func (api *HystrixClient) Prepare(method string, u *url.URL, body io.Reader) (*http.Request, error) {
    req, err := http.NewRequest(method, u.String(), body)
    if err != nil {
        return nil, api.Err("Prepare", err.Error())
    }
    if api.UserAgent == "" {
        api.UserAgent = fmt.Sprintf("Golang Ai API Hystrix Client (%s)", VERSION)
    }
    api.SetHeader(req, "User-Agent", api.UserAgent)
    if api.Auth != nil {
        api.Auth.Use(req)
    }

    return req, nil
}

func (api *HystrixClient) Exec(r *http.Request) (*http.Response, error) {
    if api.httpClient == nil {
        api.httpClient = &http.Client{}
    }
    var response *http.Response
    if e := hystrix.Do(
        api.Name,
        func() error {
            res, err := api.httpClient.Do(r)
            if err != nil {
                return err
            }
            response = res
            return nil
        },
        nil,
    ); e != nil {
        return nil, api.Err("Exec", e.Error())
    }

    return response, nil
}

func (api *HystrixClient) SetAuth(auth AuthInterface) {
    api.Auth = auth
}

func (api *HystrixClient) SetHeaders(r *http.Request, headers map[string]string) {
    for key, value := range headers {
        api.SetHeader(r, key, value)
    }
}

func (api *HystrixClient) SetHeader(r *http.Request, key, value string) {
    r.Header.Set(key, value)
}

func (api *HystrixClient) Err(method, value string) error {
    name := api.Name
    if name == "" {
        name = "Ai API Hystrix Client"
    }
    return errors.New(fmt.Sprintf("%s.%s(): %s", api.Name, method, value))
}