-- name: CategoriesForWord :many
select w.word_id,
       w.word_text,
       c.category
from dictionary.word w
inner join dictionary.word_category wc on w.word_id = wc.word_id
inner join dictionary.category c on c.category_id = wc.category_id
where w.word_id = $1;


-- name: WordsForCategory :many
select c.category,
       w.word_id,
       w.word_text
from dictionary.category c
inner join dictionary.word_category wc on c.category_id = wc.category_id
inner join dictionary.word w on wc.word_id = w.word_id
where c.category_id = $1;


-- name: SentencesForWord :many
select w.word_id,
       w.word_text,
       s.sentence_id,
       s.text as sentence_text
from dictionary.word w
inner join dictionary.word_in_context sw on w.word_id = sw.word_id
inner join dictionary.sentence s on sw.sentence_id = s.sentence_id
where w.word_id = $1;


-- name: WordsForSentence :many
select s.sentence_id,
       s.text as sentence_text,
       w.word_id,
       w.word_text
from dictionary.sentence s
inner join dictionary.word_in_context sw on s.sentence_id = sw.sentence_id
inner join dictionary.word w on sw.word_id = w.word_id
where s.sentence_id = $1;


-- name: SourashtraTranslation :many
with sourashtra as (
    select word_id as sourashtra_word_id, 
           word_text as sourashtra_text, 
           part_of_speech as sourashtra_part_of_speech,
           word_text_alt as sourashtra_text_alt
    from word
    where word_id = $1
    and language = 'Sourashtra'
),
other as (
    select word_id as target_word_id, 
           word_text as target_word_text, 
           word_text_alt as target_text_alt,
           language as target_language
    from word 
    where language != 'Sourashtra'
)
select sourashtra.*,
       other.*,
       translation.context
from sourashtra 
inner join translation
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
