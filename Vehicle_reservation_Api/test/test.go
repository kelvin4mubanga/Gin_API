
func TestCreateReservation(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reservation := Reservation{Vehicle: "Test Vehicle", Customer: "Test Customer", Date: time.Now(), Status: "Pending"}
	jsonValue, _ := json.Marshal(reservation)
	c.Request = httptest.NewRequest("POST", "/reservations", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateReservation(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdReservation Reservation
	err := json.Unmarshal(w.Body.Bytes(), &createdReservation)
	assert.Nil(t, err)
	assert.NotZero(t, createdReservation.ID)
}

func TestGetReservations(t *testing.T) {
	setup()
	defer teardown()

	// Create some reservations
	database.DB.Create(&Reservation{Vehicle: "Vehicle 1", Customer: "Customer 1", Date: time.Now(), Status: "Pending"})
	database.DB.Create(&Reservation{Vehicle: "Vehicle 2", Customer: "Customer 2", Date: time.Now(), Status: "Confirmed"})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/reservations", nil)

	GetReservations(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var reservations []Reservation
	err := json.Unmarshal(w.Body.Bytes(), &reservations)
	assert.Nil(t, err)
	assert.Len(t, reservations, 2)
}

func TestGetReservation(t *testing.T) {
	setup()
	defer teardown()

	// Create a reservation
	reservation := Reservation{Vehicle: "Test Vehicle", Customer: "Test Customer", Date: time.Now(), Status: "Pending"}
	database.DB.Create(&reservation)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/reservations/1", nil)

	GetReservation(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedReservation Reservation
	err := json.Unmarshal(w.Body.Bytes(), &retrievedReservation)
	assert.Nil(t, err)
	assert.Equal(t, reservation.ID, retrievedReservation.ID)
}

func TestUpdateReservation(t *testing.T) {
	setup()
	defer teardown()

	// Create a reservation
	reservation := Reservation{Vehicle: "Test Vehicle", Customer: "Test Customer", Date: time.Now(), Status: "Pending"}
	database.DB.Create(&reservation)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedReservation := Reservation{Vehicle: "Updated Vehicle", Customer: "Updated Customer", Date: time.Now(), Status: "Confirmed"}
	jsonValue, _ := json.Marshal(updatedReservation)
	c.Request = httptest.NewRequest("PUT", "/reservations/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateReservation(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedReservation Reservation
	err := json.Unmarshal(w.Body.Bytes(), &retrievedReservation)
	assert.Nil(t, err)
	assert.Equal(t, updatedReservation.Vehicle, retrievedReservation.Vehicle)
}

func TestDeleteReservation(t *testing.T) {
	setup()
	defer teardown()

	// Create a reservation
	reservation := Reservation{Vehicle: "Test Vehicle", Customer: "Test Customer", Date: time.Now(), Status: "Pending"}
	database.DB.Create(&reservation)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/reservations/1", nil)

	DeleteReservation(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var reservations []Reservation
	database.DB.Find(&reservations)
	assert.Len(t, reservations, 0)
}
/*
 These tests cover the following scenarios:

- Creating a new reservation
- Retrieving all reservations
- Retrieving a single reservation by ID
- Updating a reservation
- Deleting a reservation

You can save these tests in a file named reservation_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.
*/