package tmdb

func (suite *TMBDTestSuite) TestGetTVSeasonDetails() {
	got, err := suite.client.GetTVSeasonDetails(gotID, 1, nil)
	suite.Nil(err)
	suite.Equal("Season 1", got.Name)
}

func (suite *TMBDTestSuite) TestGetTVSeasonDetailsFail() {
	_, err := suite.client.GetTVSeasonDetails(0, 1, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	got, err := suite.client.GetTVSeasonDetails(gotID, 1, options)
	suite.Nil(err)
	suite.Equal("1ª Temporada", got.Name)
}

func (suite *TMBDTestSuite) TestGetTVSeasonChange() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	got, err := suite.client.GetTVSeasonChanges(gotSeasonID, options)
	suite.Nil(err)
	suite.Equal("5c423aaf925141344cb32a9d", got.Changes[0].Items[0].ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonCredits() {
	tv, err := suite.client.GetTVSeasonCredits(gotID, 7, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonCreditsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVSeasonCredits(0, 7, nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVSeasonCredits(gotID, 7, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonExternalIDs() {
	tv, err := suite.client.GetTVSeasonExternalIDs(gotID, 7, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonExternalIDsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVSeasonExternalIDs(0, 7, nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonExternalIDsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVSeasonExternalIDs(gotID, 7, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonImages() {
	tv, err := suite.client.GetTVSeasonImages(gotID, 7, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonImagesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVSeasonImages(0, 7, nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVSeasonImages(gotID, 7, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonVideos() {
	tv, err := suite.client.GetTVSeasonVideos(gotID, 7, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVSeasonVideosFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVSeasonVideos(0, 7, nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonVideosWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVSeasonVideos(gotID, 7, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}
