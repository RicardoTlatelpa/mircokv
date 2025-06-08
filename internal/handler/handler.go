package handler

import (
	"fmt"
	"net/http"

	"github.com/RicardoTlatelpa/microkv/internal/peer"
	"github.com/RicardoTlatelpa/microkv/internal/store"
)


func RegisterRoutes(store * store.Store, ring *peer.Ring) {
	http.HandleFunc("/get", func(w http.ResponseWriter, r * http.Request){
		key := r.URL.Query().Get("key")
		if !ring.IsOwner(key){
			val, err := ring.ProxyGet(key)
			if err != nil {
				http.Error(w, "proxy failed", http.StatusBadGateway)
				return
			}
			fmt.Fprint(w, val)
			return
		}
		val, ok := store.Get(key)
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, val)
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request){
		key := r.FormValue("key")
		val := r.FormValue("value")

		if !ring.IsOwner(key) {
			err := ring.ProxySet(key,val)
			if err != nil {
				http.Error(w, "proxy failed", http.StatusBadGateway)
			}
			return
		}
		store.Set(key,val)
		fmt.Fprint(w, "OK")
	})
}