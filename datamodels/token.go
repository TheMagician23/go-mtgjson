package datamodels

// CardToken describes the properties of a single token card found in a Set.
type CardToken struct {
	Artist          string         `json:"artist"`
	ArtistIds       []string       `json:"artistIds,omitempty"`
	AsciiName       string         `json:"asciiName,omitempty"`
	Availability    []string       `json:"availability,omitempty"`
	BoosterTypes    []string       `json:"boosterTypes,omitempty"`
	BorderColor     string         `json:"borderColor"`
	CardParts       []string       `json:"cardParts,omitempty"`
	ColorIdentity   []string       `json:"colorIdentity"`
	ColorIndicator  []string       `json:"colorIndicator,omitempty"`
	Colors          []string       `json:"colors"`
	EdhrecSaltiness float32        `json:"edhrecSaltiness,omitempty"`
	FaceName        string         `json:"faceName,omitempty"`
	FaceFlavorName  string         `json:faceFlavorName,omitempty"`
	Finishes        []string       `json:"finishes"`
	FlavorName      string         `json:"flavorName,omitempty"`
	FlavorText      string         `json:"flavorText,omitempty"`
	FrameEffects    []string       `json:"frameEffects,omitempty"`
	FrameVersion    string         `json:"frameVersion,omitempty"`
	HasFoil         bool           `json:"hasFoil"`    // Deprecated; use `finishes` instead
	HasNonFoil      bool           `json:"hasNonFoil"` // Deprecated; use `finishes` instead
	Identifiers     Identifiers    `json:"identifiers"`
	IsFullArt       bool           `json:"isFullArt,omitempty"`
	IsFunny         bool           `json:"isFunny,omitempty"`
	IsOnlineOnly    bool           `json:"isOnlineOnly,omitempty"`
	IsOversized     bool           `json:"isOversized,omitempty"`
	IsPromo         bool           `json:"isPromo,omitempty"`
	IsReprint       bool           `json:"isReprint,omitempty"`
	IsTextless      bool           `json:"isTextless,omitempty"`
	Keywords        []string       `json:"keywords,omitempty"`
	Language        string         `json:"language"`
	Layout          string         `json:"layout"`
	Loyalty         string         `json:"loyalty,omitempty"`
	ManaCost        string         `json:"manaCost,omitempty"`
	Name            string         `json:"name"`
	Number          string         `json:"number"`
	Orientation     string         `json:"orientation,omitempty"`
	OriginalText    string         `json:"originalText,omitempty"`
	OriginalType    string         `json:"originalType,omitempty"`
	OtherFaceIds    []string       `json:"otherFaceIds,omitempty"`
	Power           string         `json:"power,omitempty"`
	PromoTypes      []string       `json:"promoTypes,omitempty"`
	RelatedCards    RelatedCards   `json:"relatedCards,omitempty"`
	ReverseRelated  []string       `json:"reverseRelated,omitempty"` // Deprecated; use `relatedCards` instead
	SecurityStamp   string         `json:"securityStamp,omitempty"`
	SetCode         string         `json:"setCode"`
	Side            string         `json:"side,omitempty"`
	Signature       string         `json:"signature,omitempty"`
	SourceProducts  SourceProducts `json:"sourceProducts,omitempty"`
	Subsets         []string       `json:"subsets,omitempty"`
	Subtypes        []string       `json:"subtypes"`
	Supertypes      []string       `json:"supertypes"`
	Text            string         `json:"text,omitempty"`
	Toughness       string         `json:"toughness,omitempty"`
	Type            string         `json:"type"`
	Types           []string       `json:"types"`
	Uuid            string         `json:"uuid"`
	Watermark       string         `json:"watermark,omitempty"`
}
