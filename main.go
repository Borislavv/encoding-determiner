package main

import (
	"errors"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/encoding/ianaindex"
	"unicode/utf8"
)

func FindEncoding(data []byte) (encoding.Encoding, error) {
	// Проверяем, закодированы ли данные в UTF-8
	if utf8.Valid(data) {
		return encoding.Nop, nil
	}

	// Пытаемся найти кодировку, используя индекс кодировки HTML
	enc, err := htmlindex.Get("ASCII")
	if err == nil {
		if _, err := enc.NewDecoder().Bytes(data); err == nil {
			return enc, nil
		}
	}

	// Пытаемся найти кодировку, используя индекс кодировки IANA
	enc, err = ianaindex.IANA.Encoding("ASCII")
	if err == nil {
		if _, err := enc.NewDecoder().Bytes(data); err == nil {
			return enc, nil
		}
	}

	enc, err = ianaindex.MIB.Encoding("ASCII")
	if err == nil {
		if _, err := enc.NewDecoder().Bytes(data); err == nil {
			return enc, nil
		}
	}

	enc, err = ianaindex.MIME.Encoding("ASCII")
	if err == nil {
		if _, err := enc.NewDecoder().Bytes(data); err == nil {
			return enc, nil
		}
	}

	return nil, errors.New("unsupported")
}

func main() {
	data := []byte{169}
	enc, err := FindEncoding(data)
	if err != nil {
		panic(err)
	}

	str, err := enc.NewDecoder().String(string(data))
	if err != nil {
		panic(err)
	}

	println("Кодировка:" + str)
}
