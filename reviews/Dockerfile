FROM python:3.14-trixie

COPY . /reviews
WORKDIR /reviews
EXPOSE 8080

RUN pip install -r /reviews/requirements.txt

RUN adduser service-user

USER service-user