package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestDefaultValuesHaveExpectedTypes(t *testing.T) {
	t.Parallel()

	values := DefaultValues()
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{name: "decimal int", value: values.NumDecimal, want: "int"},
		{name: "octal int", value: values.NumOctal, want: "int"},
		{name: "hexadecimal int", value: values.NumHexadecimal, want: "int"},
		{name: "float64", value: values.Pi, want: "float64"},
		{name: "string", value: values.Name, want: "string"},
		{name: "bool", value: values.IsActive, want: "bool"},
		{name: "complex64", value: values.ComplexNum, want: "complex64"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := reflect.TypeOf(tt.value).String()
			if got != tt.want {
				t.Fatalf("TypeName(%v) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestConcatAsStringConvertsAllValuesInOrder(t *testing.T) {
	t.Parallel()

	values := Values{
		NumDecimal:     42,
		NumOctal:       0o52,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	}

	want := "4242423.14Golangtrue(1+2i)"
	if got := ConcatAsString(values); got != want {
		t.Fatalf("ConcatAsString(DefaultValues()) = %q, want %q", got, want)
	}
}

func TestStringToRunesPreservesUnicodeCharacters(t *testing.T) {
	t.Parallel()

	got := StringToRunes("Go язык")
	want := []rune{'G', 'o', ' ', 'я', 'з', 'ы', 'к'}

	if !slices.Equal(want, got) {
		t.Fatalf("StringToRunes() = %#v, want %#v", got, want)
	}
}

func TestHashRunesWithSaltAddsSaltInMiddleBeforeHashing(t *testing.T) {
	t.Parallel()

	runes := []rune("abcd")
	salted := "abgo-2024cd"
	want := fmt.Sprintf("%x", sha256.Sum256([]byte(salted)))

	got, err := HashRunesWithSalt(runes, "go-2024")
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("HashRunesWithSalt() = %q, want %q", got, want)
	}
}

func TestHashRunesWithSaltRejectsEmptyInputs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		runes   []rune
		salt    string
		wantErr string
	}{
		{
			name:    "empty salt",
			runes:   []rune("abcd"),
			salt:    "",
			wantErr: "hash runes must not be empty string",
		},
		{
			name:    "empty runes",
			runes:   nil,
			salt:    "go-2024",
			wantErr: "runes must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := HashRunesWithSalt(tt.runes, tt.salt)
			if err == nil {
				t.Fatalf("HashRunesWithSalt() error = nil, want %q", tt.wantErr)
			}
			if err.Error() != tt.wantErr {
				t.Fatalf("HashRunesWithSalt() error = %q, want %q", err.Error(), tt.wantErr)
			}
			if got != "" {
				t.Fatalf("HashRunesWithSalt() = %q, want empty string", got)
			}
		})
	}
}

func TestProcessDefaultValuesReturnsExpectedFinalResult(t *testing.T) {
	t.Parallel()

	values := Values{
		NumDecimal:     42,
		NumOctal:       0o52,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	}
	combined := "4242423.14Golangtrue(1+2i)"
	runes := []rune(combined)
	middle := len(runes) / 2
	salted := string(runes[:middle]) + "go-2024" + string(runes[middle:])
	wantHash := fmt.Sprintf("%x", sha256.Sum256([]byte(salted)))

	result := ProcessValues(values)

	if result.Combined != combined {
		t.Fatalf("Combined = %q, want %q", result.Combined, combined)
	}
	if !slices.Equal(result.Runes, runes) {
		t.Fatalf("Runes = %#v, want %#v", result.Runes, runes)
	}
	if result.Hash != wantHash {
		t.Fatalf("Hash = %q, want %q", result.Hash, wantHash)
	}
}
