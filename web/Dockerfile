# Stage 1: Build the React app
FROM node:16 as builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the app in production mode
RUN npm run build

# Stage 2: Serve the app with Nginx
FROM nginx:alpine

# Copy the build output to Nginx's web directory
COPY --from=builder /app/build /usr/share/nginx/html

# Expose port 80 for Nginx
EXPOSE 80

# Run Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
