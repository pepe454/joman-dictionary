# Backend Development

## Tasks

- [ ] Setup development Environment
  - [ ] Setup vscode to use python, ruff, and uv
  - [ ] Install python
  - [ ] Setup uv package manager
  - [ ] pyproject.toml
  - [ ] Install postgres server, spin it up using docker
  - [ ] Install fastapi, sqlalchemy, and pydantic

- [ ] Database Design
  - [ ] Spin up postgres and pgadmin using docker, traefik, etc etc
  - [ ] Design some of the database schema for core features:
    - [ ] words
    - [ ] sentences
    - [ ] languages
    - [ ] users
    - [ ] audio files
    - [ ] definitions
    - [ ] categories
    - [ ] Once you have the er diagram defined, generate the schema
    - [ ] Setup alembic, sqlalchemy, and pydantic to work with the schema

- [ ] API Design
  - [ ] Setup some simple routes in fastapi
    - [ ] Get all words for a given category
    - [ ] Get all sentences for a given word
    - [ ] Get the audio file for a word
