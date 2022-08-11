package trends

type TrendResponse struct{
		Title string
		Url string
		TwitterUrl string
}

type TrendInterface interface {
		FetchTrend() (TrendResponse, error)
}