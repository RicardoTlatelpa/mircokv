package peer

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/stathat/consistent"
)

type Ring struct {
	self string
	peers *consistent.Consistent
}

func NewRing(self string, peerList []string) *Ring {
	r := &Ring {
		self: self,
		peers: consistent.New(),
	}

	for _, p := range peerList {
		if p != self {
			r.peers.Add(p)
		}
	}
	return r
}

func (r *Ring) GetNode(key string) string {
	node, err := r.peers.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	return node
}

func (r *Ring) IsOwner(key string) bool {
	return r.GetNode(key) == r.self
}

func (r *Ring) ProxyGet(key string) (string,error) {
	node := r.GetNode(key)
	resp, err := http.Get("http://" + node + "/get?key=" + url.QueryEscape(key))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

func (r *Ring) ProxySet(key, value string) error {
	node := r.GetNode(key)
	resp, err := http.PostForm("http://"+node+"/set", url.Values{
		"key": {key},
		"value": {value},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}