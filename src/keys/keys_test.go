package keys

import (
	"bytes"
	"testing"

	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/encodeutils"
)

func TestKeyCreate(t *testing.T) {
	mk := cryptoutils.RandomBytes(32)
	keyLen := 32

	k := NewKey("foo", keyLen)

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

	if k.Name != sk.Name {
		t.Fatalf("Expected same name %v but found %v", k.Name, sk.Name)
	}

	k2, err := sk.Unsecure(mk)
	if err != nil {
		t.Fatal(err)
	}

	if k.Name != k2.Name {
		t.Fatalf("Expected same name %v but found %v", k.Name, k2.Name)
	}

	if !bytes.Equal(k.Value, k2.Value) {
		t.Fatalf("Expected equal key value after unsecuring: %v != %v", k2.Value, k.Value)
	}
}

func TestSecureKeyCreate(t *testing.T) {
	mk := cryptoutils.RandomBytes(32)
	keyLen := 32

	sk, err := NewSecureKey("foo", SecureKeyParams{keyLen, mk})
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
	mk := cryptoutils.RandomBytes(32)
	badMk := cryptoutils.RandomBytes(31)
	keyLen := 32

	k := NewKey("foo", keyLen)

	_, err := k.Secure(badMk)
	if err == nil {
		t.Fatal("Securing key with bad master key length should generate an error")
	}

	sk, err := NewSecureKey("foo", SecureKeyParams{keyLen, mk})
	if err != nil {
		t.Fatal(err)
	}

	_, err = sk.Unsecure(badMk)
	if err == nil {
		t.Fatal("Unecuring key with bad master key length should generate an error")
	}
}
