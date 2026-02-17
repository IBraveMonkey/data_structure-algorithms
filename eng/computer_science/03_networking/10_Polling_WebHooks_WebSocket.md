# 📡 Data Exchange: Polling, Webhooks, WebSockets


## 📑 Table of Contents
1. [Short Polling](#short-polling)
2. [Long Polling](#long-polling)
3. [Webhooks](#webhooks)
4. [WebSockets](#websockets)
5. [Server-Sent Events (SSE)](#server-sent-events-sse)
6. [Technology Comparison](#technology-comparison)

---

How does a server notify a client that something new has happened? In standard HTTP, the server remains silent until it is explicitly asked for information. Several approaches have been developed to solve this problem.

---


## 1. ⏱️ Short Polling

The simplest and most "naive" method. The client periodically "spams" the server with the question: "Is there anything new?"

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Client as Client
    participant Server as Server

    Client->>Server: Any news?
    Server-->>Client: No.
    Note over Client: Waits for 5 seconds
    Client->>Server: Anything now?
    Server-->>Client: Yes, here is your message!







```

> [!WARNING]
> **Cons**: Significant unnecessary load on the server and network. 99% of requests may return empty data, yet resources are still consumed to establish each HTTP connection.

---


## 2. ⏳ Long Polling

An optimized version of polling. The client makes a request, and rather than responding immediately, the server **holds** the request open until data becomes available (or a timeout occurs).

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Client as Client
    participant Server as Server

    Client->>Server: Any news?
    Note over Server: Server waits for data...
    Note over Server: (20 seconds pass)
    Server-->>Client: Yes, here are the updates!
    Note over Client: Client immediately sends a new request
    Client->>Server: Waiting for the next update...







```

> [!TIP]
> This is a good compromise if you don't require ultra-fast real-time synchronization but want to save on server resources.

---


## 3. ⚓ Webhooks

This is "inverted" HTTP. Instead of the client reaching out to the server, the **server reaches out to the client**.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Payment as Payment Provider
    participant App as Your Application

    Note over Payment: Payment completed
    Payment->>App: POST /webhook {status: "success"}
    App-->>Payment: 200 OK







```

> [!IMPORTANT]
> **Accessibility Requirement**: Your application must have a public IP/URL so the external service can reach it. This is commonly used for integrations with services like Stripe, GitHub, or Telegram.

---


## 4. 🔌 WebSockets

A full-duplex, bidirectional communication protocol. After an initial "handshake," the client and server act as equal participants in the conversation.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Client as Client
    participant Server as Server

    Client->>Server: HTTP Handshake (Upgrade: websocket)
    Server-->>Client: 101 Switching Protocols
    Note over Client, Server: Channel Open (TCP)
    Client->>Server: Message 1
    Server->>Client: Response 1
    Server->>Client: Urgent Notification!
    Client->>Server: Okay, received.







```

- **Pros**: Minimal latency and low overhead since headers are reduced.
- **Cons**: Difficult to scale horizontally; every client keeps an open connection, which consumes server memory and connection limits.

---


## 5. Server-Sent Events (SSE)

A unidirectional "streaming" method from the server to the client over a standard HTTP connection.

> [!NOTE]
> SSE is ideal for news feeds or stock tickers where the client only needs to listen to updates and doesn't need to send data back to the server over the same channel.

---


## 6. 📊 Technology Comparison

| Technology | Direction | Latency | Implementation Complexity |
|:---|:---|:---:|:---|
| **Short Polling** | Client -> Server | High | Low |
| **Long Polling** | Client -> Server | Medium | Medium |
| **Webhooks** | Server -> Client | Low | Medium |
| **WebSockets** | Bidirectional | Minimal | High |

---


## 🎯 Key Takeaways

- **Short Polling**: Best for MVPs or non-critical, simple tasks.
- **Long Polling**: Use when WebSockets are not feasible.
- **Webhooks**: Standard for server-to-server notifications.
- **WebSockets**: Essential for chat applications, multiplayer games, and real-time financial dashboards.

<!-- QUIZ_START 
[
    {
        "question": "Which data exchange approach implies that the server initiates an HTTP request to the client (usually another server) when an event occurs?",
        "options": ["Short Polling", "Long Polling", "Webhooks", "WebSockets"],
        "correctIndex": 2
    },
    {
        "question": "What is the primary advantage of WebSockets over standard HTTP for real-time applications?",
        "options": ["Browser caching support", "Full-duplex bidirectional communication with low latency", "Simplicity of client-side implementation", "No need for a TCP connection"],
        "correctIndex": 1
    },
    {
        "question": "Which technology is best suited for unidirectional data streaming (e.g., news feeds) from server to client over standard HTTP?",
        "options": ["Short Polling", "Server-Sent Events (SSE)", "Webhooks", "JSONP"],
        "correctIndex": 1
    }
]
QUIZ_END -->
