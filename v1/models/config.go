// author gmfan
// date 2023/11/25
package models

type (
	Config struct {
		Header Header `json:"header"`
		Links  []Link `json:"links"`
	}

	Header struct {
		Icon     string `json:"icon"`
		Title    string `json:"title"`
		SubTitle string `json:"subTitle"`
	}

	Link struct {
		Icon  string `json:"icon"`
		Title string `json:"title"`
		Url   string `json:"url"`
	}
)
