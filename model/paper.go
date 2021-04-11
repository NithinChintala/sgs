package model

type Paper struct {
	Id      int     `json:"id"`
	Year    int     `json:"year"`
	Title   string  `json:"title"`
	Journal *string  `json:"journal"`
	Volume  *int    `json:"volume"`
	Issue   *int    `json:"issue"`
	Pages   *string `json:"pages"`
	Doi     *string `json:"doi"`
}
