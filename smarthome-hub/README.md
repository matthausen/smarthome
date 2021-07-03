## Smarthome Hub

The smarthome hub is a MQTT broker that listens to a topic 'home/+/+' and reports temperature, humidity, date.

This information is served via a flask REST API to a frontend client

### How to run
- `python3 mqtt.py`