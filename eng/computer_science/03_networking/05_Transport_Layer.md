# 🚛 OSI L4: The Transport Layer


## 📑 Table of Contents
1. [TCP: Order and Reliability](#transmission-control-protocol-tcp-reliability)
2. [UDP: Speed and Risk](#user-datagram-protocol-udp-speed-vs-reliability)
3. [Ports and Sockets](#ports-and-sockets-how-applications-communicate)
4. [The TCP 3-Way Handshake](#three-way-handshake)

---

The Transport Layer is the "delivery service" inside your OS. It connects two specific processes (programs) running on two computers, rather than just the computers themselves.

---


## 1. 🤝 TCP — The "Reliable Contract"

**TCP (Transmission Control Protocol)** is a connection-oriented protocol. It guarantees that data arrives in full, in the correct order, and without duplicates.


### Flow Control
Imagine a high-speed server sending data to an old smartphone. The smartphone can't process it fast enough.
- TCP uses a **Sliding Window**. The receiver tells the sender: "I have room for only 10 KB." The sender won't send more until it receives confirmation that space has been cleared.


### Congestion Control
If the internet itself is "congested" (routers are overloaded), TCP senses this through packet loss:
- **Slow Start**: Start slowly and double the speed as long as everything is going well.
- **Congestion Avoidance**: If a packet is lost, TCP sharply reduces its speed and then begins to grow more cautiously. This prevents the internet from collapsing under its own weight.

---


## 🏁 UDP — "Fire and Forget"

**UDP (User Datagram Protocol)** doesn't waste time on handshakes or confirmations.
- **Overhead**: A UDP header is only 8 bytes (TCP's is at least 20 bytes).
- **Use Cases**: Critical when speed is more important than perfect accuracy. If one frame is lost in a video call, you won't even notice. But if you wait 500ms for it to be re-requested, the call becomes a slideshow.

---


## 🚪 Ports and Sockets: Where to Knock?

An IP address leads to the "house" (the server). But inside the house, there are hundreds of "doors" (ports).
- **System Ports (0-1023)**: HTTP (80), HTTPS (443), SSH (22). These typically require administrative privileges.
- **Ephemeral (Dynamic) Ports**: When your browser connects to a server, it opens a random port (e.g., 54321) so the server knows exactly where to send the response.

> [!TIP]
> **TIME_WAIT**: A socket state after a TCP connection is closed. The OS keeps it "busy" for a few minutes so that "lost" packets still in the network don't accidentally join a new connection. This is a common reason for `Address already in use` errors during frequent application restarts.

---


## 🚀 Modernity: QUIC and Reliable UDP

Modern giants (Google, Facebook) are moving away from TCP in favor of **QUIC**. This protocol runs on top of UDP but implements reliability and encryption (TLS 1.3) directly. This allows connections to be established twice as fast.

---


## 📊 TCP vs. UDP: Quick Reference

| Property | TCP | UDP |
|:---|:---|:---|
| **Connection** | Required (3-way handshake) | Not required |
| **Reliability** | Guaranteed (Retransmission) | Not guaranteed |
| **Ordering** | Strict (Sequencing) | Best effort |
| **Speed** | Lower due to checks | Maximum possible speed |
| **Typical Scenarios** | Databases, APIs, Files | Streaming, Online Games, DNS, QUIC |

---


## 🎯 Key Takeaways

- **TCP** is used for critical data where accuracy is paramount (e.g., financial transactions, source code).
- **UDP** is used for real-time data where speed is more important than perfect recovery (e.g., voice, video).
- **Ports** allow a single computer with one IP address to run hundreds of different network applications simultaneously.

<!-- QUIZ_START 
[
    {
        "question": "What is the process of establishing a connection in the TCP protocol, consisting of the exchange of SYN, SYN-ACK, and ACK flags, called?",
        "options": ["Translation", "Handshake", "Encapsulation", "Session"],
        "correctIndex": 1
    },
    {
        "question": "For which of the following tasks is the UDP protocol best suited?",
        "options": ["Transferring bank transactions", "Loading a web page (HTML/JS)", "Online gaming and real-time video calls", "Sending emails"],
        "correctIndex": 2
    },
    {
        "question": "What is a network 'Socket'?",
        "options": ["A type of network cable", "A software library for JSON parsing", "A combination of an IP address, protocol type, and port number", "A physical connector on the motherboard"],
        "correctIndex": 2
    }
]
QUIZ_END -->
