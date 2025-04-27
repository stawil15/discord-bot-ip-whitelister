#!/bin/bash

set -e  # Exit immediately if a command fails

echo "🧹 Uninstalling the Discord Whitelist Bot..."

# 1. Stop and disable the systemd service if it exists
if systemctl list-units --full -all | grep -Fq "ip_whitelister_bot.service"; then
    echo "🛑 Stopping and disabling the service..."
    systemctl stop ip_whitelister_bot
    systemctl disable ip_whitelister_bot
    systemctl daemon-reload
else
    echo "  Service not found, skipping stop/disable."
fi

# 2. Remove the systemd service file
if [ -f /etc/systemd/system/ip_whitelister_bot.service ]; then
    echo "🗑️ Removing systemd service file..."
    rm -f /etc/systemd/system/ip_whitelister_bot.service
    systemctl daemon-reload
else
    echo "  Systemd service file already removed."
fi

# 3. Remove the bot binary
if [ -f /usr/local/bin/ip_whitelister_bot ]; then
    echo "🗑️ Removing bot executable..."
    rm -f /usr/local/bin/ip_whitelister_bot
fi

# 4. Remove environment and configuration files
if [ -d /etc/ip_whitelister_bot ]; then
    echo "🗑️ Removing configuration directory..."
    rm -rf /etc/ip_whitelister_bot
fi

# 5. Preserve logs: Do NOT delete /var/log/ip_whitelister_bot.log
echo "🛑 Skipping log file removal (preserving /var/log/ip_whitelister_bot.log)."

# 6. Remove bot data directory (only /var/lib/ip_whitelister_bot)
if [ -d /var/lib/ip_whitelister_bot ]; then
    echo "🗑️ Removing bot data directory..."
    rm -rf /var/lib/ip_whitelister_bot
fi

# 7. Remove sudoers permission
if [ -f /etc/sudoers.d/ip_whitelister_bot ]; then
    echo "🗑️ Removing sudoers file..."
    rm -f /etc/sudoers.d/ip_whitelister_bot
fi

# 8. Remove system user and group
if id "whitelistbot" &>/dev/null; then
    echo "👤 Deleting system user whitelistbot..."
    userdel whitelistbot || true
else
    echo "  User whitelistbot not found."
fi

echo "✅ Uninstallation completed successfully!"