-- check out words.
select *
from word;

-- -- join words to categories 
select word.word_text,
       category.category
from word
join word_category wc
on word.word_id = wc.word_id
join category
on category.category_id = wc.category_id;


-- -- join words to each other
with sourashtra_words as (
    select word_text, part_of_speech 
    from word
    where language = 'Sourashtra'
),

other_words as (
    select word_text, word_text_alt
    from word 
    where language != 'Sourashtra'
)

select sourashtra_words.*,
       other_words.*



-- select * 
