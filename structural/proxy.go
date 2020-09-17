package structural

import "fmt"

type ILocalClient interface {
	Get(url string) (resp string, err error)
}

type LocalHost struct{}

func (lh *LocalHost) Get(url string) (resp string, err error) {
	switch url {
	case "https://github.com/user":
		resp = "Github User Account"
	case "https://github.com/repo":
		resp = "Github Repo"
	default:
		err = fmt.Errorf("Unknown URL")
	}

	return
}

type LocalHostProxy struct {
	lh *LocalHost
}

func NewLocalHostProxy() *LocalHostProxy {
	return &LocalHostProxy{lh: &LocalHost{}}
}

func (proxy *LocalHostProxy) Get(url string) (resp string, err error) {
	if url == "" {
		err = fmt.Errorf("url cant be empty")
	}

	if url == "localhost/user" {
		resp, err = proxy.lh.Get("https://github.com/user")
	}

	if url == "localhost/repo" {
		resp, err = proxy.lh.Get("https://github.com/repo")
	}

	return
}
