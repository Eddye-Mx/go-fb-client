# Connection and use of Facebook Insights API V15
## Getting token and IDs to connect

1. In order to start pulling insights from Facebook API you need to prepare:
- Your Facebook account with access to Facebook Ads Accounts

   ![Captura de Pantalla 2022-10-07 a la(s) 12 07 23](https://user-images.githubusercontent.com/99297157/194613035-0a7df68c-e86c-4a25-90ad-2d956f476b51.png)
   
- Create a new app on facebook developers https://developers.facebook.com/apps/

   ![Captura de Pantalla 2022-10-07 a la(s) 12 10 43](https://user-images.githubusercontent.com/99297157/194612996-0759209e-dd56-482e-9baa-f4d95c495092.png)
   
- Enable the marketing API in the products section 
   
   ![Captura de Pantalla 2022-10-07 a la(s) 12 14 39](https://user-images.githubusercontent.com/99297157/194613836-ddd3fc1b-6675-4a97-aef5-9acf3dfa40c9.png)

- Generate an API token from https://developers.facebook.com/apps > Marketing API > Tools > Generate Token, save the token cause it won't be saved by facebook

   ![Captura de Pantalla 2022-10-07 a la(s) 12 18 01](https://user-images.githubusercontent.com/99297157/194614210-c2aabe6a-1e3c-4192-a848-15a1f126ec83.png)

- Get your Facebook Ads Account ID on https://business.facebook.com/

2. For debugging you can start using [Graph API](https://developers.facebook.com/tools/explorer/2038369806366978/) from FB tools
   
![Captura de Pantalla 2022-10-07 a la(s) 12 25 25](https://user-images.githubusercontent.com/99297157/194615615-2823cab1-1ef4-4390-953b-a6e36c4ebc3e.png)

## Making Requests
Requests have an specific path to call insights from each level, see [documentation](https://developers.facebook.com/docs/marketing-api/insights/)

![Captura de Pantalla 2022-10-07 a la(s) 12 39 33](https://user-images.githubusercontent.com/99297157/194617527-a474f7b6-f261-4a22-bf19-372c85ea8a99.png)

[Here](https://developers.facebook.com/docs/marketing-api/reference/ads-insights/) is the documentation of each Field and Param you can pull from API

For this code example the response is given in the format
```[map[
  account_id:XXXXXXXXXXXX account_name:XXXXXXXXXX action_values: [
    map[action_type: offsite_conversion.fb_pixel_purchase value:611156954.09] 
    map[action_type:omni_purchase value:611156954.09]] 
    actions:[
      map[action_type:photo_view value:231] 
      map[action_type:like value:8] 
      map[action_type:comment value:1751] 
      map[action_type:onsite_conversion.post_save value:878] 
      map[action_type:link_click value:170227] 
      map[action_type:post value:1881] 
      map[action_type:post_reaction value:28507] 
      map[action_type:video_view value:1691134] 
      map[action_type:landing_page_view value:82813] 
      map[action_type:offsite_conversion.fb_pixel_add_to_cart value:525221]
      map[action_type:offsite_conversion.fb_pixel_initiate_checkout value:398657]
      map[action_type:offsite_conversion.fb_pixel_purchase value:153230] 
      map[action_type:offsite_conversion.fb_pixel_search value:847662] 
      map[action_type:offsite_conversion.fb_pixel_view_content value:2022296] 
      map[action_type:post_engagement value:1894609] 
      map[action_type:page_engagement value:1894617] 
      map[action_type:omni_add_to_cart value:525221] 
      map[action_type:omni_initiated_checkout value:398657] 
      map[action_type:omni_purchase value:153230] 
      map[action_type:omni_search value:847662] 
      map[action_type:omni_view_content value:2022296] 
      map[action_type:add_to_cart value:525221] 
      map[action_type:initiate_checkout value:398657] 
      map[action_type:purchase value:153230] 
      map[action_type:search value:847662] 
      map[action_type:view_content value:2022296]] 
    clicks:347894 
    cost_per_unique_click:2.891627 
    cpc:2.128745 
    cpm:15.998378 
    cpp:56.395731 
    ctr:0.75154 
    date_start:2022-09-01 
    date_stop:2022-10-01 
    frequency:3.525091 
    impressions:46290788 
    objective:MULTIPLE 
    reach:13131801 
    spend:740577.52 
    unique_clicks:256111 
    unique_ctr:1.950311]
  ]
```