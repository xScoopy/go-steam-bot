# go-steam-bot

## Project Description 

This project was created to provide just a tiny bit of relief to those gamers grinding it out for 8+ hours a day. Through this go steam bot, a list of discounted top selling games on steam can be pulled up with just a few characters typed into a designated slack channel. Since most offices expect people to have high utilization rates, we can't have everyone surfing the steam store at work. This bot allows a quick scrape of steam store discounts so my fellow gamers can keep an eye on any deals too good to pass up while on their lunch break, waiting for a build to finish, etc., all while still in the slack workspace. 

## Install instructions 

1. Create a new slack app in your slack workspace. (may need admin privileges to do this)
2. Under the 'add features and functionality' section, choose Bots. 
3. Provide permissions to your bot based on the choices and slack documentation. It must at least have the read:write permissions for a channel it is a member of for this application. 
4. Keep note of the bot token provided to you after creation, this is needed later. 
5. Enable SocketMode in the app settings in the Slack UI. 
6. Enable Event Subscriptions in the app settings in the Slack UI. (ensure connections:write permissions are granted)
7. An app-level token will be generated, keep note of that as well so we can communicate with the slack events API. 
8. Clone this repo into a new directory. 
9. Retrieve the channel ID in which you want this bot to work in. The channel ID can be found by looking at channel details in the slack app, or in the url if using the web ui for slack. 
10. Invite the bot to the desired channel. 
11. In the project directory root, create a *.env* file. within that file, place the following key names with their corresponding tokens.   

GOSECRET=botToken

CHANNELID=channelID 

STEAMBOT-WEBSOCKET=appToken

## How to use

After install instructions have been completed above, navigate to the project directory and run the following command : 

> go run scrapeBot.go 

Use @ to mention the bot by name in the same channel whose ID you used in the .env file. Once mentioned, the bot will scrape and serve you the top selling discounted games on the steam store, and print them out to the slack channel. 

### Gif of working slackbot 
![alt text](Mar-01-2022 20-31-58.gif "Demo Giphy")
