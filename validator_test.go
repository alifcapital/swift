package swift_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateBic(t *testing.T) {
	testData := []struct {
		input string
		fails bool
	}{
		{
			input: "BASACHGG",
			fails: false,
		},
		{
			input: "123",
			fails: true,
		},
		{
			input: "",
			fails: true,
		},
		{
			input: "       ",
			fails: true,
		},
		{
			input: "ALDLTJ21",
			fails: false,
		},
	}

	for _, v := range testData {
		err := ValidateBic(v.input)
		if v.fails {
			assert.Error(t, err, v)
		} else {
			assert.Nil(t, err, v)
		}
	}
}

func TestValidateBicFi(t *testing.T) {
	testData := []struct {
		input string
		fails bool
	}{
		{
			input: "BASACHGG",
			fails: false,
		},
		{
			input: "123",
			fails: true,
		},
		{
			input: "",
			fails: true,
		},
		{
			input: "       ",
			fails: true,
		},
		{
			input: "ALDLTJ21",
			fails: false,
		},
	}

	for _, v := range testData {
		err := ValidateBicFi(v.input)
		if v.fails {
			assert.Error(t, err, v)
		} else {
			assert.Nil(t, err, v)
		}
	}
}

func TestValidateCreditorAccount(t *testing.T) {
	testData := []struct {
		input string
		fails bool
	}{
		{
			input: "GB3112000000001987426375",
			fails: false,
		},
		{
			input: "",
			fails: true,
		},
		{
			input: "GB3112000000001987426375GB3112000000001987426375",
			fails: true,
		},
	}

	for _, v := range testData {
		err := ValidateCreditorAccount(v.input)
		if v.fails {
			assert.Error(t, err, v)
		} else {
			assert.Nil(t, err, v)
		}
	}
}

func TestValidateCreditorName(t *testing.T) {
	testData := []struct {
		input string
		fails bool
	}{
		{
			input: "GB3112000000001987426375",
			fails: false,
		},
		{
			input: "",
			fails: true,
		},
		{
			input: "GB3112000000001987426375GB3112000000001987426375GB3112000000001987426375GB3112000000001987426375GB3112000000001987426375GB3112000000001987426375",
			fails: true,
		},
	}

	for _, v := range testData {
		err := ValidateCreditorName(v.input)
		if v.fails {
			assert.Error(t, err, v)
		} else {
			assert.Nil(t, err, v)
		}
	}
}

func TestValidateCountryCode(t *testing.T) {
	testData := []struct {
		input string
		fails bool
	}{
		{
			input: "GB",
			fails: false,
		},
		{
			input: "",
			fails: true,
		},
		{
			input: "TJK",
			fails: true,
		},
	}

	for _, v := range testData {
		err := ValidateCountryCode(v.input)
		if v.fails {
			assert.Error(t, err, v)
		} else {
			assert.Nil(t, err, v)
		}
	}
}
