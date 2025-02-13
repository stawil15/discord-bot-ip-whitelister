
# Discord Bot IP Whitelister

A Discord bot that allows users to manage their IP whitelisting and bans directly from Discord. With the `/whitelist` command, players can securely add their IP addresses to the UFW firewall of a game server, preventing unauthorized access and mitigating DDoS attacks.
Admins have access to the `/ban` command, which allows them to ban a user by their Discord ID, preventing them from whitelisting their IP. This bot is designed for servers using IP-based whitelisting, such as FiveM, Minecraft, Rust, and more.


## Features
- **Slash command `/whitelist (ip)`**: Players can whitelist their IPs securely.
- **Slash command `/ban (user_id)`**: Admins can prevent specific users from adding their IP.
- **Prevents unauthorized access**: Only approved users can connect.
- **DDoS Mitigation**: Reduces attack surface by limiting access to verified players.
- **Multi-Game Support**: Works with FiveM, Minecraft, Rust, and any server using UFW.
- **SQLite Database**: Keeps a log of whitelisted users and their IPs for traceability.
- **Systemd Service Support**: Runs as a background service for automatic startup and reliability.
- **Secure & Configurable**: Uses environment variables for easy customization.



## Installation

### 1. One-Line Install Command  
Run this command to **automatically download and install** the bot:

```sh
curl -sSL https://raw.githubusercontent.com/geekloper/discord-bot-ip-whitelister/main/install.sh  | sudo bash
```

This script will:
- Checks if **UFW** and **SQLite3** are installed (prompts if missing).
- Creates a dedicated system user (`whitelistbot`).
- Downloads the latest bot binary from GitHub.
- Installs the bot to `/usr/local/bin/`.
- Sets up configuration files in `/etc/ip_whitelister_bot/`.
- Creates a SQLite database in `/var/lib/ip_whitelister_bot/`.
- Configures and enables a **systemd service** for auto-restart.

---

### 2. Configure the Bot
Edit the `.env` file to add your Discord bot credentials:

```sh
sudo vi /etc/whitelist_bot/.env
```

Example `.env` file:
```ini
BOT_TOKEN=your-bot-token
BOT_GUILD_ID=your-guild-id
ADMIN_IDS=1234567890,0987654321
DELETE_COMMANDS=true
SERVICES=80/tcp,443/tcp
DEBUG=false
UFW_PATH=/usr/sbin/ufw
DB_PATH=/var/lib/ip_whitelister_bot/whitelist.db
```

---

### 3. Start the Bot
Once configured, start and enable the bot:

```sh
sudo systemctl restart ip_whitelister_bot
sudo systemctl enable ip_whitelister_bot
```

Check if the bot is running:

```sh
sudo systemctl status ip_whitelister_bot
```

View logs:

```sh
sudo journalctl -u ip_whitelister_bot -f
```

## Usage

## 🛠️ Usage

### ✅ Whitelisting an IP  
Users can whitelist their IP via the Discord command:

```
/whitelist 192.168.1.1
```

This will automatically add the IP to UFW and store it in the database.

### ❌ Banning a User  
Admins can **ban a user by Discord ID** to prevent them from using `/whitelist`:

```
/ban 1234567890
```

Banning removes their ability to whitelist IPs & deny their ips in UFW.



## Uninstallation
To remove the bot completely:

```sh
sudo systemctl stop ip_whitelister_bot
sudo systemctl disable ip_whitelister_bot
sudo rm -rf /usr/local/bin/ip_whitelister_bot /etc/ip_whitelister_bot /var/lib/ip_whitelister_bot /etc/systemd/system/ip_whitelister_bot.service
sudo systemctl daemon-reload
```

---

## Contributing
Contributions are welcome! If you find a bug or want to improve the bot, feel free to submit an issue or pull request.

---

## License
This project is open-source under the MIT License.
