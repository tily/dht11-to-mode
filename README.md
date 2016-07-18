# dht11-to-mode

Simple program for Rapsberry Pi to collect dht11 sensor data and publishing them as events to [MODE](http://www.tinkermode.com).

## Usage

First, you have to connect your Raspberry Pi and dht11 sensor apppropriately.

* [Temperature monitoring with Raspberry Pi and DHT11/22 temperature and humidity sensor - Documentation - Documentation](http://docs.gadgetkeeper.com/pages/viewpage.action?pageId=7700673)
* [DHT11 Humidity & Temperature Sensor Module | UUGear](http://www.uugear.com/portfolio/dht11-humidity-temperature-sensor-module/)

Then, download [dht-to-mode](https://dht11-to-mode.jp-east-2.os.cloud.nifty.com/dht11-to-mode) binary to `/home/pi/`.

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

## Run as a service with `systemd`

Add `/etc/systemd/system/dht11-to-mode.service`:

```
[Unit]
Description=DHT11 to MODE
Wants=network-online.target
After=network.target network-online.target

[Service]
Type=simple
EnvironmentFile=/etc/default/dht11-to-mode
ExecStart=/home/pi/dht11-to-mode
Restart=always

[Install]
WantedBy=multi-user.target
```

And add: `/etc/default/dht11-to-mode`:

```
MODE_ENDPOINT=<Your MODE endpoint>
MODE_DEVICE_ID=<Your device's ID>
MODE_DEVICE_API_KEY=<Your device's API Key>
INTERVAL=<Interval in seconds>
```

Then you can start service as follows:

```
$ sudo systemctl start dht11-to-mode
```

Enable service to start automatically after a crash or a server reboot.

```
$ sudo systemctl enable dht11-to-mode
```
