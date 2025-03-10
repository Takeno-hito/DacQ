package model

type Score struct {
	User string  `gorm:"primaryKey"`
	Loss float64 `gorm:"not null,index"`
}

func CreateScore(score Score) error {
	return db.Create(&score).Error
}

func GetScoreByUser(user string) (Score, error) {
	var score Score
	err := db.Where("user = ?", user).First(&score).Error
	return score, err
}

func GetRankings() ([]Score, error) {
	var scores []Score
	err := db.Order("loss asc").Find(&scores).Error
	return scores, err
}

func UpdateScore(user string, val float64) error {
	return db.Model(&Score{}).Where("user = ?", user).Update("loss", val).Error
}
