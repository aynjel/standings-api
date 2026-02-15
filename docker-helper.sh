#!/bin/bash
# Docker Helper Script for Tabulation API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get command from arguments
COMMAND=${1:-help}

# Functions
print_help() {
    echo -e "${BLUE}=== Tabulation API Docker Helper ===${NC}"
    echo ""
    echo "Usage: ./docker-helper.sh <command>"
    echo ""
    echo "Commands:"
    echo -e "  ${GREEN}build${NC}           Build Docker image"
    echo -e "  ${GREEN}up${NC}              Start containers (development)"
    echo -e "  ${GREEN}down${NC}            Stop containers"
    echo -e "  ${GREEN}logs${NC}            View application logs"
    echo -e "  ${GREEN}logs-postgres${NC}   View PostgreSQL logs"
    echo -e "  ${GREEN}shell${NC}           Open shell in running container"
    echo -e "  ${GREEN}test${NC}            Run quick health check"
    echo -e "  ${GREEN}restart${NC}         Restart containers"
    echo -e "  ${GREEN}clean${NC}           Remove containers and volumes"
    echo -e "  ${GREEN}postgres${NC}        Start with PostgreSQL database"
    echo -e "  ${GREEN}sqlite${NC}          Start with SQLite (default)"
    echo -e "  ${GREEN}help${NC}            Show this help message"
    echo ""
    echo "Examples:"
    echo "  ./docker-helper.sh build"
    echo "  ./docker-helper.sh up"
    echo "  ./docker-helper.sh logs"
    echo "  ./docker-helper.sh postgres"
}

build_image() {
    echo -e "${YELLOW}Building Docker image...${NC}"
    docker-compose build
    echo -e "${GREEN}Build complete!${NC}"
}

start_dev() {
    echo -e "${YELLOW}Starting development containers (SQLite)...${NC}"
    docker-compose up -d
    sleep 2
    echo -e "${GREEN}Containers started!${NC}"
    echo -e "${BLUE}Application URL: http://localhost:4300${NC}"
    echo -e "View logs with: ./docker-helper.sh logs"
}

start_postgres() {
    echo -e "${YELLOW}Starting containers with PostgreSQL...${NC}"
    docker-compose --profile postgres up -d
    sleep 3
    echo -e "${GREEN}Containers started!${NC}"
    echo -e "${BLUE}Application URL: http://localhost:4300${NC}"
    echo -e "${BLUE}PostgreSQL URL: localhost:5432${NC}"
    echo -e "View logs with: ./docker-helper.sh logs"
}

stop_containers() {
    echo -e "${YELLOW}Stopping containers...${NC}"
    docker-compose down
    echo -e "${GREEN}Containers stopped!${NC}"
}

view_logs() {
    echo -e "${BLUE}=== Application Logs ===${NC}"
    docker-compose logs -f app
}

view_postgres_logs() {
    echo -e "${BLUE}=== PostgreSQL Logs ===${NC}"
    docker-compose logs -f postgres
}

open_shell() {
    echo -e "${BLUE}Opening shell in app container...${NC}"
    docker-compose exec app sh
}

health_check() {
    echo -e "${YELLOW}Running health check...${NC}"
    sleep 1
    if curl -s http://localhost:4300/ > /dev/null; then
        echo -e "${GREEN}✓ Application is running!${NC}"
        echo -e "Response: $(curl -s http://localhost:4300/)"
    else
        echo -e "${RED}✗ Application is not responding${NC}"
        exit 1
    fi
}

restart_containers() {
    echo -e "${YELLOW}Restarting containers...${NC}"
    docker-compose restart
    sleep 2
    echo -e "${GREEN}Containers restarted!${NC}"
}

clean_all() {
    echo -e "${RED}WARNING: This will remove containers and volumes!${NC}"
    read -p "Continue? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}Cleaning up...${NC}"
        docker-compose down -v
        docker-compose --profile postgres down -v
        echo -e "${GREEN}Cleanup complete!${NC}"
    else
        echo "Cancelled."
    fi
}

# Execute command
case $COMMAND in
    build)
        build_image
        ;;
    up|dev|sqlite)
        start_dev
        ;;
    postgres|pg)
        start_postgres
        ;;
    down|stop)
        stop_containers
        ;;
    logs)
        view_logs
        ;;
    logs-postgres|logs-pg)
        view_postgres_logs
        ;;
    shell|bash|sh)
        open_shell
        ;;
    test|check|health)
        health_check
        ;;
    restart)
        restart_containers
        ;;
    clean)
        clean_all
        ;;
    help|--help|-h)
        print_help
        ;;
    *)
        echo -e "${RED}Unknown command: $COMMAND${NC}"
        echo "Run './docker-helper.sh help' for usage information"
        exit 1
        ;;
esac
