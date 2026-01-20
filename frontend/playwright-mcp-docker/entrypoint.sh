#!/bin/sh
set -e

# Use MCP_PORT from environment or default to 8931
INTERNAL_PORT=3302

echo "Internal MCP port (if using SSE): $INTERNAL_PORT"

# Execute @playwright/mcp using npx, passing arguments
exec npx @playwright/mcp@ --port "$INTERNAL_PORT"