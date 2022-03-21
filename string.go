package gf

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"strings"
	"unsafe"
)

func StringJoin(variables ...string) string {
	var (
		n int
		b strings.Builder
	)
	for i := range variables {
		n += len(variables[i])
	}
	b.Grow(n)
	for i := range variables {
		b.WriteString(variables[i])
	}
	return b.String()
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringCreateGzip(str string) ([]byte, error) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.DefaultCompression)
	if err != nil {
		return nil, err
	}
	if _, err = gz.Write(StringToBytes(str)); err != nil {
		return nil, err
	}
	if err = gz.Flush(); err != nil {
		return nil, err
	}
	if err = gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func StringParseGzip(data []byte) (string, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return BytesToString(b), nil
}
