// Package uniqueid implements the logic for generating those 14
package uniqueId

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

const (
	// base62 character set
	base string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Random integer ceil value
	maxRandomIntCeil int64 = 9999999999999

	// firstJan2014EpochTs is the Timestamp of 1st Jan 2014 in nanosecond precision.
	// It represents the Unix timestamp in nanoseconds for the reference point
	// used in the unique ID generation process. It provides a specific epoch
	// from which time-based IDs are calculated.
	// WARNING:
	// 1. Do not change this value as it can disrupt understanding of
	//    existing data and cause compatibility issues with other components.
	// 2. We want to keep IDs across systems consistent and keeping this fixed
	//    across services helps us in understanding when was it created
	//    (timestamp) just by looking at the ID.
	firstJan2014EpochTs int64 = 1728691200 * 1000 * 1000 * 1000
)

func init() {
	rand.Seed(int64(randUint32()))
}

// New returns a 14 character UUID.
func New() string {
	nanotime := time.Now().UnixNano()

	random := rand.Int63n(maxRandomIntCeil)
	base62Rand := base62Encode(random)

	// We need exactly 4 chars. If greater than 4, strip and use the last 4 chars
	if len(base62Rand) > 4 {
		base62Rand = base62Rand[len(base62Rand)-4:]
	}

	// If less than 4, left pad with '0'
	base62Rand = fmt.Sprintf("%04s", base62Rand)

	b62 := base62Encode(nanotime - firstJan2014EpochTs)
	id := b62 + base62Rand

	return id
}

func base62Encode(num int64) string {
	index := base
	res := ""

	for {
		res = string(index[num%62]) + res
		num = int64(num / 62)
		if num == 0 {
			break
		}
	}
	return res
}

// randUint32 returns a random uint32 using crypto/rand which should in turn be
// used for seeding math/rand.
func randUint32() uint32 {
	buf := make([]byte, 4)
	// This panic is very very unlikely(refer crypto/rand). Anyway this
	// function should not be used regularly but for one time seeding etc.
	if _, err := cryptorand.Reader.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v;", err))
	}
	// Using BigEndian or LittleEndian does not matter here.
	return binary.BigEndian.Uint32(buf)
}
