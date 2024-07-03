FROM python:3.9

WORKDIR /app

COPY requirements.txt requirements.txt
COPY app.py app.py
COPY config.py config.py
COPY db.py db.py
COPY templates/ templates/

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5000

ENTRYPOINT ["flask", "run", "-d", "5000"]
