package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func newNumericDateFromSeconds(seconds float64) *jwt.NumericDate {
	timeFromSeconds := time.Unix(int64(seconds), 0)
	return &jwt.NumericDate{Time: timeFromSeconds}
}

func newClaimStringsFromStrings(strings []string) *jwt.ClaimStrings {
	cs := jwt.ClaimStrings(strings)
	return &cs
}

type Claims struct {
	Subject  string           `json:"sub"`
	Iat      int64            `json:"iat"`
	Nbf      int64            `json:"nbf"`
	Issuer   string           `json:"iss"`
	Audience jwt.ClaimStrings `json:"aud"`
	Exp      int64            `json:"exp"`
}

func (c *Claims) GetSubject() (string, error) {
	return c.Subject, nil
}

func (c *Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return newNumericDateFromSeconds(float64(c.Exp)), nil
}

func (c *Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return newNumericDateFromSeconds(float64(c.Nbf)), nil
}

func (c *Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return newNumericDateFromSeconds(float64(c.Iat)), nil
}

func (c *Claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.Audience, nil
}

func (c *Claims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

func newClaimsFromJwtClaims(claims jwt.Claims) (*Claims, bool) {
	subject, z := claims.GetSubject()
	issuer, x := claims.GetIssuer()
	audience, c := claims.GetAudience()
	iat, v := claims.GetIssuedAt()
	nbf, b := claims.GetNotBefore()
	exp, n := claims.GetExpirationTime()
	y := &Claims{
		Subject:  subject,
		Iat:      iat.Unix(),
		Nbf:      nbf.Unix(),
		Issuer:   issuer,
		Audience: audience,
		Exp:      exp.Unix(),
	}
	return y, z == nil && x == nil && c == nil && v == nil && b == nil && n == nil
}

func getJwtAlgorithm(input string) jwt.SigningMethod {
	switch input {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	default:
		return jwt.SigningMethodHS256
	}
}

func SignJWT(c *Claims) string {
	secret := os.Getenv("JWT_SECRET")
	algorithm := os.Getenv("JWT_ALGORITHM")
	signingMethod := getJwtAlgorithm(algorithm)
	token := jwt.NewWithClaims(signingMethod, c)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func VerifyJWT(token string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	algorithm := os.Getenv("JWT_ALGORITHM")
	signingMethod := getJwtAlgorithm(algorithm)
	t, err := jwt.Parse(token, func(jwt *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if t.Method != signingMethod {
		fmt.Printf("Format: %s, Method: %s", t.Method, signingMethod)
		return nil, jwt.ErrSignatureInvalid
	}

	claims, ok := newClaimsFromJwtClaims(t.Claims)

	if !ok || !t.Valid {
		fmt.Printf("ok: %v, valid: %v \n", ok, t.Valid)
		fmt.Printf("Claims: %v \n", claims)
		fmt.Printf("t.Claims: %v \n", t.Claims)
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func RefreshJWT(token string, addTime time.Duration) (string, error) {
	claims, err := VerifyJWT(token)
	if err != nil {
		return "", err
	}
	claims.Iat = time.Now().UTC().Unix()
	claims.Nbf = time.Now().UTC().Unix()
	claims.Exp = time.Now().UTC().Add(time.Hour * addTime).Unix()
	return SignJWT(claims), nil
}
