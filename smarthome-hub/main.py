from flask import Flask
import paho.mqtt.client as mqtt

app = Flask(__name__)

MQTT_ADDRESS = '192.168.0.129'
MQTT_USER = 'username'
MQTT_PASSWORD = 'passowrd'
MQTT_TOPIC = 'home/+/+'


def on_connect(client, userdata, flags, rc):
   print('Connected with result code ' + str(rc))
   client.subscribe(MQTT_TOPIC)


def on_message(client, userdata, msg):
   print(msg.topic + ' ' + str(msg.payload))
   return str(msg.payload)

def handleMQTT():    
   mqtt_client = mqtt.Client()
   mqtt_client.username_pw_set(MQTT_USER, MQTT_PASSWORD)
   mqtt_client.on_connect = on_connect
   mqtt_client.on_message = on_message

   mqtt_client.connect(MQTT_ADDRESS, 1883)
   mqtt_client.loop_forever() 

@app.route('/home-weather-station')
def hello_world():
   payload = on_message()
   return payload

if __name__ == '__main__':
   print('MQTT to InfluxDB bridge')
   handleMQTT()
   app.run()