# Joman Sourashtra Language Dictionary

## Screenshots Here

## Core Features

### Dictionary and Translations

- Database of sourashtra words and sentences with metadata including part of speech and category.
- Translations of sourashtra words to Tamil and English (other languages can be supported if there is data)
- Sentence translations of sourashtra to Tamil and Enlgish, with linkage to sourashtra words in context. This can include sourashtra idioms that native speakers will use.
- Audio recordings of sourashtra words.
- Bulk .csv uploads allowing admins to mass upload corpus texts.
- Verb conjugations in all tenses for the verbs
- Pluralized noun endings

### Frontend

- Search database for an english, sourashtra, or tamil word and see its translation to other languages.
- Clicking on a sourashtra word will display:
  - english translation
  - tamil translation
  - conjugations if the word is a verb
  - plural form if the word is a noun
  - audio player if the word has an associated audio file.
- Words by Category, like "nature", "family", "food", "education", ...
- Sentences by Category, same categories as above.

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

## Note

- Because sourashtra script is not known by many native speakers, the Harvard Kyoto transliteration into roman characters will allow everyone to read sourashtra words. However, the sourashtra lipi will be supported and stored for those who can read it.
- Project stack heavily inspired by [Full Stack FastAPI Template](https://github.com/fastapi/full-stack-fastapi-template)
