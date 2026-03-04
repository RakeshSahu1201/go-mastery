package kudohandlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"main/kudoroutes"
)

// setupRouter initializes the Gin router with our actual routes
// so that our tests exercise the exact same routing logic as production.
func setupRouter() *gin.Engine {
	// Switch to test mode so Gin doesn't spam standard output
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	kudoroutes.RegisterUserRoutes(r)
	return r
}

// TestMain runs once before any tests in this package start.
// We use it to load the .env file so kudostore.NewClient() inside the
// handlers can successfully connect to the local test database.
func TestMain(m *testing.M) {
	// The test runs from the kudohandlers/ folder, so .env is one level up
	_ = godotenv.Load("../.env")
	os.Exit(m.Run())
}

// TestGetUsers executes an integration test against GET /users/
func TestGetUsers(t *testing.T) {
	router := setupRouter()

	// 1. Create a dummy HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users/", nil)
	require.NoError(t, err)

	// 2. Create an HTTP ResponseRecorder (acts like a fake browser to capture the response)
	w := httptest.NewRecorder()

	// 3. Serve the HTTP request to our router
	router.ServeHTTP(w, req)

	// 4. Assert the results!
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK")

	// We expect a JSON array back. It might be empty `[]`, or it might have users
	// depending on the state of your local database.
	var responseBody []map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	require.NoError(t, err, "Response should be valid JSON")
}

// TestCreateUser executes an integration test against POST /users/
func TestCreateUser(t *testing.T) {
	router := setupRouter()

	// We use a random email so the test doesn't fail with "email already exists"
	// if we run it multiple times.
	payload := []byte(`{"name":"Integration Test", "email":"test-random-` + randomString() + `@example.com"}`)

	req, err := http.NewRequest(http.MethodPost, "/users/", bytes.NewBuffer(payload))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP 201 Created")

	// Verify the response contains the new user ID
	var responseBody map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	require.NoError(t, err)

	assert.NotNil(t, responseBody["id"], "Response should include the created user's ID")
	assert.Equal(t, "Integration Test", responseBody["name"])
}

// Helper to prevent duplicate email constraint errors across test runs
func randomString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
