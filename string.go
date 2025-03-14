package clair

import (
	"bytes"
	"regexp"
)

func ToSnakeCase(s string) string {
	r := []byte{}
	for _, b := range []byte(s) {
		if b >= 0x41 && b <= 0x5a {
			r = append(r, []byte{0x5f, b + 0x20}...)
		} else if b == 0x2d {
			r = append(r, 0x5f)
		} else {
			r = append(r, b)
		}
	}
	r = bytes.ToLower(bytes.Trim(r, "_ \r\n\t\v"))
	return string(r)
}

func ToPascalCase(s string) string {
	re := regexp.MustCompile(`([ _\-\r\n\t\v]+)`)
	b := re.ReplaceAll([]byte(s), []byte{0x20})
	b = bytes.Trim(b, " ")
	t := bytes.Split(b, []byte{0x20})
	r := []byte{}
	for _, v := range t {
		v = bytes.ToLower(v)
		if v[0] >= 0x61 && v[0] <= 0x7a {
			v[0] -= 0x20
		}
		r = append(r, v...)
	}
	return string(r)
}

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`([ _\-\r\n\t\v]+)`)
	b := re.ReplaceAll([]byte(s), []byte{0x20})
	b = bytes.Trim(b, " ")
	t := bytes.Split(b, []byte{0x20})
	r := []byte{}
	for i, v := range t {
		v = bytes.ToLower(v)
		if i != 0 && v[0] >= 0x61 && v[0] <= 0x7a {
			v[0] -= 0x20
		}
		r = append(r, v...)
	}
	return string(r)
}
