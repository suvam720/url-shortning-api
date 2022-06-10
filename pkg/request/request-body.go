package request

type RequestBody struct {
	ReqUrl string `json:"url" gorm:"unique"`
}
