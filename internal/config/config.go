// Package config - 프로그램 설정을 위한 구조체와 메서드를 포함하고 있음
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	DefaultAPIBaseURL   = "https://open-api.bser.io"
	DefaultAPIAuthKey   = ""
	DefaultAPIRateLimit = 1
)

type Config struct {
	API API `mapstructure:"api"`
}

type API struct {
	BaseURL   string `mapstructure:"base_url"`   // 이터널 리턴 API 주소
	AuthKey   string `mapstructure:"auth_key"`   // 이터널 리턴 API 인증 키
	RateLimit int    `mapstructure:"rate_limit"` // 이터널 리턴 API 제한량 (초당 허용 요청 횟수)
}

// Load - 설정 파일과 환경변수에서 값을 읽어 Config 구조체를 반환함
func Load() (*Config, error) {
	v := viper.New()

	// 환경 변수 접두사 설정 및 키 바인드
	v.BindEnv("api.base_url", "API_BASE_URL")
	v.BindEnv("api.auth_key", "API_AUTH_KEY")
	v.BindEnv("api.rate_limit", "API_RATE_LIMIT")

	// 기본 값 설정
	v.SetDefault("api.base_url", DefaultAPIBaseURL)
	v.SetDefault("api.auth_key", DefaultAPIAuthKey)
	v.SetDefault("api.rate_limit", DefaultAPIRateLimit)

	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Config 구조체 생성 및 설정 값 언마샬링
	c := new(Config)
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, validate(c)
}

func validate(c *Config) error {
	if c.API.BaseURL == "" {
		return fmt.Errorf("API base URL is required")
	}
	if c.API.AuthKey == "" {
		return fmt.Errorf("API auth key is required")
	}
	if c.API.RateLimit <= 0 {
		return fmt.Errorf("API rate limit must be greater than 0")
	}
	return nil
}
