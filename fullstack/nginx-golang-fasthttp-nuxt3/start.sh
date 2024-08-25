#!/bin/sh

# Start the Go backend
/usr/local/bin/backend &

# Check and start Nuxt app
if [ -f /app/frontend/first-frontend/server/index.mjs ]; then
    echo "Starting Nuxt app..."
    cd /app/frontend/first-frontend && NODE_ENV=production HOST=0.0.0.0 PORT=3000 node server/index.mjs &
else
    echo "Error: Nuxt app file not found at /app/frontend/first-frontend/server/index.mjs"
    echo "Contents of /app/frontend/first-frontend/server:"
    ls -la /app/frontend/first-frontend/server
fi

# Check and start Nuxt app
if [ -f /app/frontend/second-frontend/server/index.mjs ]; then
    echo "Starting Nuxt app..."
    cd /app/frontend/second-frontend && NODE_ENV=production HOST=0.0.0.0 PORT=3001 node server/index.mjs &
else
    echo "Error: Nuxt app file not found at /app/frontend/second-frontend/server/index.mjs"
    echo "Contents of /app/frontend/second-frontend/server:"
    ls -la /app/frontend/second-frontend/server
fi

# Wait for all services to start
sleep 5

# Start Nginx
nginx -g 'daemon off;'