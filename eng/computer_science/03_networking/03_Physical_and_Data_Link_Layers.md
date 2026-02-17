# 🔌 OSI L1 & L2: Physical and Data Link Layers


## 📑 Table of Contents
1. [Physical Layer (L1): Bits and Signals](#physical-signals)
2. [Data Link Layer (L2): Frames and MAC addresses](#data-link-frames)
3. [Switching](#switching-and-collisions)
4. [The Ethernet Frame](#ethernet-frame-structure)

---

These layers are the "foundation" of the network. They are responsible for ensuring that an electrical impulse or a beam of light is transformed into a structured data format.

---


## 1. ⚡ L1: Physical Layer

At this layer, there are no names or addresses—only volts, hertz, and bits. The primary goal is to deliver a bit from point A to point B as accurately as possible.


### Transmission Media:
- **Copper (Ethernet)**: Uses electrical signals. It's cheap, but the signal fades after 100 meters, and the cable picks up interference from microwave ovens or elevators.
- **Fiber Optics**: Uses light. Incredible speeds (Tbps), distances of tens of kilometers, and complete immunity to electrical interference. The "gold standard" for backbones.
- **Radio (Wi-Fi/5G)**: The medium is shared. All devices "shout" into the same air, which creates challenges for access control.


### Important L1 Concepts:
- **Duplex**: 
  - **Half-Duplex**: Like a walkie-talkie—we take turns speaking. 
  - **Full-Duplex**: Like a telephone—we speak simultaneously (modern Ethernet).
- **Bandwidth**: The width of the channel in Hz, which determines its throughput (Mbps).

---


## 2. 🧱 L2: Data Link Layer

At this stage, bits are grouped into **Frames**. If L1 is the road, L2 is the traffic rules at a single intersection.


### MAC Address: The Name in the Passport
Every device has a unique physical address like `AA:BB:CC:DD:EE:FF`. It's needed so that neighbors on a local network can recognize each other.


### The Switch
This device operates at L2. It's "smarter" than old hubs because it builds a **MAC Table**. It knows: "Computer with MAC A is on port 1, MAC B is on port 2." It sends data specifically to the destination, not to everyone.

---


## 🚦 How they work together (The ARP Example)

When your PC knows a neighbor's IP but doesn't know their MAC:
1. Your PC shouts to the whole network (Broadcast): "Hey, who has IP 192.168.1.5? Tell me your MAC!"
2. The node with that IP responds: "That's me, here's my MAC."
3. Your PC remembers this in the **ARP Table** and starts sending frames.

---


## 💡 Why is this important to understand?

1. **MTU (Maximum Transmission Unit)**: The frame limit (usually 1500 bytes). Sending packets larger than this causes fragmentation at L3, which puts a heavy load on the router's CPU.
2. **Wi-Fi Interference**: Understanding that Wi-Fi is a shared medium explains why speeds drop when neighbors turn on their routers on the same channel.
3. **Cloud Networking**: In AWS/GCP, Layer 2 is often virtualized (SDN), allowing IP addresses to be moved between servers without changing physical settings.

---


## 🎯 Key Takeaways

- **L1** is "physics" and signals.
- **L2** is the logic of "face-to-face" communication (in one subnet).
- **The Switch** sends data only to the intended recipient.

<!-- QUIZ_START 
[
    {
        "question": "Which device at the Data Link layer (L2) learns the MAC addresses of devices on its ports for efficient frame forwarding?",
        "options": ["Hub", "Router", "Switch", "Repeater"],
        "correctIndex": 2
    },
    {
        "question": "What is the name for the maximum frame size (typically 1500 bytes), exceeding which leads to data fragmentation?",
        "options": ["TTL", "MTU", "Bandwidth", "RTT"],
        "correctIndex": 1
    },
    {
        "question": "Where is a device's MAC address typically 'burned in'?",
        "options": ["In RAM", "In the Operating System", "In the network interface card (NIC) at the factory", "In the router's configuration file"],
        "correctIndex": 2
    }
]
QUIZ_END -->
