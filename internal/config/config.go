// Package config - 프로그램 설정을 위한 구조체와 메서드를 포함하고 있음
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	API API
}

type API struct {
	BaseURL   string // API 요청을 보낼 URL
	AuthKey   string // API 사용을 위해 필요한 인증 키, 해당 값은 API 요청 단계에서 자동으로 `x-api-key` 헤더에 포함됨
	RateLimit int    // API 1초당 요청 제한 개수 (ex. 해당 값이 1일 경우 초당 1회 요청)
}

// Load - 환경변수에서 값을 읽어 Config 구조체를 반환함
func Load() (*Config, error) {
	viper.AutomaticEnv()

	// 환경 변수 접두사 설정 및 키 바인드 (RIO)
	viper.SetEnvPrefix("RIO")
	if err := viper.BindEnv("api.auth_key", "API_AUTH_KEY"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("api.rate_limit", "RATE_LIMIT"); err != nil {
		return nil, err
	}

	// 기본 값 설정
	viper.SetDefault("api.base_url", "https://open-api.bser.io")
	viper.SetDefault("api.auth_key", "")
	viper.SetDefault("api.rate_limit", 1)

	// Config 구조체 생성 및 설정 값 언마샬링
	c := new(Config)
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
