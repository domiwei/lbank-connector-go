package pkg

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	mrand "math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HmacHashing(s string) string {
	d := []byte(s)
	h := md5.New()
	h.Write(d)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func HmacSHA256(paramStr, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(paramStr))
	return hex.EncodeToString(h.Sum(nil))
}

func PrettyPrint(str []byte) (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, str, "", " "); err != nil {
		return string(str), err
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

//func PrettyPrint(i interface{}) string {
//	s, _ := json.MarshalIndent(i, "", "\t")
//	return string(s)
//}

func Map2JsonString(param map[string]interface{}) string {
	data, _ := json.Marshal(param)
	return string(data)
}

func CurrentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func Timestamp() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}

func RandomStr() string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	arr := strings.Split(s, "")
	mrand.Shuffle(len(s), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return strings.Join(arr[:35], "")
}

func Interface2Str(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float32:
		return strconv.FormatFloat(i.(float64), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

func ParsePKCS1PrivateKey(secret []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(secret)
	if block == nil {
		return nil, errors.New("secret key error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey.(*rsa.PrivateKey), err
}

func FormatStringBySign(kwargs map[string]string) string {
	var keys []string

	for k := range kwargs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	paramsSortStr := ""
	for _, key := range keys {
		paramsSortStr += key + "=" + kwargs[key] + "&"
	}
	paramsSortStr = paramsSortStr[:len(paramsSortStr)-1]

	return paramsSortStr
}

func RSASign(params string, privateKey *rsa.PrivateKey) string {
	md5Hash := md5.Sum([]byte(params))
	md5String := strings.ToUpper(hex.EncodeToString(md5Hash[:]))

	paramsSha256 := sha256.New()
	paramsSha256.Write([]byte(md5String))
	sha256Hash := paramsSha256.Sum(nil)

	sigMsg, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sha256Hash)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(sigMsg)
}

func HmacSha256Base64Signer(params string, secretKey string) (string, error) {
	md5Hash := md5.Sum([]byte(params))
	md5String := strings.ToUpper(hex.EncodeToString(md5Hash[:]))

	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(md5String))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(mac.Sum(nil)), nil
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RandomUUID() string {
	id := uuid.New()
	strID := id.String()
	return strID
}
