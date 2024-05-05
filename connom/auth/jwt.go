package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/go-halo/utils"
	"github.com/golang-jwt/jwt/v5"

	"github.com/fengjx/lucky/connom/config"
)

var (
	secret  []byte
	version string
)

func init() {
	secret = []byte(config.GetConfig().Auth.Secret)
	version = config.GetConfig().Auth.Version
}

type LoginPayload struct {
	UID int64
}

type TokenClaims struct {
	LoginPayload
	jwt.RegisteredClaims
}

type header struct {
	Version string `json:"version"`
}

func GenToken(payload LoginPayload) (string, error) {
	now := utils.Now
	exp := now.Add(time.Hour * 24 * 7)
	claims := TokenClaims{
		LoginPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	h := header{
		Version: version,
	}
	tk := encode(h, jwtToken)
	return tk, nil
}

// Parse 解析 token
func Parse(token string) (payload LoginPayload, expiresAt int64, err error) {
	_, jwtToken, err := decode(token)
	if err != nil {
		return
	}
	tk, err := jwt.ParseWithClaims(jwtToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return
	}
	if claims, ok := tk.Claims.(*TokenClaims); ok && tk.Valid {
		expiresAt = claims.ExpiresAt.Unix()
		payload = claims.LoginPayload
		return
	}
	return
}

func encode(h header, jwtToken string) string {
	hd, _ := json.ToJson(h)
	hs := base64.RawURLEncoding.EncodeToString([]byte(hd))
	ts := base64.RawURLEncoding.EncodeToString([]byte(jwtToken))
	token := strings.Join([]string{hs, ts}, ".")
	return token
}

func decode(token string) (h header, jwtToken string, err error) {
	arr := strings.Split(token, ".")
	if len(arr) < 2 {
		return
	}
	hb, err := base64.RawURLEncoding.DecodeString(arr[0])
	if err != nil {
		return
	}
	err = json.FromBytes(hb, &h)
	if err != nil {
		return
	}
	tb, err := base64.RawURLEncoding.DecodeString(arr[1])
	if err != nil {
		return
	}
	jwtToken = string(tb)
	return
}
