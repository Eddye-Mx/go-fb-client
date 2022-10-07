package main

import (
	"fmt"
	"os"

	//import go package for FB client
	fb "github.com/huandu/facebook"
)

//Fields constant contains a default series of metrics to be pulled on functions
const Fields = "account_id,account_name,clicks,cost_per_unique_click,cpc,cpm,cpp,ctr,frequency,impressions,objective,reach,spend,unique_clicks,unique_ctr,action_values,actions"

func main() {
	//call a function to get that level metrics on a specific date range
	results, err := GetAccount("2022-09-01", "2022-10-01")
	if err != nil {
		//print errors
		fmt.Println(err)
	} else {
		//print results
		fmt.Println(results)
	}
}

//GetAccount function returns metrics on Account level for on specific date range
func GetAccount(date_start, date_finish string) (interface{}, error) {
	//Asking insights to Facebook API using package function Get with the API path for account insigths and the Params to be Queried, ACCOUNT_ID is an enviroment variable with the ID of Facebook Ads account ID in the format 'act_XXXXXXXXXXX'
	res, err := fb.Get("/v15.0/"+os.Getenv("ACCOUNT_ID")+"/insights", fb.Params{
		//time_range contains since and until where their values must be given in format "YYYY-MM-DD" as strings
		"time_range": map[string]interface{}{
			"since": date_start,
			"until": date_finish,
		},
		//Fields to be called from API, Readme*
		"fields": Fields,
		//API Access Token
		"access_token": os.Getenv("API_TOKEN"),
	})
	//Looking for errors
	if err != nil {
		fmt.Println("Error", err)
	}
	//Returns data using Get method for fb.Result variable and returns error
	return res.Get("data"), err
}

//GetCampaign function returns metrics on Campaign level for an specified date range
func GetCampaign(date_start, date_finish string) (interface{}, error) {
	//Asking insights to Facebook API using package function Get with the API path for campaign insigths and the Params to be Queried, CAMPAIGN_ID is an enviroment variable with the ID of Facebook Ads Campaign ID in the format 'XXXXXXXXXXX'
	res, err := fb.Get("/v15.0/"+os.Getenv("CAMPAIGN_ID")+"/insights", fb.Params{
		//time_range contains since and until where their values must be given in format "YYYY-MM-DD" as strings
		"time_range": map[string]interface{}{
			"since": date_start,
			"until": date_finish,
		},
		//Fields to be called from API, Readme*
		"fields": Fields,
		//API Access Token
		"access_token": os.Getenv("API_TOKEN"),
	})
	//Looking for errors
	if err != nil {
		fmt.Println("Error", err)
	}
	//Returns data using Get method for fb.Result variable and returns error
	return res.Get("data"), err
}

//GetAdset function returns metrics on Adset level for an specified date range
func GetAdset(date_start, date_finish string) (interface{}, error) {
	//Asking insights to Facebook API using package function Get with the API path for adset insigths and the Params to be Queried, ADSET_ID is an enviroment variable with the ID of Facebook Ads Adset ID in the format 'XXXXXXXXXXX'
	res, err := fb.Get("/v15.0/"+os.Getenv("ADSET_ID")+"/insights", fb.Params{
		//time_range contains since and until where their values must be given in format "YYYY-MM-DD" as strings
		"time_range": map[string]interface{}{
			"since": date_start,
			"until": date_finish,
		},
		//Fields to be called from API, Readme*
		"fields": Fields,
		//API Access Token
		"access_token": os.Getenv("API_TOKEN"),
	})
	//Looking for errors
	if err != nil {
		fmt.Println("Error", err)
	}
	//Returns data using Get method for fb.Result variable and returns error
	return res.Get("data"), err
}

//GetAds function returns metrics on Ad level for an specified date range
func GetAds(date_start, date_finish string) (interface{}, error) {
	//Asking insights to Facebook API using package function Get with the API path for ad insigths and the Params to be Queried, AD_ID is an enviroment variable with the ID of Facebook Ads Ad ID in the format 'XXXXXXXXXXX'
	res, err := fb.Get("/v15.0/"+os.Getenv("AD_ID")+"/insights", fb.Params{
		//time_range contains since and until where their values must be given in format "YYYY-MM-DD" as strings
		"time_range": map[string]interface{}{
			"since": date_start,
			"until": date_finish,
		},
		//Fields to be called from API, Readme*
		"fields": Fields,
		//API Access Token
		"access_token": os.Getenv("API_TOKEN"),
	})
	//Looking for errors
	if err != nil {
		fmt.Println("Error", err)
	}
	//Returns data using Get method for fb.Result variable and returns error
	return res.Get("data"), err
}
