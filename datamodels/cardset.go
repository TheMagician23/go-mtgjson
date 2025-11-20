package datamodels

// CardSet has information of single card in a Set
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
	EdhrecSaltiness         float32          `json:"edhrecSaltiness,omitempty"`
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
