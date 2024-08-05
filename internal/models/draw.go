package models

type LotteryResult struct {
	DrawID   int    `json:"draw_id"`
	DrawDate string `json:"draw_date"`
	Number1  string `json:"number_1"`
	Number2  string `json:"number_2"`
	Number3  string `json:"number_3"`
	Number4  string `json:"number_4"`
	Number5  string `json:"number_5"`
}
