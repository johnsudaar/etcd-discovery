package service

import (
	"errors"
	"fmt"
)

type Ports map[string]string

type Host struct {
	Name     string `json:"name"`
	Ports    Ports  `json:"ports"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewHost(hostname string, ports Ports, params ...string) (*Host, error) {
	h := &Host{Name: hostname, Ports: ports}
	if ports == nil || len(ports) == 0 {
		return nil, errors.New("ports is nil, an interface should be defined")
	}
	if len(params) == 1 {
		return nil, errors.New("if user is defined, password should be too")
	} else if len(params) == 2 {
		h.User = params[0]
		h.Password = params[1]
	}

	return h, nil
}

func (h *Host) Url(scheme, path string) (string, error) {
	var url string
	var port string
	var ok bool
	if port, ok = h.Ports[scheme]; !ok {
		return "", errors.New("unknown scheme")
	}
	if h.User != "" {
		url = fmt.Sprintf("%s://%s:%s@%s:%s%s",
			scheme, h.User, h.Password, h.Name, port, path,
		)
	} else {
		url = fmt.Sprintf("%s://%s:%s%s",
			scheme, h.Name, port, path,
		)
	}
	return url, nil
}
