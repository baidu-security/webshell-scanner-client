package scanner

type EnqueueResponse struct {
	Status int    `json:"status"`
	Descr  string `json:"descr"`
	Md5    string `json:"md5"`
	Url    string `json:"url"`
}

type ResultDataResponse struct {
	Path  string `json:"path"`
	Descr string `json:"descr"`
}

type ResultResponse struct {
	Md5      string               `json:"md5"`
	Status   string               `json:"status"`
	Total    int                  `json:"total"`
	Scanned  int                  `json:"scanned"`
	Detected int                  `json:"detected"`
	Data     []ResultDataResponse `json:"data"`
}
