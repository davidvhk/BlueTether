package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	content := `{
		"TargetMAC": "AA:BB:CC:DD:EE:FF",
		"ThresholdRSSI": -80,
		"CheckInterval": "5s",
		"LockCommand": "echo locked"
	}`
	tmpFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Test loading the config
	config, err := loadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("loadConfig failed: %v", err)
	}

	if config.TargetMAC != "AA:BB:CC:DD:EE:FF" {
		t.Errorf("Expected TargetMAC AA:BB:CC:DD:EE:FF, got %s", config.TargetMAC)
	}
	if config.ThresholdRSSI != -80 {
		t.Errorf("Expected ThresholdRSSI -80, got %d", config.ThresholdRSSI)
	}
	if config.CheckInterval != "5s" {
		t.Errorf("Expected CheckInterval 5s, got %s", config.CheckInterval)
	}
	if config.LockCommand != "echo locked" {
		t.Errorf("Expected LockCommand 'echo locked', got %s", config.LockCommand)
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := loadConfig("non-existent.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "invalid*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(`{invalid: json}`)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	_, err = loadConfig(tmpFile.Name())
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
