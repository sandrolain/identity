package keys

import (
	"bytes"
	"testing"

	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/encodeutils"
)

func TestKeyCreate(t *testing.T) {
	byt, err := cryptoutils.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	mk := MasterKeyFromBytes(byt)
	keyLen := 32

	k, _ := NewKey(keyLen)

	if len(k.Value) != keyLen {
		t.Fatalf("Expected key value length of %v but found %v", keyLen, len(k.Value))
	}

	vh := k.ValueHex()
	vfh, err := encodeutils.HexDecode(vh)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(k.Value, vfh) {
		t.Fatalf("Expected key value decoded from hex of %v but found %v", k.Value, vfh)
	}

	vb64 := k.ValueBase64()
	vfb64, err := encodeutils.Base64Decode(vb64)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(k.Value, vfb64) {
		t.Fatalf("Expected key value decoded from b64 of %v but found %v", k.Value, vfb64)
	}

	sk, err := k.Secure(mk)
	if err != nil {
		t.Fatal(err)
	}

	k2, err := sk.Unsecure(mk)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(k.Value, k2.Value) {
		t.Fatalf("Expected equal key value after unsecuring: %v != %v", k2.Value, k.Value)
	}
}

func TestSecureKeyCreate(t *testing.T) {
	byt, err := cryptoutils.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	mk := MasterKeyFromBytes(byt)
	keyLen := 32

	sk, err := NewSecureKey(keyLen, mk)
	if err != nil {
		t.Fatal(err)
	}

	k, err := sk.Unsecure(mk)
	if err != nil {
		t.Fatal(err)
	}

	if len(k.Value) != keyLen {
		t.Fatalf("Expected key value length of %v but found %v", keyLen, len(k.Value))
	}
}

func TestSecuringErrors(t *testing.T) {
	byt, err := cryptoutils.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	mk := MasterKeyFromBytes(byt)
	keyLen := 32

	_, err = NewSecureKey(keyLen, mk)
	if err != nil {
		t.Fatal(err)
	}
}
