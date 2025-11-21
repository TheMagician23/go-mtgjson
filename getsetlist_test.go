package mtgjson

/*
type dummyGetSetListAPI struct{}

func (api *dummyGetSetListAPI) doRequest(path string) (io.Reader, error) {
	dummy := `{"meta":{"date":"2025-11-17","version":"5.2.2+20251117"},"data":[{"baseSetSize":0,"booster":null,"code":"P15A","decks":null,"languages":null,"name":"15th Anniversary Cards","releaseDate":"2008-04-01","totalSetSize":2,"type":"promo"},{"baseSetSize":0,"booster":null,"code":"HTR","decks":null,"languages":null,"name":"2016 Heroes of the Realm","releaseDate":"2017-09-20","totalSetSize":3,"type":"memorabilia"}]}`
	return strings.NewReader(dummy), nil
}

func TestGetSetList(t *testing.T) {
	c := &Client{&dummyGetSetListAPI{}}
	sets, err := c.GetSetList()
	if err != nil {
		t.Fatalf("failed to get set list")
	}

	var expectedSets SetList
	expectedSets = SetList{
		Meta: SetListMeta{
			Date:    "2025-11-17",
			Version: "5.2.2+20251117",
		},
		Sets: []Set{
			Set{
				BaseSetSize:        0,
				Block:              "",
				Booster:            nil,
				Cards:              nil,
				CardsphereSetId:    0,
				Code:               "P15A",
				CodeV3:             "",
				Decks:              nil,
				IsForeignOnly:      false,
				IsFoilOnly:         false,
				IsNonFoilOnly:      false,
				IsOnlineOnly:       false,
				IsPaperOnly:        false,
				IsPartialPreview:   false,
				KeyruneCode:        "",
				Languages:          nil,
				CardmarketId:       0,
				CardmarketIdExtras: 0,
				CardmarketName:     "",
				MTGOCode:           "",
				Name:               "15th Anniversary Cards",
				ReleaseDate:        "2008-04-01",
				SealedProduct:      nil,
				TCGPlayerGroupID:   0,
				Tokens:             nil,
				TokenSetCode:       "",
				TotalSetSize:       2,
				Translations:       nil,
				Type:               "promo",
			},
			Set{
				BaseSetSize:        0,
				Block:              "",
				Booster:            nil,
				Cards:              nil,
				CardsphereSetId:    0,
				Code:               "HTR",
				CodeV3:             "",
				Decks:              nil,
				IsForeignOnly:      false,
				IsFoilOnly:         false,
				IsNonFoilOnly:      false,
				IsOnlineOnly:       false,
				IsPaperOnly:        false,
				IsPartialPreview:   false,
				KeyruneCode:        "",
				Languages:          nil,
				CardmarketId:       0,
				CardmarketIdExtras: 0,
				CardmarketName:     "",
				MTGOCode:           "",
				Name:               "2016 Heroes of the Realm",
				ReleaseDate:        "2017-09-20",
				SealedProduct:      nil,
				TCGPlayerGroupID:   0,
				Tokens:             nil,
				TokenSetCode:       "",
				TotalSetSize:       3,
				Translations:       nil,
				Type:               "memorabilia",
			},
		},
	}

	if !reflect.DeepEqual(sets, expectedSets) {
		t.Fatalf("got wrong sets")
	}
}
*/
