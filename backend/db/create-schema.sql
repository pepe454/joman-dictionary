BEGIN;


CREATE SCHEMA IF NOT EXISTS dictionary;


CREATE TABLE IF NOT EXISTS dictionary.language
(
    language_id serial NOT NULL,
    language VARCHAR(16) NOT NULL,
    language_alt VARCHAR(16) NULL,
    CONSTRAINT pk_language PRIMARY KEY (language_id)
);


CREATE TABLE IF NOT EXISTS dictionary.part_of_speech 
(
    part_of_speech VARCHAR(16) NOT NULL,
    CONSTRAINT pk_part_of_speech PRIMARY KEY (part_of_speech)
);


CREATE TABLE IF NOT EXISTS dictionary.word
(
    word_id serial NOT NULL,
    word_text text NOT NULL,
    language_id integer NOT NULL,
    part_of_speech VARCHAR(16) NOT NULL,
    word_text_alt text NULL,
    audio_file_path text,
    CONSTRAINT pk_word PRIMARY KEY (word_id),
    CONSTRAINT word_text_distinct UNIQUE (word_text)
);


CREATE TABLE IF NOT EXISTS dictionary.translation
(
    sourashtra_word_id integer NOT NULL,
    other_word_id integer NOT NULL,
    context text,
    CONSTRAINT pk_translation_id PRIMARY KEY (sourashtra_word_id, other_word_id)
);


CREATE TABLE IF NOT EXISTS dictionary.sentence
(
    sentence_id serial NOT NULL,
    text text NOT NULL,
    language_id integer NOT NULL,
    audio_file_path text,
    PRIMARY KEY (sentence_id)
);


CREATE TABLE IF NOT EXISTS dictionary.word_in_context
(
    word_id integer NOT NULL,
    sentence_id integer NOT NULL,
    PRIMARY KEY (word_id, sentence_id)
);


CREATE TABLE IF NOT EXISTS dictionary.sentence_translation
(
    sourashtra_sentence_id integer NOT NULL,
    other_sentence_id integer NOT NULL,
    PRIMARY KEY (sourashtra_sentence_id, other_sentence_id)
);


CREATE TABLE IF NOT EXISTS dictionary.category
(
    category_id serial NOT NULL,
    category text NOT NULL,
    PRIMARY KEY (category_id)
);


CREATE TABLE IF NOT EXISTS dictionary.word_category
(
    word_id integer NOT NULL,
    category_id integer NOT NULL,
    PRIMARY KEY (word_id, category_id)
);


/* The following are FK Constraints between tables. */

ALTER TABLE IF EXISTS dictionary.word
    ADD CONSTRAINT fk_word_language FOREIGN KEY (language_id)
    REFERENCES dictionary.language (language_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE RESTRICT;
CREATE INDEX IF NOT EXISTS fki_fk_word_language
    ON dictionary.word(language_id);


ALTER TABLE IF EXISTS dictionary.word
    ADD CONSTRAINT fk_word_part_of_speech FOREIGN KEY (part_of_speech)
    REFERENCES dictionary.part_of_speech (part_of_speech) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE RESTRICT;
CREATE INDEX IF NOT EXISTS fki_fk_word_part_of_speech
    ON dictionary.word(part_of_speech);


ALTER TABLE IF EXISTS dictionary.translation
    ADD CONSTRAINT fk_sourashtra_word_id FOREIGN KEY (sourashtra_word_id)
    REFERENCES dictionary.word (word_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_sourashtra_word_index
    ON dictionary.translation(sourashtra_word_id);


ALTER TABLE IF EXISTS dictionary.translation
    ADD CONSTRAINT fk_other_word_id FOREIGN KEY (other_word_id)
    REFERENCES dictionary.word (word_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_other_word_id
    ON dictionary.translation(other_word_id);


ALTER TABLE IF EXISTS dictionary.sentence
    ADD CONSTRAINT fk_sentence_language FOREIGN KEY (language_id)
    REFERENCES dictionary.language (language_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;
CREATE INDEX IF NOT EXISTS fki_fk_sentence_language
    ON dictionary.sentence(language_id);


ALTER TABLE IF EXISTS dictionary.word_in_context
    ADD CONSTRAINT fk_word_word_in_ctxt FOREIGN KEY (word_id)
    REFERENCES dictionary.word (word_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_word_word_in_ctxt
    ON dictionary.word_in_context(word_id);


ALTER TABLE IF EXISTS dictionary.word_in_context
    ADD CONSTRAINT fk_sentence_word_in_ctxt FOREIGN KEY (sentence_id)
    REFERENCES dictionary.sentence (sentence_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_sentence_word_in_ctxt
    ON dictionary.word_in_context(sentence_id);


ALTER TABLE IF EXISTS dictionary.sentence_translation
    ADD CONSTRAINT fk_sourashtra_sentence_id  FOREIGN KEY (sourashtra_sentence_id)
    REFERENCES dictionary.sentence (sentence_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_sourashtra_sentence_id
    ON dictionary.sentence_translation(sourashtra_sentence_id);


ALTER TABLE IF EXISTS dictionary.sentence_translation
    ADD CONSTRAINT fk_other_sentence_id  FOREIGN KEY (other_sentence_id)
    REFERENCES dictionary.sentence (sentence_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_other_sentence_id
    ON dictionary.sentence_translation(other_sentence_id);


ALTER TABLE IF EXISTS dictionary.word_category
    ADD CONSTRAINT fk_word_id_to_category FOREIGN KEY (word_id)
    REFERENCES dictionary.word (word_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_word_id_to_category
    ON dictionary.word_category(word_id);


ALTER TABLE IF EXISTS dictionary.word_category
    ADD CONSTRAINT fk_category_id_to_word FOREIGN KEY (category_id)
    REFERENCES dictionary.category (category_id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE;
CREATE INDEX IF NOT EXISTS fki_fk_category_id_to_word
    ON dictionary.word_category(category_id);

END;