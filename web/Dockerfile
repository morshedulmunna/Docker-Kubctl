# Stage 1: Build Stage
FROM node:lts-alpine AS builder

# Set working directory
WORKDIR /app

# Copy only package files for dependency installation
COPY package*.json ./

# Install dependencies with legacy peer deps and strict peer deps disabled
RUN npm config set legacy-peer-deps true && \
    npm config set strict-peer-deps false && \
    npm install --no-audit --no-fund

# Copy the rest of the application code
COPY . .

# Build the Next.js application
RUN npm run build

# Stage 2: Production Stage
FROM node:lts-alpine AS production

# Set working directory
WORKDIR /app

# Install production-only dependencies
COPY package.json ./
RUN npm install --legacy-peer-deps

# Copy the build files from the builder stage
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/next.config.mjs ./next.config.mjs

# Expose the application port
EXPOSE 3000

# Use a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Start the Next.js application
CMD ["npm", "start"]