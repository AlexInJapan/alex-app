# alex-app

A full-stack application with Go backend, Next.js frontend, and PostgreSQL database.

## Services

- **Backend**: Go API server (port 8080)
- **Frontend**: Next.js application (port 3000)
- **Database**: PostgreSQL (port 5432)
- **Playwright**: End-to-end testing (profile: test)
- **Playwright MCP**: Model Context Protocol server for testing automation (profile: mcp)

## Quick Start

### Start all services
```bash
docker-compose up -d
```

### Run Playwright tests
```bash
# Run tests in Docker
docker-compose --profile test up playwright --build --abort-on-container-exit

# Or use the npm script
cd frontend && npm run test:docker
```

### Start Playwright MCP server
```bash
# Start MCP server in Docker
docker-compose --profile mcp up playwright-mcp -d

# Or pull and run the official MCP Playwright image directly
docker pull mcp/playwright
docker run -p 3001:3001 mcp/playwright
```

## Playwright MCP Server

The Playwright MCP server provides automated browser testing capabilities through the Model Context Protocol using the official `mcp/playwright` image. It runs on port 3001 and can communicate with the frontend container.

### Available Tools

The official MCP Playwright image provides various browser automation tools including:
- Browser launching and management
- Page navigation and interaction
- Screenshot capture
- Form filling and validation
- And many more browser automation capabilities

### Usage

The MCP server can be used with any MCP client to automate browser testing against your frontend application. The official image includes a comprehensive set of tools for web automation and testing.

## Development

### Frontend Development
```bash
cd frontend
npm install
npm run dev
```

### Backend Development
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

### Database Migrations
```bash
cd backend
make migrate
```