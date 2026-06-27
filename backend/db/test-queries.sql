-- get all words with all categories
select w.word_id,
       w.word_text,
       c.category
from dictionary.word w
inner join dictionary.word_category wc on w.word_id = wc.word_id
inner join dictionary.category c on c.category_id = wc.category_id


-- -- name: Word Translations :many
with sourashtra_words as (
    select word_id, word_text, part_of_speech 
    from word
    where language = 'Sourashtra'
),
other_words as (
    select word_id, word_text, word_text_alt
    from word 
    where language != 'Sourashtra'
)

select sourashtra_words.*,
       other_words.*
from sourashtra_words 
join translation wt
on sourashtra_words.word_id = wt.sourashtra_word_id
join other_words
on other_words.word_id = wt.other_word_id
