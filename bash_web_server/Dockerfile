FROM python:3.10-alpine
COPY requirements.txt /opt/app/requirements.txt
WORKDIR /opt/app
RUN pip install -r requirements.txt
COPY . /opt/app
ARG CI_JOB_TOKEN
RUN pip install setuptools==70.0.0 packaging==24.1
RUN apk add git