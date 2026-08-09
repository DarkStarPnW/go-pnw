package pnw

import "context"

// ----------------------------------------------------------------------------
// Filter types
// ----------------------------------------------------------------------------

// NationFilter filters the nations query.
type NationFilter struct {
	ID           []int
	AllianceID   []int
	Color        []Color
	CreatedAfter string   // RFC 3339 / ISO 8601 DateTime string
	LeaderName   []string // partial match, case-insensitive
	NationName   []string // partial match, case-insensitive
	VMode        *bool    // vacation mode filter
	MinScore     *float64
	MaxScore     *float64
	MinCities    *int
	MaxCities    *int
	First        int // max per page (default/max 500)
	Page         int
}

func (f NationFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.AllianceID) > 0 {
		v["alliance_id"] = f.AllianceID
	}
	if len(f.Color) > 0 {
		v["color"] = f.Color
	}
	if f.CreatedAfter != "" {
		v["created_after"] = f.CreatedAfter
	}
	if len(f.LeaderName) > 0 {
		v["leader_name"] = f.LeaderName
	}
	if len(f.NationName) > 0 {
		v["nation_name"] = f.NationName
	}
	if f.VMode != nil {
		v["vmode"] = *f.VMode
	}
	if f.MinScore != nil {
		v["min_score"] = *f.MinScore
	}
	if f.MaxScore != nil {
		v["max_score"] = *f.MaxScore
	}
	if f.MinCities != nil {
		v["min_cities"] = *f.MinCities
	}
	if f.MaxCities != nil {
		v["max_cities"] = *f.MaxCities
	}
	if f.First > 0 {
		v["first"] = f.First
	}
	if f.Page > 0 {
		v["page"] = f.Page
	}
	return v
}

// AllianceFilter filters the alliances query.
type AllianceFilter struct {
	ID    []int
	First int
	Page  int
}

func (f AllianceFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if f.First > 0 {
		v["first"] = f.First
	}
	if f.Page > 0 {
		v["page"] = f.Page
	}
	return v
}

// WarFilter filters the wars query.
type WarFilter struct {
	ID         []int
	NationID   []int
	AllianceID []int
	Active     *bool
	DaysAgo    *int // max 14; default 1 when Active not set
}

func (f WarFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if f.Active != nil {
		v["active"] = *f.Active
	}
	if f.DaysAgo != nil {
		v["days_ago"] = *f.DaysAgo
	}
	return v
}

// TradeFilter filters the trades query.
type TradeFilter struct {
	ID            []int
	NationID      []int
	Accepted      *bool
	BuyOrSell     string // "buy" or "sell"
	OfferResource string
	Type          TradeType
}

func (f TradeFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.NationID) > 0 {
		v["nation_id"] = f.NationID
	}
	if f.Accepted != nil {
		v["accepted"] = *f.Accepted
	}
	if f.BuyOrSell != "" {
		v["buy_or_sell"] = f.BuyOrSell
	}
	if f.OfferResource != "" {
		v["offer_resource"] = f.OfferResource
	}
	if f.Type != "" {
		v["type"] = f.Type
	}
	return v
}

// BountyFilter filters the bounties query.
type BountyFilter struct {
	ID       []int
	NationID []int
	Type     WarType
}

func (f BountyFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.NationID) > 0 {
		v["nation_id"] = f.NationID
	}
	if f.Type != "" {
		v["type"] = f.Type
	}
	return v
}

// CityFilter filters the cities query.
type CityFilter struct {
	ID       []int
	NationID []int
	First    int
	Page     int
}

func (f CityFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.NationID) > 0 {
		v["nation_id"] = f.NationID
	}
	if f.First > 0 {
		v["first"] = f.First
	}
	if f.Page > 0 {
		v["page"] = f.Page
	}
	return v
}

// WarAttackFilter filters the warattacks query.
type WarAttackFilter struct {
	ID    []int
	WarID []int
}

func (f WarAttackFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.WarID) > 0 {
		v["war_id"] = f.WarID
	}
	return v
}

// BankRecFilter filters the bankrecs query. The API filters by the parties on
// the record rather than by nation or alliance, so use SenderID / ReceiverID.
type BankRecFilter struct {
	ID         []int
	SenderID   []int
	ReceiverID []int
	// MinID returns only records newer than this ID.
	MinID int
	First int
	Page  int
}

func (f BankRecFilter) toVars() map[string]any {
	v := map[string]any{}
	if len(f.ID) > 0 {
		v["id"] = f.ID
	}
	if len(f.SenderID) > 0 {
		v["sid"] = f.SenderID
	}
	if len(f.ReceiverID) > 0 {
		v["rid"] = f.ReceiverID
	}
	if f.MinID > 0 {
		v["min_id"] = f.MinID
	}
	if f.First > 0 {
		v["first"] = f.First
	}
	if f.Page > 0 {
		v["page"] = f.Page
	}
	return v
}

// ----------------------------------------------------------------------------
// Query methods
// ----------------------------------------------------------------------------

// Nations returns a paginated list of nations matching the filter.
// Fields controls which Nation fields are fetched; pass nil to use defaults.
func (c *Client) Nations(ctx context.Context, filter NationFilter, fields ...string) (*NationPaginator, error) {
	if len(fields) == 0 {
		fields = defaultNationFields
	}
	q := `query Nations(
		$id: [Int], $alliance_id: [Int], $color: [String],
		$created_after: DateTime, $leader_name: [String], $nation_name: [String],
		$vmode: Boolean, $min_score: Float, $max_score: Float,
		$min_cities: Int, $max_cities: Int, $first: Int, $page: Int
	) {
		nations(
			id: $id, alliance_id: $alliance_id, color: $color,
			created_after: $created_after, leader_name: $leader_name, nation_name: $nation_name,
			vmode: $vmode, min_score: $min_score, max_score: $max_score,
			min_cities: $min_cities, max_cities: $max_cities, first: $first, page: $page
		) {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		Nations NationPaginator `json:"nations"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return &out.Nations, nil
}

// Alliances returns a paginated list of alliances matching the filter.
func (c *Client) Alliances(ctx context.Context, filter AllianceFilter, fields ...string) (*AlliancePaginator, error) {
	if len(fields) == 0 {
		fields = defaultAllianceFields
	}
	q := `query Alliances($id: [Int], $first: Int, $page: Int) {
		alliances(id: $id, first: $first, page: $page) {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		Alliances AlliancePaginator `json:"alliances"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return &out.Alliances, nil
}

// WarPaginator wraps paginated War results.
type WarPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []War         `json:"data"`
}

// Wars returns a paginated list of wars matching the filter.
func (c *Client) Wars(ctx context.Context, filter WarFilter, fields ...string) (*WarPaginator, error) {
	if len(fields) == 0 {
		fields = defaultWarFields
	}
	q := `query Wars(
		$id: [Int], $nation_id: [Int], $alliance_id: [Int],
		$active: Boolean, $days_ago: Int
	) {
		wars(
			id: $id, nation_id: $nation_id, alliance_id: $alliance_id,
			active: $active, days_ago: $days_ago
		) {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		Wars WarPaginator `json:"wars"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return &out.Wars, nil
}

// Trades returns trades matching the filter.
func (c *Client) Trades(ctx context.Context, filter TradeFilter, fields ...string) ([]Trade, error) {
	if len(fields) == 0 {
		fields = defaultTradeFields
	}
	q := `query Trades(
		$id: [Int], $nation_id: [Int], $accepted: Boolean,
		$buy_or_sell: String, $offer_resource: String, $type: String
	) {
		trades(
			id: $id, nation_id: $nation_id, accepted: $accepted,
			buy_or_sell: $buy_or_sell, offer_resource: $offer_resource, type: $type
		) { ` + joinFields(fields) + ` }
	}`

	var out struct {
		Trades []Trade `json:"trades"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return out.Trades, nil
}

// TradepricePaginator wraps paginated Tradeprice results.
type TradepricePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Tradeprice  `json:"data"`
}

// Tradeprices returns the current market resource prices.
func (c *Client) Tradeprices(ctx context.Context, fields ...string) (*TradepricePaginator, error) {
	if len(fields) == 0 {
		fields = defaultTradepriceFiels
	}
	q := `query {
		tradeprices {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		Tradeprices TradepricePaginator `json:"tradeprices"`
	}
	if err := c.do(ctx, q, nil, &out); err != nil {
		return nil, err
	}
	return &out.Tradeprices, nil
}

// TopTradeInfo returns the market summary, including the average price of each resource.
func (c *Client) TopTradeInfo(ctx context.Context) (*TopTradeInfo, error) {
	q := `query {
		top_trade_info {
			market_index
			resources { resource average_price }
		}
	}`

	var out struct {
		TopTradeInfo TopTradeInfo `json:"top_trade_info"`
	}
	if err := c.do(ctx, q, nil, &out); err != nil {
		return nil, err
	}
	return &out.TopTradeInfo, nil
}

// Treasures returns all active treasures.
func (c *Client) Treasures(ctx context.Context, fields ...string) ([]Treasure, error) {
	if len(fields) == 0 {
		fields = defaultTreasureFields
	}
	q := `query { treasures { ` + joinFields(fields) + ` } }`

	var out struct {
		Treasures []Treasure `json:"treasures"`
	}
	if err := c.do(ctx, q, nil, &out); err != nil {
		return nil, err
	}
	return out.Treasures, nil
}

// Colors returns all game color blocs and their trade bonuses.
func (c *Client) Colors(ctx context.Context) ([]ColorInfo, error) {
	q := `query { colors { color bloc_name turn_bonus } }`

	var out struct {
		Colors []ColorInfo `json:"colors"`
	}
	if err := c.do(ctx, q, nil, &out); err != nil {
		return nil, err
	}
	return out.Colors, nil
}

// Bounties returns bounties matching the filter.
func (c *Client) Bounties(ctx context.Context, filter BountyFilter, fields ...string) ([]Bounty, error) {
	if len(fields) == 0 {
		fields = defaultBountyFields
	}
	q := `query Bounties($id: [Int], $nation_id: [Int], $type: String) {
		bounties(id: $id, nation_id: $nation_id, type: $type) {
			` + joinFields(fields) + `
		}
	}`

	var out struct {
		Bounties []Bounty `json:"bounties"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return out.Bounties, nil
}

// Cities returns a paginated list of cities matching the filter.
func (c *Client) Cities(ctx context.Context, filter CityFilter, fields ...string) (*CityPaginator, error) {
	if len(fields) == 0 {
		fields = defaultCityFields
	}
	q := `query Cities($id: [Int], $nation_id: [Int], $first: Int, $page: Int) {
		cities(id: $id, nation_id: $nation_id, first: $first, page: $page) {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		Cities CityPaginator `json:"cities"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return &out.Cities, nil
}

// WarAttacks returns attacks matching the filter.
func (c *Client) WarAttacks(ctx context.Context, filter WarAttackFilter, fields ...string) ([]WarAttack, error) {
	if len(fields) == 0 {
		fields = defaultWarAttackFields
	}
	q := `query WarAttacks($id: [Int], $war_id: [Int]) {
		warattacks(id: $id, war_id: $war_id) { ` + joinFields(fields) + ` }
	}`

	var out struct {
		WarAttacks []WarAttack `json:"warattacks"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return out.WarAttacks, nil
}

// BankRecPaginator wraps paginated BankRec results.
type BankRecPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []BankRec     `json:"data"`
}

// BankRecs returns bank records matching the filter.
func (c *Client) BankRecs(ctx context.Context, filter BankRecFilter, fields ...string) ([]BankRec, error) {
	if len(fields) == 0 {
		fields = defaultBankRecFields
	}
	q := `query BankRecs(
		$id: [Int], $sid: [Int], $rid: [Int],
		$min_id: Int, $first: Int, $page: Int
	) {
		bankrecs(
			id: $id, sid: $sid, rid: $rid,
			min_id: $min_id, first: $first, page: $page
		) {
			paginatorInfo { ` + paginatorInfoFields + ` }
			data { ` + joinFields(fields) + ` }
		}
	}`

	var out struct {
		BankRecs BankRecPaginator `json:"bankrecs"`
	}
	if err := c.do(ctx, q, filter.toVars(), &out); err != nil {
		return nil, err
	}
	return out.BankRecs.Data, nil
}

// Query executes a raw GraphQL query and decodes the top-level data into out.
// Use this when the typed helpers don't cover your use case.
func (c *Client) Query(ctx context.Context, query string, variables map[string]any, out any) error {
	return c.do(ctx, query, variables, out)
}

// ----------------------------------------------------------------------------
// Default field sets
// ----------------------------------------------------------------------------

const paginatorInfoFields = `count currentPage firstItem hasMorePages lastItem lastPage perPage total`

func joinFields(fields []string) string {
	result := ""
	for i, f := range fields {
		if i > 0 {
			result += " "
		}
		result += f
	}
	return result
}

var defaultNationFields = []string{
	"id", "alliance_id", "alliance_position",
	"nation_name", "leader_name", "continent",
	"wars_won", "wars_lost", "score", "color",
	"num_cities", "soldiers", "tanks", "aircraft",
	"ships", "missiles", "nukes",
	"vacation_mode_turns", "beige_turns",
	"domestic_policy", "government_type",
	"last_active", "date",
}

var defaultAllianceFields = []string{
	"id", "name", "acronym", "score", "color",
	"flag", "forum_link", "discord_link", "date",
}

var defaultWarFields = []string{
	"id", "date", "reason", "war_type",
	"att_id", "att_alliance_id", "def_id", "def_alliance_id",
	"att_resistance", "def_resistance",
	"att_peace", "def_peace",
	"winner_id",
}

var defaultTradeFields = []string{
	"id", "type", "date",
	"sender_id", "receiver_id",
	"offer_resource", "offer_amount",
	"buy_or_sell", "price_per_unit",
	"accepted", "date_accepted",
}

var defaultTradepriceFiels = []string{
	"id", "date", "tax_id",
	"coal", "oil", "uranium", "iron", "bauxite", "lead",
	"gasoline", "munitions", "steel", "aluminum", "food", "credits",
}

var defaultTreasureFields = []string{
	"name", "color", "continent", "bonus", "spawndate", "nation_id",
}

var defaultBountyFields = []string{
	"id", "date", "nation_id", "amount", "type",
}

var defaultCityFields = []string{
	"id", "nation_id", "name", "date",
	"infrastructure", "land", "powered",
	"coal_power", "oil_power", "nuclear_power", "wind_power",
	"coal_mine", "oil_well", "uranium_mine", "iron_mine",
	"bauxite_mine", "lead_mine", "farm",
	"oil_refinery", "aluminum_refinery", "steel_mill", "munitions_factory",
	"barracks", "factory", "hangar", "drydock",
}

var defaultWarAttackFields = []string{
	"id", "date", "att_id", "def_id", "type", "war_id",
	"att_casualties", "def_casualties",
	"infra_destroyed", "infra_destroyed_value",
	"money_stolen", "success",
}

var defaultBankRecFields = []string{
	"id", "date", "tax_id",
	"sid", "stype", "rid", "rtype",
	"note",
	"money", "coal", "oil", "uranium", "iron", "bauxite", "lead",
	"gasoline", "munitions", "steel", "aluminum", "food",
}
