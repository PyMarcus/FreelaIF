import logging
import time


import os
import requests
from dotenv import load_dotenv

load_dotenv()
CONFIG_URL = os.getenv('CONFIG_URL')

def get_config():
    while True:
        try:
            resp = requests.get(CONFIG_URL, timeout=5)
            if resp.status_code == 200:
                return resp.json()
            else:
                logging.warning(f'Config server returned {resp.status_code}')
        except Exception as e:
            logging.warning(f'Waiting for config-server: {e}')
        time.sleep(3)
