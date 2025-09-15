# Makefile for common operations
install:
	npm install

build:
	npm run build

dev:
	npm run dev

test:
	npm run test

setup:
	docker-compose up -d db

ci:
	npm run lint && npm run test
