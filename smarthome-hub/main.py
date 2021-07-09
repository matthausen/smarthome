from flask import Flask
import os
from dotenv import load_dotenv
import redis

load_dotenv()

redis_host = os.getenv('REDIS_HOST')
redis_port = os.getenv('REDIS_PORT')
redis_password = os.getenv('REDIS_PASSWORD')

app = Flask(__name__)

@app.route('/home/temperature')
def temperature():
   try:
        r = redis.StrictRedis(host=redis_host, port=redis_port, password=redis_password, decode_responses=True)

        temp = r.get("Temp")
       
        return temp   
   
   except Exception as e:
        print(e)
        return e

@app.route('/home/humidity')
def humidity():
   try:
        r = redis.StrictRedis(host=redis_host, port=redis_port, password=redis_password, decode_responses=True)

        hum = r.get("Hum")
        return hum        
   
   except Exception as e:
        print(e)
        return e

@app.route('/home/timestamp')
def timestamp():
   try:
        r = redis.StrictRedis(host=redis_host, port=redis_port, password=redis_password, decode_responses=True)

        timestamp = r.get("Timestamp")
        return timestamp        
   
   except Exception as e:
        print(e)
        return e

if __name__ == '__main__':
   app.run()