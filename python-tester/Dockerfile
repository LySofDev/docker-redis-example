FROM python:3.7.1-alpine3.8
RUN mkdir /python
COPY requirements.txt /python/
WORKDIR /python
RUN pip install -r requirements.txt
COPY . /python/
