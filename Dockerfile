# syntax = docker/dockerfile:1

FROM golang:1.22.0
# Node.js app lives here
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/go-htmx


# Set production environment
#ENV NODE_ENV="production"


# Throw-away build stage to reduce size of final image
#FROM base as build

# Install packages needed to build node modules
#RUN apt-get update -qq && \
#    apt-get install --no-install-recommends -y build-essential node-gyp pkg-config python-is-python3

# Install node modules
#COPY --link package-lock.json package.json ./
#RUN npm ci

# Copy application code
#COPY --link . .


# Final stage for app image
#FROM base

# Copy built application
#COPY --from=build /app /app

# Start the server by default, this can be overwritten at runtime
EXPOSE 3000
CMD [ "./bin/go-htmx" ]
