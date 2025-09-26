Here are some unit tests for your product handler functions:

func TestCreateProduct(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	product := Product{Name: "Test Product", Price: 19.99, Description: "This is a test product", Stock: 10}
	jsonValue, _ := json.Marshal(product)
	c.Request = httptest.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateProduct(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdProduct Product
	err := json.Unmarshal(w.Body.Bytes(), &createdProduct)
	assert.Nil(t, err)
	assert.NotZero(t, createdProduct.ID)
}

func TestGetProducts(t *testing.T) {
	setup()
	defer teardown()

	// Create some products
	database.DB.Create(&Product{Name: "Product 1", Price: 9.99, Description: "This is product 1", Stock: 5})
	database.DB.Create(&Product{Name: "Product 2", Price: 14.99, Description: "This is product 2", Stock: 10})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/products", nil)

	GetProducts(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var products []Product
	err := json.Unmarshal(w.Body.Bytes(), &products)
	assert.Nil(t, err)
	assert.Len(t, products, 2)
}

func TestGetProduct(t *testing.T) {
	setup()
	defer teardown()

	// Create a product
	product := Product{Name: "Test Product", Price: 19.99, Description: "This is a test product", Stock: 10}
	database.DB.Create(&product)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/products/1", nil)

	GetProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedProduct Product
	err := json.Unmarshal(w.Body.Bytes(), &retrievedProduct)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, retrievedProduct.ID)
}

func TestUpdateProduct(t *testing.T) {
	setup()
	defer teardown()

	// Create a product
	product := Product{Name: "Test Product", Price: 19.99, Description: "This is a test product", Stock: 10}
	database.DB.Create(&product)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedProduct := Product{Name: "Updated Product", Price: 24.99, Description: "This is an updated product", Stock: 15}
	jsonValue, _ := json.Marshal(updatedProduct)
	c.Request = httptest.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedProduct Product
	err := json.Unmarshal(w.Body.Bytes(), &retrievedProduct)
	assert.Nil(t, err)
	assert.Equal(t, updatedProduct.Name, retrievedProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	setup()
	defer teardown()

	// Create a product
	product := Product{Name: "Test Product", Price: 19.99, Description: "This is a test product", Stock: 10}
	database.DB.Create(&product)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/products/1", nil)

	DeleteProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var products []Product
	database.DB.Find(&products)
	assert.Len(t, products, 0)
}

These tests cover the following scenarios:

- Creating a new product
- Retrieving all products
- Retrieving a single product by ID
- Updating a product
- Deleting a product

You can save these tests in a file named product_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.