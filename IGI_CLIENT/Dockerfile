 
# Stage 1: Build Stage
FROM node:20 AS builder

# Set working directory
WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Global webpack installation
RUN npm install -g webpack webpack-cli

# Copy source code and build
COPY . .
RUN npm install && npm run build

# Stage 2: Serve Stage
FROM nginx:alpine

# Copy built files to Nginx directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose application port
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
