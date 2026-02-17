# 🧠 Raspberry Pi and Arduino in Smart Home: DIY Creation

## 📑 Contents
1. [What are Raspberry Pi and Arduino?](#what-are-raspberry-pi-and-arduino)
2. [Differences Between Raspberry Pi and Arduino](#differences-between-raspberry-pi-and-arduino)
3. [Application in Smart Home](#application-in-smart-home)
4. [How to Choose Between Raspberry Pi and Arduino](#how-to-choose-between-raspberry-pi-and-arduino)
5. [Practical Project Examples](#practical-project-examples)
6. [Connection and Setup](#connection-and-setup)
7. [Programming and Firmware](#programming-and-firmware)
8. [Integration with IoT Platforms](#integration-with-iot-platforms)
9. [Security of Homemade Devices](#security-of-homemade-devices)

---

## 1. 🤖 What are Raspberry Pi and Arduino?

### 🧩 Raspberry Pi
**Raspberry Pi** is a **single-board computer** (SBC - Single Board Computer) that operates under an operating system (usually Linux, such as Raspbian/Debian).

#### 🧠 Key Characteristics:
*   **Processor:** ARM (various generations, from 1.2GHz to 2.4GHz)
*   **RAM:** from 512MB to 8GB
*   **OS:** Linux (Raspberry Pi OS, Ubuntu, Raspbian, etc.)
*   **GPIO ports:** 40-pin GPIO header for connecting peripherals
*   **Connectivity:** Ethernet, Wi-Fi, Bluetooth
*   **Video:** HDMI output, 4K support (depending on model)

**Practical Example:**
Raspberry Pi can perform complex tasks such as:
- Managing multiple sensors and actuators
- Running a web server for remote control
- Processing images from a camera for facial recognition
- Serving as the central hub of a smart home

### 🛠️ Arduino
**Arduino** is a **microcontroller board** designed for simple prototyping and creating electronic projects.

#### 🧠 Key Characteristics:
*   **Microcontroller:** ATmega328P (in Arduino Uno), other variants
*   **Memory:** Flash 32KB, SRAM 2KB (in Arduino Uno)
*   **OS:** None (program loaded directly into microcontroller)
*   **GPIO ports:** 14 digital and 6 analog pins (in Arduino Uno)
*   **Connectivity:** USB, Serial, I2C, SPI
*   **Power:** 5V or 7-12V via external adapter

**Practical Example:**
Arduino is ideal for tasks such as:
- Reading data from temperature, humidity, and motion sensors
- Controlling LEDs, motors, relays
- Simple automation without complex computations
- Collecting and transmitting data to another controller (e.g., Raspberry Pi)

---

## 2. ⚖️ Differences Between Raspberry Pi and Arduino

### 📊 Comparison Table:

| Characteristic | Raspberry Pi | Arduino |
|----------------|--------------|---------|
| **Device Type** | Single-board computer | Microcontroller |
| **OS** | Linux | None (program in ROM) |
| **Computational Power** | High (multi-core processor) | Low (single-core microcontroller) |
| **Memory** | GB RAM | KB RAM |
| **Purpose** | Complex tasks, network services | Simple tasks, GPIO work |
| **Programming** | Python, C++, Java, etc. | C/C++ (Arduino IDE) |
| **Price** | Higher ($35-100+) | Lower ($10-30) |
| **Power** | 5V (USB) | 5V or 7-12V |
| **Internet Connectivity** | Built-in (Wi-Fi, Ethernet) | Through additional modules |

### 🔄 Combined Usage
In real smart home projects, a **combination** of Raspberry Pi and Arduino is often used:
*   **Raspberry Pi** — as the "brain" of the system, handling complex tasks, connecting to the internet, managing databases
*   **Arduino** — as "sensors and actuators," collecting data from sensors and controlling simple devices
*   They are connected via Serial, I2C, or SPI interfaces

**Practical Example:**
In a smart home system:
- Arduino Uno collects data from temperature, humidity, and motion sensors
- Arduino sends data to Raspberry Pi via Serial
- Raspberry Pi analyzes data, saves to database, sends notifications to phone, and controls smart outlets via Wi-Fi

---

## 3. 🏡 Application in Smart Home

### 🧠 Raspberry Pi in Smart Home:

#### 🏠 Central Hub:
*   **Home Assistant** — full-featured smart home management platform
*   **OpenHAB** — flexible automation system
*   **Node-RED** — visual programming of scenarios

#### 🎥 Security:
*   **Surveillance cameras** — with motion detection support
*   **Intercoms** — with video call support
*   **Alarm systems** — with notifications and event recording

#### 🌡️ Climate Control:
*   **Smart thermostat** — with schedule and weather forecast support
*   **Humidifier/Air purifier** — with automatic control based on sensors

#### 🎛️ Multimedia:
*   **Media center** — for music and video playback
*   **Voice assistant** — with Alexa, Google Assistant support

### 🛠️ Arduino in Smart Home:

#### 🔧 Simple Sensors:
*   **Temperature/humidity sensors** — DHT11, DHT22, DS18B20
*   **Motion sensors** — HC-SR501 PIR
*   **Water level sensors** — for smart irrigation
*   **Air quality sensors** — MQ-2, MQ-135

#### ⚙️ Actuators:
*   **Relays** — for turning high-voltage devices on/off
*   **Servos** — for controlling locks, blinds
*   **Stepper motors** — for automatic curtains
*   **LED strips** — addressable (WS2812B, NeoPixel)

#### 🌐 Communication:
*   **Wi-Fi modules** — ESP8266, ESP32 (sometimes as standalone solutions)
*   **RF modules** — 433MHz, 315MHz for wireless sensors
*   **Bluetooth** — for local control

**Practical Example:**
Arduino-based irrigation system:
- Soil moisture sensor measures water level
- If level is below threshold, Arduino turns on pump via relay
- Status is transmitted to Raspberry Pi for database logging and owner notification

---

## 4. 🤔 How to Choose Between Raspberry Pi and Arduino

### 🧮 Selection Criteria:

#### 📌 Use Raspberry Pi if:
*   You need an **operating system** (Linux)
*   You require **internet connection** and network protocol work
*   You need **high computational power** (image processing, video)
*   You require **database work**
*   You need **lots of memory** for complex programs
*   You plan a **graphical interface** or web server
*   You require **multitasking** (simultaneous execution of multiple processes)

#### 📌 Use Arduino if:
*   You need **simple sensor reading** and actuator control
*   You require **low power consumption**
*   You need **reliability and simplicity** (no operating system)
*   You require **real-time** (precise timing control)
*   You need **low voltage** (3.3V, 5V)
*   You require **low price** for simple tasks
*   You need **simplicity of connection** to electronic components

#### 📌 Use an old laptop if:
*   You **have an old laptop** that is no longer in use
*   You need **high computational power** and memory
*   You require **working with a full operating system** (Windows, Linux)
*   You plan **complex integration** with existing programs
*   You need **more ports** for connecting devices
*   You require **keyboard and screen** for system control
*   You need a **temporary solution** until purchasing Raspberry Pi
*   You plan **experimentation** with various platforms and programs

### 🎯 Scenario Examples:

#### 🏠 Smart Outlet:
*   **Arduino** — perfectly suited for a simple smart outlet with relay
*   **Raspberry Pi** — excessive for this task

#### 🏠 Smart Home Control Center:
*   **Raspberry Pi** — perfectly suited for running Home Assistant
*   **Arduino** — unsuitable due to computational power limitations

#### 🏠 Surveillance System:
*   **Raspberry Pi** — with camera and motion detection
*   **Arduino** — unsuitable due to need for image processing

#### 🏠 Temperature Sensor:
*   **Arduino** — perfectly suited for reading from DHT22 sensor
*   **Raspberry Pi** — excessive for this task

---

## 5. 🛠️ Practical Project Examples

### 🌡️ Arduino-based Smart Thermostat:
```
Temperature Sensor (DS18B20) → Arduino → Relay → Heater
                    ↓
              (Serial) → Raspberry Pi → Home Assistant
```

**Functionality:**
1. Arduino continuously measures temperature
2. If temperature is below set threshold, heater is turned on via relay
3. Arduino sends temperature data to Raspberry Pi
4. Raspberry Pi displays data in Home Assistant and allows threshold adjustment

### 🚪 Raspberry Pi-based Smart Door:
```
Camera → Raspberry Pi → Facial Recognition → Unlock (GPIO)
   ↓
Telegram Bot → Visitor notifications
```

**Functionality:**
1. Camera captures video from front door
2. Raspberry Pi recognizes faces using face_recognition library
3. If face is known, door lock opens via GPIO
4. If face is unknown, notification is sent to Telegram

### 🌱 Arduino-based Automatic Irrigation:
```
Moisture Sensor → Arduino → Relay → Water Pump
                    ↓
              (Serial) → Raspberry Pi → Notifications
```

**Functionality:**
1. Arduino checks soil moisture every 30 minutes
2. If moisture is below threshold, pump is activated for 5 minutes
3. Irrigation status is sent to Raspberry Pi
4. Raspberry Pi sends notifications and keeps irrigation logs

### 🏠 Raspberry Pi-based Control Center:
```
Raspberry Pi → Home Assistant → MQTT → Devices
                ↓
         Web Interface ← Phone control
```

**Functionality:**
1. Raspberry Pi runs Home Assistant
2. Home Assistant connects to various devices via MQTT
3. User controls devices through web interface or mobile app
4. System supports automation scenarios

---

## 6. 🔧 Connection and Setup

### 🧰 Required Components:

#### 📦 For Raspberry Pi:
*   **Raspberry Pi** (recommended 3B+, 4, or Zero W)
*   **MicroSD card** (at least 16GB, Class 10)
*   **Power supply** (5V/3A for Pi 4, 5V/2.5A for other models)
*   **Case** (optional)
*   **HDMI cable** (for initial setup)
*   **Keyboard and mouse** (for initial setup)

#### 📦 For Arduino:
*   **Arduino Uno** (or other board)
*   **USB cable** (A-Male to B-Male)
*   **Breadboard**
*   **Connecting wires** (jumper wires)
*   **Resistors** (various values)
*   **LEDs** (various colors)
*   **Sensors** (depending on project)

### 🖥️ Raspberry Pi Setup:

#### 1. Installing Operating System:
1. Download Raspberry Pi OS image from official website
2. Use **Raspberry Pi Imager** to write image to SD card
3. Insert SD card into Raspberry Pi and connect power

#### 2. Network Connection:
*   **Wi-Fi:** Configure via GUI or wpa_supplicant.conf file
*   **Ethernet:** Simply connect cable

#### 3. SSH and VNC (optional):
*   Enable SSH: `sudo raspi-config` → Interface Options → SSH
*   Enable VNC: `sudo raspi-config` → Interface Options → VNC

### 🛠️ Arduino Setup:

#### 1. Installing Arduino IDE:
1. Download Arduino IDE from official website
2. Install on your computer
3. Connect Arduino to computer via USB

#### 2. Selecting Board and Port:
1. In Arduino IDE: Tools → Board → select your board (e.g., Arduino Uno)
2. Tools → Port → select appropriate COM port

#### 3. Test Connection:
```cpp
void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
}

void loop() {
  digitalWrite(LED_BUILTIN, HIGH);
  delay(1000);
  digitalWrite(LED_BUILTIN, LOW);
  delay(1000);
}
```

---

## 7. 💻 Programming and Firmware

### 🐍 Raspberry Pi Programming:

#### Python (most popular language):
```python
import RPi.GPIO as GPIO
import time

# GPIO setup
GPIO.setmode(GPIO.BCM)
GPIO.setup(18, GPIO.OUT)

# LED blinking
try:
    while True:
        GPIO.output(18, GPIO.HIGH)
        time.sleep(1)
        GPIO.output(18, GPIO.LOW)
        time.sleep(1)
except KeyboardInterrupt:
    GPIO.cleanup()
```

#### Installing required libraries:
```bash
sudo apt update
sudo apt install python3-pip
pip3 install RPi.GPIO
```

### 🛠️ Arduino Programming:

#### Arduino Program Structure:
```cpp
// Variable and constant declarations
const int ledPin = 13;

void setup() {
  // Initialization (executed once)
  pinMode(ledPin, OUTPUT);
  Serial.begin(9600);  // Serial initialization
}

void loop() {
  // Main loop (executed continuously)
  digitalWrite(ledPin, HIGH);
  delay(1000);
  digitalWrite(ledPin, LOW);
  delay(1000);
  
  // Send data to Serial
  Serial.println("LED toggled");
}
```

#### Example: Arduino as Temperature Sensor:
```cpp
#include <DHT.h>

#define DHT_PIN 2
#define DHT_TYPE DHT22

DHT dht(DHT_PIN, DHT_TYPE);

void setup() {
  Serial.begin(9600);
  dht.begin();
}

void loop() {
  float humidity = dht.readHumidity();
  float temperature = dht.readTemperature();

  if (!isnan(humidity) && !isnan(temperature)) {
    Serial.print("Temp: ");
    Serial.print(temperature);
    Serial.print("°C, Humidity: ");
    Serial.print(humidity);
    Serial.println("%");
  }

  delay(2000);
}
```

### 🔄 Communication Between Arduino and Raspberry Pi:

#### Serial (UART):
```python
# Raspberry Pi (Python)
import serial
import time

ser = serial.Serial('/dev/ttyUSB0', 9600)  # or '/dev/ttyACM0'

while True:
    if ser.in_waiting > 0:
        line = ser.readline().decode('utf-8').rstrip()
        print(f"Received: {line}")
```

```cpp
// Arduino
void setup() {
  Serial.begin(9600);
}

void loop() {
  // Send data every 5 seconds
  Serial.println("Hello from Arduino!");
  delay(5000);
}
```

---

## 8. 🌐 Integration with IoT Platforms

### 🏠 Home Assistant:

#### Connecting Arduino via MQTT:
1. Install PubSubClient library on Arduino
2. Configure Arduino to publish data to MQTT broker
3. In Home Assistant, add MQTT integration

```cpp
#include <ESP8266WiFi.h>
#include <PubSubClient.h>

const char* ssid = "YOUR_WIFI_SSID";
const char* password = "YOUR_WIFI_PASSWORD";
const char* mqtt_server = "YOUR_MQTT_SERVER_IP";

WiFiClient espClient;
PubSubClient client(espClient);

void setup_wifi() {
  delay(10);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
  }
}

void callback(char* topic, byte* payload, unsigned int length) {
  // Handle commands from Home Assistant
}

void reconnect() {
  while (!client.connected()) {
    if (client.connect("arduinoClient")) {
      client.subscribe("home/arduino/command");
    }
    delay(5000);
  }
}

void setup() {
  setup_wifi();
  client.setServer(mqtt_server, 1883);
  client.setCallback(callback);
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  // Publish temperature data
  client.publish("home/arduino/temperature", "23.5");
  delay(10000);
}
```

### 🌐 Cloud Platforms:

#### Blynk:
*   Platform for creating mobile apps to control Arduino/Raspberry Pi
*   Simple integration via Blynk library

#### ThingSpeak:
*   Platform for collecting and visualizing sensor data
*   Supports HTTP API for sending data

#### Cayenne:
*   Platform from MyDevices for IoT device prototyping
*   Supports Arduino, Raspberry Pi, ESP8266, and other platforms

---

## 9. 🔒 Security of Homemade Devices

### 🚨 Potential Threats:

#### Devices:
*   **Insecure passwords** — using default or simple passwords
*   **Outdated software** — non-updatable firmware and libraries
*   **Open ports** — unprotected network connection
*   **Lack of encryption** — data transmission in plain text

#### Network:
*   **Unsecured Wi-Fi network** — using WEP or no password
*   **Lack of segmentation** — IoT devices on same network as main devices

### 🔐 Security Recommendations:

#### For Raspberry Pi:
1. **System updates:**
   ```bash
   sudo apt update && sudo apt upgrade
   ```

2. **Change default credentials:**
   ```bash
   passwd  # change user password
   sudo raspi-config  # change username
   ```

3. **Firewall setup:**
   ```bash
   sudo apt install ufw
   sudo ufw enable
   sudo ufw allow ssh
   sudo ufw allow http
   ```

4. **Disable unnecessary services:**
   ```bash
   sudo systemctl disable bluetooth  # if unused
   sudo systemctl disable avahi-daemon  # if unused
   ```

5. **SSH security:**
   ```bash
   # Editing /etc/ssh/sshd_config
   Port 2222  # change default port
   PermitRootLogin no  # disallow root login
   PasswordAuthentication no  # use keys only
   ```

#### For Arduino:
1. **Code security:**
   * Don't store passwords in plain text in code
   * Use encryption for sensitive data transmission
   * Validate incoming data for correctness

2. **Physical security:**
   * Place devices in protected locations
   * Use enclosures with unauthorized access protection

3. **Firmware updates:**
   * Regularly update used libraries
   * Check for updates for used modules (e.g., ESP8266)

#### 🌐 Network Security:
1. **Create separate IoT network:**
   * Use VLAN or separate SSID for IoT devices
   * Limit access between IoT network and main network

2. **VPN for remote access:**
   * Use VPN instead of opening ports for remote access
   * Configure two-factor authentication

3. **Traffic monitoring:**
   * Use tools for monitoring network activity
   * Configure alerts for suspicious activity

---

## 10. 🚀 Conclusion

Raspberry Pi and Arduino are powerful tools for creating custom solutions in smart homes. They allow implementing both simple tasks (such as controlling LEDs and sensors) and complex systems (such as smart home control centers with facial recognition).

**Raspberry Pi** is ideal for tasks requiring an operating system, network connections, and computational power. **Arduino** is for simple tasks involving sensor reading and actuator control.

By combining these platforms, you can create a flexible and powerful smart home system tailored to specific needs. The key is to consider security when creating and operating homemade devices.

The next materials will discuss other aspects of IoT and smart homes, including communication protocols, cloud platforms, and advanced automation scenarios.

<!-- QUIZ_START
[
    {
        "question": "What is the main difference between Raspberry Pi and Arduino?",
        "options": [
            "Raspberry Pi is more expensive than Arduino",
            "Raspberry Pi is a computer with OS, Arduino is a microcontroller",
            "Arduino has more GPIO pins",
            "Raspberry Pi cannot connect to the internet"
        ],
        "correctIndex": 1
    },
    {
        "question": "Which programming language is used for Arduino?",
        "options": [
            "Python",
            "Java",
            "C/C++",
            "JavaScript"
        ],
        "correctIndex": 2
    },
    {
        "question": "Which component is better for processing images from a camera?",
        "options": [
            "Arduino Uno",
            "Raspberry Pi",
            "ESP8266",
            "ATtiny85"
        ],
        "correctIndex": 1
    }
]
QUIZ_END -->