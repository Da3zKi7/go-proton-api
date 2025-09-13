package server

import "github.com/Da3zKi7/go-proton-api/server/backend"

func init() {
	backend.GenerateKey = backend.FastGenerateKey
}
