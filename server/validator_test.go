package server

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	type result struct {
		v ValidationResponse
		e bool
	}

	testCases := []struct {
		input    []byte
		expected result
	}{
		{
			[]byte(`{"iban":"123","valid": false}`),
			result{
				v: ValidationResponse{Iban: "123", Valid: false},
				e: false,
			},
		},
		{
			[]byte(`{"iban":"IT123456","valid": true}`),
			result{
				v: ValidationResponse{Iban: "IT123456", Valid: true},
				e: false,
			},
		},
		{
			[]byte(`{IT123456", true}`),
			result{
				v: ValidationResponse{},
				e: true,
			},
		},
	}

	for _, test := range testCases {
		r := bytes.NewReader(test.input)

		observed, err := decode(r)
		if err != nil {
			if !test.expected.e {
				t.Fatalf("failed to run the test: %s", err)
			}
		} else {
			if test.expected.e {
				t.Errorf("for input '%s', expected error to be '%t', observed '%t'",
					test.input, test.expected.e, false)
			}
		}

		if observed.Iban != test.expected.v.Iban {
			t.Errorf("for input '%s', expected iban to be '%s', observed '%s'",
				test.input, test.expected.v.Iban, observed.Iban)
		}

		if observed.Valid != test.expected.v.Valid {
			t.Errorf("for input '%s', expected valid to be '%t', observed '%t'",
				test.input, test.expected.v.Valid, observed.Valid)
		}

	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := []byte(`{"iban":"IT123456","valid": true}`)
		r := bytes.NewReader(d)
		decode(r)
	}
}

func TestSanitize(t *testing.T) {
	testCases := []struct {
		input, expected string
	}{
		{"", ""},
		{"£", ""},
		{"/", ""},
		{"IT82347293498724", "IT82347293498724"},
		{"I#T8234729349/1111", "IT82347293491111"},
		{")(IT8234729|'349/1111", "IT82347293491111"},
	}

	for _, test := range testCases {
		observed, err := sanitize(test.input)
		if err != nil {
			t.Fatalf("failed to run the test: %s", err)
		}

		if observed != test.expected {
			t.Errorf("for input '%s', expected '%s', observed '%s'",
				test.input, test.expected, observed)
		}
	}
}

func BenchmarkSanitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sanitize("IT8234729349/1111")
	}
}
