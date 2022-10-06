package main

import (
	"fmt"
	"os"

	fb "github.com/huandu/facebook"
)

//5044019125631592
func main() {

	fmt.Println(GetAds())

}

func GetAccount() interface{} {
	res, err := fb.Get("/v15.0/"+os.Getenv("ACCOUNT_ID")+"/insights", fb.Params{
		"time_range": map[string]interface{}{
			"since": "2022-09-25",
			"until": "2022-10-01",
		},
		"fields":       "impressions,clicks",
		"access_token": os.Getenv("API_TOKEN"),
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	return res.Get("data")
}

func GetCampaign() interface{} {
	res, err := fb.Get("/v15.0/"+os.Getenv("CAMPAIGN_ID")+"/insights", fb.Params{
		"time_range": map[string]interface{}{
			"since": "2022-09-25",
			"until": "2022-10-01",
		},
		"fields":       "impressions,reach",
		"access_token": os.Getenv("API_TOKEN"),
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	return res.Get("data")
}

func GetAdset() interface{} {
	res, err := fb.Get("/v15.0/"+os.Getenv("ADSET_ID")+"/insights", fb.Params{
		"time_range": map[string]interface{}{
			"since": "2022-09-25",
			"until": "2022-10-01",
		},
		"fields":       "impressions,reach",
		"access_token": os.Getenv("API_TOKEN"),
	})
	if err != nil {
		fmt.Println("Error", err)
	}

	return res.Get("data")
}

func GetAds() interface{} {
	res, err := fb.Get("/v15.0/"+os.Getenv("AD_ID")+"/insights", fb.Params{
		"time_range": map[string]interface{}{
			"since": "2022-09-25",
			"until": "2022-10-01",
		},
		"fields":       "impressions,reach",
		"access_token": os.Getenv("API_TOKEN"),
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	return res.Get("data")
}
