Here are some unit tests for your loan handler functions:

func TestCreateLoan(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	loan := Loan{Borrower: "John Doe", Amount: 1000.0, Status: "pending", DueDate: time.Now().AddDate(0, 0, 30)}
	jsonValue, _ := json.Marshal(loan)
	c.Request = httptest.NewRequest("POST", "/loans", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateLoan(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdLoan Loan
	err := json.Unmarshal(w.Body.Bytes(), &createdLoan)
	assert.Nil(t, err)
	assert.NotZero(t, createdLoan.ID)
}

func TestGetLoans(t *testing.T) {
	setup()
	defer teardown()

	// Create some loans
	database.DB.Create(&Loan{Borrower: "John Doe", Amount: 1000.0, Status: "pending", DueDate: time.Now().AddDate(0, 0, 30)})
	database.DB.Create(&Loan{Borrower: "Jane Doe", Amount: 500.0, Status: "approved", DueDate: time.Now().AddDate(0, 0, 60)})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/loans", nil)

	GetLoans(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var loans []Loan
	err := json.Unmarshal(w.Body.Bytes(), &loans)
	assert.Nil(t, err)
	assert.Len(t, loans, 2)
}

func TestGetLoan(t *testing.T) {
	setup()
	defer teardown()

	// Create a loan
	loan := Loan{Borrower: "John Doe", Amount: 1000.0, Status: "pending", DueDate: time.Now().AddDate(0, 0, 30)}
	database.DB.Create(&loan)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/loans/1", nil)

	GetLoan(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedLoan Loan
	err := json.Unmarshal(w.Body.Bytes(), &retrievedLoan)
	assert.Nil(t, err)
	assert.Equal(t, loan.ID, retrievedLoan.ID)
}

func TestUpdateLoan(t *testing.T) {
	setup()
	defer teardown()

	// Create a loan
	loan := Loan{Borrower: "John Doe", Amount: 1000.0, Status: "pending", DueDate: time.Now().AddDate(0, 0, 30)}
	database.DB.Create(&loan)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedLoan := Loan{Borrower: "Jane Doe", Amount: 500.0, Status: "approved", DueDate: time.Now().AddDate(0, 0, 60)}
	jsonValue, _ := json.Marshal(updatedLoan)
	c.Request = httptest.NewRequest("PUT", "/loans/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateLoan(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedLoan Loan
	err := json.Unmarshal(w.Body.Bytes(), &retrievedLoan)
	assert.Nil(t, err)
	assert.Equal(t, updatedLoan.Borrower, retrievedLoan.Borrower)
}

func TestDeleteLoan(t *testing.T) {
	setup()
	defer teardown()

	// Create a loan
	loan := Loan{Borrower: "John Doe", Amount: 1000.0, Status: "pending", DueDate: time.Now().AddDate(0, 0, 30)}
	database.DB.Create(&loan)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/loans/1", nil)

	DeleteLoan(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var loans []Loan
	database.DB.Find(&loans)
	assert.Len(t, loans, 0)
}

These tests cover the following scenarios:

- Creating a new loan
- Retrieving all loans
- Retrieving a single loan by ID
- Updating a loan
- Deleting a loan

You can save these tests in a file named loan_test.go and run them using the go test command. Make sure to update