-- name: CategoriesForWord :many
select w.word_id,
       w.word_text,
       c.category,
       c.category_id
from dictionary.word w
inner join dictionary.word_category wc on w.word_id = wc.word_id
inner join dictionary.category c on c.category_id = wc.category_id
where w.word_id = $1;


-- name: ListCategories :many
select category_id, 
       category
from dictionary.category
order by category;


-- name: SentencesForWord :many
select w.word_id,
       w.word_text,
       s.sentence_id,
       s.sentence_text
from dictionary.word w
inner join dictionary.word_in_context sw on w.word_id = sw.word_id
inner join dictionary.sentence s on sw.sentence_id = s.sentence_id
where w.word_id = $1;


-- name: WordsForSentence :many
select s.sentence_id,
       s.sentence_text,
       w.word_id,
       w.word_text
from dictionary.sentence s
inner join dictionary.word_in_context wc on s.sentence_id = wc.sentence_id
inner join dictionary.word w on wc.word_id = w.word_id
where s.sentence_id = $1;


-- name: TranslationsForCategory :many
with sourashtra as (
    select w.word_id as sourashtra_word_id,
           w.word_text as sourashtra_text,
           w.part_of_speech,
           w.word_text_alt as sourashtra_text_alt
    from dictionary.word w
    inner join dictionary.word_category wc on w.word_id = wc.word_id
    where wc.category_id = $1
    and w.language = 'Sourashtra'
),
other as (
    select word_id as target_word_id,
           word_text as target_word_text,
           word_text_alt as target_text_alt,
           language as target_language
    from dictionary.word
    where language != 'Sourashtra'
)
select sourashtra.*,
       other.*,
       translation.context
from sourashtra
inner join dictionary.translation on sourashtra.sourashtra_word_id = translation.sourashtra_word_id
inner join other on translation.other_word_id = other.target_word_id;


-- name: SourashtraTranslation :many
with sourashtra as (
    select word_id as sourashtra_word_id, 
           word_text as sourashtra_text, 
           part_of_speech,
           word_text_alt as sourashtra_text_alt
    from dictionary.word
    where dictionary.word.word_id = $1
    and language = 'Sourashtra'
),
other as (
    select word_id as target_word_id, 
           word_text as target_word_text, 
           word_text_alt as target_text_alt,
           language as target_language
    from dictionary.word 
    where language != 'Sourashtra'
)
select sourashtra.*,
       other.*,
       translation.context
from sourashtra 
inner join dictionary.translation
on   sourashtra.sourashtra_word_id = translation.sourashtra_word_id
inner join other
on   translation.target_word_id = other.target_word_id;


-- name: TranslationToSourashtra :many
with other as (
    select word_id as source_word_id, 
           word_text as source_word_text, 
           part_of_speech,
           language as source_language
    from dictionary.word 
    where dictionary.word.word_id = $1
),
sourashtra as (
    select word_id as sourashtra_word_id, 
           word_text as sourashtra_text
    from dictionary.word
    where language = 'Sourashtra'
)
select sourashtra.*,
       other.*,
       translation.context
from sourashtra 
inner join dictionary.translation
on   sourashtra.sourashtra_word_id = translation.sourashtra_word_id
inner join other
on   translation.target_word_id = other.target_word_id;


-- name: WordSearch :many
select word_id,
       word_text,
       part_of_speech,
       word_text_alt,
       similarity(word_text, $2) as score
from dictionary.word
where language = $1
and   (word_text % $2 or word_text_alt % $2) 
order by score desc;
