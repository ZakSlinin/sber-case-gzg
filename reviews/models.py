from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column
from sqlalchemy import UUID, Text, String
from pydantic import BaseModel
from uuid import uuid4
from datetime import datetime

class Base(DeclarativeBase):
	pass

class Review(Base):
	__tablename__ = "reviews"
	
	id: Mapped[uuid4] = mapped_column(UUID(as_uuid=True), primary_key=True)
	username: Mapped[str] = mapped_column(String(30))
	email: Mapped[str] = mapped_column(String(100), unique=True)
	mark: Mapped[int] = mapped_column()
	comment: Mapped[str] = mapped_column(Text())
	created_at: Mapped[datetime] = mapped_column()
	confirmed: Mapped[bool] = mapped_column()


class CreateReviewRequest(BaseModel):
	username: str
	email: str
	mark: int
	comment: str

class ConfirmRequest(BaseModel):
	email: str

class ConfirmResponse(BaseModel):
	message: str

class ReviewResponse(BaseModel):
	username: str
	email: str 
	mark: int
	comment: str
	created_at: datetime