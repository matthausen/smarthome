import os
import paho.mqtt.client as mqtt
import redis
from dotenv import load_dotenv

load_dotenv()


redis_host = os.getenv('REDIS_HOST')
redis_port = os.getenv('REDIS_PORT') #port provided by `minikube service redis-master --url`
redis_password = os.getenv('REDIS_PASSWORD')

MQTT_ADDRESS = os.getenv('MQTT_ADDRESS')
MQTT_USER = os.getenv('MQTT_USER')
MQTT_PASSWORD = os.getenv('MQTT_PASSWORD')
MQTT_TOPIC = os.getenv('MQTT_TOPIC')


def on_connect(client, userdata, flags, rc):
    """ The callback for when the client receives a CONNACK response from the server."""
    print('Connected with result code ' + str(rc))
    client.subscribe(MQTT_TOPIC)


def on_message(client, userdata, msg):
    """The callback for when a PUBLISH message is received from the server."""
    print(msg.topic + ' ' + str(msg.payload))
    try:
   
        r = redis.StrictRedis(host=redis_host, port=redis_port, password=redis_password, decode_responses=True)

        if msg.topic == 'home/garden/temperature': r.set("Temp", str(msg.payload))
        if msg.topic == 'home/garden/humidity': r.set("Hum", str(msg.payload))
        if msg.topic == 'home/garden/datetime': r.set("Timestamp", str(msg.payload))
      
    except Exception as e:
        print(e)


def main():
    mqtt_client = mqtt.Client()
    mqtt_client.username_pw_set(MQTT_USER, MQTT_PASSWORD)
    mqtt_client.on_connect = on_connect
    mqtt_client.on_message = on_message

    mqtt_client.connect(MQTT_ADDRESS, 1883)
    mqtt_client.loop_forever()


if __name__ == '__main__':
    print('MQTT to InfluxDB bridge')
    main()