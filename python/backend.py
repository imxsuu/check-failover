import json
import requests

import uvicorn
from fastapi import FastAPI, Request
from fastapi.logger import logger
import logging
import time

import os

app = FastAPI()

@app.post("/")
async def main(request: Request):

    msg_file = await request.json()
    print('[x] message received')

    print(msg_file['status'])

if __name__ == '__main__':
    uvicorn.run("backend:app", host='0.0.0.0', port=9010, debug=True)