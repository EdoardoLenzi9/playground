# Proof of Concept 1 - Interfaces

Make in communication a go microservice `m1` with:
1. another go microservice `m2`
2. a message broker `mb`
3. a database `db`

## [Ogen](https://ogen.dev/docs/intro)

```sh
go mod init m2
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

## [Mosquitto](https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang)

```sh
mkdir ./mosquitto
mkdir ./mosquitto/config
mkdir ./mosquitto/data
mkdir ./mosquitto/log
```

```sh
cat ./mosquitto/config/mosquitto.conf
...
```

```sh
mosquitto_passwd -b /mosquitto/passwd root root # add user root (pw root)
```

```sh
mosquitto_pub -h 127.0.0.1 -p 1883 -t "test/topic" -m "Hello, MQTT" -u root -P root
```

```sh
mosquitto_sub -h 127.0.0.1 -p 1883 -t "test/topic" -u root -P root
```

## [MQTTX](https://mqttx.app/downloads?os=docker)

| key | value|
|:--  |:--   |
| name | root |
| host | ws://127.0.0.1 |
| port | 9001 |
| path | /mqtt |
| username | root |
| password | root |
| SSL/TSL | false |

## [Paho]

```sh
go get github.com/eclipse/paho.mqtt.golang
```