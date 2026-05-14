package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const salt = "go-2024"

type Values struct {
	NumDecimal     int
	NumOctal       int
	NumHexadecimal int
	Pi             float64
	Name           string
	IsActive       bool
	ComplexNum     complex64
}

type Result struct {
	Combined string
	Runes    []rune
	Hash     string
}

func DefaultValues() Values {
	return Values{
		NumDecimal:     42,
		NumOctal:       0o52,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	}
}

func TypeName(value any) string {
	return reflect.TypeOf(value).String()
}

func ConcatAsString(v Values) string {
	var b strings.Builder

	b.WriteString(strconv.Itoa(v.NumDecimal))
	b.WriteString(strconv.Itoa(v.NumOctal))
	b.WriteString(strconv.Itoa(v.NumHexadecimal))
	b.WriteString(strconv.FormatFloat(v.Pi, 'f', -1, 64))
	b.WriteString(v.Name)
	b.WriteString(strconv.FormatBool(v.IsActive))
	b.WriteString(fmt.Sprint(v.ComplexNum))
	return b.String()
}

func StringToRunes(value string) []rune {
	return []rune(value)
}

func HashRunesWithSalt(runes []rune, salt string) (string, error) {
	if len(salt) == 0 {
		return "", errors.New("hash runes must not be empty string")
	}
	if len(runes) == 0 {
		return "", errors.New("runes must not be empty")
	}

	middle := len(runes) / 2
	salted := string(runes[:middle]) + salt + string(runes[middle:])
	hash := sha256.Sum256([]byte(salted))

	return fmt.Sprintf("%x", hash), nil
}

func ProcessValues(values Values) Result {
	combined := ConcatAsString(values)
	runes := StringToRunes(combined)
	hash, err := HashRunesWithSalt(runes, salt)
	if err != nil {
		log.Fatal(err)
	}
	return Result{
		Combined: combined,
		Runes:    runes,
		Hash:     hash,
	}
}

func main() {
	values := DefaultValues()
	result := ProcessValues(values)

	fmt.Printf("numDecimal: %v, type: %s\n", values.NumDecimal, TypeName(values.NumDecimal))
	fmt.Printf("numOctal: %v, type: %s\n", values.NumOctal, TypeName(values.NumOctal))
	fmt.Printf("numHexadecimal: %v, type: %s\n", values.NumHexadecimal, TypeName(values.NumHexadecimal))
	fmt.Printf("pi: %v, type: %s\n", values.Pi, TypeName(values.Pi))
	fmt.Printf("name: %v, type: %s\n", values.Name, TypeName(values.Name))
	fmt.Printf("isActive: %v, type: %s\n", values.IsActive, TypeName(values.IsActive))
	fmt.Printf("complexNum: %v, type: %s\n", values.ComplexNum, TypeName(values.ComplexNum))
	fmt.Printf("combined string: %s\n", result.Combined)
	fmt.Printf("runes: %v\n", result.Runes)
	fmt.Printf("sha256 with salt %q: %s\n", salt, result.Hash)
}
