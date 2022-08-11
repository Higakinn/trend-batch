package trends

type Zenn struct {
	Name string
}

func (z *Zenn) FetchTrends() ([]TrendResponse, error){
	var results []TrendResponse
	res := TrendResponse{
		Title: "zenn",
		Url: "http://zenn.com",
		TwitterUrl: "https://twitter.com",
	}
	results = append(results, res)
	return results,nil
}
