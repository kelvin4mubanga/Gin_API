
func TestCreateTodo(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	todo := Todo{Task: "Test Todo", Completed: false}
	jsonValue, _ := json.Marshal(todo)
	c.Request = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateTodo(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdTodo Todo
	err := json.Unmarshal(w.Body.Bytes(), &createdTodo)
	assert.Nil(t, err)
	assert.NotZero(t, createdTodo.ID)
}

func TestGetTodos(t *testing.T) {
	setup()
	defer teardown()

	// Create some todos
	database.DB.Create(&Todo{Task: "Todo 1", Completed: false})
	database.DB.Create(&Todo{Task: "Todo 2", Completed: true})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/todos", nil)

	GetTodos(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var todos []Todo
	err := json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Nil(t, err)
	assert.Len(t, todos, 2)
}

func TestGetTodo(t *testing.T) {
	setup()
	defer teardown()

	// Create a todo
	todo := Todo{Task: "Test Todo", Completed: false}
	database.DB.Create(&todo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/todos/1", nil)

	GetTodo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedTodo Todo
	err := json.Unmarshal(w.Body.Bytes(), &retrievedTodo)
	assert.Nil(t, err)
	assert.Equal(t, todo.ID, retrievedTodo.ID)
}

func TestUpdateTodo(t *testing.T) {
	setup()
	defer teardown()

	// Create a todo
	todo := Todo{Task: "Test Todo", Completed: false}
	database.DB.Create(&todo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedTodo := Todo{Task: "Updated Todo", Completed: true}
	jsonValue, _ := json.Marshal(updatedTodo)
	c.Request = httptest.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateTodo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedTodo Todo
	err := json.Unmarshal(w.Body.Bytes(), &retrievedTodo)
	assert.Nil(t, err)
	assert.Equal(t, updatedTodo.Task, retrievedTodo.Task)
}

func TestDeleteTodo(t *testing.T) {
	setup()
	defer teardown()

	// Create a todo
	todo := Todo{Task: "Test Todo", Completed: false}
	database.DB.Create(&todo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/todos/1", nil)

	DeleteTodo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var todos []Todo
	database.DB.Find(&todos)
	assert.Len(t, todos, 0)
}
/*
These tests cover the following scenarios:

- Creating a new todo
- Retrieving all todos
- Retrieving a single todo by ID
- Updating a todo
- Deleting a todo
*/
