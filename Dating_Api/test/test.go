Here are some unit tests for your profile API handlers:

func TestCreateProfile(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	profile := Profile{}
	jsonValue, _ := json.Marshal(profile)
	c.Request = httptest.NewRequest("POST", "/profiles", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateProfile(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdProfile Profile
	err := json.Unmarshal(w.Body.Bytes(), &createdProfile)
	assert.Nil(t, err)
	assert.NotZero(t, createdProfile.ID)
}

func TestGetProfiles(t *testing.T) {
	setup()
	defer teardown()

	// Create some profiles
	database.DB.Create(&Profile{})
	database.DB.Create(&Profile{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/profiles", nil)

	GetProfiles(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var profiles []Profile
	err := json.Unmarshal(w.Body.Bytes(), &profiles)
	assert.Nil(t, err)
	assert.Len(t, profiles, 2)
}

func TestGetProfile(t *testing.T) {
	setup()
	defer teardown()

	// Create a profile
	profile := Profile{}
	database.DB.Create(&profile)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/profiles/1", nil)

	GetProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedProfile Profile
	err := json.Unmarshal(w.Body.Bytes(), &retrievedProfile)
	assert.Nil(t, err)
	assert.Equal(t, profile.ID, retrievedProfile.ID)
}

func TestUpdateProfile(t *testing.T) {
	setup()
	defer teardown()

	// Create a profile
	profile := Profile{}
	database.DB.Create(&profile)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedProfile := Profile{}
	jsonValue, _ := json.Marshal(updatedProfile)
	c.Request = httptest.NewRequest("PUT", "/profiles/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedProfile Profile
	err := json.Unmarshal(w.Body.Bytes(), &retrievedProfile)
	assert.Nil(t, err)
	assert.Equal(t, profile.ID, retrievedProfile.ID)
}

func TestDeleteProfile(t *testing.T) {
	setup()
	defer teardown()

	// Create a profile
	profile := Profile{}
	database.DB.Create(&profile)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/profiles/1", nil)

	DeleteProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var profiles []Profile
	database.DB.Find(&profiles)
	assert.Len(t, profiles, 0)
}

These tests cover the following scenarios:

- Creating a new profile
- Retrieving all profiles
- Retrieving a single profile by ID
- Updating a profile
- Deleting a profile

You can save these tests in a file named profile_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.

You may also want to test the RegisterRoutes function to ensure that the routes are registered correctly.

func TestRegisterRoutes(t *testing.T) {
	router := gin.New()
	RegisterRoutes(router)

	routes := router.Routes()
	assert.Len(t, routes, 5)
}