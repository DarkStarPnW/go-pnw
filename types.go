package pnw

// ----------------------------------------------------------------------------
// Enumerations
// ----------------------------------------------------------------------------

type AlliancePosition string

const (
	AlliancePositionNoAlliance AlliancePosition = "NOALLIANCE"
	AlliancePositionApplicant  AlliancePosition = "APPLICANT"
	AlliancePositionMember     AlliancePosition = "MEMBER"
	AlliancePositionOfficer    AlliancePosition = "OFFICER"
	AlliancePositionHeir       AlliancePosition = "HEIR"
	AlliancePositionLeader     AlliancePosition = "LEADER"
)

type WarType string

const (
	WarTypeOrdinary  WarType = "ORDINARY"
	WarTypeRaid      WarType = "RAID"
	WarTypeAttrition WarType = "ATTRITION"
)

type AttackType string

const (
	AttackTypeGround      AttackType = "GROUND"
	AttackTypeAirInfra    AttackType = "AIRVINFRA"
	AttackTypeAirSoldiers AttackType = "AIRVSOLDIERS"
	AttackTypeAirTanks    AttackType = "AIRVTANKS"
	AttackTypeAirShips    AttackType = "AIRVSHIPS"
	AttackTypeAirAir      AttackType = "AIRVAIR"
	AttackTypeNaval       AttackType = "NAVAL"
	AttackTypeMissile     AttackType = "MISSILE"
	AttackTypeMissileHQ   AttackType = "MISSILEFAIL"
	AttackTypeNuke        AttackType = "NUKE"
	AttackTypeNukeFail    AttackType = "NUKEFAIL"
	AttackTypeFortify     AttackType = "FORTIFY"
	AttackTypePeace       AttackType = "PEACE"
	AttackTypeAllyPeace   AttackType = "ALLIED"
	AttackTypeNIT         AttackType = "IT" // Incognito
	AttackTypeSurrender   AttackType = "SURRENDER"
)

type TradeType string

const (
	TradeTypePersonal TradeType = "PERSONAL"
	TradeTypeAlliance TradeType = "ALLIANCE"
	TradeTypeGlobal   TradeType = "GLOBAL"
)

type Color string

const (
	ColorAqua   Color = "aqua"
	ColorBlack  Color = "black"
	ColorBlue   Color = "blue"
	ColorBrown  Color = "brown"
	ColorGreen  Color = "green"
	ColorLime   Color = "lime"
	ColorMaroon Color = "maroon"
	ColorOlive  Color = "olive"
	ColorOrange Color = "orange"
	ColorPink   Color = "pink"
	ColorPurple Color = "purple"
	ColorRed    Color = "red"
	ColorTeal   Color = "teal"
	ColorWhite  Color = "white"
	ColorYellow Color = "yellow"
	ColorBeige  Color = "beige"
	ColorGray   Color = "gray"
)

// ----------------------------------------------------------------------------
// Pagination
// ----------------------------------------------------------------------------

// PaginatorInfo contains pagination metadata returned by paginated queries.
type PaginatorInfo struct {
	Count        int  `json:"count"`
	CurrentPage  int  `json:"currentPage"`
	FirstItem    int  `json:"firstItem"`
	HasMorePages bool `json:"hasMorePages"`
	LastItem     int  `json:"lastItem"`
	LastPage     int  `json:"lastPage"`
	PerPage      int  `json:"perPage"`
	Total        int  `json:"total"`
}

// ----------------------------------------------------------------------------
// Nation
// ----------------------------------------------------------------------------

type Nation struct {
	// Discord is the username set on the nation's account page; DiscordID is the
	// numeric account ID when the player has linked their account.
	Discord          string           `json:"discord"`
	DiscordID        string           `json:"discord_id"`
	ID               ID               `json:"id"`
	AllianceID       ID               `json:"alliance_id"`
	AlliancePosition AlliancePosition `json:"alliance_position"`
	NationName       string           `json:"nation_name"`
	LeaderName       string           `json:"leader_name"`
	ContinentName    string           `json:"continent"`
	WarsWon          int              `json:"wars_won"`
	WarsLost         int              `json:"wars_lost"`
	Score            float64          `json:"score"`
	Color            Color            `json:"color"`
	NumCities        int              `json:"num_cities"`
	Soldiers         int              `json:"soldiers"`
	Tanks            int              `json:"tanks"`
	Aircraft         int              `json:"aircraft"`
	Ships            int              `json:"ships"`
	Missiles         int              `json:"missiles"`
	Nukes            int              `json:"nukes"`
	Spies            int              `json:"spies"`
	VacationMode     int              `json:"vacation_mode_turns"`
	Beige            int              `json:"beige_turns"`
	DomesticPolicy   string           `json:"domestic_policy"`
	GovernmentType   string           `json:"government_type"`
	Religion         string           `json:"religion"`
	Population       int              `json:"population"`
	WarPolicy        string           `json:"war_policy"`
	Espionage        bool             `json:"espionage_available"`
	LastActive       string           `json:"last_active"`
	Date             string           `json:"date"`

	// Resources (only returned when querying your own nation or alliance bank)
	Money     float64 `json:"money"`
	Coal      float64 `json:"coal"`
	Oil       float64 `json:"oil"`
	Uranium   float64 `json:"uranium"`
	Iron      float64 `json:"iron"`
	Bauxite   float64 `json:"bauxite"`
	Lead      float64 `json:"lead"`
	Gasoline  float64 `json:"gasoline"`
	Munitions float64 `json:"munitions"`
	Steel     float64 `json:"steel"`
	Aluminum  float64 `json:"aluminum"`
	Food      float64 `json:"food"`

	// Relations
	Alliance *Alliance `json:"alliance,omitempty"`
	Cities   []City    `json:"cities,omitempty"`
	Wars     []War     `json:"wars,omitempty"`
	Bankrecs []BankRec `json:"bankrecs,omitempty"`
	Taxes    []BankRec `json:"taxrecs,omitempty"`
}

// NationPaginator wraps paginated Nation results.
type NationPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Nation      `json:"data"`
}

// ----------------------------------------------------------------------------
// Alliance
// ----------------------------------------------------------------------------

type Alliance struct {
	ID         ID      `json:"id"`
	Name       string  `json:"name"`
	Acronym    string  `json:"acronym"`
	Score      float64 `json:"score"`
	Color      Color   `json:"color"`
	Flag       string  `json:"flag"`
	ForumURL   string  `json:"forum_link"`
	IRCURL     string  `json:"irc_link"`
	DiscordURL string  `json:"discord_link"`
	WikiURL    string  `json:"wiki_link"`
	Date       string  `json:"date"`

	// Alliance bank (restricted to alliance members/officers)
	Money     float64 `json:"money"`
	Coal      float64 `json:"coal"`
	Oil       float64 `json:"oil"`
	Uranium   float64 `json:"uranium"`
	Iron      float64 `json:"iron"`
	Bauxite   float64 `json:"bauxite"`
	Lead      float64 `json:"lead"`
	Gasoline  float64 `json:"gasoline"`
	Munitions float64 `json:"munitions"`
	Steel     float64 `json:"steel"`
	Aluminum  float64 `json:"aluminum"`
	Food      float64 `json:"food"`

	// Relations
	Nations  []Nation  `json:"nations,omitempty"`
	Treaties []Treaty  `json:"treaties,omitempty"`
	Bankrecs []BankRec `json:"bankrecs,omitempty"`
	Taxrecs  []BankRec `json:"taxrecs,omitempty"`
}

// AlliancePaginator wraps paginated Alliance results.
type AlliancePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Alliance    `json:"data"`
}

// Treaty represents a diplomatic treaty between two alliances.
type Treaty struct {
	ID          ID        `json:"id"`
	Date        string    `json:"date"`
	TreatyType  string    `json:"treaty_type"`
	TurnsLeft   int       `json:"turns_left"`
	Alliance1ID ID        `json:"alliance1_id"`
	Alliance2ID ID        `json:"alliance2_id"`
	Alliance1   *Alliance `json:"alliance1,omitempty"`
	Alliance2   *Alliance `json:"alliance2,omitempty"`
}

// ----------------------------------------------------------------------------
// City
// ----------------------------------------------------------------------------

type City struct {
	ID               ID      `json:"id"`
	NationID         ID      `json:"nation_id"`
	Name             string  `json:"name"`
	Date             string  `json:"date"`
	Infrastructure   float64 `json:"infrastructure"`
	Land             float64 `json:"land"`
	Powered          bool    `json:"powered"`
	CoalPower        int     `json:"coal_power"`
	OilPower         int     `json:"oil_power"`
	NuclearPower     int     `json:"nuclear_power"`
	WindPower        int     `json:"wind_power"`
	CoalMine         int     `json:"coal_mine"`
	OilWell          int     `json:"oil_well"`
	UraniumMine      int     `json:"uranium_mine"`
	IronMine         int     `json:"iron_mine"`
	BauxiteMine      int     `json:"bauxite_mine"`
	LeadMine         int     `json:"lead_mine"`
	Farm             int     `json:"farm"`
	GasRefinery      int     `json:"oil_refinery"`
	AluminumRefinery int     `json:"aluminum_refinery"`
	SteelMill        int     `json:"steel_mill"`
	MunitionsFactory int     `json:"munitions_factory"`
	Barracks         int     `json:"barracks"`
	Factory          int     `json:"factory"`
	HangarFacility   int     `json:"hangar"`
	Drydock          int     `json:"drydock"`
	Supermarket      int     `json:"supermarket"`
	Bank             int     `json:"bank"`
	ShoppingMall     int     `json:"shopping_mall"`
	Stadium          int     `json:"stadium"`
	Hospital         int     `json:"hospital"`
	RecyclingCenter  int     `json:"recycling_center"`
	Subway           int     `json:"subway"`
	Superstore       int     `json:"superstore"`
	Zoo              int     `json:"zoo"`
	TreatmentPlant   int     `json:"pollution_control_center"`
	NuclearPlant     int     `json:"nuclear_power_plant"`
	Research         int     `json:"research_lab"`
	Embassy          int     `json:"embassy"`
	ImprovementBonus int     `json:"maintenance_improvements"`
}

// CityPaginator wraps paginated City results.
type CityPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []City        `json:"data"`
}

// ----------------------------------------------------------------------------
// War
// ----------------------------------------------------------------------------

type War struct {
	ID             ID      `json:"id"`
	Date           string  `json:"date"`
	Reason         string  `json:"reason"`
	WarType        WarType `json:"war_type"`
	GroundControl  ID      `json:"ground_control"`
	AirSuperiority ID      `json:"air_superiority"`
	NavalBlockade  ID      `json:"naval_blockade"`
	WinnerID       ID      `json:"winner_id"`
	AttID          ID      `json:"att_id"`
	AttAllianceID  ID      `json:"att_alliance_id"`
	DefID          ID      `json:"def_id"`
	DefAllianceID  ID      `json:"def_alliance_id"`
	AttPoints      int     `json:"att_points"`
	DefPoints      int     `json:"def_points"`
	AttPeace       bool    `json:"att_peace"`
	DefPeace       bool    `json:"def_peace"`
	AttResistance  int     `json:"att_resistance"`
	DefResistance  int     `json:"def_resistance"`
	AttFortify     bool    `json:"att_fortify"`
	DefFortify     bool    `json:"def_fortify"`
	// Military losses
	AttSoldiersLost int `json:"att_soldiers_lost"`
	DefSoldiersLost int `json:"def_soldiers_lost"`
	AttTanksLost    int `json:"att_tanks_lost"`
	DefTanksLost    int `json:"def_tanks_lost"`
	AttAircraftLost int `json:"att_aircraft_lost"`
	DefAircraftLost int `json:"def_aircraft_lost"`
	AttShipsLost    int `json:"att_ships_lost"`
	DefShipsLost    int `json:"def_ships_lost"`
	AttMissilesUsed int `json:"att_missiles_used"`
	DefMissilesUsed int `json:"def_missiles_used"`
	AttNukesUsed    int `json:"att_nukes_used"`
	DefNukesUsed    int `json:"def_nukes_used"`
	// Infrastructure destroyed
	AttInfraDestroyed float64 `json:"att_infra_destroyed"`
	DefInfraDestroyed float64 `json:"def_infra_destroyed"`
	// Money looted
	AttMoneyLooted float64 `json:"att_money_looted"`
	DefMoneyLooted float64 `json:"def_money_looted"`
	// Relations
	Attacks  []WarAttack `json:"attacks,omitempty"`
	Attacker *Nation     `json:"attacker,omitempty"`
	Defender *Nation     `json:"defender,omitempty"`
}

// WarAttack represents a single attack within a war.
type WarAttack struct {
	ID                  ID         `json:"id"`
	Date                string     `json:"date"`
	AttID               ID         `json:"att_id"`
	DefID               ID         `json:"def_id"`
	Type                AttackType `json:"type"`
	WarID               ID         `json:"war_id"`
	AttCasualties       int        `json:"att_casualties"`
	DefCasualties       int        `json:"def_casualties"`
	InfraDestroyed      float64    `json:"infra_destroyed"`
	InfraDestroyedValue float64    `json:"infra_destroyed_value"`
	MoneyStolen         float64    `json:"money_stolen"`
	Success             int        `json:"success"`
	AttEquipmentLost    int        `json:"att_equipment_lost"`
	DefEquipmentLost    int        `json:"def_equipment_lost"`
	CityID              ID         `json:"city_id"`
}

// ----------------------------------------------------------------------------
// Trade
// ----------------------------------------------------------------------------

type Trade struct {
	ID            ID        `json:"id"`
	Type          TradeType `json:"type"`
	Date          string    `json:"date"`
	SenderID      ID        `json:"sender_id"`
	ReceiverID    ID        `json:"receiver_id"`
	OfferResource string    `json:"offer_resource"`
	OfferAmount   int       `json:"offer_amount"`
	BuyOrSell     string    `json:"buy_or_sell"`
	PricePerUnit  int       `json:"price_per_unit"`
	Accepted      bool      `json:"accepted"`
	DateAccepted  string    `json:"date_accepted"`
	OriginalID    ID        `json:"original_trade_id"`
}

// TradeInfoResource holds average market price data for a single resource.
type TradeInfoResource struct {
	Resource     string  `json:"resource"`
	AveragePrice float64 `json:"average_price"`
}

// TopTradeInfo is the market summary returned by the top_trade_info query.
type TopTradeInfo struct {
	MarketIndex float64             `json:"market_index"`
	Resources   []TradeInfoResource `json:"resources"`
}

// Tradeprice represents the current market price for a resource.
type Tradeprice struct {
	ID        ID      `json:"id"`
	Date      string  `json:"date"`
	Coal      float64 `json:"coal"`
	Oil       float64 `json:"oil"`
	Uranium   float64 `json:"uranium"`
	Iron      float64 `json:"iron"`
	Bauxite   float64 `json:"bauxite"`
	Lead      float64 `json:"lead"`
	Gasoline  float64 `json:"gasoline"`
	Munitions float64 `json:"munitions"`
	Steel     float64 `json:"steel"`
	Aluminum  float64 `json:"aluminum"`
	Food      float64 `json:"food"`
	Credits   float64 `json:"credits"`
}

// ----------------------------------------------------------------------------
// Treasure
// ----------------------------------------------------------------------------

type Treasure struct {
	Name      string  `json:"name"`
	Color     Color   `json:"color"`
	Continent string  `json:"continent"`
	Bonus     int     `json:"bonus"`
	SpawnDate string  `json:"spawndate"`
	NationID  ID      `json:"nation_id"`
	Nation    *Nation `json:"nation,omitempty"`
}

// ----------------------------------------------------------------------------
// Bounty
// ----------------------------------------------------------------------------

type Bounty struct {
	ID       ID      `json:"id"`
	Date     string  `json:"date"`
	NationID ID      `json:"nation_id"`
	Amount   int     `json:"amount"`
	Type     WarType `json:"type"`
	Nation   *Nation `json:"nation,omitempty"`
}

// ----------------------------------------------------------------------------
// Bank Record
// ----------------------------------------------------------------------------

type BankRec struct {
	ID   ID     `json:"id"`
	Date string `json:"date"`
	// TaxID is non-zero when the record is an automatic tax collection.
	TaxID        ID      `json:"tax_id"`
	EnteringID   ID      `json:"eid"`
	EnteringType int     `json:"etype"`
	SenderID     ID      `json:"sid"`
	SenderType   int     `json:"stype"`
	ReceiverID   ID      `json:"rid"`
	ReceiverType int     `json:"rtype"`
	Note         string  `json:"note"`
	Money        float64 `json:"money"`
	Coal         float64 `json:"coal"`
	Oil          float64 `json:"oil"`
	Uranium      float64 `json:"uranium"`
	Iron         float64 `json:"iron"`
	Bauxite      float64 `json:"bauxite"`
	Lead         float64 `json:"lead"`
	Gasoline     float64 `json:"gasoline"`
	Munitions    float64 `json:"munitions"`
	Steel        float64 `json:"steel"`
	Aluminum     float64 `json:"aluminum"`
	Food         float64 `json:"food"`
}

// ----------------------------------------------------------------------------
// Color info
// ----------------------------------------------------------------------------

// ColorInfo describes a game color bloc and its trade bonus.
type ColorInfo struct {
	Color     Color  `json:"color"`
	BloCName  string `json:"bloc_name"`
	TurnBonus int    `json:"turn_bonus"`
}
