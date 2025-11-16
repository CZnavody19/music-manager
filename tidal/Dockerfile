FROM python:3.12-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
RUN pip install -U 'watchdog[watchmedo]'
RUN apt-get update && \
    apt-get install -y --no-install-recommends ffmpeg && \
    rm -rf /var/lib/apt/lists/*
COPY . .

CMD ["watchmedo", "auto-restart", "--directory=./", "--pattern='*.py'", "--recursive", "python", "main.py"]
