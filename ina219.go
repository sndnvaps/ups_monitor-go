package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/ina219"
	"periph.io/x/host/v3"
)

//shutdown the Raspberry Pi
func shutdownRpi() (string, error) {
	cmd := exec.Command("poweroff", []string{}...)
	o, err := cmd.CombinedOutput()
	var out string
	if err == nil {
		out = string(o)
	}
	return out, err
}

func mainImpl() error {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		return err
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		return fmt.Errorf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// create the sensor Opts
	sensorOpts := ina219.Opts{
		Address:       0x42,
		SenseResistor: 100 * physic.MilliOhm,     // 0.1Ohm
		MaxCurrent:    1000 * physic.MilliAmpere, //1A
	}

	// Create a new power sensor.
	sensor, err := ina219.New(bus, &sensorOpts)
	if err != nil {
		return fmt.Errorf("failed to open new sensor: %v", err)
	}

	// Read values from sensor every 3 second.
	everySecond := time.NewTicker(time.Second * 3).C
	var halt = make(chan os.Signal, 1)
	signal.Notify(halt, syscall.SIGTERM)
	signal.Notify(halt, syscall.SIGINT)

	fmt.Println("ctrl+c to exit")
	for {
		select {
		case <-everySecond:
			p, err := sensor.Sense()
			if err != nil {
				return fmt.Errorf("sensor reading error: %v", err)
			}
			busVoltage := p.Voltage.String()
			busVoltage = busVoltage[:len(busVoltage)-1] //remove "v" from the last string slice
			busVoltageNum, _ := strconv.ParseFloat(busVoltage, 64)
			BatVoltagePersent := (float64)((busVoltageNum - 6.0) / 2.4 * 100.0)
			if BatVoltagePersent > 100.0 {
				BatVoltagePersent = 100.0
			}
			//shutdown happens when battery is below 25%
			if BatVoltagePersent < 25.0 {
				shutdownRpi()
			}
			if BatVoltagePersent < 0.0 {
				BatVoltagePersent = 0.0
			}

			fmt.Println(p)
			tmpstr := fmt.Sprintf("BatVoltagePersent %.2f%%", BatVoltagePersent)
			fmt.Println(tmpstr)
		case <-halt:
			return nil
		}
	}

}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "ina219: %s.\n", err)
		return
	}
}
