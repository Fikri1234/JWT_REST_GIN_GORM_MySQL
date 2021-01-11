package util

import (
	"JWT_REST_GIN_GORM_MySQL/model"
	"JWT_REST_GIN_GORM_MySQL/repository"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/segmentio/ksuid"

	"github.com/dgrijalva/jwt-go"
)

var accessSecret string
var refreshSecret string
var timeToken int
var timeRefreshToken int

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUUID   string `json:"-"`
	RefreshUUID  string `json:"-"`
	AtExpires    int64  `json:"atExpires"`
	RtExpires    int64  `json:"rtExpires"`
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

// CreateToken ...
func CreateToken(u model.MUser) (*TokenDetails, error) {

	var err error

	accessSecret = viper.GetString("JWT.ACCESS_SECRET")
	refreshSecret = viper.GetString("JWT.REFRESH_SECRET")

	var expJwt, err1 = repository.GetMParamByKey("EXPIRES_JWT")
	if err1 != nil {
		return nil, err1
	}
	var expRefJwt, err2 = repository.GetMParamByKey("EXPIRES_REFRESH_JWT")
	if err2 != nil {
		return nil, err2
	}

	timeToken, err := strconv.Atoi(expJwt.ParamValue)
	if err != nil {
		return nil, err
	}
	timeRefreshToken, err := strconv.Atoi(expRefJwt.ParamValue)
	if err != nil {
		return nil, err
	}

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(timeToken)).Unix()
	td.AccessUUID = ksuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * time.Duration(timeRefreshToken)).Unix()
	td.RefreshUUID = td.AccessUUID + "++" + strconv.FormatInt(u.ID, 10)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":         td.AtExpires,
		"access_uuid": td.AccessUUID,
		"user_id":     u.ID,
		"name":        u.UserName,
		"authorized":  true,
	})
	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":          td.RtExpires,
		"refresh_uuid": td.RefreshUUID,
		"user_id":      u.ID,
		"name":         u.UserName,
	})
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// ExtractFromCookies ...
func ExtractFromCookies(r *http.Request) (*AccessDetails, error) {
	tokenStr := ExtractToken(r)
	// verify token
	token, err := VerifyToken(r, tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID := claims["user_id"].(float64)

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     int64(userID),
		}, nil
	}
	return nil, err
}

// ExtractToken ...
func ExtractToken(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")
	if len(bearToken) == 0 {
		return ""
	}

	// extract token
	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) == 2 {
		return tokenArr[1]
	}

	return ""
}

// VerifyToken ...
func VerifyToken(r *http.Request, tokenStr string) (*jwt.Token, error) {

	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, err
		}

		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(s)
		}

		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
