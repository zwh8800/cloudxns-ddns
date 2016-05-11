package cloudxns

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	"github.com/zwh8800/cloudxns-ddns/conf"
)

const mineType = "application/json"
const successCode = 1

type dynamicDnsRequestPayload struct {
	Domain string `json:"domain"`
}

type dynamicDnsResponsePayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func MD5(src string) string {
	sumData := md5.Sum([]byte(src))
	return hex.EncodeToString(sumData[:])
}

func hmac(url, body, date string) string {
	return MD5(conf.Conf.CloudXNS.APIKey + url + body + date + conf.Conf.CloudXNS.SecureKey)
}

func dynamicDns(domain string) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	reqPayload := dynamicDnsRequestPayload{
		Domain: domain,
	}
	if err := encoder.Encode(reqPayload); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", dynamicDnsApiUrl, buf)
	if err != nil {
		return err
	}
	now := time.Now()
	req.Header.Set("Content-Type", mineType)
	req.Header.Set("API-KEY", conf.Conf.CloudXNS.APIKey)
	req.Header.Set("API-REQUEST-DATE", now.Format(time.RFC1123Z))
	req.Header.Set("API-HMAC", hmac(dynamicDnsApiUrl, buf.String(), now.Format(time.RFC1123Z)))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	resPayload := &dynamicDnsResponsePayload{}
	if err := json.Unmarshal(responseData, resPayload); err != nil {
		return err
	}
	if resPayload.Code != successCode {
		return fmt.Errorf("%d: %s", resPayload.Code, resPayload.Message)
	}

	return nil
}

func DynamicDns(domains []string) error {
	for _, domain := range domains {
		if err := dynamicDns(domain); err != nil {
			return err
		}
	}
	return nil
}
