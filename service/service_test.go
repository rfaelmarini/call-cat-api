package service

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rfaelmarini/call-cat-api/entity"
	"github.com/rfaelmarini/call-cat-api/repository"
)

func TestService(t *testing.T) {
	os.Setenv("API_KEY", "1a9c1e22-9dc7-48fa-844c-5d137e80694")
	os.Setenv("JWT_SECRET", "1a9c1e22-9dc7-48fa-844c-5d137e80694")
	os.Setenv("DB_NAME", "callcatapitestdb")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_ADDRESS", "127.0.0.1:3306")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

const (
	REQUESTED_URL = "https://api.thecatapi.com/v1/breeds/search"
	BODY          = "Response body"
	STATUS_CODE   = 200
)

var testResponse entity.Response = entity.Response{
	RequestedURL: REQUESTED_URL,
	Body:         BODY,
	StatusCode:   STATUS_CODE,
}

var _ = Describe("Response Service", func() {
	var (
		responseRepository repository.ResponseRepository
		responseService    ResponseService
	)

	BeforeEach(func() {
		responseRepository = repository.NewResponseRepository()
		responseService = NewResponseService(responseRepository)
	})

	Context("Save and retrieve a new response", func() {
		BeforeEach(func() {
			responseService.Save(testResponse)
		})

		It("should new entry exits", func() {
			result := responseService.Find(testResponse.RequestedURL)

			Ω(result.RequestedURL).ShouldNot(BeEmpty())
		})

		AfterEach(func() {
			responseRepository.Delete(testResponse)
		})
	})
})

var _ = Describe("Login Service", func() {
	Context("Login user", func() {
		It("should pass trough login action", func() {
			loginService := NewLoginService()
			ok := LoginService.Login(loginService, "admin", "@#$RF@!718")
			Ω(ok).Should(BeTrue())
		})

		It("should not pass trough login action", func() {
			loginService := NewLoginService()
			ok := LoginService.Login(loginService, "user", "@#$RF@!918")
			Ω(ok).Should(BeFalse())
		})
	})
})

var _ = Describe("JWT Service", func() {
	Context("Create and validate", func() {
		It("should create a valid JWT Token", func() {
			jwtService := JWTAuthService()
			token := jwtService.GenerateToken("admin", true)
			Ω(token).ShouldNot(BeEmpty())

			result, _ := jwtService.ValidateToken(token)
			Ω(result.Valid).Should(BeTrue())
		})
	})
})
