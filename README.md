# platebot
gyms are closed. covid-19 has caused a frenzy for gym equipment. i need a bot to tell me when plates become available for purchase.

Currently alerting is only available through twilio. Go to twilio.com and sign up for a free account. Use the sid and token to use at launch. If you don't want sms or dont want to sign up for twilio, just put fake info into the flags. 

# Command line usage

```
go build && ./platebot -sid <twilio sid> -token <twilio token> -to <your phone number> -from <your twilio number>
```

# Example stdout

```
25LB Machined Olympic - Pair https://www.roguefitness.com/rogue-machined-olympic-plates
35LB Machined Olympic - Pair https://www.roguefitness.com/rogue-machined-olympic-plates
alert successfully sent
completed search in 15.692020901s
25LB Machined Olympic - Pair https://www.roguefitness.com/rogue-machined-olympic-plates
35LB Machined Olympic - Pair https://www.roguefitness.com/rogue-machined-olympic-plates
message not sent, throttled
```
