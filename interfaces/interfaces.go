package interfaces

import "net/http"

type HttpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}
