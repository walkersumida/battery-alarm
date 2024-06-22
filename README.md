## Battery Alarm

This is a simple script that will check the battery level of your Mac and notify if it is below a certain threshold. 

## Usage

#### 1. Clone the repository

#### 2. Run the Script

Run the script with the desired threshold as an argument. For example, to set the threshold to 20% run the following command:

```bash
go run main.go -t 20
```

or

```bash
go build main.go -o batteryalarm
./batteryalarm -t 20
```

#### 3. Register a Scheduled Job

TODO: Add a description

https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/Chapters/ScheduledJobs.html
