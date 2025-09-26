Here are some unit tests for your job handler functions:

func TestCreateJob(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	job := Job{}
	jsonValue, _ := json.Marshal(job)
	c.Request = httptest.NewRequest("POST", "/jobs", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateJob(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdJob Job
	err := json.Unmarshal(w.Body.Bytes(), &createdJob)
	assert.Nil(t, err)
	assert.NotZero(t, createdJob.ID)
}

func TestGetJobs(t *testing.T) {
	setup()
	defer teardown()

	// Create some jobs
	database.DB.Create(&Job{})
	database.DB.Create(&Job{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/jobs", nil)

	GetJobs(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var jobs []Job
	err := json.Unmarshal(w.Body.Bytes(), &jobs)
	assert.Nil(t, err)
	assert.Len(t, jobs, 2)
}

func TestGetJob(t *testing.T) {
	setup()
	defer teardown()

	// Create a job
	job := Job{}
	database.DB.Create(&job)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/jobs/1", nil)

	GetJob(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedJob Job
	err := json.Unmarshal(w.Body.Bytes(), &retrievedJob)
	assert.Nil(t, err)
	assert.Equal(t, job.ID, retrievedJob.ID)
}

func TestUpdateJob(t *testing.T) {
	setup()
	defer teardown()

	// Create a job
	job := Job{}
	database.DB.Create(&job)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedJob := Job{}
	jsonValue, _ := json.Marshal(updatedJob)
	c.Request = httptest.NewRequest("PUT", "/jobs/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdateJob(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedJob Job
	err := json.Unmarshal(w.Body.Bytes(), &retrievedJob)
	assert.Nil(t, err)
	assert.Equal(t, job.ID, retrievedJob.ID)
}

func TestDeleteJob(t *testing.T) {
	setup()
	defer teardown()

	// Create a job
	job := Job{}
	database.DB.Create(&job)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/jobs/1", nil)

	DeleteJob(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var jobs []Job
	database.DB.Find(&jobs)
	assert.Len(t, jobs, 0)
}

These tests cover the following scenarios:

- Creating a new job
- Retrieving all jobs
- Retrieving a single job by ID
- Updating a job
- Deleting a job

You can save these tests in a file named job_test.go and run them using the go test command. Make sure to update the database.DB variable to point to your actual database connection or a test database.