package main

import (
    "fmt"
		"os"

		. "github.com/Higakinn/trend-batch/trends"
)


func main() {
		// Argsが何件あるか出力します
		fmt.Println("count:", len(os.Args))

		// Argsの中身を一件ずつ出力します
		for i, v := range os.Args {
			fmt.Printf("args[%d] -> %s\n", i, v)
		}
    trends := map[string]TrendInterface{
			"qiita": &Qiita{Name: "qiita"},
			"zenn": &Zenn{Name: "zenn"},
    }

    for siteName, trend := range trends {
			fmt.Println(fmt.Sprintf("%s trend fetching start!\n", siteName))
			results, err := trend.FetchTrends()
			if err != nil {
				// エラー時の処理
				fmt.Println("error!!")
			}
			fmt.Println(results)
			fmt.Println(fmt.Sprintf("%s trend fetching end!\n", siteName))
    }
}