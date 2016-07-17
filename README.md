# dht11-to-mode

Simple program for collecting dht11 sensor data and publishing them as events to MODE.

## Usage

First, you have to connect your Raspberry Pi and dht11 sensor apppropriately.

Then, download dht-to-mode binary to `/home/pi/dht-to-mode`.

```
$ cd /home/pi/
$ wget https://dht11-to-mode.jp-east-2.os.cloud.nifty.com/dht11-to-mode
$ chmod a+x dht11-to-mode
```

Write your configuration to `/home/pi/.env`:

```
MODE_ENDPOINT=<Your MODE endpoint>
MODE_DEVICE_ID=<Your device's ID>
MODE_DEVICE_API_KEY=<Your device's API Key>
INTERVAL=<Interval in seconds>
```

And, start the program:

```
$ sudo ./dht11-to-mode
2016/07/17 06:24:51 [dht11] Triggering event: {HomeId:0 Timestamp:0001-01-01 00:00:00 +0000 UTC EventType:dht11-start EventData:map[value:1] OriginDeviceId:0 OriginDeviceClass: OriginDeviceIp:}
2016/07/17 06:24:52 [dht11] Triggering event: {HomeId:0 Timestamp:0001-01-01 00:00:00 +0000 UTC EventType:dht11-temperature EventData:map[value:30] OriginDeviceId:0 OriginDeviceClass: OriginDeviceIp:}
2016/07/17 06:24:52 [dht11] Triggering event: {HomeId:0 Timestamp:0001-01-01 00:00:00 +0000 UTC EventType:dht11-humidity EventData:map[value:45] OriginDeviceId:0 OriginDeviceClass: OriginDeviceIp:}
2016/07/17 06:24:52 [dht11] Triggering event: {HomeId:0 Timestamp:0001-01-01 00:00:00 +0000 UTC EventType:dht11-retried EventData:map[value:0] OriginDeviceId:0 OriginDeviceClass: OriginDeviceIp:}
```

## Systemd integration

Write `/etc/systemd/system/dht11-to-mode.service`:

```
[Unit]
Description=DHT11 to MODE

[Service]
Type=simple
EnvironmentFile=/etc/sysconfig/dht11-to-mode
ExecStart=/home/pi/dht11-to-mode
Restart=always

[Install]
WantedBy=multi-user.target
```

And write: `/etc/sysconfig/dht11-to-mode`:

```
MODE_ENDPOINT=<Your MODE endpoint>
MODE_DEVICE_ID=<Your device's ID>
MODE_DEVICE_API_KEY=<Your device's API Key>
INTERVAL=<Interval in seconds>
```

Then you can start service as follows:

```
systemctl start dht11-to-mode
```
