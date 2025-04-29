package stalcraft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	DemoDomain       = "dapi.stalcraft.net"
	ProductionDomain = "eapi.stalcraft.net"
)

const (
	RegionNA  = "NA"
	RegionEU  = "EU"
	RegionSEA = "SEA"
	RegionRU  = "RU"
)

const (
	DemoAppAccessToken  = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwibmJmIjoxNjczNzk3ODM4LCJleHAiOjQ4MjczOTc4MzgsImlhdCI6MTY3Mzc5NzgzOCwianRpIjoiYXhwbzAzenJwZWxkMHY5dDgzdzc1N2x6ajl1MmdyeHVodXVlb2xsZ3M2dml1YjVva3NwZTJ3eGFrdjJ1eWZxaDU5ZDE2ZTNlN2FqdW16Z3gifQ.ZNSsvwAX72xT5BzLqqYABuH2FGbOlfiXMK5aYO1H5llG51ZjcPvOYBDRR4HUoPZVLFY8jyFUsEXNM7SYz8qL9ePmLjJl6pib8FEtqVPmf9ldXvKkbaaaSp4KkJzsIEMY_Z5PejB2Vr-q-cL13KPgnLGUaSW-2X_sHPN7VZJNMjRgjw4mPiRZTe4CEpQq0BEcPrG6OLtU5qlZ6mLDJBjN2xtK0DI6xgmYriw_5qW1mj1nqF_ewtUiQ1KTVhDgXnaNUdkGsggAGqyicTei0td6DTKtnl3noD5VkipWn_CwSqb2Mhm16I9BPfX_d5ARzWrnrwPRUf6PA_7LipNU6KkkW0mhZfmwEPTm_sXPus0mHPENoVZArdFT3L5sOYBcpqwvVIEtxRUTdcsKp-y-gSzao5muoyPVoCc2LEeHEWx0cIi9spsZ46SPRQpN4baVFp7y5rp5pjRsBKHQYUJ0lTmh1_vyfzOzbtNN2v6W_5w9JTLrN1U6fhmifvKHppFSEqD6DameL1TC59kpIdufRkEU9HE4O-ErEf1GuJFRx-Dew6XDvb_ExhvEqcw31yNvKzpVqLYJfLazqn6tUbVuAiPwpy6rP9tYO2taT1vj5TGn_vxwDu9zoLWe796tFMPS-kmbCglxB5C9L4EbpfWNbWxYjUkTvjT2Ml9OnrB0UbYo1jI"
	DemoUserAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwic3ViIjoiMSIsIm5iZiI6MTY3Mzc5NzgzOCwiZXhwIjo0ODI3Mzk3ODM4LCJpYXQiOjE2NzM3OTc4MzgsImp0aSI6IjJlamRwOG54a3A1djRnZWdhbWVyeWlkMW5ic24zZDhpZ2oyejgzem1vMDYzNjNoaXFkNWhwOTY1MHZwdWh4OXEybXBmd2hnbnUxNHR5cmp2In0.Ocw4CzkkuenkAOjkAR1RuFgLqix7VJ-8vWVS3KAJ1T3SgIWJG145xqG2qms99knu5azn_oaoeyMOXhyG_fuMQFGOju317GiS6pAXAFGOKvxcUCfdpFcEHO6TWGM8191-tlfV-0rAqCi62gprKyr-SrUG3nUJhv6XKegja_vYVujRVx0ouAaDvDKawiOssG5If_hXGhdhnmb3_7onnIc4hFsm4i9QVkWXe8GO6OsS999ZIX0ClNhTk2kKKTl2dDVIiKha_HB1aghm_LOYoRgb3i3B_DH4UO312rHYR5I4qO43c8x-TW7NwovItDSzhiCmcxZuUUeAUF3yFr5ovaR4fMj1LEy3y3V2piQDKPwmBOpI9S6OzWUIBJYcRYlT2HIrWCRc0YvM7AOGoxcH2Gf4ncqcF_M8fw7IMKf3pdnuxf1EbdEpzOapBD1Pw065em-U8PN4LVzw9lhIHx_Yj69qaFEx7Bhw3BCwsrx-o9hgg7T1TOV6kF11YfR99lIuj9z96XBLg5ipt-M_j7nHRoHWhM0Rc6uLIKPg0In0xYkybSfWG6v3Hs6kwgB7wkqpXpoVQltJvlqjtlf9Pp4zmkqlWQHx9as4xsgoTAQyCgaC0kisICNC58_g3QrJAfoFXW68x-OHlRKCAPqoR9V-0cVs-B83szaFmsEGegAttFLlDhE"
)

type Client struct {
	client *http.Client

	domain      string
	region      string
	accessToken string
}

type ClientOption func(*Client)

func NewClient(region, accessToken string, opt ...ClientOption) (*Client, error) {
	c := &Client{
		region:      region,
		accessToken: accessToken,
	}
	for _, o := range opt {
		o(c)
	}
	if c.region == "" {
		return nil, errors.New("no region provided")
	}
	if c.accessToken == "" {
		return nil, errors.New("no access token provided")
	}
	if c.client == nil {
		c.client = http.DefaultClient
	}
	if c.domain == "" {
		c.domain = ProductionDomain
	}

	return c, nil
}

func WithHTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.client = client
	}
}

func WithDemoDomain() func(*Client) {
	return func(c *Client) {
		c.domain = DemoDomain
	}
}

func roundtrip[R any](ctx context.Context, c *Client, path string) (R, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path)
	if err != nil {
		return *new(R), fmt.Errorf("creating request: %v", err)
	}

	for {
		res, err := do[R](c.client, req)
		if err != nil {
			var erl *RateLimitError
			if errors.As(err, &erl) {
				next := time.UnixMilli(erl.Reset).Sub(time.Now())
				<-time.After(next)
				continue
			} else {
				return *new(R), err
			}
		}
		return res, nil
	}
}

func do[R any](c *http.Client, r *http.Request) (R, error) {
	res, err := c.Do(r)
	if err != nil {
		return *new(R), fmt.Errorf("sending request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			limit, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Limit"))
			if err != nil {
				return *new(R), fmt.Errorf("parsing X-Ratelimit-Limit header: %v", err)
			}
			remaining, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Remaining"))
			if err != nil {
				return *new(R), fmt.Errorf("parsing X-Ratelimit-Remaining header: %v", err)
			}
			reset, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Reset"))
			if err != nil {
				return *new(R), fmt.Errorf("parsing X-Ratelimit-Reset header: %v", err)
			}
			return *new(R), &RateLimitError{
				Limit:     limit,
				Remaining: remaining,
				Reset:     int64(reset),
			}
		}

		var e *Err
		err = json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return *new(R), fmt.Errorf("decoding error response body: %v", err)
		}
		return *new(R), e
	}

	var resp R
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return *new(R), fmt.Errorf("decoding response body: %v", err)
	}
	return resp, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string) (*http.Request, error) {
	uri := fmt.Sprintf("https://%s/%s/%s", c.domain, c.region, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequestWithContext(ctx, method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	req.Header.Set("Accept", "application/json")
	return req, nil
}

type Err struct {
	Title   string         `json:"title"`
	Status  int            `json:"status"`
	Details map[string]any `json:"details"`
}

func (e *Err) Error() string {
	msg := fmt.Sprintf("%s (status %d)", e.Title, e.Status)
	if len(e.Details) > 0 {
		msg += fmt.Sprintf(": %s", e.Details)
	}
	return msg
}

type RateLimitError struct {
	Limit     int
	Remaining int
	Reset     int64
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit reached, limit %d, remaining %d, reset %d", e.Limit, e.Remaining, e.Reset)
}
