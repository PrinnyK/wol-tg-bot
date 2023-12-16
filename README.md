# wol-tg-bot
## 介绍
基于Golang实现的Telegram Bot，监听特定消息触发
Wake-on-LAN（WoL）唤醒目标设备。例如，可部署家用路由器或其他设备上，在外远程唤醒家中PC。

## 使用
### 创建Bot
[Telegram Bot 申请教程](https://core.telegram.org/bots/features#botfather)
### 部署Bot
```bash
# 下载代码
git clone https://github.com/PrinnyK/wol-tg-bot.git
cd wol-tg-bot

# 编译本机平台执行
go build -o wol-tg-bot
# 或者交叉编译给其他平台执行 如arm芯片的linux设备
GOOS=linux GOARCH=arm go build -o wol-tg-bot

# 准备配置文件 config.json
{
  "botToken": "your-telegram-bot-api-token",
  "validUserNameList": ["telegram-username-with-permission"],
  "targetMacAddr": "wake-target-mac-address",
  "targetIpAddr": "255.255.255.255",  // 非必填 默认255.255.255.255
  "targetPort": 9  // 非必填 默认9
}

# 将config.json和编译产物wol-tg-bot放同一目录下，在对应平台运行
path/to/your/wol-tg-bot
```
> 另提供 Dockerfile 构建镜像，注意改下编译的目标平台。docker运行前将 config.json 映射到 /app 目录下即可
### 使用Bot
在 Telegram 给你的 Bot 发送 /power 即可唤醒目标设备，Bot 触发后会回复 Done
