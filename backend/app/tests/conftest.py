from collections.abc import Generator
import pytest

from app.core.db import engine
from sqlalchemy.orm import Session


@pytest.fixture(scope="session", autouse=True)
def db() -> Generator[Session, None, None]:
    with Session(engine) as session:
        yield session
