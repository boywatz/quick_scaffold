# Use the official Nginx image as the base image
FROM nginx:latest

# Copy a custom Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80 to make the server accessible
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
