package keys

import (
	"time"

	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/encodeutils"
)

type MasterKey [32]byte

func MasterKeyFromBytes(in []byte) (res MasterKey) {
	copy(res[:], in[:32])
	return
}

func NewKey(length int) (*Key, error) {
	keyValue, err := cryptoutils.RandomBytes(length)
	if err != nil {
		return nil, err
	}
	return &Key{
		Value:   keyValue,
		Created: time.Now(),
	}, nil
}

func NewSecureKey(length int, mk MasterKey) (*SecuredKey, error) {
	keyValue, err := cryptoutils.RandomBytes(length)
	if err != nil {
		return nil, err
	}
	key := &Key{
		Value:   keyValue,
		Created: time.Now(),
	}
	return key.Secure(mk)
}

type Key struct {
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
		Value:   dec,
		Created: k.Created,
	}, nil
}
