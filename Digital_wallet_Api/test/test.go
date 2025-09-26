Here are some unit tests for your wallet handler functions:

func TestCreateWallet(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	wallet := Wallet{}
	jsonValue, _ := json.Marshal(wallet)
	c.Request = httptest.NewRequest("POST", "/wallets", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateWallet(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdWallet Wallet
	err := json.Unmarshal(w.Body.Bytes(), &createdWallet)
	assert.Nil(t, err)
	assert.NotZero(t, createdWallet.ID)
}

func TestGetWallets(t *testing.T) {
	setup()
	defer teardown()

	// Create some wallets
	database.DB.Create(&Wallet{})
	database.DB.Create(&Wallet{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/wallets", nil)

	GetWallets(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var wallets []Wallet
	err := json.Unmarshal(w.Body.Bytes(), &wallets)
	assert.Nil(t, err)
	assert.Len(t, wallets, 2)
}

func TestGetWallet(t *testing.T) {
	setup()
	defer teardown()

	// Create a wallet
	wallet := Wallet{}
	database.DB.Create(&wallet)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/wallets/1", nil)

	GetWallet(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedWallet Wallet
	err := json.Unmarshal(w.Body.Bytes(), &retrievedWallet)
	assert.Nil(t, err)
	assert.Equal(t, wallet.ID, retrievedWallet.ID)
}

func TestUpdateWallet(t *testing.T) {
	setup()
	defer teardown()

	// Create a wallet
	wallet := Wallet{}
	database.DB.Create(&wallet)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedWallet := Wallet{}
	jsonValue, _ := json.Marshal(updatedWallet)
	c.Request = httptest.NewRequest("PUT", "/wallets/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateWallet(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedWallet Wallet
	err := json.Unmarshal(w.Body.Bytes(), &retrievedWallet)
	assert.Nil(t, err)
	assert.Equal(t, wallet.ID, retrievedWallet.ID)
}

func TestDeleteWallet(t *testing.T) {
	setup()
	defer teardown()

	// Create a wallet
	wallet := Wallet{}
	database.DB.Create(&wallet)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/wallets/1", nil)

	DeleteWallet(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var wallets []Wallet
	database.DB.Find(&wallets)
	assert.Len(t, wallets, 0)
}

These tests cover the following scenarios:

- Creating a new wallet
- Retrieving all wallets
- Retrieving a single wallet by ID
- Updating a wallet
- Deleting a wallet

You can save these tests in a file named wallet_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.