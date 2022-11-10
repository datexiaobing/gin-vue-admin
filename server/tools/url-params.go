package tools

import (
	"errors"
	"net/url"
	"strings"
)

// 返回URL中的sign 和path
func GetSignPath(s string) (sign string, path string) {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)
	// fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	return m["sign"][0], u.Path
}

// 返回URL中的path
func GetPath(s string) (path string) {

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return u.Path
}

func GetPathFromEnc(s string) ([]string, error) {

	a := strings.Split(s, "?")

	if len(a) != 3 {
		return []string{}, errors.New("wrong encrypt data")
	}
	return a, nil
}
