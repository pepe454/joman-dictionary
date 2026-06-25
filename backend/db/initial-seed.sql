BEGIN;

/* This script simply adds the word for mother in all 3 languages with metadata. */

INSERT INTO dictionary.language (language, language_alt) VALUES
    ('Sourashtra', 'ꢱꣃꢬꢵꢰ꣄ꢜ꣄ꢬ'),
    ('English', NULL),
    ('தமிழ்', 'Tamil');


INSERT INTO dictionary.part_of_speech (part_of_speech) VALUES
    ('noun'),
    ('pronoun'),
    ('verb'),
    ('adjective'),
    ('adverb'),
    ('preposition'),
    ('conjunction'),
    ('interjection'),
    ('numeral'),
    ('article'),
    ('interrogative');


INSERT INTO dictionary.word (word_text, language, part_of_speech, word_text_alt) VALUES
    ('ambO', 'Sourashtra', 'noun', 'ꢂꢪ꣄ꢨꣁ'),
    ('mother', 'English', 'noun', NULL),
    ('அம்மா', 'தமிழ்', 'noun', 'amma');


INSERT INTO dictionary.translation (sourashtra_word_id, other_word_id, context) VALUES
    (1, 2, NULL),
    (1, 3, NULL);


INSERT INTO dictionary.category (category) VALUES
    ('family'), -- always put family first!
    -- these are all parts of speech
    ('adjectives'),
    ('adverbs'),
    ('pronouns'),
    ('verbs'),
    ('questions'),
    ('numbers'),
    ('conjunctions'),
    ('prepositions'),
    ('animals'),
    ('body'),
    ('clothing'),
    ('cognition'),
    ('colors'),
    ('feelings'),
    ('food'),
    ('house'),
    ('nature'),
    ('people'),
    ('places'),
    ('taste'),
    ('time');


INSERT INTO dictionary.word_category (word_id, category_id) VALUES
    (1, 1), -- sourashtra ambO -> family
    (2, 1), -- english    mother -> family
    (3, 1); -- tamil      ammo -> family

END;