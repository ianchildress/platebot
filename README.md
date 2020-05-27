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

## follow up
Using this bot I was alerted one morning that the Ohio Power Bar was in stock. I ran downstairs and successfully ordered! It also showed plates were in stock despite the website showing they were not in stock. Believing in the bot, I continued to refresh and eventually the page showed items in stock and I was able to place multiple orders during the small availability window. It's my suspicion that the products page updates in advance of the items already being available to place in the cart. For example, on the Olympic Plates product page it showed all items as unavailable to be placed in the cart but the bot kept alerting that the items were there. After inspecting the source I found that items that I was being alerted to had this setting:

```
optionStockStatus[7179] = 1;
```

but other items that also showed unavailable had the following setting:

```
optionStockStatus[7181] = 0;
```

So the page appears to update prior to inventory being available for placement in cart. 
