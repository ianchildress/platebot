# platebot
gyms are closed. covid-19 has caused a frenzy for gym equipment. i need a bot to tell me when plates become available for purchase.

Currently alerting is only available through twilio. Go to twilio.com and sign up for a free account. Use the sid and token to use at launch. If you don't want sms or dont want to sign up for twilio, just put fake info into the flags. 

```
go build && ./platebot -sid <twilio sid> -token <twilio token> -to <your phone number> -from <your twilio number>
```
