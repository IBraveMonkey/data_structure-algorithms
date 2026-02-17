# 🛡️ Network Security: Infrastructure


## 📑 Contents

1. [Secure Access: Proxy(Private Network) & Tunneling](#secure-access-proxyprivate-network--tunneling)
2. [Channel Encryption (TLS/SSL)](#channel-encryption-tlsssl)
3. [Network Layer Attacks](#network-layer-attacks)
4. [Infrastructure Protection Checklist](#infrastructure-protection-checklist)

---

Network security is the foundation. If the "pipe" through which data flows is leaking, no application-level security will help.
In this section, we will cover **infrastructure and connectivity protection**.

> [!NOTE]
> A deep dive into cryptography, authentication (JWT, OAuth), and application code security can be found in the [**07_Security**](../07_security/README.md) section.

---


## 1. 🚇 Secure Access: Proxy(Private Network) & Tunneling

How to transmit secret data over the "dirty" internet so that no one sees it? Use tunnels.


### Proxy(Private Network) (PPN)
Connects remote devices into a single **virtual private network**.
*   **Scenario**: You are in a cafe, but the database server is only accessible from the office network. Connect to PPN — and you are "virtually" in the office.
*   **WireGuard**: Modern, fast, and lightweight protocol (state-of-the-art).
*   **OpenVPN**: Industry standard, reliable, but slower.


### Tunneling and Encapsulation
This is the process of packing one protocol inside another.
*   **SSH Tunnel (Port Forwarding)**: Forwarding a local port to a remote server via an encrypted SSH connection.
    *   `ssh -L 5432:localhost:5432 user@server` — now your local DB points to the remote one.

> [!TIP]
> Proxy servers hide the client or protect the server, but do not always encrypt traffic like PPNs do. Read more in [**09_Proxy, Gateway, LoadBalancer**](./09_Proxy_Gateway_LoadBalancer.md).

---


## 2. 🔐 Channel Encryption (TLS/SSL)

The main goal at the network layer is **Confidentiality** (no eavesdropping) and **Integrity** (no packet tampering).

**TLS (Transport Layer Security)** is the de-facto standard for secure data transmission (HTTPS, FTPS, SMTPS).
It ensures that you have connected specifically to `google.com`, not a hacker, and that your traffic is encypted.

> [!IMPORTANT]
> **SSL is dead** and insecure. All modern systems use TLS 1.2 or 1.3.


### Learn More:
*   📜 **How the Handshake and Certificates work?**
*   🔑 **What is the difference between symmetric and asymmetric encryption?**
*   👉 Read the detailed guide in [**TLS and SSL: Deep Dive**](../07_security/2.TLS_SSL.md).

---


## 3. ⚔️ Network Layer Attacks

Here we look at attacks specifically targeting the **infrastructure**.


### ARP Spoofing / Poisoning (L2 Attack)
An attack inside a local network. A hacker tells your computer: "I am the router." And all your traffic goes through them.
*   *Defense*: Port isolation on switches (Port Security), using PPN in public networks.


### DDoS (Distributed Denial of Service)
Attack on availability.
1.  **Volumetric (L3/L4)**: Clogging the channel with garbage (UDP Flood, SYN Flood).
2.  **Application (L7)**: Smart requests loading the CPU (more in the Security section).

> [!WARNING]
> For a full breakdown of phishing, social engineering, and malware, read [**Cyber Threats and Scams**](../07_security/0.Cyber_Threats_Scams.md).


### IP Spoofing
Spoofing the sender's IP address. A hacker sends a packet supposedly from a "trusted" server.
*   *Defense*: Traffic filtering at the network edge (Ingress/Egress filtering).

---


## 4. 🛡️ Infrastructure Protection Checklist

What is important to do when setting up a server?

- [ ] **Firewall (UFW/iptables)**: Close ALL ports except necessary ones (usually 80, 443, and 22).
- [ ] **SSH Hardening**:
    -   Disable password login (`PasswordAuthentication no`).
    -   Only use SSH keys.
    -   Disable root login (`PermitRootLogin no`).
    -   Change the default port 22 (protection against scanners).
- [ ] **Private Network (VPC)**: The database MUST NOT be exposed to the internet. It should only be accessible to the application within a private network.
- [ ] **Fail2Ban**: Automatic banning of IPs that fail authentication too often.
- [ ] **HSTS**: Configure the web server to enforce HTTPS.

> [!TIP]
> The security checklist for **application code** (SQL Injection, XSS, CSRF) is in [**Best Practices**](../07_security/3.Security_Best_Practices.md).

<!-- QUIZ_START 
[
    {
        "question": "Which version of the TLS protocol is currently considered the most modern and secure?",
        "options": ["SSL 3.0", "TLS 1.0", "TLS 1.1", "TLS 1.3"],
        "correctIndex": 3
    },
    {
        "question": "What is the primary difference between an L7 (Application) DDoS attack and an L3/L4 (Network/Transport) attack?",
        "options": ["L7 is always larger in terms of traffic volume", "L7 targets application logic (e.g., heavy SQL queries) rather than just clogging the network pipe", "L7 attacks are impossible to mitigate", "L7 is only used against mobile apps"],
        "correctIndex": 1
    },
    {
        "question": "What is the primary purpose of SSH Tunneling (Port Forwarding)?",
        "options": ["To increase internet speed", "To forward a local port to a remote server via an encrypted connection", "To automate OS updates", "To create graphical user interfaces"],
        "correctIndex": 1
    }
]
QUIZ_END -->

