package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Config represents the application configuration
type Config struct {
	TargetMAC     string `json:"TargetMAC"`
	ThresholdRSSI int    `json:"ThresholdRSSI"`
	CheckInterval string `json:"CheckInterval"`
	LockCommand   string `json:"LockCommand"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func checkDependencies(config *Config) error {
	if _, err := exec.LookPath("hcitool"); err != nil {
		return errors.New("hcitool not found. Please install bluez-deprecated or similar package.")
	}

	parts := strings.Fields(config.LockCommand)
	if len(parts) == 0 {
		return errors.New("LockCommand is empty in configuration.")
	}
	if _, err := exec.LookPath(parts[0]); err != nil {
		return fmt.Errorf("LockCommand binary '%s' not found in PATH.", parts[0])
	}

	out, err := exec.Command("hcitool", "dev").Output()
	if err != nil {
		return fmt.Errorf("failed to check Bluetooth status: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 2 {
		return errors.New("no Bluetooth adapter found or Bluetooth is powered off. Please enable Bluetooth.")
	}

	return nil
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "Path to configuration file")
	flag.StringVar(&configPath, "c", "config.json", "Path to configuration file (shorthand)")
	flag.Parse()

	config, err := loadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error loading config from %s: %v\n", configPath, err)
		os.Exit(1)
	}

	// Run startup checks
	if err := checkDependencies(config); err != nil {
		fmt.Fprintf(os.Stderr, "[-] Startup validation failed: %v\n", err)
		os.Exit(1)
	}

	checkInterval, err := time.ParseDuration(config.CheckInterval)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error parsing CheckInterval: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[*] Guardian Active. Watching for: %s\n", config.TargetMAC)
	fmt.Printf("[*] Config path: %s\n", configPath)

	isAway := false
	for {
		// Use hcitool to check the signal strength (RSSI)
		out, err := exec.Command("hcitool", "name", config.TargetMAC).Output()

		deviceMissing := err != nil || len(strings.TrimSpace(string(out))) == 0

		if deviceMissing {
			if !isAway {
				fmt.Println("[-] Device out of range! Triggering Lock...")
				executeLock(config.LockCommand)
				isAway = true
			}
		} else {
			if isAway {
				fmt.Printf("[+] Device %s returned. Resuming active monitoring.\n", config.TargetMAC)
				isAway = false
			}
		}

		time.Sleep(checkInterval)
	}
}

func executeLock(command string) {
	exec.Command("sh", "-c", command).Run()
}
