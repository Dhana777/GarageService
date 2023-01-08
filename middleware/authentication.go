package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func GenerateToken(userid string) string {
	fmt.Println("user id is", userid)
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 10).Unix()
	td.AccessUuid = userid

	// td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	// td.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	// os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims["authorized"] = true
//	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	fmt.Println("access token is", at)
	var err error
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		fmt.Println(err, ".........")
	}

	return td.AccessToken
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	fmt.Println("token is ", token.Claims)
	res, ok := token.Claims.(jwt.Claims)
	fmt.Printf("lllllllllllllllllll %+v", res)
	if !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("token is", token)
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
