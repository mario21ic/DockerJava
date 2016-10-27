FROM ruby:2.3.1-slim
MAINTAINER Mario Inga <mario21ic@gmail.com>

RUN mkdir /app
COPY ./Gemfile /app/
RUN cd /app && bundler install

WORKDIR /app
