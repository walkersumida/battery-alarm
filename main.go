package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var threshold int
	var rootCmd = &cobra.Command{
		Use:   "batteryalarm",
		Short: "",
		Long:  ``,
		RunE: func(_ *cobra.Command, args []string) error {
			return cmdDo(threshold)
		},
	}

	rootCmd.Flags().IntVarP(&threshold, "threshold", "t", 25, "Threshold of battery percent")

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func cmdDo(threshold int) error {
	shell := exec.Command("ioreg", "-l")
	stdout, err := shell.Output()
	if err != nil {
		return err
	}

	output := string(stdout)
	rows := strings.Split(output, "\n")

	devices, err := findDevices(rows)
	if err != nil {
		return err
	}

	for _, device := range devices {
		if device.BatteryPercent != nil && *device.BatteryPercent <= threshold {
			mack.Notify(fmt.Sprintf("Battery of %s is low: %d%%\n", device.Name, *device.BatteryPercent), "Battery Alarm")
		}
	}

	return nil
}

type Device struct {
	Name           string
	BatteryPercent *int // nil if a device has no battery
}

func findDevices(rows []string) ([]Device, error) {
	var isSection bool
	secStart := "<class AppleDeviceManagementHIDEventService"
	secEndRegexp, err := regexp.Compile(`^[\|\s]*\}$`)
	if err != nil {
		return nil, err
	}
	projectRegexp, err := regexp.Compile(`^[\|\s]*\"Product\" = \"([\w\s\d\/_-]*)\"$`)
	if err != nil {
		return nil, err
	}
	batteryRegexp, err := regexp.Compile(`^[\|\s]*\"BatteryPercent\" = ([\d]*)$`)
	if err != nil {
		return nil, err
	}

	var devices []Device
	var deviceIdx int
	for _, row := range rows {
		sMatched, err := regexp.MatchString(secStart, row)
		if err != nil {
			return nil, err
		}
		if sMatched {
			devices = append(devices, Device{})
			isSection = true
		}
		if isSection {
			if res := projectRegexp.FindSubmatch([]byte(row)); res != nil {
				devices[deviceIdx].Name = string(res[1])
			} else if res := batteryRegexp.FindSubmatch([]byte(row)); res != nil {
				p, err := strconv.Atoi(string(res[1]))
				if err != nil {
					return nil, err
				}
				devices[deviceIdx].BatteryPercent = &p
			} else if secEndRegexp.MatchString(row) {
				isSection = false
				deviceIdx++
			}
		}
	}

	return devices, nil
}
