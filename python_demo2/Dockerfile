FROM python:3.5-slim
MAINTAINER Mario Inga <mario21ic@gmail.com>

RUN mkdir /app
COPY ./src/requirements.txt /app/
RUN pip install -r /app/requirements.txt

WORKDIR /app
