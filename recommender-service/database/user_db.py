from sqlalchemy import create_engine, Column, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

Base = declarative_base()

def get_user_model(table_name):
    class User(Base):
        __tablename__ = table_name
        user_id = Column(String, primary_key=True)
        email = Column(String, nullable=False)
        nome = Column(String, nullable=False)
    return User

def get_engine(db_url):
    return create_engine(db_url)

def get_session(engine):
    Session = sessionmaker(bind=engine)
    return Session()
