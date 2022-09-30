package test

import (
	"log"
	"strconv"
	"testing"

	"github.com/bitwyre/template-golang/pkg/datastore/postgres/entity"
	"github.com/bitwyre/template-golang/pkg/datastore/postgres/seeder"
	"github.com/bitwyre/template-golang/pkg/test/config"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"net/http"
	"net/http/httptest"
)

type GetUserTestSuite struct {
	suite.Suite
	baseUrl string
	client  test_config.ClientTestSuite
	entity  *entity.User
}

func TestGetOTPTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserTestSuite))
}

func (suite *GetUserTestSuite) SetupSuite() {
	app := test_config.BootstrapAppTest()

	suite.client.Db = app.Db
	suite.client.Gin = app.Gin
	suite.baseUrl = "/user/"
}

func (suite *GetUserTestSuite) SetupTest() {
	suite.entity = seeder.UserDataSeeder(suite.client.Db).CreateOne()
}

func (suite *GetUserTestSuite) AfterTest(_, _ string) {
	// Hard delete to prevent leaving garbage data
	err := suite.client.Db.Unscoped().Delete(&suite.entity).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *GetUserTestSuite) TestErrorUserNotFound() {
	userId := "9999"
	req, err := http.NewRequest("GET", suite.baseUrl+userId, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	suite.client.Gin.ServeHTTP(w, req)

	errorMsg := gjson.Get(w.Body.String(), "error")
	suite.Equal("NOT_FOUND", errorMsg.Get("code").Str)
	suite.Equal("User is not found", errorMsg.Get("message").Str)
}

func (suite *GetUserTestSuite) TestSuccessResendOTP() {
	userId := strconv.Itoa(int(suite.entity.ID))
	req, err := http.NewRequest("GET", suite.baseUrl+userId, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	suite.client.Gin.ServeHTTP(w, req)
	jsonData := gjson.Get(w.Body.String(), "results")

	suite.Equal(jsonData.Get("ID").Raw, userId)
	suite.NotEmpty(jsonData.Get("Email").Raw)
	suite.NotEmpty(jsonData.Get("Status").Raw)
}
