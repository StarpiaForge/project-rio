package model

import "time"

// Season - 시즌 메타 정보
//
// [GET] https://open-api.bser.io/v2/data/Season
type Season struct {
	ID        int       `json:"seasonID"`
	Name      string    `json:"seasonName"`
	Start     time.Time `json:"seasonStart"`
	End       time.Time `json:"seasonEnd"`
	IsCurrent int       `json:"isCurrent"`
}

// Character - 캐릭터 메타 정보, 불필요한 수치 값들은 최적화를 위해 제거됨
//
// [GET] https://open-api.bser.io/v2/data/Character
type Character struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}
