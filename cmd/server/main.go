package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RicardoTlatelpa/microkv/internal/handler"
	"github.com/RicardoTlatelpa/microkv/internal/peer"
	"github.com/RicardoTlatelpa/microkv/internal/store"
)


func main() {
	self := os.Getenv("SELF_ADDRESS")
	if self == ""{
		log.Fatal("SELF_ADDRESS must be set")		
	}
	peerEnv := os.Getenv("PEERS")
	if peerEnv == "" {
		log.Fatal("PEERS must be set")
	}
	peerList := strings.Split(peerEnv, ",")

	store := store.NewStore()

	ring := peer.NewRing(self, peerList)

	handler.RegisterRoutes(store, ring)

	log.Printf("starting server at %s", self)
	log.Fatal(http.ListenAndServe(self,nil))
}