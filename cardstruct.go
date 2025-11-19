package mtgjson

type ForeignData struct {
	FaceName    string      `json:"faceName,omitempty"`
	FlavorText  string      `json:"flavorText,omitempty"`
	Identifiers Identifiers `json:"identifiers"`
	Language    string      `json:"language"`
	Name        string      `json:"name"`
	Text        string      `json:"text,omitempty"`
	Type        string      `json:"type,omitempty"`
	Uuid        string      `json:"uuid"`
}

type LeadershipSkills struct {
	Brawl       bool `json:"brawl"`
	Commander   bool `json:"commander"`
	Oathbreaker bool `json:"oathbreaker"`
}

type Legalities struct {
	Alchemy         string `json:"alchemy,omitempty"`
	Brawl           string `json:"brawl,omitempty"`
	Commander       string `json:"commander,omitempty"`
	Duel            string `json:"duel,omitempty"`
	Explorer        string `json:"explorer,omitempty"`
	Future          string `json:"future,omitempty"`
	Gladiator       string `json:"gladiator,omitempty"`
	Historic        string `json:"historic,omitempty"`
	HistoricBrawl   string `json:"historicbrawl,omitempty"`
	Legacy          string `json:"legacy,omitempty"`
	Modern          string `json:"modern,omitempty"`
	Oathbreaker     string `json:"oathbreaker,omitempty"`
	Oldschool       string `json:"oldschool,omitempty"`
	Pauper          string `json:"pauper,omitempty"`
	PauperCommander string `json:"paupercommander,omitempty"`
	Penny           string `json:"penny,omitempty"`
	Pioneer         string `json:"pioneer,omitempty"`
	Predh           string `json:"predh,omitempty"`
	Premodern       string `json:"premodern,omitempty"`
	Standard        string `json:"standard,omitempty"`
	StandardBrawl   string `json:"standardbrawl,omitempty"`
	Timeless        string `json:"timeless,omitempty"`
	Vintage         string `json:"vintage,omitempty"`
}

type RelatedCards struct {
	ReverseRelated []string `json:"reverseRelated,omitempty"`
	Spellbook      []string `json:"spellbook,omitempty"`
}

type Rulings struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

type SourceProducts struct {
	Etched  []string `json:"etched,omitempty"`
	Foil    []string `json:"foil,omitempty"`
	NonFoil []string `json:"nonfoil,omitempty"`
}

// Card has information of single card
// Based on MTGJSON v5.5.2+20251117
type CardSet struct {
	Artist                  string           `json:"artist"`
	ArtistIds               []string         `json:"artistIds,omitempty"`
	AsciiName               string           `json:"asciiName,omitempty"`
	AttractionLights        []int            `json:"attractionLights,omitempty"`
	Availability            []string         `json:"availability,omitempty"`
	BoosterTypes            []string         `json:"boosterTypes,omitempty"`
	BorderColor             string           `json:"borderColor"`
	CardParts               []string         `json:"cardParts,omitempty"`
	ColorIdentity           []string         `json:"colorIdentity"`
	ColorIndicator          []string         `json:"colorIndicator,omitempty"`
	Colors                  []string         `json:"colors"`
	ConvertedManaCost       float32          `json:"convertedManaCost"` // Deprecated; use `manaValue` instead
	Defense                 string           `json:"defense,omitempty"`
	DuelDeck                string           `json:"duelDeck,omitempty"`
	EdhrecRank              int              `json:"edhrecRank"`
	EdhrecSaltiness         int              `json:"edhrecSaltiness,omitempty"`
	FaceConvertedManaCost   float32          `json:"faceConvertedManaCost,omitempty"` // Deprecated; use `faceManaValue` instead
	FaceFlavorName          string           `json:"faceFlavorName,omitempty"`
	FaceManaValue           float32          `json:"faceManaValue,omitempty"`
	FaceName                string           `json:"faceName,omitempty"`
	Finishes                []string         `json:"finishes"`
	FlavorName              string           `json:"flavorName,omitempty"`
	FlavorText              string           `json:"flavorText,omitempty"`
	ForeignData             []ForeignData    `json:"foreignData,omitempty"`
	FrameEffects            []string         `json:"frameEffects,omitempty"`
	FrameVersion            string           `json:"frameVersion"`
	Hand                    string           `json:"hand,omitempty"`
	HasAlternativeDeckLimit bool             `json"hasAlternativeDeckLimit,omitempty"`
	HasContentWarning       bool             `json:"hasContentWarning,omitempty"`
	HasFoil                 bool             `json:"hasFoil"`    // Deprecated; use `finishes` instead
	HasNonFoil              bool             `json:"hasNonFoil"` // Deprecated; use `finishes` instead
	Identifiers             Identifiers      `json:"identifiers"`
	IsAlternative           bool             `json:"isAlternative,omitempty"`
	IsFullArt               bool             `json:"isFullArt,omitempty"`
	IsFunny                 bool             `json:"isFunny,omitempty"`
	IsGameChanger           bool             `json:"isGameChanger,omitempty"`
	IsOnlineOnly            bool             `json:"isOnlineOnly,omitempty"`
	IsOversized             bool             `json:"isOversized,omitempty"`
	IsPromo                 bool             `json:"isPromo,omitempty"`
	IsRebalanced            bool             `json:"isRebalanced,omitempty"`
	IsReprint               bool             `json:"isReprint,omitempty"`
	IsReserved              bool             `json:"isReserved,omitempty"`
	IsStarter               bool             `json:"isStarter,omitempty"` // Deprecated
	IsStorySpotlight        bool             `json:"isStorySpotlight,omitempty"`
	IsTextless              bool             `json:"isTextless,omitempty"`
	IsTimeshifted           bool             `json:"isTimeshifted,omitempty"`
	Keywords                []string         `json:"keywords,omitempty"`
	Language                string           `json:"language"`
	Layout                  string           `json:"layout"`
	LeadershipSkills        LeadershipSkills `json:"leadershipSkills,omitempty"`
	Legalities              Legalities       `json:"legalities"`
	Life                    string           `json:"life,omitempty"`
	Loyalty                 string           `json:"loyalty,omitempty"`
	ManaCost                string           `json:"manaCost,omitempty"`
	ManaValue               float32          `json:"manaValue"`
	Name                    string           `json:"name"`
	Number                  string           `json:"number"`
	OriginalPrintings       []string         `json:"originalPrintings,omitempty"`
	OriginalReleaseDate     string           `json:"originalReleaseDate,omitempty"`
	OriginalText            string           `json:"originalText,omitempty"`
	OriginalType            string           `json:"originalType,omitempty"`
	OtherFaceIds            []string         `json:"otherFaceIds,omitempty"`
	Power                   string           `json:"power,omitempty"`
	Printings               []string         `json:"printings,omitempty"`
	PromoTypes              []string         `json:"promoTypes,omitempty"`
	PurchaseUrls            PurchaseUrls     `json:"purchaseUrls"`
	Rarity                  string           `json:"rarity"`
	RelatedCards            RelatedCards     `json:"relatedCards,omitempty"`
	RebalancedPrintings     []string         `json:"rebalancedPrintings,omitempty"`
	Rulings                 []Rulings        `json:"rulings,omitempty"`
	SecurityStamp           string           `json:"securityStamp,omitempty"`
	SetCode                 string           `json:"setCode"`
	Side                    string           `json:"side,omitempty"`
	Signature               string           `json:"signature,omitempty"`
	SourceProducts          SourceProducts   `json:"sourceProducts,omitempty"`
	Subsets                 []string         `json:"subsets,omitempty"`
	Subtypes                []string         `json:"subtypes"`
	Supertypes              []string         `json:"supertypes"`
	Text                    string           `json:"text,omitempty"`
	Toughness               string           `json:"toughness,omitempty"`
	Type                    string           `json:"type"`
	Types                   []string         `json:"types"`
	Uuid                    string           `json:"uuid"`
	Variations              []string         `json:"variations,omitempty"`
	Watermark               string           `json:"watermark,omitempty"`
}
