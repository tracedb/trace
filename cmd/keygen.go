package main

import (
	"fmt"

	"github.com/unit-io/unitd/message/security"
	"github.com/unit-io/unitd/pkg/hash"
)

const (
	// KeySize is the size of the key used by this AEAD, in bytes.
	KeySize = 32

	// NonceSize is the size of the nonce used with the standard variant of this
	// AEAD, in bytes.
	//
	// Note that this is too short to be safely generated at random if the same
	// key is reused more than 2³² times.
	NonceSize = 24
	// NonceSizeX is the size of the nonce used with the XChaCha20-Poly1305
	// variant of this AEAD, in bytes.
	NonceSizeX = 24
)

type keyGenRequest struct {
	Key   string
	Topic string
	Type  string
}

func (m *keyGenRequest) access() uint32 {
	required := security.AllowNone

	for i := 0; i < len(m.Type); i++ {
		switch c := m.Type[i]; c {
		case 'r':
			required |= security.AllowRead
		case 'w':
			required |= security.AllowWrite
		}
	}

	return required
}

func generateKey() {
	fmt.Println("connectionstore", hash.WithSalt([]byte("connectionstore"), uint32(3376684800)))
	fmt.Println("keygen: ", hash.WithSalt([]byte("keygen"), uint32(3376684800)))
	fmt.Println("presence: ", hash.WithSalt([]byte("presence"), uint32(3376684800)))
	fmt.Println("clientid: ", hash.WithSalt([]byte("clientid"), uint32(3376684800)))
	fmt.Println("...: ", hash.WithSalt([]byte("..."), uint32(3376684800)))
	fmt.Println("*", hash.WithSalt([]byte("*"), uint32(3376684800)))
	message := keyGenRequest{
		//Key:   master,
		Topic: "dev20.world",
		Type:  "rwp",
	}

	key, _ := security.GenerateKey(uint32(3376684800), []byte(message.Topic), message.access())
	fmt.Println("Key: ", key)
}
