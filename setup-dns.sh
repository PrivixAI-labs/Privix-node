#!/bin/bash

# Usage: setup-dns.sh --domain <your-domain> --ip <your-ip> --p2p-addr <your-peer-id> [--port <custom-port>]
# Example: sudo ./setup-dns.sh --domain node1.yourdomain.com --ip 37.60.239.84 --p2p-addr 16Uiu2HAmUkuKgbDRaZzjE8SzNp3ptZcJCD8pFV2XwkZ9Tc2jpxaZ --port 20001

# Default values
PORT=10001

# Parse arguments
for arg in "$@"
do
    case $arg in
        --domain=*)
        DOMAIN="${arg#*=}"
        shift
        ;;
        --ip=*)
        IP="${arg#*=}"
        shift
        ;;
        --p2p-addr=*)
        P2P_ADDR="${arg#*=}"
        shift
        ;;
        --port=*)
        PORT="${arg#*=}"
        shift
        ;;
        *)
        echo "Unknown argument: $arg"
        exit 1
        ;;
    esac
done

# Check for required arguments
if [ -z "$DOMAIN" ] || [ -z "$IP" ] || [ -z "$P2P_ADDR" ]; then
    echo "Usage: $0 --domain=<your-domain> --ip=<your-ip> --p2p-addr=<your-peer-id> [--port=<custom-port>]"
    exit 1
fi

# Step 1: Install necessary packages
echo "Installing necessary packages (nginx, certbot)..."
sudo apt update
sudo apt install -y nginx certbot python3-certbot-nginx

# Step 2: Set up Nginx configuration for the domain with reverse proxy
NGINX_CONF="/etc/nginx/sites-available/$DOMAIN"

if [ ! -f "$NGINX_CONF" ]; then
    echo "Creating Nginx configuration for $DOMAIN..."
    sudo tee "$NGINX_CONF" > /dev/null <<EOL
server {
    listen 443 ssl;
    server_name $DOMAIN;

    ssl_certificate /etc/letsencrypt/live/$DOMAIN/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/$DOMAIN/privkey.pem;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;
    ssl_ciphers "ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384";

    location / {
        proxy_pass http://$IP:$PORT;  # Forward to internal port (default or custom)
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}

server {
    listen 80;
    server_name $DOMAIN;
    return 301 https://\$host\$request_uri;  # Redirect all HTTP to HTTPS
}
EOL

    # Enable the configuration and restart Nginx
    sudo ln -s /etc/nginx/sites-available/$DOMAIN /etc/nginx/sites-enabled/
else
    echo "Nginx configuration for $DOMAIN already exists."
fi

# Step 3: Obtain SSL certificate using Let's Encrypt
echo "Obtaining SSL certificate for $DOMAIN..."
sudo certbot --nginx -d $DOMAIN --non-interactive --agree-tos --email admin@$DOMAIN

# Step 4: Test Nginx and reload
echo "Testing and reloading Nginx..."
sudo nginx -t
sudo systemctl reload nginx

# Step 5: Add cron job for automatic certificate renewal
echo "Adding cron job for SSL renewal..."
(crontab -l 2>/dev/null; echo "0 3 * * * certbot renew --quiet && systemctl reload nginx") | crontab -

# Step 6: Output the bootnode configuration
echo "Here is your bootnode configuration for Msc-node:"
echo "\"/dns4/$DOMAIN/p2p/$P2P_ADDR\""

echo "Setup complete! Your Msc-node is now accessible via DNS with SSL and hidden ports."
