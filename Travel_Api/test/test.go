Here are some unit tests for your destination handler functions:

func TestCreateDestination(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	destination := Destination{Name: "Test Destination", Country: "Test Country", Description: "This is a test destination"}
	jsonValue, _ := json.Marshal(destination)
	c.Request = httptest.NewRequest("POST", "/destinations", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateDestination(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdDestination Destination
	err := json.Unmarshal(w.Body.Bytes(), &createdDestination)
	assert.Nil(t, err)
	assert.NotZero(t, createdDestination.ID)
}

func TestGetDestinations(t *testing.T) {
	setup()
	defer teardown()

	// Create some destinations
	database.DB.Create(&Destination{Name: "Destination 1", Country: "Country 1", Description: "Description 1"})
	database.DB.Create(&Destination{Name: "Destination 2", Country: "Country 2", Description: "Description 2"})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/destinations", nil)

	GetDestinations(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var destinations []Destination
	err := json.Unmarshal(w.Body.Bytes(), &destinations)
	assert.Nil(t, err)
	assert.Len(t, destinations, 2)
}

func TestGetDestination(t *testing.T) {
	setup()
	defer teardown()

	// Create a destination
	destination := Destination{Name: "Test Destination", Country: "Test Country", Description: "This is a test destination"}
	database.DB.Create(&destination)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/destinations/1", nil)

	GetDestination(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedDestination Destination
	err := json.Unmarshal(w.Body.Bytes(), &retrievedDestination)
	assert.Nil(t, err)
	assert.Equal(t, destination.ID, retrievedDestination.ID)
}

func TestUpdateDestination(t *testing.T) {
	setup()
	defer teardown()

	// Create a destination
	destination := Destination{Name: "Test Destination", Country: "Test Country", Description: "This is a test destination"}
	database.DB.Create(&destination)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedDestination := Destination{Name: "Updated Destination", Country: "Updated Country", Description: "This is an updated destination"}
	jsonValue, _ := json.Marshal(updatedDestination)
	c.Request = httptest.NewRequest("PUT", "/destinations/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateDestination(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedDestination Destination
	err := json.Unmarshal(w.Body.Bytes(), &retrievedDestination)
	assert.Nil(t, err)
	assert.Equal(t, updatedDestination.Name, retrievedDestination.Name)
}

func TestDeleteDestination(t *testing.T) {
	setup()
	defer teardown()

	// Create a destination
	destination := Destination{Name: "Test Destination", Country: "Test Country", Description: "This is a test destination"}
	database.DB.Create(&destination)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/destinations/1", nil)

	DeleteDestination(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var destinations []Destination
	database.DB.Find(&destinations)
	assert.Len(t, destinations, 0)
}

These tests cover the following scenarios:

- Creating a new destination
- Retrieving all destinations
- Retrieving a single destination by ID
- Updating a destination
- Deleting a destination

You can save these tests in a file named destination_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.