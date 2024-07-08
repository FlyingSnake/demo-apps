FROM python:3.9

WORKDIR /app

COPY requirements.txt requirements.txt
COPY * .
COPY templates/ templates/

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5000

ENTRYPOINT ["flask", "run", "--host", "0.0.0.0", "--port", "5000"]
