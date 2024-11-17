package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestLoad_WithEnvVars(t *testing.T) {
	var (
		TestBaseURL   = "https://test.api.com"
		TestAuthKey   = "123456789"
		TestRateLimit = 30
	)
	// 환경 변수 설정
	os.Setenv("API_BASE_URL", TestBaseURL)
	os.Setenv("API_AUTH_KEY", TestAuthKey)
	os.Setenv("API_RATE_LIMIT", strconv.Itoa(TestRateLimit))

	// 테스트 종료 후 환경 변수 해제
	defer os.Unsetenv("API_BASE_URL")
	defer os.Unsetenv("API_AUTH_KEY")
	defer os.Unsetenv("API_RATE_LIMIT")

	// LoadConfig 호출
	config, err := Load()
	assert.NoError(t, err, "Error in LoadConfig")

	// 환경 변수 기반 설정 값 검증
	assert.Equal(t, TestBaseURL, config.API.BaseURL)
	assert.Equal(t, TestAuthKey, config.API.AuthKey)
	assert.Equal(t, TestRateLimit, config.API.RateLimit)
}

func TestLoad_WithDefaults(t *testing.T) {
	// 환경 변수 없이 기본값만으로 설정

	// LoadConfig 호출
	config, err := Load()
	assert.NoError(t, err, "LoadConfig 함수가 오류 없이 실행되어야 합니다")

	// 기본값 검증
	assert.Equal(t, DefaultAPIBaseURL, config.API.BaseURL)
	assert.Equal(t, DefaultAPIAuthKey, config.API.AuthKey)
	assert.Equal(t, DefaultAPIRateLimit, config.API.RateLimit)
}
