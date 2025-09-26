
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Post{})
}

func teardown() {
	db.Exec("DELETE FROM posts")
}

func TestCreatePost(t *testing.T) {
	setup()
	defer teardown()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	post := Post{Title: "Test Post", Content: "This is a test post", Author: "John Doe"}
	jsonValue, _ := json.Marshal(post)
	c.Request = httptest.NewRequest("POST", "/posts", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreatePost(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdPost Post
	err := json.Unmarshal(w.Body.Bytes(), &createdPost)
	assert.Nil(t, err)
	assert.NotZero(t, createdPost.ID)
}

func TestGetPosts(t *testing.T) {
	setup()
	defer teardown()

	// Create some posts
	db.Create(&Post{Title: "Post 1", Content: "Content 1", Author: "Author 1"})
	db.Create(&Post{Title: "Post 2", Content: "Content 2", Author: "Author 2"})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/posts", nil)

	GetPosts(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var posts []Post
	err := json.Unmarshal(w.Body.Bytes(), &posts)
	assert.Nil(t, err)
	assert.Len(t, posts, 2)
}

func TestGetPost(t *testing.T) {
	setup()
	defer teardown()

	// Create a post
	post := Post{Title: "Test Post", Content: "This is a test post", Author: "John Doe"}
	db.Create(&post)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("GET", "/posts/1", nil)

	GetPost(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedPost Post
	err := json.Unmarshal(w.Body.Bytes(), &retrievedPost)
	assert.Nil(t, err)
	assert.Equal(t, post.ID, retrievedPost.ID)
}

func TestUpdatePost(t *testing.T) {
	setup()
	defer teardown()

	// Create a post
	post := Post{Title: "Test Post", Content: "This is a test post", Author: "John Doe"}
	db.Create(&post)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	updatedPost := Post{Title: "Updated Post", Content: "This is an updated post", Author: "John Doe"}
	jsonValue, _ := json.Marshal(updatedPost)
	c.Request = httptest.NewRequest("PUT", "/posts/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	UpdatePost(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedPost Post
	err := json.Unmarshal(w.Body.Bytes(), &retrievedPost)
	assert.Nil(t, err)
	assert.Equal(t, updatedPost.Title, retrievedPost.Title)
}

func TestDeletePost(t *testing.T) {
	setup()
	defer teardown()

	// Create a post
	post := Post{Title: "Test Post", Content: "This is a test post", Author: "John Doe"}
	db.Create(&post)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	c.Request = httptest.NewRequest("DELETE", "/posts/1", nil)

	DeletePost(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var posts []Post
	db.Find(&posts)
	assert.Len(t, posts, 0)
}

/*These tests cover the following scenarios:

- Creating a new post
- Retrieving all posts
- Retrieving a single post by ID
- Updating a post
- Deleting a post

Each test sets up a test database, creates a test context, and calls the handler function being tested. The test then asserts
*/