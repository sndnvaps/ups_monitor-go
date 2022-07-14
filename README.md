# ups monitor
 ups status monitor for Raspberry Pi Openwrt
 
# 图片

![UPS-HAT](https://www.waveshare.net/photo/accBoard/UPS-HAT/UPS-HAT-1.jpg)

# connect (openwrt raspbeery pi需要先安装i2c内核驱动的支持)
    Raspberry Pi I2C interface

# hack （如果获取ups连接到树莓派后，用的是哪个地址）
```bash
opkg install  i2c-tools

i2cdetect -y 1
```
![ups-hat-i2c-address](https://www.waveshare.net/w/upload/f/f1/UPS_HAT_I2C.png)

默认设置的地址为 0x42(需要修改 [/ina219.go#L44](/ina219.go#L44))
```go
	// create the sensor Opts
	sensorOpts := ina219.Opts{
		Address:       0x42, // ina219 ic2 address
		SenseResistor: 100 * physic.MilliOhm,     // 0.1Ohm
		MaxCurrent:    1000 * physic.MilliAmpere, //1A
	}
```

# depend


# how to build
## Install golang in the host (the host system must by  Linux or Mac)

## just run make to build the ups_monitor
```bash
make
```

# Install ups_monitor to Raspbery Pi
192.168.xx.xx is the ip of you RPI
```bash
scp ups_monitor root@192.168.xx.xx:/tmp
```

# Run ups_monitor on Raspberry Pi
connect RPI with ssh
```bash
cd /tmp
mv /tmp/ups_monitor /usr/bin/ups_monitor
ups_monitor
```
# wiki
## ups-hat介绍文档
https://www.waveshare.net/wiki/UPS_HAT

## Raspberry Pi OpenWRT打开 I2C支持 
https://www.icode9.com/content-4-1367375.html
    
# python support from waveshare
https://www.waveshare.net/w/upload/d/d9/UPS_HAT.7z


# 声明
 此代码，目前只测试于 <strong>微雪的UPS_HAT</strong>

 编译的时候，需要<strong>go</strong>支持

