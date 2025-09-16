# Developer Notes

Project stack heavily inspired by [Full Stack FastAPI Template](https://github.com/fastapi/full-stack-fastapi-template)

## Technology Stack

- Docker container deployment on Linux
- Backend:
  - PostgreSQL database
  - FastAPI to serve requests from the frontend (python)
  - SQLAlchemy ORM
  - Alembic for database migrations
  - Pydantic for data validation
  - uv package manager
  - devtools: pytest, ruff, uv
- Frontend:
  - Vue 3 (typescript)
  - Pinia for state management
  - Vue Router for routing
  - Vuetify for UI components
  - devtools: eslint, prettier, vitest, playwright

## Developer Setup

- Ubuntu 24.04 on Linux or [wsl2 on windows running Ubuntu 24.04](https://learn.microsoft.com/en-us/windows/wsl/install)
- [Install docker](https://docs.docker.com/get-started/get-docker/)
- [Install python 3.13](https://docs.astral.sh/uv/guides/install-python/)
- [Install node and npm](https://nodejs.org/en/download)
- [Install uv](https://docs.astral.sh/uv/getting-started/installation/)
- Copy .env.example to .env, set environment variables manually, and save them in your password manager.
