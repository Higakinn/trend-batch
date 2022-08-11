package trends

type Zenn struct {
	Name string
}

func (z *Zenn) FetchTrend() (TrendResponse, error){
	res := TrendResponse{
		Title: "zenn",
		Url: "http://zenn.com",
		TwitterUrl: "https://twitter.com",
	}
	return res,nil
}
