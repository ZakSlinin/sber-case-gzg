from fastapi import FastAPI, HTTPException
from fastapi_sqlalchemy import DBSessionMiddleware 
from fastapi_sqlalchemy import db 
from models import CreateReviewRequest, ReviewResponse, Review, ConfirmResponse, ConfirmRequest
import aiohttp
from datetime import datetime
import json
import os
import uuid

DB_USER = os.environ.get("DB_USER")
DB_HOST = os.environ.get("DB_HOST")
DB_PORT = os.environ.get("DB_PORT")
DB_PASSWORD = os.environ.get("DB_PASSWORD")
DB_NAME = os.environ.get("DB_NAME")

app = FastAPI()  

app.add_middleware(DBSessionMiddleware, db_url=f"postgresql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}")



@app.post("/api/review/create", response_model=ReviewResponse)
async def create_review(review: CreateReviewRequest):
	if not "@" in review.email:
		raise HTTPException(status_code=400, detail="Email is invalid")
	if not(review.mark >= 1 and review.mark <= 5):
		raise HTTPException(status_code=400, detail="Mark must be 1 <= mark <= 5")
	try:	
		cr = Review(
			id=uuid.uuid4(),
			username=review.username,
			email=review.email,
			mark=review.mark,
			comment=review.comment,
			created_at=datetime.now(),
			confirmed=False
			)
	except Exception:
		raise HTTPException(status_code=400, detail="Review with this email is already exist")
	db.session.add(cr)
	db.session.commit()
	async with aiohttp.ClientSession() as session:
		async with session.post('http://email-verification:8080/api/email-verification/verify', data=json.dumps({"email": review.email})) as response: pass

	return {
		"id": cr.id,
		"username": cr.username,
		"email": cr.email,
		"mark": cr.mark,
		"comment": cr.comment,
		"created_at": cr.created_at
		}

@app.get("/api/review/get", response_model=list[ReviewResponse])
async def get_reviews():
	return db.session.query(Review).filter(Review.confirmed==True).all()


@app.post("/api/review/confirm", response_model=ConfirmResponse)
async def confirm(request: ConfirmRequest):
	review = db.session.query(Review).filter(Review.email == request.email).one()
	review.confirmed = True
	db.session.add(review)
	db.session.commit()
	return {"message": "OK"}

@app.get("/api/email-verification/verify")
async def verify(token: str):
	async with aiohttp.ClientSession() as session:
		async with session.get(f"http://email-verification:8080/api/email-verification/verify?{token}") as response:
			return respons.text