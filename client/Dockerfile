# Build stage
FROM node:20-alpine AS builder
WORKDIR /app

# Reduce timeouts on slow/unstable networks (optional: pass NPM_REGISTRY for a mirror)
ARG NPM_REGISTRY=
RUN if [ -n "$NPM_REGISTRY" ]; then npm config set registry "$NPM_REGISTRY"; fi && \
    npm config set fetch-timeout 300000 && \
    npm config set fetch-retries 5

COPY package.json package-lock.json* ./
RUN npm install

COPY . .
# API URL is inlined at build time for client-side requests
ARG NEXT_PUBLIC_URL=http://localhost:8082
ENV NEXT_PUBLIC_URL=${NEXT_PUBLIC_URL}
ARG NEXT_PUBLIC_EXPIRE_TOKEN=1
ENV NEXT_PUBLIC_EXPIRE_TOKEN=${NEXT_PUBLIC_EXPIRE_TOKEN}

RUN npm run build

# Run stage
FROM node:20-alpine
WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

COPY --from=builder /app/package.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 3000
CMD ["npm", "start"]
