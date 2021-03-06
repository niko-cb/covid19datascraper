package dialogflow

const (
	// Prefecture Message
	DateSuffix               = "までの情報です\n"
	PrefecturePrefix         = "\n都道府県名:   "
	CasesPrefix              = "\n検査陽性者(今まで):   "
	ActiveCasesPrefix        = "\n検査陽性者(現在):   "
	TodayConfirmedPrefix     = "\n今日の検査陽性者:  "
	YesterdayConfirmedPrefix = "\n昨日の検査陽性者:  "
	RecoveredPrefix          = "\n回復者:   "
	DeathsPrefix             = "\n死者:   "
	CitiesPrefix             = "\n\n[市町村の公開データ]\n"

	// Symptoms Message
	Header       = "症状\n"
	CommonPrefix = "\n初期症状: "
	RarePrefix   = "\n\n人によっての症状: "
	SeverePrefix = "\n\n重篤な症状: "
)
