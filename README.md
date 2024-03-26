The [Pimoroni Blinkt!](https://shop.pimoroni.com/products/blinkt) is an 8x RGB LED add-on for the Raspberry Pi.

When mounting on Orange Pi, remember pinout is different, you need to turn the circuit 180 degrees.


### Install Go

```
sudo apt-get install golang-go gcc
? export GOPATH=$HOME/go
```

### Install Armbian-Config
```
sudo apt-get install armbian-config
```
### Enable GPIO on Armbian-Config
```
sudo armbian-config
```
Then go to System, Hardware, check "w1-gpio" and Save

### Install WiringOP ?
```
git clone https://github.com/zhaolei/WiringOP.git -b h3
cd WiringOP
chmod +x ./build
sudo ./build
```
