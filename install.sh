#!/bin/bash

set -e  # Stops the script if a command fails

echo "📦 Installing the Discord Whitelist Bot..."


# Checking if UFW is installed
if ! command -v ufw &> /dev/null; then
    echo "❌ UFW is not installed. Please install it using:"
    echo "   sudo apt update && sudo apt install ufw -y"
    exit 1
fi

# 2. Checking if SQLite3 is installed
if ! command -v sqlite3 &> /dev/null; then
    echo "❌ SQLite3 is not installed. Please install it using:"
    echo "   sudo apt update && sudo apt install sqlite3 -y"
    exit 1
fi

# Creating the system user (without shell)
if ! id "whitelistbot" &>/dev/null; then
    useradd -r -m -d /var/lib/ip_whitelister_bot -s /usr/sbin/nologin whitelistbot
fi


# Creating the necessary directories

mkdir -p /etc/ip_whitelister_bot
mkdir -p /var/lib/ip_whitelister_bot
mkdir -p /var/log



# Download the latest release from GitHub
GITHUB_REPO="geekloper/discord-bot-ip-whitelister"
LATEST_RELEASE_URL=$(curl -s https://api.github.com/repos/$GITHUB_REPO/releases/latest | grep "browser_download_url" | cut -d '"' -f 4)

if [[ -z "$LATEST_RELEASE_URL" ]]; then
    echo "❌ Failed to fetch the latest release. Check your GitHub repo and tags."
    exit 1
fi

echo "⬇️ Downloading latest release from $LATEST_RELEASE_URL..."
curl -L $LATEST_RELEASE_URL -o /usr/local/bin/ip_whitelister_bot

# Set permission
chmod +x /usr/local/bin/ip_whitelister_bot

# Creating the default .env file if it doesn't exist
if [ ! -f /etc/ip_whitelister_bot/.env ]; then
    cat <<EOL > /etc/ip_whitelister_bot/.env
BOT_TOKEN=<your_bot_token>
BOT_GUILD_ID=<server_guild_id>
ADMIN_IDS=<admin_id>,<admin_id>,...
DELETE_COMMANDS=true
SERVICES=80/tcp,443/tcp
DEBUG=0
UFW_PATH=/usr/sbin/ufw
DB_PATH=/var/lib/ip_whitelister_bot/whitelist.db
EOL
    echo "✅ Configuration file created in /etc/ip_whitelister_bot/.env"
fi

# Creating the systemd service
cat <<EOL > /etc/systemd/system/ip_whitelister_bot.service
[Unit]
Description=Discord Whitelist Bot
After=network.target

[Service]
Type=simple
User=whitelistbot
Group=whitelistbot
WorkingDirectory=/var/lib/ip_whitelister_bot
ExecStart=/usr/local/bin/ip_whitelister_bot
Restart=always
EnvironmentFile=/etc/ip_whitelister_bot/.env
StandardOutput=append:/var/log/ip_whitelister_bot.log
StandardError=append:/var/log/ip_whitelister_bot.log

[Install]
WantedBy=multi-user.target
EOL

# Setting the correct permissions
chown -R whitelistbot:whitelistbot /var/lib/ip_whitelister_bot
chmod 600 /etc/ip_whitelister_bot/.env
chown whitelistbot:whitelistbot /etc/ip_whitelister_bot/.env
chmod 755 /var/lib/ip_whitelister_bot
touch /var/lib/ip_whitelister_bot/whitelist.db
chown whitelistbot:whitelistbot /var/lib/ip_whitelister_bot/whitelist.db
chmod 660 /var/lib/ip_whitelister_bot/whitelist.db


# Give bot sudo permission for UFW
cat <<EOL > /etc/sudoers.d/ip_whitelister_bot
whitelistbot ALL=(ALL) NOPASSWD: /usr/sbin/ufw
EOL
chmod 440 /etc/sudoers.d/ip_whitelister_bot


# Reloading systemd and enabling the service
systemctl daemon-reload
systemctl enable ip_whitelister_bot
systemctl restart ip_whitelister_bot

echo "✅ Installation completed!"
echo "👉 Remember to edit /etc/ip_whitelister_bot/.env and restart the bot using: systemctl restart ip_whitelister_bot"
