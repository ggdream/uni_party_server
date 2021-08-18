package captcha

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	verifyApi = "https://c.dun.163yun.com/api/v2/verify"
	version   = "v2"
)

var (
	CaptchaIdIsEmpty    = errors.New("captchaId is empty")
	SecretIdIsEmpty     = errors.New("secretId is empty")
	SecretKeyIsEmpty    = errors.New("secretKey is empty")
	ValidateDataIsEmpty = errors.New("validate data is empty")
)

type VerifyResult struct {
	Err       int    `json:"error"`
	Msg       string `json:"msg"`
	Result    bool   `json:"result"`
	Phone     string `json:"phone"`
	ExtraData string `json:"extraData"`
}

type dunShell struct {
	client *http.Client

	captchaId string
	secretId  string
	secretKey string
}

func NewDun(captchaId, secretId, secretKey string) (*dunShell, error) {
	if captchaId == "" {
		return nil, CaptchaIdIsEmpty
	}
	if secretId == "" {
		return nil, SecretIdIsEmpty
	}
	if secretKey == "" {
		return nil, SecretKeyIsEmpty
	}

	return &dunShell{
		client: &http.Client{
			Timeout: 6 * time.Second,
		},
		captchaId: captchaId,
		secretId:  secretId,
		secretKey: secretKey,
	}, nil
}

func (s *dunShell) postForm(url string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(s.convertToQueryParams(params)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (s *dunShell) convertToQueryParams(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for k, v := range params {
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}

func (s *dunShell) Verify(validate, user string) (*VerifyResult, error) {
	if validate == "" || user == "" {
		return nil, ValidateDataIsEmpty
	}
	params := map[string]string{}
	params["captchaId"] = s.captchaId
	params["validate"] = validate
	params["user"] = user
	params["secretId"] = s.secretId
	params["version"] = version
	params["timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	params["nonce"] = s.random(20)
	params["signature"] = s.genSignature(s.secretKey, params)

	data, err := s.postForm(verifyApi, params)
	if err != nil {
		return nil, err
	}

	verifyResult := &VerifyResult{}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, verifyResult)
	return verifyResult, err
}

func (s *dunShell) genSignature(secretKey string, params map[string]string) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	buf := bytes.NewBufferString("")
	for _, key := range keys {
		buf.WriteString(key + params[key])
	}
	buf.WriteString(secretKey)
	has := md5.Sum(buf.Bytes())
	return fmt.Sprintf("%x", has)
}

func (s *dunShell) random(l int) string {
	strBytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, strBytes[r.Intn(len(strBytes))])
	}
	return string(result)
}
