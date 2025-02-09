# 🚀 Discord Bot IP Whitelister - Linux Service Setup

This guide explains how to **install, configure, and run** the Discord bot as a **systemd service** on a Linux server.


## 📌 Prerequisites
Before proceeding, ensure you have:
- A **Linux server** (Ubuntu, Debian, CentOS, etc.)
- **Go installed** (for building the binary)
- **Systemd** (default in most modern Linux distributions)
- **A valid `.env` file** with bot credentials



## 🛠️ Step 1: Build the Go Application

If you haven’t already compiled the bot, run:

```sh
go build -o bot-app
```

Move the binary to a suitable directory:

```sh
sudo mv bot-app /usr/local/bin/
```

Ensure the binary has execution permissions:

```sh
sudo chmod +x /usr/local/bin/bot-app
```

---

## 📁 Step 2: Create an Environment Variables File

To securely store environment variables, create a dedicated environment file:

```sh
sudo nano /etc/default/discord-bot
```

Add your required environment variables:

```ini
BOT_TOKEN=your_bot_token_here
BOT_GUILD_ID=your_guild_id
SERVICE_PORTS=80,443
DELETE_COMMADS=true
```

Ensure the file is readable by the service user:

```sh
sudo chmod 600 /etc/default/discord-bot
```

---

## ⚙️ Step 3: Create a Systemd Service File

To manage the bot as a Linux service, create a `systemd` unit file:

```sh
sudo nano /etc/systemd/system/discord-bot.service
```

Paste the following configuration:

```ini
[Unit]
Description=Discord Bot IP Whitelister
After=network.target

[Service]
Type=simple
User=your_linux_user  # Change this to the appropriate user
Group=your_linux_user
WorkingDirectory=/usr/local/bin
ExecStart=/usr/local/bin/bot-app
Restart=always
RestartSec=5
EnvironmentFile=/etc/default/discord-bot  # Load environment variables

[Install]
WantedBy=multi-user.target
```

Save and exit.

Reload `systemd` to apply the changes:

```sh
sudo systemctl daemon-reload
```

---

## 🚀 Step 4: Start and Enable the Service

Start the bot service:

```sh
sudo systemctl start discord-bot
```

Enable it to run on boot:

```sh
sudo systemctl enable discord-bot
```

Check if the service is running:

```sh
sudo systemctl status discord-bot
```

You should see output similar to:

```sh
● discord-bot.service - Discord Bot IP Whitelister
   Loaded: loaded (/etc/systemd/system/discord-bot.service; enabled; vendor preset: enabled)
   Active: active (running) since Sun 2025-02-08 14:00:00 UTC; 2s ago
   Main PID: 12345 (bot-app)
   CGroup: /system.slice/discord-bot.service
```

---

## 📊 Step 5: Viewing Logs

To monitor logs in real time:

```sh
sudo journalctl -u discord-bot -f
```

To view logs for the last hour:

```sh
sudo journalctl -u discord-bot --since "1 hour ago"
```

---

## 🔄 Managing the Service

Restart the bot:

```sh
sudo systemctl restart discord-bot
```

Stop the bot:

```sh
sudo systemctl stop discord-bot
```

Disable it from starting on boot:

```sh
sudo systemctl disable discord-bot
```

---

## 🛠️ Troubleshooting

### ❌ The service doesn’t start
- Check logs using:  
  ```sh
  sudo journalctl -u discord-bot -xe
  ```
- Ensure the binary is executable:
  ```sh
  sudo chmod +x /usr/local/bin/bot-app
  ```
- Make sure the environment file exists and has correct permissions:
  ```sh
  ls -l /etc/default/discord-bot
  ```

### 🔄 Changes not applying after updating `.env`
- Restart the service:
  ```sh
  sudo systemctl restart discord-bot
  ```
- If still not working, reload `systemd`:
  ```sh
  sudo systemctl daemon-reload
  sudo systemctl restart discord-bot
  ```

