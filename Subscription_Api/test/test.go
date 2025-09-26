//Here are some unit tests for your subscription handler functions:

func TestCreateSubscription(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	subscription := Subscription{User: "John Doe", Plan: "Monthly", Status: "active", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0)}
	jsonValue, _ := json.Marshal(subscription)
	c.Request = httptest.NewRequest("POST", "/subscriptions", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateSubscription(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdSubscription Subscription
	err := json.Unmarshal(w.Body.Bytes(), &createdSubscription)
	assert.Nil(t, err)
	assert.NotZero(t, createdSubscription.ID)
}

func TestGetSubscriptions(t *testing.T) {
	setup()
	defer teardown()

	// Create some subscriptions
	database.DB.Create(&Subscription{User: "John Doe", Plan: "Monthly", Status: "active", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0)})
	database.DB.Create(&Subscription{User: "Jane Doe", Plan: "Yearly", Status: "inactive", StartDate: time.Now(), EndDate: time.Now().AddDate(1, 0, 0)})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/subscriptions", nil)

	GetSubscriptions(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var subscriptions []Subscription
	err := json.Unmarshal(w.Body.Bytes(), &subscriptions)
	assert.Nil(t, err)
	assert.Len(t, subscriptions, 2)
}

func TestGetSubscription(t *testing.T) {
	setup()
	defer teardown()

	// Create a subscription
	subscription := Subscription{User: "John Doe", Plan: "Monthly", Status: "active", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0)}
	database.DB.Create(&subscription)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/subscriptions/1", nil)

	GetSubscription(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedSubscription Subscription
	err := json.Unmarshal(w.Body.Bytes(), &retrievedSubscription)
	assert.Nil(t, err)
	assert.Equal(t, subscription.ID, retrievedSubscription.ID)
}

func TestUpdateSubscription(t *testing.T) {
	setup()
	defer teardown()

	// Create a subscription
	subscription := Subscription{User: "John Doe", Plan: "Monthly", Status: "active", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0)}
	database.DB.Create(&subscription)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedSubscription := Subscription{User: "Jane Doe", Plan: "Yearly", Status: "inactive", StartDate: time.Now(), EndDate: time.Now().AddDate(1, 0, 0)}
	jsonValue, _ := json.Marshal(updatedSubscription)
	c.Request = httptest.NewRequest("PUT", "/subscriptions/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateSubscription(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedSubscription Subscription
	err := json.Unmarshal(w.Body.Bytes(), &retrievedSubscription)
	assert.Nil(t, err)
	assert.Equal(t, updatedSubscription.User, retrievedSubscription.User)
}

func TestDeleteSubscription(t *testing.T) {
	setup()
	defer teardown()

	// Create a subscription
	subscription := Subscription{User: "John Doe", Plan: "Monthly", Status: "active", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 1, 0)}
	database.DB.Create(&subscription)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/subscriptions/1", nil)

	DeleteSubscription(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var subscriptions []Subscription
	database.DB.Find(&subscriptions)
	assert.Len(t, subscriptions, 0)
}
/*
 These tests cover the following scenarios:

- Creating a new subscription
- Retrieving all subscriptions
- Retrieving a single subscription by ID
- Updating a subscription
- Deleting a subscription


*/