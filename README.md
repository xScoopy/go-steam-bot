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
9. In the project directory root, create a *.env* file. within that file, place the following key names with their corresponding tokens. \n GOSECRET=botToken \n CHANNELID=channelID \n STEAMBOT-WEBSOCKET=appToken

## How to use

### Gif of working slackbot 
