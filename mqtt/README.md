# Starting a local MQTT broker

From the `mqtt` folder, run the following command:

```
docker run -it --rm -p 1883:1883 -p 9001:9001 -v $(pwd)/mosquitto:/mosquitto/ eclipse-mosquitto
```