package keys

import (
	"testing"
	"time"
)

func TestExpiringKeyCreate(t *testing.T) {
	keyLen := 32
	dur, _ := time.ParseDuration("30m")
	key := NewExpiringKey(keyLen, dur, dur)

	if len(key.Value) != keyLen {
		t.Fatalf("Expected value length of %v but found %v", keyLen, len(key.Value))
	}

	exp := time.Now().Add(time.Minute * 30)
	if key.Expire.Unix() != exp.Unix() {
		t.Fatalf("Expected Expire time of %v but found %v", exp, key.Expire)
	}

	exp = exp.Add(time.Minute * 30)
	if key.EOL.Unix() != exp.Unix() {
		t.Fatalf("Expected EOL time of %v but found %v", exp, key.EOL)
	}

	if key.IsExpired() {
		t.Fatalf("Key shouldn't have expired with time diff %v", key.Expire.Sub(time.Now()))
	}

	if key.IsDead() {
		t.Fatalf("Key shouldn't be dead with time diff %v", key.EOL.Sub(time.Now()))
	}

	key = NewExpiringKey(keyLen, time.Minute*-20, time.Minute*30)

	if !key.IsExpired() {
		t.Fatalf("Key should have expired with time diff %v", key.Expire.Sub(time.Now()))
	}

	if key.IsDead() {
		t.Fatalf("Key shouldn't be dead with time diff %v", key.EOL.Sub(time.Now()))
	}

	key = NewExpiringKey(keyLen, time.Minute*-40, time.Minute*30)

	if !key.IsExpired() {
		t.Fatalf("Key should have expired with time diff %v", key.Expire.Sub(time.Now()))
	}

	if !key.IsDead() {
		t.Fatalf("Key should be dead with time diff %v", key.EOL.Sub(time.Now()))
	}
}

func TestKeyRegeneration(t *testing.T) {
	keyLen := 32
	keyValid := NewExpiringKey(keyLen, time.Minute*30, time.Minute*30)
	keyExpired := NewExpiringKey(keyLen, time.Minute*-20, time.Minute*30)
	keyDead := NewExpiringKey(keyLen, time.Minute*-40, time.Minute*30)

	valueValid := keyValid.ValueBase64()
	valueExpired := keyExpired.ValueBase64()
	valueDead := keyDead.ValueBase64()

	list := NewExpiringKeyList(keyDead, keyExpired, keyValid)

	valuesList := list.ValuesBase64()

	if !containsString(valuesList, valueValid) {
		t.Fatalf("Keys list %v not contains valid key %v", valuesList, valueValid)
	}
	if !containsString(valuesList, valueExpired) {
		t.Fatalf("Keys list %v not contains expired key %v", valuesList, valueExpired)
	}
	if !containsString(valuesList, valueDead) {
		t.Fatalf("Keys list %v not contains dead key %v", valuesList, valueDead)
	}

	changed, keyValid2, list := list.EvaluateKeys(keyLen, time.Minute*30, time.Minute*30)

	if !changed {
		t.Fatalf("Expect keys to be changed")
	}

	if keyValid2.ValueBase64() != valueValid {
		t.Fatalf("Valid key %v shouldn't be dropped and replaced by %v", valueValid, keyValid2.ValueBase64())
	}

	valuesNewList := list.ValuesBase64()

	if len(list.Keys) != 2 {
		t.Fatalf("Dead key %v should be dropped from %v", valueDead, valuesNewList)
	}

	if !containsString(valuesNewList, valueValid) {
		t.Fatalf("New keys list %v not contains valid key %v", valuesNewList, valueValid)
	}
	if !containsString(valuesNewList, valueExpired) {
		t.Fatalf("New keys list %v not contains expired key %v", valuesNewList, valueExpired)
	}
	if containsString(valuesNewList, valueDead) {
		t.Fatalf("New keys list %v contains dead key %v", valuesNewList, valueDead)
	}

	changed, keyValid3, list := list.EvaluateKeys(keyLen, time.Minute*30, time.Minute*30)

	if changed {
		t.Fatalf("Expect keys to not be changed")
	}

	if keyValid3.ValueBase64() != valueValid {
		t.Fatalf("Valid key %v shouldn't be dropped and replaced by %v", valueValid, keyValid2.ValueBase64())
	}

	keyExpired = NewExpiringKey(keyLen, time.Minute*-20, time.Minute*30)
	valueExpired = keyExpired.ValueBase64()

	list = NewExpiringKeyList(keyExpired)

	changed, validKey, newList := list.EvaluateKeys(keyLen, time.Minute*30, time.Minute*30)

	if !changed {
		t.Fatalf("Expect keys to be changed")
	}

	if len(newList.Keys) != 2 {
		t.Fatalf("Expect new list to have two keys: %v", len(newList.Keys))
	}

	if containsString(list.ValuesBase64(), validKey.ValueBase64()) {
		t.Fatalf("Expected old list %v to not contains new valid key %v", list.ValuesBase64(), validKey.ValueBase64())
	}
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
