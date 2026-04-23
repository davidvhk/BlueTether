# 🔒 BlueTether

**BlueTether** is a lightweight Linux security and automation daemon written in Go that acts as a 
proximity-based Dead Man's Switch. It monitors the Bluetooth signal of your 
trusted device (Phone/Watch) and automatically executes a **custom shell command** if the 
connection is lost.

While commonly used to lock your session, it can be configured to trigger any automation (e.g., pausing music, sending an alert, or stopping services).

## ✨ Features
- **Zero-Touch Automation:** Automatically trigger actions when you walk away.
- **State-Aware Monitoring:** Only triggers the command once per departure and resumes monitoring when you return.
- **Fully Customizable:** Execute **any** shell command, script, or binary via JSON configuration.
- **Modern Linux Support:** Uses `loginctl` by default for reliable session locking on Ubuntu and systemd-based distributions.
- **Safety Checks:** Automatically verifies Bluetooth status and dependencies at startup.

## 📋 Prerequisites
BlueTether depends on the following tools:
- `bluez` (for `hcitool`). On modern Ubuntu, you may need to install `bluez-deprecated` if `hcitool` is missing.
- `loginctl` (part of `systemd`, installed by default on most modern distros).
- A Bluetooth adapter that is **powered on**.

## 🚀 Installation & Build
1. **Install Go:** Ensure you have Go installed on your machine.
2. **Clone the repository:**
   ```bash
   git clone https://github.com/your-repo/BlueTether.git
   cd BlueTether
   ```
3. **Build the binary:**
   ```bash
   make build
   ```
   *Alternatively, you can build manually using `go build -o bluetether bluetether.go`.*

## 📦 Packaging (RPM & DEB)
This project uses [nFPM](https://nfpm.goreleaser.com/) to generate native Linux packages.

### 1. Install nFPM
If you have Go installed, you can install nFPM directly:
```bash
go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
```

### 2. Update your PATH
Ensure your Go binary directory is in your shell's PATH to run `nfpm` and the `Makefile` packaging targets:
```bash
# Add to your ~/.bashrc or ~/.zshrc
export PATH=$PATH:$(go env GOPATH)/bin
```
*Don't forget to run `source ~/.bashrc` after updating.*

### 3. Build Packages
Once nFPM is installed and in your PATH, you can generate packages using:
- `make deb`: Build Debian package.
- `make rpm`: Build RPM package.
- `make package`: Build both.

## 🛠 Makefile Targets
The project includes a `Makefile` for convenience:
- `make build`: Compiles the application.
- `make run`: Builds and executes the application.
- `make clean`: Removes the compiled binary.
- `make fmt`: Formats the Go source code.
- `make vet`: Runs Go static analysis.
- `make help`: Shows available commands.

## 🤝 Pairing Your Device
For BlueTether to work reliably, your device should be **paired** and **trusted** by your Linux machine. You can do this using the `bluetoothctl` interactive tool:

1. **Open the Bluetooth control tool:**
   ```bash
   bluetoothctl
   ```
2. **Scan for devices:**
   ```bash
   scan on
   ```
   *Look for your device's MAC address (e.g., `C0:1C:6A:76:B1:51`).*
3. **Pair the device:**
   ```bash
   pair [MAC_ADDRESS]
   ```
4. **Trust the device (Crucial):**
   ```bash
   trust [MAC_ADDRESS]
   ```
   *Trusting the device ensures that it can reconnect automatically without manual intervention.*
5. **Exit:**
   ```bash
   exit
   ```

## ⚙️ Configuration
Create a `config.json` file in the same directory as the `bluetether` binary. You can use the following template:

```json
{
    "TargetMAC": "C0:1C:6A:76:B1:51",
    "ThresholdRSSI": -90,
    "CheckInterval": "3s",
    "LockCommand": "loginctl lock-session"
}
```

### Configuration Options:
- `TargetMAC`: The Bluetooth MAC address of your trusted device (e.g., Phone or Watch).
- `CheckInterval`: How often to scan for the device (e.g., "3s", "5s", "500ms").
- `LockCommand`: The shell command to execute when the device is lost. 
  - *Recommended for Ubuntu:* `"loginctl lock-session"`
  - *Legacy:* `"gnome-screensaver-command -l"`

## 🛠 Usage
Run the application from your terminal:
```bash
./bluetether
```

### Custom Configuration
Specify a custom configuration file path with the `-config` flag:
```bash
./bluetether -config my-custom-config.json
```

### CLI Flags
- `-config`, `-c`: Path to the configuration file (default: `config.json`).
- `-help`, `-h`: Show usage information.

## 🔍 Tips
- **Find your MAC address:** Run `bluetoothctl devices` to see a list of paired devices and their MAC addresses.
- **Check Bluetooth Status:** Run `hcitool dev` to ensure your adapter is detected and active.
