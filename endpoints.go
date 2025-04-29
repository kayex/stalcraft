package stalcraft

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListRegions returns a list of regions that can be accessed via the API.
//
// https://eapi.stalcraft.net/reference#/paths/regions/get
func (c *Client) ListRegions(ctx context.Context) ([]RegionInfo, error) {
	uri := fmt.Sprintf("https://%s/regions", c.domain)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	return do[[]RegionInfo](c.client, req)
}

// EmissionStatus returns information about current emission, if any, and recorded time of the previous one.
//
// https://eapi.stalcraft.net/reference#/paths/region--emission/get
func (c *Client) EmissionStatus(ctx context.Context) (*EmissionResponse, error) {
	return roundtrip[*EmissionResponse](ctx, c, "/emission")
}

// FriendList returns list of character names who are friends with the specified character. Requires
// user authentication.
//
// https://eapi.stalcraft.net/reference#/paths/region--friends--character/get
func (c *Client) FriendList(ctx context.Context, name string) ([]string, error) {
	return roundtrip[[]string](ctx, c, fmt.Sprintf("/character/by-name/%s/profile", url.PathEscape(name)))
}

// ItemPriceHistory returns history of prices for lots of the given item which were bought from auction.
// Prices are sorted in descending order by recorded time of purchase.
//
// https://eapi.stalcraft.net/reference#/paths/region--auction--item--history/get
func (c *Client) ItemPriceHistory(ctx context.Context, itemID string, page LotPage) (*PricesListing, error) {
	// TODO: Debug, total >= 0 but prices array is empty.
	return roundtrip[*PricesListing](ctx, c, fmt.Sprintf("/auction/%s/lots?%s", url.PathEscape(itemID), page.Query().Encode()))
}

// ActiveItemLots returns a list of currently active lots on auction for the given item.
//
// https://eapi.stalcraft.net/reference#/paths/region--auction--item--lots/get
func (c *Client) ActiveItemLots(ctx context.Context, itemID string, page ActiveLotPage) (*LotListing, error) {
	return roundtrip[*LotListing](ctx, c, fmt.Sprintf("/auction/%s/lots?%s", url.PathEscape(itemID), page.Query().Encode()))
}

// CharacterProfile returns information about a player's profile.
//
// https://eapi.stalcraft.net/reference#/paths/region--character-by-name--character--profile/get
func (c *Client) CharacterProfile(ctx context.Context, name string) (*CharacterProfileData, error) {
	// TODO: Check if this actually works. Returns 404 for all the demo API user names.
	return roundtrip[*CharacterProfileData](ctx, c, fmt.Sprintf("/character/by-name/%s/profile", url.PathEscape(name)))
}

// ListCharacters returns a list of characters created by the authenticated user. Requires user authentication.
//
// https://eapi.stalcraft.net/reference#/paths/region--characters/get
func (c *Client) ListCharacters(ctx context.Context) ([]FullCharacterInfo, error) {
	return roundtrip[[]FullCharacterInfo](ctx, c, "/characters")
}

// ClanInformation returns information about the given clan.
//
// https://eapi.stalcraft.net/reference#/paths/region--clan--clan-id--info/get
func (c *Client) ClanInformation(ctx context.Context, region, clanID string) (*ClanInfo, error) {
	return roundtrip[*ClanInfo](ctx, c, fmt.Sprintf("/clan/%s/info", url.PathEscape(clanID)))
}

// ClanMembers returns list of members in the given clan. Requires user authentication.
// Can only be used when authenticated user has at least one character in the clan.
//
// https://eapi.stalcraft.net/reference#/paths/region--clan--clan-id--members/get
func (c *Client) ClanMembers(ctx context.Context, region, clanID string) ([]ClanMember, error) {
	return roundtrip[[]ClanMember](ctx, c, fmt.Sprintf("/clan/%s/members", url.PathEscape(clanID)))
}

// ListClans returns all clans which are currently registered in the game in the specified region.
//
// https://eapi.stalcraft.net/reference#/paths/region--clans/get
func (c *Client) ListClans(ctx context.Context, region string) (*ClansListResponse, error) {
	return roundtrip[*ClansListResponse](ctx, c, "/clans")
}
