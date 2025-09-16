from typing import List, Optional

from sqlalchemy import Column, ForeignKeyConstraint, Integer, PrimaryKeyConstraint, Table, Text
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship


class Base(DeclarativeBase):
    pass


class Category(Base):
    __tablename__ = "category"
    __table_args__ = (
        PrimaryKeyConstraint("category_id", name="category_pkey"),
        {"schema": "dictionary"},
    )

    category_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    category: Mapped[str] = mapped_column(Text)

    word: Mapped[List["Word"]] = relationship(
        "Word", secondary="dictionary.word_category", back_populates="category"
    )


class Language(Base):
    __tablename__ = "language"
    __table_args__ = (
        PrimaryKeyConstraint("language_id", name="pk_language"),
        {"schema": "dictionary"},
    )

    language_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    language: Mapped[str] = mapped_column(Text)

    sentence: Mapped[List["Sentence"]] = relationship("Sentence", back_populates="language")
    word: Mapped[List["Word"]] = relationship("Word", back_populates="language")


class Sentence(Base):
    __tablename__ = "sentence"
    __table_args__ = (
        ForeignKeyConstraint(
            ["language_id"], ["dictionary.language.language_id"], name="fk_sentence_language"
        ),
        PrimaryKeyConstraint("sentence_id", name="sentence_pkey"),
        {"schema": "dictionary"},
    )

    sentence_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    text: Mapped[str] = mapped_column(Text)
    language_id: Mapped[Optional[int]] = mapped_column(Integer)
    audio_file_path: Mapped[Optional[str]] = mapped_column(Text)

    language: Mapped[Optional["Language"]] = relationship("Language", back_populates="sentence")
    word: Mapped[List["Word"]] = relationship(
        "Word", secondary="dictionary.word_in_context", back_populates="sentence"
    )


class Word(Base):
    __tablename__ = "word"
    __table_args__ = (
        ForeignKeyConstraint(
            ["language_id"], ["dictionary.language.language_id"], name="fk_word_language"
        ),
        PrimaryKeyConstraint("word_id", name="pk_word"),
        {
            "comment": "This is the written text of the word, and the building block of "
            "the entire system. ",
            "schema": "dictionary",
        },
    )

    word_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    text: Mapped[str] = mapped_column(Text)
    language_id: Mapped[int] = mapped_column(Integer)
    audio_file_path: Mapped[Optional[str]] = mapped_column(Text)

    category: Mapped[List["Category"]] = relationship(
        "Category", secondary="dictionary.word_category", back_populates="word"
    )
    sentence: Mapped[List["Sentence"]] = relationship(
        "Sentence", secondary="dictionary.word_in_context", back_populates="word"
    )
    language: Mapped["Language"] = relationship("Language", back_populates="word")
    this_word: Mapped[List["Word"]] = relationship(
        "Word",
        secondary="dictionary.translation",
        primaryjoin=lambda: Word.word_id == t_translation.c.other_word_id,
        secondaryjoin=lambda: Word.word_id == t_translation.c.this_word_id,
        back_populates="other_word",
    )
    other_word: Mapped[List["Word"]] = relationship(
        "Word",
        secondary="dictionary.translation",
        primaryjoin=lambda: Word.word_id == t_translation.c.this_word_id,
        secondaryjoin=lambda: Word.word_id == t_translation.c.other_word_id,
        back_populates="this_word",
    )
    definition: Mapped[List["Definition"]] = relationship("Definition", back_populates="word")


class Definition(Base):
    __tablename__ = "definition"
    __table_args__ = (
        ForeignKeyConstraint(["word_id"], ["dictionary.word.word_id"], name="fk_word"),
        PrimaryKeyConstraint("definition_id", "word_id", name="definition_pkey"),
        {"schema": "dictionary"},
    )

    definition_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    word_id: Mapped[int] = mapped_column(Integer, primary_key=True)
    text: Mapped[str] = mapped_column(Text)

    word: Mapped["Word"] = relationship("Word", back_populates="definition")


t_translation = Table(
    "translation",
    Base.metadata,
    Column("this_word_id", Integer, primary_key=True, nullable=False),
    Column("other_word_id", Integer, primary_key=True, nullable=False),
    ForeignKeyConstraint(["other_word_id"], ["dictionary.word.word_id"], name="fk_other_word"),
    ForeignKeyConstraint(["this_word_id"], ["dictionary.word.word_id"], name="fk_this_word"),
    PrimaryKeyConstraint("this_word_id", "other_word_id", name="translation_pkey"),
    schema="dictionary",
)


t_word_category = Table(
    "word_category",
    Base.metadata,
    Column("word_id", Integer, primary_key=True, nullable=False),
    Column("category_id", Integer, primary_key=True, nullable=False),
    ForeignKeyConstraint(["category_id"], ["dictionary.category.category_id"], name="fk_category"),
    ForeignKeyConstraint(["word_id"], ["dictionary.word.word_id"], name="fk_word"),
    PrimaryKeyConstraint("word_id", "category_id", name="word_category_pkey"),
    schema="dictionary",
)


t_word_in_context = Table(
    "word_in_context",
    Base.metadata,
    Column("word_id", Integer, primary_key=True, nullable=False),
    Column("sentence_id", Integer, primary_key=True, nullable=False),
    ForeignKeyConstraint(["sentence_id"], ["dictionary.sentence.sentence_id"], name="fk_sentence"),
    ForeignKeyConstraint(["word_id"], ["dictionary.word.word_id"], name="fk_word"),
    PrimaryKeyConstraint("word_id", "sentence_id", name="word_in_context_pkey"),
    schema="dictionary",
)
