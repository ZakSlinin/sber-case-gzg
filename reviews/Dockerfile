FROM python:3.14-trixie

COPY . /
WORKDIR / 
EXPOSE 8080

RUN pip install -r /requirements.txt

RUN adduser service-user

USER service-user