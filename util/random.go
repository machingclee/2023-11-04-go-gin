package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10).String)
}

func RandomString(n int) sql.NullString {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sql.NullString{String: sb.String(), Valid: true}
}

func RandomOwner() sql.NullString {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(400, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
