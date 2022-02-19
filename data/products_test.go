package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "sdfdsa",
		Price: 1,
		SKU:   "abcd-efgh-jkl",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
