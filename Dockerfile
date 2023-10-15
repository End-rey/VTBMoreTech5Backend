FROM python:3.9

WORKDIR /app

COPY /internal/utils/flask.py /app/flask.py

COPY /internal/utils/Weights/lean /app/Weights/lean

RUN pip install fastapi uvicorn torch numpy pydantic

CMD ["uvicorn", "flask:app", "--reload"]