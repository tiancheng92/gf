package gf

import (
	"errors"
	"net/url"
	"sort"
	"strings"
)

func UrlFormat(urlStr string) (string, error) {
	obj, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if obj.Host == "" {
		return "", errors.New("host is empty")
	}

	for {
		if strings.Contains(obj.Path, "//") {
			obj.Path = strings.ReplaceAll(obj.Path, "//", "/")
		} else {
			break
		}
	}

	query := obj.Query()
	var keyList []string
	for k := range query {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	queryNew := make(url.Values, len(keyList))
	for i := range keyList {
		values := ArrayDeduplicate(query[keyList[i]])
		sort.Strings(values)
		for j := range values {
			if values[j] != "" {
				queryNew.Add(keyList[i], values[j])
			}
		}
	}

	obj.RawQuery = queryNew.Encode()

	return obj.String(), nil
}
