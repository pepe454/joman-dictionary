# Joman Sourashtra Language Dictionary

## Core Features
- Database of sourashtra words and sentences with metadata like part of speech and category.
- Translation from sourashtra to and from english. Ability to add additional language in the future.
- Admin user can upload new words and sentences, as well as edit existing data. Bulk upload in file or csv text format also supported.
- Website to view dictionary in categories.
- Search database for sourashtra words or an english translation.


## Technology Stack
- Database in postgresql
- Containerized in Docker, deployed on ubuntu
- Vue.js 3 
- Languages: Typescript and SQL
- Formatter: prettier 
- Testing: vitest
- e2e Testing: playwright


## Developer Setup
- Using Ubuntu 24.04, in windows install wsl 2
- Install docker, if not already available.  
- Install npm
- Run make install