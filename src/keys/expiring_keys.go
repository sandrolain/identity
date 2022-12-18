package keys

import (
	"time"

	"github.com/sandrolain/go-utilities/pkg/encodeutils"
)

func NewExpiringKey(length int, toExpire time.Duration, toEol time.Duration) *ExpiringKey {
	now := time.Now()
	expire := now.Add(toExpire)
	eol := expire.Add(toEol)
	return &ExpiringKey{
		Value:   GenerateKeyValue(length),
		Created: now,
		Expire:  expire,
		EOL:     eol,
	}
}

type ExpiringKey struct {
	Value   []byte
	Created time.Time
	Expire  time.Time
	EOL     time.Time
}

func (k *ExpiringKey) Valid() bool {
	return k.Expire.After(time.Now())
}

func (k *ExpiringKey) ValidAtTime(t time.Time) bool {
	return k.Created.UTC().Before(t) && k.Expire.After(t)
}

func (k *ExpiringKey) IsExpired() bool {
	return k.Expire.Before(time.Now())
}

func (k *ExpiringKey) IsDead() bool {
	return k.EOL.Before(time.Now())
}

func (k *ExpiringKey) ValueHex() string {
	return encodeutils.HexEncode(k.Value)
}

func (k *ExpiringKey) ValueBase64() string {
	return encodeutils.Base64Encode(k.Value)
}

func NewExpiringKeyList(list ...*ExpiringKey) *ExpiringKeyList {
	res := ExpiringKeyList{}
	if len(list) > 0 {
		res.Keys = append(res.Keys, list...)
	}
	return &res
}

type ExpiringKeyList struct {
	Keys []*ExpiringKey
}

func (l *ExpiringKeyList) EvaluateKeys(length int, toExpire time.Duration, toEol time.Duration) (bool, *ExpiringKey, *ExpiringKeyList) {
	newList := ExpiringKeyList{}
	changed := false
	lk := (*l).Keys
	var validKey *ExpiringKey
	for _, key := range lk {
		if !key.IsDead() {
			newList.Keys = append(newList.Keys, key)
			if !key.IsExpired() {
				validKey = key
			}
		} else {
			changed = true
		}
	}
	if validKey == nil {
		validKey = NewExpiringKey(length, toExpire, toEol)
		newList.Keys = append(newList.Keys, validKey)
		changed = true
	}
	return changed, validKey, &newList
}

func (l *ExpiringKeyList) GetValidKey() *ExpiringKey {
	lk := (*l).Keys
	for _, key := range lk {
		if !key.IsExpired() {
			return key
		}
	}
	return nil
}

func (l *ExpiringKeyList) GetValidKeyAtTime(t time.Time) *ExpiringKey {
	lk := (*l).Keys
	for _, key := range lk {
		if key.ValidAtTime(t) {
			return key
		}
	}
	return nil
}

func (l *ExpiringKeyList) Values() [][]byte {
	lk := (*l).Keys
	res := make([][]byte, len(lk))
	for i, key := range lk {
		res[i] = key.Value
	}
	return res
}

func (l *ExpiringKeyList) ValuesHex() []string {
	lk := (*l).Keys
	res := make([]string, len(lk))
	for i, key := range lk {
		res[i] = key.ValueHex()
	}
	return res
}

func (l *ExpiringKeyList) ValuesBase64() []string {
	lk := (*l).Keys
	res := make([]string, len(lk))
	for i, key := range lk {
		res[i] = key.ValueBase64()
	}
	return res
}
