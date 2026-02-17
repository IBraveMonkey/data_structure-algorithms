# 17. 🔐 SSH and Remote Access

## 📑 Contents
1. [What is SSH?](#what-is-ssh)
2. [How SSH Works](#how-ssh-works)
3. [SSH Keys (Public/Private Key Authentication)](#ssh-keys-publicprivate-key-authentication)
4. [SSH Tunneling (Port Forwarding)](#ssh-tunneling-port-forwarding)
5. [Practical Usage Examples](#practical-usage-examples)
6. [SSH Security](#ssh-security)
7. [SSH Alternatives](#ssh-alternatives)

---

## 1. 🤔 What is SSH?

**SSH (Secure Shell)** is a network protocol that provides secure remote login and command execution over an insecure network.

### 🔄 Analogy
Think of SSH as a **secure tunnel** between two cities. You can safely transport any documents from one city to another, even if the territory between them is controlled by bandits.

### 🎯 Main SSH Capabilities:
*   **Remote command-line access** (Terminal/Shell)
*   **Secure file transfer** (SCP, SFTP)
*   **Port forwarding** (Tunneling)
*   **Task automation** on remote servers

---

## 2. 🔄 How SSH Works?

### 🏗️ SSH Connection Architecture
```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    subgraph Client ["Client (Your Computer)"]
        A[SSH Client]
    end

    subgraph Internet ["Internet (insecure network)"]
        B[Encrypted Connection]
    end

    subgraph Server ["Server (remote computer)"]
        C[SSH Server]
    end

    A --- B --- C
    style A fill:#99ff99,stroke:#333
    style C fill:#99ff99,stroke:#333
    style B fill:#009688,stroke:#333,stroke-width:2px
```

### 🔄 Connection Establishment Process:
1.  **Connection:** Client connects to server via TCP port (usually 22).
2.  **Encryption negotiation:** Client and server agree on encryption algorithm.
3.  **Authentication:** Client proves identity to server (password or keys).
4.  **Secure channel:** All data is now transmitted encrypted.

### 🧪 Connection Example:
```bash
# Connect to server by IP address
ssh username@server_ip_address

# Example:
ssh admin@192.168.1.100

# Connect with custom port (if not standard)
ssh -p 2222 username@server_ip_address
```

---

## 3. 🔑 SSH Keys (Public/Private Key Authentication)

### 🤝 How It Works?
Instead of passwords, you can use a **key pair**:
*   **Private Key** — stored on your machine (never transmitted over network!)
*   **Public Key** — stored on the server

### 🔄 Authentication Process:
```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Client as Client
    participant Server as Server

    Client->>Server: Connecting to SSH
    Server-->>Client: Send me proof of your identity
    Client->>Server: Encrypt random message with my Private Key
    Server->>Server: Decrypt message with Public Key (if matches - this is indeed the key owner!)
    Server-->>Client: Welcome!
```

### 🛠️ Creating a Key Pair:
```bash
# Generate a new key pair
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"

# Or more modern algorithm
ssh-keygen -t ed25519 -C "your_email@example.com"
```

### 📁 File Structure:
*   `~/.ssh/id_rsa` or `~/.ssh/id_ed25519` — Private Key (stored locally)
*   `~/.ssh/id_rsa.pub` or `~/.ssh/id_ed25519.pub` — Public Key (copied to server)

### 🧾 Examples of Keys:

#### Private Key (id_rsa or id_ed25519):
```
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXAAAEBAAAAAC3NzaC
1lZDI1NTE5AAAAIEB7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXAAAAIHYWRtaW
4BZ3Vlc3QB
-----END OPENSSH PRIVATE KEY-----

#### Public Key (id_rsa.pub or id_ed25519.pub):

```
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEB7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VXx7VX admin@computer.local
```

### 🔐 What Keys Look Like:
*   **Private Key** — long encoded text starting with `-----BEGIN` and ending with `-----END`. 
    *   **NEVER** share this with anyone!
    *   Stored only on your computer
    *   Has file permissions 600 (only you can read/write)
*   **Public Key** — starts with key type (`ssh-rsa`, `ssh-ed25519`, `ssh-dss`)
    *   Safe to share with others
    *   Placed in `~/.ssh/authorized_keys` file on server
    *   Format: `key_type encoded_key comment`

### 📤 Copying Public Key to Server:
```bash
# Automatic copy
ssh-copy-id username@server_ip_address

# Or manual copy
cat ~/.ssh/id_rsa.pub | ssh username@server_ip_address "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"
```

### 🛡️ Key Security:
*   **Private Key** must be protected:
    *   Set proper permissions: `chmod 600 ~/.ssh/id_rsa`
    *   Use passphrase when generating (optional but safer)
    *   Never transmit over unsecured channels

### 🧩 Files in ~/.ssh directory:
*   `id_rsa` / `id_ed25519` — your private key
*   `id_rsa.pub` / `id_ed25519.pub` — your public key
*   `authorized_keys` — list of trusted public keys (on server)
*   `known_hosts` — list of servers you've connected to

### 🔄 How SSH automatically uses keys:
*   When connecting, SSH automatically looks for keys in `~/.ssh/`
*   Tries to use `id_rsa`, `id_ed25519` and other standard names
*   If key is protected with passphrase, you'll be prompted once per session
*   After successful authentication, keys are used for subsequent connections in the session
*   SSH-agent can store keys in memory to avoid entering passphrase repeatedly

### 🤖 SSH Agent (ssh-agent):
*   Special service that stores your keys in memory
*   Starts automatically on most systems
*   When adding key to agent, you enter passphrase only once
*   Example usage:
  ```bash
  # Start agent (if not running)
  eval "$(ssh-agent -s)"
  
  # Add key to agent
  ssh-add ~/.ssh/id_rsa
  
  # Now you can connect without entering passphrase
  ssh user@server
  ```

---

## 4. 🚇 SSH Tunneling (Port Forwarding)

### 🤔 What Is It?
SSH Tunneling allows you to **forward traffic** through an encrypted SSH connection.

### 🧩 Types of Tunneling:

#### 🔍 Local Port Forwarding (L)
Forwards a local port to a remote server.
```bash
# Example: localhost:8080 -> server:3306 (MySQL)
ssh -L 8080:localhost:3306 username@server_ip_address
```
*   **Analogy:** You open a door in your house (8080), and behind it is a room on the remote server (3306)*

#### 🌐 Remote Port Forwarding (R)
Forwards a remote port to your local machine.
```bash
# Example: server:9000 -> localhost:80 (local web server)
ssh -R 9000:localhost:80 username@server_ip_address
```
*   **Analogy:** You create a door on the remote server (9000) that leads to your home (80)*

#### 🔄 Dynamic Port Forwarding (D)
Creates a SOCKS proxy through SSH.
```bash
# Example: Create SOCKS proxy on port 1080
ssh -D 1080 username@server_ip_address
```
*   **Analogy:** You create a universal portal through which any traffic can pass*

---

## 5. 🛠️ Practical Usage Examples

### 🖥️ Server Management
```bash
# Connect to cloud server
ssh ubuntu@myserver.example.com

# Execute command on remote server
ssh admin@server "df -h && free -m"
```

### 📂 File Transfer
```bash
# Copy file to server (SCP)
scp file.txt admin@server:/home/admin/

# Copy file from server
scp admin@server:/var/log/app.log ./

# Copy directory
scp -r ./local_folder admin@server:/remote/path/
```

### 🌐 Secure Access to Internal Services
```bash
# Access internal database via SSH tunnel
ssh -L 3307:internal-db:3306 admin@jumpserver

# Now connect to localhost:3307 instead of internal server
mysql -h localhost -P 3307 -u user -p
```

### ⚙️ Task Automation
```bash
# Run script on multiple servers
for server in server1 server2 server3; do
    ssh admin@$server "uptime && df -h"
done
```

---

## 6. 🛡️ SSH Security

### 🔒 Security Best Practices:
1.  **Use keys instead of passwords** — more secure and convenient
2.  **Disable root login** — prevents direct superuser access
3.  **Change default port** — reduces automated attack attempts
4.  **Restrict user list** — allow SSH only for needed users
5.  **Use Fail2Ban** — blocks IPs after multiple failed attempts

### ⚙️ Configuring `/etc/ssh/sshd_config`:
```bash
# Not recommended to use root user directly
PermitRootLogin no

# Change default port (e.g., to 2222)
Port 2222

# Allow only specific users
AllowUsers admin deploy

# Disable password authentication (keys only)
PasswordAuthentication no

# Enable key authentication
PubkeyAuthentication yes
```

### 🚨 Common Threats:
*   **Brute Force attacks** — password guessing
*   **MITM (Man-in-the-Middle)** — connection interception
*   **Private Key Compromise** — loss of private key

---

## 7. 🔄 SSH Alternatives

### 🌐 Telnet
*   **Old protocol** for remote access
*   **Insecure** — transmits data in plain text
*   **Rarely used**, mainly for diagnostics

### 🖱️ RDP (Remote Desktop Protocol)
*   **For Windows** — allows remote desktop control
*   **Graphical interface** instead of command line
*   **Heavier protocol**

### 💻 VNC (Virtual Network Computing)
*   **Remote screen access** to any computer
*   **Platform-independent** (Windows, Linux, macOS)
*   **Can be slow** on weak connections

### 🌍 HTTP(S) API
*   **For automation** — execute commands via REST API
*   **Secure** when properly implemented
*   **Scalable** — can manage thousands of devices

---

## 8. 🚀 Summary

SSH is a powerful tool for secure remote access to computers. It provides:

*   **Encryption** of all traffic between client and server
*   **Reliable authentication** via passwords or keys
*   **Flexible capabilities** for tunneling and file transfer
*   **Security** when properly configured

<!-- QUIZ_START
[
    {
        "question": "What does the acronym SSH stand for?",
        "options": [
            "Secure System Hosting",
            "Simple Shell Handler",
            "Secure Shell",
            "System Security Hub"
        ],
        "correctIndex": 2
    },
    {
        "question": "What is the default port used by SSH?",
        "options": [
            "21",
            "22",
            "23",
            "80"
        ],
        "correctIndex": 1
    },
    {
        "question": "What is an SSH Key Pair?",
        "options": [
            "Two identical keys for double protection",
            "A password split into two parts",
            "Public and private keys used for authentication",
            "Keys for encrypting and decrypting data"
        ],
        "correctIndex": 2
    }
]
QUIZ_END -->