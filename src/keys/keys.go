package keys

import (
	"crypto/rand"
	"time"

	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/encodeutils"
)

type MasterKey [32]byte

func MasterKeyFromBytes(in []byte) (res MasterKey) {
	copy(res[:], in[:32])
	return
}

func GenerateKeyValue(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}

func NewKey(name string, length int) *Key {
	// TODO: name validation
	return &Key{
		Name:    name,
		Value:   GenerateKeyValue(length),
		Created: time.Now(),
	}
}

type SecureKeyParams struct {
	Length    int
	MasterKey MasterKey
}

func NewSecureKey(name string, p SecureKeyParams) (*SecuredKey, error) {
	// TODO: name validation
	key := &Key{
		Name:    name,
		Value:   GenerateKeyValue(p.Length),
		Created: time.Now(),
	}
	return key.Secure(p.MasterKey)
}

type Key struct {
	Name    string    `json:"name" bson:"name"`
	Value   []byte    `json:"value" bson:"value"`
	Created time.Time `json:"created" bson:"created"`
}

func (k *Key) ValueHex() string {
	return encodeutils.HexEncode(k.Value)
}

func (k *Key) ValueBase64() string {
	return encodeutils.Base64Encode(k.Value)
}

func (k *Key) Secure(masterKey MasterKey) (*SecuredKey, error) {
	sec, hash, err := cryptoutils.EncryptWithHash(k.Value, masterKey)
	if err != nil {
		return nil, err
	}
	return &SecuredKey{
		Name:    k.Name,
		Value:   sec,
		Hash:    hash,
		Created: k.Created,
	}, nil
}

type SecuredKey struct {
	Name    string    `json:"name" bson:"name"`
	Value   []byte    `json:"value" bson:"value"`
	Hash    []byte    `json:"hash" bson:"hash"`
	Created time.Time `json:"created" bson:"created"`
}

func (k *SecuredKey) Unsecure(masterKey [32]byte) (*Key, error) {
	dec, err := cryptoutils.DecryptAndVerify(k.Value, masterKey, k.Hash)
	if err != nil {
		return nil, err
	}
	return &Key{
		Name:    k.Name,
		Value:   dec,
		Created: k.Created,
	}, nil
}
