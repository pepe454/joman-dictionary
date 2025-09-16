"""Create-Read-Update-Delete (CRUD) operations here for models."""

from sqlalchemy.orm import Session
from app import models


# Word CRUD


def create_word(*, session: Session, word: models.Word) -> models.Word:
    session.add(word)
    session.commit()
    session.refresh(word)
    return word


def get_words_by_category(*, session: Session, category: models.Category) -> list[models.Word]:
    """Join the Word and Category tables using t_word_category translation table"""
    return []


def delete_word(*, session: Session, word: models.Word) -> models.Word:
    session.delete(word)
    session.commit()
    return word
