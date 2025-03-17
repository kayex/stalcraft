package stalcraft_test

import (
	"net/url"
	"reflect"
	"testing"

	. "github.com/kayex/stalcraft"
)

func TestPage_Next(t *testing.T) {
	cases := []struct {
		page Page
		next Page
	}{
		{
			page: Page{Offset: 0, Limit: 20},
			next: Page{Offset: 20, Limit: 20},
		},
		{
			page: Page{Offset: 12, Limit: 20},
			next: Page{Offset: 32, Limit: 20},
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			next := c.page.Next()
			if next != c.next {
				t.Errorf("Next() = %v; want %v", next, c.next)
			}
		})
	}
}

func TestPage_Query(t *testing.T) {
	cases := []struct {
		page  Page
		query url.Values
	}{
		{
			page:  Page{Offset: 0, Limit: 20},
			query: url.Values{"offset": {"0"}, "limit": {"20"}},
		},
		{
			page:  Page{Offset: 20, Limit: 20},
			query: url.Values{"offset": {"20"}, "limit": {"20"}},
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			query := c.page.Query()
			if !reflect.DeepEqual(query, c.query) {
				t.Errorf("Query() = %v; want %v", query, c.query)
			}
		})
	}
}

func TestPage_Number(t *testing.T) {
	cases := []struct {
		page  Page
		index int
	}{
		{page: Page{Offset: 0, Limit: 20}, index: 0},
		{page: Page{Offset: 320, Limit: 100}, index: 3},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			index := c.page.Number()
			if index != c.index {
				t.Errorf("Number() = %v; want %v", index, c.index)
			}
		})
	}
}
