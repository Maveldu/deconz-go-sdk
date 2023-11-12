package sdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/maxatome/go-testdeep/td"

	"github.com/Maveldu/deconz-go-sdk"
)

func TestGet(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.DeactivateAndReset)

	client := sdk.NewClient("https://test.bidule")

	httpmock.RegisterResponder("GET", "https://test.bidule/lel", httpmock.NewStringResponder(200, `{"lul": 12}`))
	httpmock.RegisterResponder("GET", "https://test.bidule/lel?lal=13", httpmock.NewStringResponder(200, `{"lul": 14}`))
	httpmock.RegisterResponder("GET", "https://test.bidule/lil", httpmock.NewStringResponder(500, "nope"))

	type lul struct {
		Lul int64 `json:"lul"`
	}
	type lel struct {
		Lal int `url:"lal"`
	}
	result := &lul{}

	td.CmpNil(t, client.Get("/lel", nil, result))
	td.Cmp(t, result, &lul{Lul: 12})

	td.CmpNil(t, client.Get("/lel", &lel{Lal: 13}, result))
	td.Cmp(t, result, &lul{Lul: 14})

	td.Cmp(t, client.Get("/lil", nil, result), &sdk.Error{Status: "500", StatusCode: 500})
}
