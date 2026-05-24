# Joman Sourashtra Language Dictionary

## Core Features

### Words and Sentences

- Word files will be stored raw in .csv files in the words/ directory for easy reading and writing. 
  For multiple-meaning words, the .csv file can have multiple rows for the same word.
- Sentences will similarly be stored in .csv files in the sentences/ directory.


### Dictionary and Translations

- Database of sourashtra words and sentences with metadata including part of speech and category.
- Translations of sourashtra words to English (eventually support Tamil language with support from Tamil speakers.)
- Sentence translations of sourashtra to English, with linkage to sourashtra words in context. 
- Audio recordings of sourashtra words.
- Bulk .csv uploads allowing admins to mass upload corpus texts.
- Verb conjugations in all tenses for the verbs
- Pluralized noun endings

### Frontend
- Search database for an English word and find the corresponding sourashtra word(s) and contextual sentences. 
- Clicking on a sourashtra word will display:
  - english translation
  - conjugations if the word is a verb
  - plural form if the word is a noun
  - audio player if the word has an associated audio file.
- Words by Category, like "nature", "family", "food", "education", ...
- Sentences by Category, same categories as above.

## Note
Because sourashtra script is not known by many native speakers, the Harvard Kyoto transliteration into roman characters 
will allow everyone to read sourashtra words. However, the sourashtra lipi can eventually be supported and stored for those who can read it.


## Tech Stack (Developer Docs)
- OS: _Linux_
- Arch: _Ubuntu 24.04_
- Backend: _Golang_
- Database: _Postgres_
- Frontend: _Vue.js_
- Containers: _Docker_
- Hosting: _AWS_
- CI/CD: _GitHub Actions_
- Reverse Proxy: _Traefik_


## Sourashtra Resources
TODO: Add a page for this compilation.
- [Thinnal.org - Community Forum](https://thinnal.org/)
- [Palkar.org - Community and Culture](http://palkar.org/)
- [Sourashtra Association](https://sourashtraassociation.org/)
- [Sourashtra Association Literary Works](https://sourashtraassociation.org/poltam-homepage/poltam-literary-works/)
- [Sourashtra Dictionary](https://dictionary.thinnal.org/)
- [Sourashtra Dictionary 2](http://sourashtradictionary.com/)
- [Noto Sans Saurashtra Font](https://www.google.com/get/noto/)
- [OS Subramanian Blog](https://subramanian-obula.blogspot.com/)
- [Sourashtra Article](https://www.endangeredalphabets.net/alphabets/sourashtra/)