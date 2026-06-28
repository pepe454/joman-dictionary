# Tasks TODO

# Backend 

### CRUD API
#### C - Create
- [ ] Add language
- [ ] Add word
- [ ] Add translation
- [ ] Add sentence
- [ ] From sentence - add links to word via word in context. 
- [ ] Add category 
- [ ] Add category for word - add link to word via word_category

#### R - Read
- [ ] Get word
- [ ] Get translation for sourashtra word.
- [ ] Get translation for english word.
- [ ] Get example sentences for word in context (could potentially contain derivatives)
- [ ] Get categories for word
- [ ] Get words for category
- [ ] Get conjugations for verb.

#### U - Update
- [ ] Update spelling for word
- [ ] Update category name
- [ ] Update sentence
- [ ] Update translation (context)

#### D - Delete
- [ ] Delete word - cascade to words in context, translations, categories, word derivatives
- [ ] Delete sentence - cascade to words in context. 
- [ ] Delete category - cascade to word categories.
- [ ] Delete translation


### Data Management

#### Seed database script with data from words and sentences:
- [ ] Populate languages table
- [ ] Populate words table
    - [ ] Sourashtra words
    - [ ] English words
- [ ] Populate translations table mapping sourashtra words to english words
- [ ] Populate sentences table and identify words in context
#### Export database to .csv files
- [ ] Export words to .csv files based on category (choose one category per word.)
#### Add lipi translations for words and sentences (big task)
#### Add Tamil words in tamil script for additional translations (yet another big task.)

### API Routes
- [ ] Define go backend, Gin framework (?) 
- [ ] Initial route - hello route 
- [ ] GET /word/query route -
    - [ ] Search sourashtra word (fuzzy)
    - [ ] Search english word (fuzzy)
- [ ] GET /word/allMetadata  -
    - [ ] Show words in context

### Schema
- [ ] Add sentence to sentence translation
- [ ] Add lipi font column to words (leave it null for now.)
- [ ] Add derivatives , derivations of a word so that word in context can still reference the original word. 
    - [ ] Plurals 
    - [ ] Nouns with suffix (for example, in)
        - um,indicates location for example park-um sEte
        - n,plural noun modifier (pillO -> pillan)
        - ni,negative modifier on a verb 
        - k,indicates a noun is a receiver of an action
- [ ] Define conjugation schema for verbs
    - Future tenses (I will go to the store tomorrow.), (I always go to the store on Mondays)
    - Past tense (Yesterday you played outside.)
    - Perfect tense (I have cooked chicken.)
    - Impersonal tense (It is raining.)  
    - Imperative / Command tense (Clean your room.)
    - Affirmative tense (I will clean my room.)
    - Negative tense (Do not talk with your mouth full.)
    - Conditional tense (If you bring an umbrella, you won't get wet)
    - Gerund (She is watching tv right now)
    - Capability tense - positive (She can play the piano)
    - Capability tense - negative (He cannot speak Chinese)
    - Completion tense er, indicating completion of verb (After we eat dinner, we can eat dessert.) 