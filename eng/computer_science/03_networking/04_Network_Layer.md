# 📍 OSI L3: The Network Layer


## 📑 Table of Contents
1. [IP Protocol (v4 and v6)](#-the-ip-protocol--global-addressing)
2. [Packet Routing](#-how-routing-works-hop-by-hop)
3. [NAT: Living with IP Scarcity](#-nat-network-address-translation)
4. [Router vs. Switch](#the-role-of-routers)

---

 The Network Layer is the "post office" of the internet. Its task is to deliver a packet from Point A (e.g., London) to Point B (e.g., Tokyo), even if there are hundreds of intermediate networks and thousands of routers in between.

---


## 1. 📬 The IP Protocol — Global Addressing

While a MAC address (L2) is like a person's name, an IP address is like their current mailing address, which can change when they move.


### IPv4 vs. IPv6: More Than Just Length
- **IPv4**: 32 bits (`192.168.1.1`). The world ran out of these addresses long ago.
- **IPv6**: 128 bits (`2001:0db8:85a3:0000:0000:8a2e:0370:7334`). This is more than just more addresses—it includes built-in security (IPsec), no fragmentation by routers, and simplified headers for faster processing.

---


## 2. 🗺️ How Routing Works (Hop-by-Hop)

A packet doesn't know the whole path. It simply jumps (hops) from one router to another, guided by a **Routing Table**.


### Core Concepts:
- **Default Gateway**: The "exit door." If a router doesn't know where to send a packet, it sends it here (usually to the ISP).
- **Metric**: The "cost" of a path. Routers choose the path with the lowest metric (the fastest or cheapest).
- **TTL (Time To Live)**: A packet cannot live forever in the network. After each hop, the TTL decreases by 1. If TTL = 0, the packet is discarded. This saves the internet from "infinite loops" of traffic.

---


## 📡 ICMP: The Network's Service Notes

ICMP is the "voice" of Layer 3. it doesn't carry user data—it reports on the network's state and problems.

1. **Ping (Echo Request/Reply)**: "Are you there?" — "Yes, I'm here." Testing reachability.
2. **Traceroute**: Uses TTL. It first sends a packet with TTL=1 (the 1st router responds), then TTL=2 (the 2nd responds), and so on, until it reaches the goal. This reveals the entire chain.
3. **Destination Unreachable**: A router's response if it cannot find a path to the target.

---


## 3. 🌐 NAT (Network Address Translation)

NAT (Network Address Translation) is a technology that allows the conversion of IP addresses when transmitting network packets through a router or firewall. The main purpose of NAT is to conserve global IP addresses by using private IP addresses within local networks and one or more public IP addresses for internet access.

### Why is NAT needed?

Imagine this situation: at home you have several devices (laptop, phone, tablet, smart speaker), all connected to the same Wi-Fi router. Each device has its own IP address in your home network (for example, 192.168.1.10, 192.168.1.11, etc.), but your router has only one public IP address provided by your internet service provider.

When you open a website, your laptop sends a request to that website. But how does the website know who to send the response to if all devices in your network share the same external IP address? That's where NAT comes into play.

### Private and Public IP Addresses

#### Private IP Addresses:
- Not routed on the internet
- Used within local networks
- Ranges:
  - 10.0.0.0 - 10.255.255.255 (10.0.0.0/8)
  - 172.16.0.0 - 172.31.255.255 (172.16.0.0/12)
  - 192.168.0.0 - 192.168.255.255 (192.168.0.0/16)

#### Public IP Addresses:
- Routed on the internet
- Globally unique
- Provided by internet service providers

### How NAT Works?

When a device in a private network sends a packet to an external network:

1. **Packet Capture**: Router intercepts the outgoing packet
2. **Address Replacement**: Source IP address (private) is replaced with the router's public IP address
3. **Port Replacement**: Source port is replaced with a unique port on the router
4. **Table Entry**: Router stores the mapping (internal IP:port → external IP:port) in its NAT table
5. **Forwarding**: Packet is sent further to the external network

When a response packet arrives:

1. **Table Lookup**: Router checks its NAT table
2. **Reverse Translation**: Packet is directed back to the internal device with restoration of the original IP address and port
3. **Delivery**: Packet is delivered to the appropriate device in the private network

### Types of NAT

#### 1. Static NAT
- Fixed mapping of one internal IP to one external IP
- Used for servers that need a constant external IP
- Example: Internal IP 192.168.1.10 ↔ External IP 203.0.113.10

#### 2. Dynamic NAT
- Mapping from a pool of available external IPs
- External IPs are assigned dynamically
- Less common, often used as an intermediate step to PAT

#### 3. PAT (Port Address Translation) or Overloading
- Most common type
- Multiple internal addresses are mapped to one external IP using different ports
- Saves the most public IP addresses
- Example: 192.168.1.10:5000 → 203.0.113.1:10000, 192.168.1.11:5000 → 203.0.113.1:10001

### Benefits of NAT

1. **Conserves Public IP Addresses**: Allows many devices to use one public IP
2. **Simplifies Network Migration**: When changing providers, only the external IP needs to be changed
3. **Enhanced Security**: Hides internal network topology from the outside world
4. **Simplified Management**: Easier to manage internal IP addresses

### Working with NAT for Developers

When developing network applications, it's important to understand how NAT works:

- **Servers**: Usually behind NAT but need to be accessible from the outside (static NAT or port forwarding is used)
- **P2P Applications**: Require special methods for establishing connections through NAT (STUN, TURN, ICE)
- **Sessions**: NAT may close connections if there is no activity for a certain period of time

---


## 💡 Why is this important to understand?

1. **Private vs. Public IP**: Your servers in a data center usually have "gray" (local) IPs. Access from the world goes through NAT or a Load Balancer.
2. **Anycast IP**: A technology where one IP address belongs to multiple servers in different parts of the world (used by Cloudflare/Google DNS). The packet will fly to the physically closest server.
3. **Latencies**: Understanding "hops" helps explain why a server in the USA will have a slow response time, even if you have 1 Gbps bandwidth—the speed of light in fiber optics is finite.

---


## 🎯 Key Takeaways

- **L3** turns many small networks into one giant one.
- A **Router** is a dispatcher choosing the best path (Hop).
- **TTL** is the packet's expiration date, preventing chaos.

<!-- QUIZ_START
[
    {
        "question": "What is the TTL (Time To Live) field in an IP packet header used for?",
        "options": ["To measure internet speed", "To prevent a packet from circling the network indefinitely in case of loops", "To encrypt data", "To determine the content type"],
        "correctIndex": 1
    },
    {
        "question": "Which technology allows an entire private network to access the internet using just one single public IP address?",
        "options": ["DNS", "DHCP", "NAT", "BGP"],
        "correctIndex": 2
    },
    {
        "question": "Which version of the IP protocol was created to solve the address shortage problem and uses 128-bit addresses?",
        "options": ["IPv4", "IPv5", "IPv6", "IPX"],
        "correctIndex": 2
    },
    {
        "question": "What does the acronym NAT stand for in the context of networking technologies?",
        "options": ["Network Authentication Token", "Network Address Translation", "Network Allocation Table", "Network Access Terminal"],
        "correctIndex": 1
    },
    {
        "question": "Which type of NAT allows multiple internal devices to use one external IP address with different port numbers?",
        "options": ["Static NAT", "Dynamic NAT", "PAT (Port Address Translation)", "None of the above"],
        "correctIndex": 2
    }
]
QUIZ_END -->
