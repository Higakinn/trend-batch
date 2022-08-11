package trends

type Qiita struct {
	Name string
}

func (q *Qiita) FetchTrend() (TrendResponse, error){
	res := TrendResponse{
		Title: "test",
		Url: "http://qiita.com",
		TwitterUrl: "https://twitter.com",
	}
	return res,nil
}


