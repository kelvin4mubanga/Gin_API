Here are some unit tests for your booking handler functions:

func TestCreateBooking(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	booking := Booking{Customer: "John Doe", Destination: "New York", Date: time.Now(), Status: "pending"}
	jsonValue, _ := json.Marshal(booking)
	c.Request = httptest.NewRequest("POST", "/bookings", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateBooking(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdBooking Booking
	err := json.Unmarshal(w.Body.Bytes(), &createdBooking)
	assert.Nil(t, err)
	assert.NotZero(t, createdBooking.ID)
}

func TestGetBookings(t *testing.T) {
	setup()
	defer teardown()

	// Create some bookings
	database.DB.Create(&Booking{Customer: "John Doe", Destination: "New York", Date: time.Now(), Status: "pending"})
	database.DB.Create(&Booking{Customer: "Jane Doe", Destination: "Los Angeles", Date: time.Now(), Status: "confirmed"})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/bookings", nil)

	GetBookings(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var bookings []Booking
	err := json.Unmarshal(w.Body.Bytes(), &bookings)
	assert.Nil(t, err)
	assert.Len(t, bookings, 2)
}

func TestGetBooking(t *testing.T) {
	setup()
	defer teardown()

	// Create a booking
	booking := Booking{Customer: "John Doe", Destination: "New York", Date: time.Now(), Status: "pending"}
	database.DB.Create(&booking)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/bookings/1", nil)

	GetBooking(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedBooking Booking
	err := json.Unmarshal(w.Body.Bytes(), &retrievedBooking)
	assert.Nil(t, err)
	assert.Equal(t, booking.ID, retrievedBooking.ID)
}

func TestUpdateBooking(t *testing.T) {
	setup()
	defer teardown()

	// Create a booking
	booking := Booking{Customer: "John Doe", Destination: "New York", Date: time.Now(), Status: "pending"}
	database.DB.Create(&booking)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedBooking := Booking{Customer: "Jane Doe", Destination: "Los Angeles", Date: time.Now(), Status: "confirmed"}
	jsonValue, _ := json.Marshal(updatedBooking)
	c.Request = httptest.NewRequest("PUT", "/bookings/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateBooking(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedBooking Booking
	err := json.Unmarshal(w.Body.Bytes(), &retrievedBooking)
	assert.Nil(t, err)
	assert.Equal(t, updatedBooking.Customer, retrievedBooking.Customer)
}

func TestDeleteBooking(t *testing.T) {
	setup()
	defer teardown()

	// Create a booking
	booking := Booking{Customer: "John Doe", Destination: "New York", Date: time.Now(), Status: "pending"}
	database.DB.Create(&booking)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/bookings/1", nil)

	DeleteBooking(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var bookings []Booking
	database.DB.Find(&bookings)
	assert.Len(t, bookings, 0)
}

These tests cover the following scenarios:

- Creating a new booking
- Retrieving all bookings
- Retrieving a single booking by ID
- Updating a booking
- Deleting a booking

You can save these tests in a file named booking_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.