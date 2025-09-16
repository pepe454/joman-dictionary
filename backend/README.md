# Backend Development

## Tasks

- [ ] Setup development Environment
  - [ ] Setup vscode to use python, ruff, and uv
  - [x] Install python
  - [ ] Setup uv package manager
  - [x] pyproject.toml
  - [ ] Install postgres server, spin it up using docker
  - [x] Install fastapi, sqlalchemy, and pydantic using uv

- [ ] Database Design
  - [x] Get the database docker container to run
  - [x] Get PGAdmin to run
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
