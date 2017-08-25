package mobile

// SearchResponse defines structure of searching request response
type SearchResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Content struct {
		Data struct {
			Custom struct {
				City         string `json:"city"`
				PositionName string `json:"positionName"`
			} `json:"custom"`
			Page struct {
				PageNo     int        `json:"pageNo"`
				PageSize   int        `json:"pageSize"`
				Start      string     `json:"start"`
				TotalCount string     `json:"totalCount"`
				Result     []Position `json:"result"`
			} `json:"page"`
		} `json:"data"`
	} `json:"content"`
}
