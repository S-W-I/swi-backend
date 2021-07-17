package server

import (
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
)

func Serve() {

	hash := crypto.Keccak256(make([]byte, 32))
	_ = hash

	m := &http.ServeMux{}
	m.HandleFunc("/k", )

	http.ListenAndServe(":80", m)
}