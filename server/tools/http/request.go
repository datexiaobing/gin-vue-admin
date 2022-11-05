package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/flipped-aurora/gin-vue-admin/server/go-admin-core/logger"
)

type Request struct {
	Url    string
	Header map[string]string
	Method string
	Data   map[string]interface{}
}

func (config *Request) Run() ([]byte, error) {
	Url := config.Url
	Method := config.Method
	Data := config.Data
	Header := make(map[string]string)

	if len(config.Header) > 0 {
		for key, value := range config.Header {
			key = strings.ToTitle(key)
			Header[key] = value
		}
	}

	var request *http.Request
	var err error

	if strings.ToTitle(Method) == "GET" || len(Method) == 0 {
		if len(Data) > 0 {
			Url = UrlAddParams(Url, MaptoUrlParams(Data))
		}
		request, err = http.NewRequest(Method, Url, nil)
	} else if strings.ToTitle(Method) == "POST" {
		var reader io.Reader
		if len(Data) > 0 {
			if value, ok := Header["CONTENT-TYPE"]; ok && strings.ToLower(value) == "application/json" {
				bytesData, err := json.Marshal(Data)
				if err != nil {
					return nil, err
				}
				reader = bytes.NewReader(bytesData)
			} else {
				reader = strings.NewReader(MaptoPosrUrlParams(Data))
				log.Debug(reader)
			}
		}
		request, err = http.NewRequest(Method, Url, reader)
	}

	if err != nil {
		return nil, err
	}

	if len(config.Header) > 0 {
		for key, value := range config.Header {
			request.Header.Set(key, value)
		}
	}

	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func UrlAddParams(Url string, Params string) string {
	if strings.Index(Params, "?") != -1 {
		return Url + "&" + Params
	} else {
		return Url + "?" + Params
	}
}

func MaptoUrlParams(Data map[string]interface{}) string {
	var Params string
	for key, value := range Data {

		Params += key + "=" + url.QueryEscape(Strval(value)) + "&"
	}
	if Params != "" {
		if strings.HasSuffix(Params, "&") {
			Params = Params[:len(Params)-len("&")]
		}
	}
	return Params
}

func MaptoPosrUrlParams(Data map[string]interface{}) string {
	var Params string
	for key, value := range Data {

		Params += key + "=" + Strval(value) + "&"
	}
	if Params != "" {
		if strings.HasSuffix(Params, "&") {
			Params = Params[:len(Params)-len("&")]
		}
	}
	return Params
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = fmt.Sprintf("%f", ft)
	case float32:
		ft := value.(float32)
		key = fmt.Sprintf("%f", ft)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
