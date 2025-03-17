package stalcraft

import (
	"net/url"
	"strconv"
)

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

type LotSort string

const (
	LotSortTimeCreated  LotSort = "time_created"
	LotSortTimeLeft     LotSort = "time_left"
	LotSortCurrentPrice LotSort = "current_price"
	LotSortBuyoutPrice  LotSort = "buyout_price"
)

const MaxLotPageLimit = 200

type Page struct {
	Offset int
	Limit  int
}

func (p Page) Query() url.Values {
	values := url.Values{}
	values.Add("limit", strconv.Itoa(p.Limit))
	values.Add("offset", strconv.Itoa(p.Offset))
	return values
}

func (p Page) Next() Page {
	return Page{
		Offset: p.Offset + p.Limit,
		Limit:  p.Limit,
	}
}

func (p Page) LastIndex() int {
	return p.Offset + p.Limit
}

func (p Page) Number() int {
	return (p.Offset / p.Limit) + 1
}

type ActiveLotPage struct {
	LotPage
	Order Order
	Sort  LotSort
}

type LotPage struct {
	Page
	Additional bool
}

func (p LotPage) Next() LotPage {
	return LotPage{
		Page:       p.Page.Next(),
		Additional: p.Additional,
	}
}

func (p LotPage) Query() url.Values {
	values := p.Page.Query()
	if p.Additional {
		values.Add("additional", "true")
	}
	return values
}

func (p ActiveLotPage) Query() url.Values {
	values := p.LotPage.Query()
	values.Add("order", string(p.Order))
	values.Add("sort", string(p.Sort))
	return values
}

func (p ActiveLotPage) Next() ActiveLotPage {
	return ActiveLotPage{
		LotPage: p.LotPage.Next(),
		Order:   p.Order,
		Sort:    p.Sort,
	}
}
