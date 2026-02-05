# 🧠 Memcached

## 📑 Table of Contents
1. [What is it? (Simple Cache)](#what-is-it-simple-cache)
2. [Architecture (Multithreading)](#architecture-multithreading)
3. [Memcached vs. Redis](#memcached-vs-redis)

---

## 1. 🤔 What is it? (Simple Cache)

**Memcached** is the "grandfather" of all caching systems. it is a distributed **Key-Value** store designed to reside entirely in RAM.

*   **Primitive**: It supports only basic operations like `SET`, `GET`, and `DELETE`. It has no support for data structures like lists, sets, or complex sorting.
*   **LRU (Least Recently Used)**: When memory reaches its limit, Memcached automatically evicts the oldest, least-frequently accessed data.
*   **Volatility**: All data is stored in volatile memory. If the server is powered down or restarted, all data is lost. It does not support persistence to disk (unlike Redis's RDB/AOF).

---

## 2. ⚡ Architecture (Multithreading)

The defining architectural difference between Memcached and Redis is **Multithreading**.

*   **Redis**: Single-threaded. It utilizes only one CPU core to 100%.
*   **Memcached**: Multi-threaded. It can leverage every core on a high-end server (e.g., 64 cores).

This makes Memcached ideal for **simple caching of small objects** (HTML fragments, session data, SQL query results) at extremely high loads where Redis might become CPU-bound.

---

## 3. 🥊 Memcached vs. Redis

| Feature | 🧠 Memcached | 🍒 Redis |
| :--- | :--- | :--- |
| **Data Types** | String only | String, List, Set, Hash, Stream, etc. |
| **Execution Model** | Multi-threaded | Single-threaded |
| **Persistence** | None (RAM-only) | Possible (RDB, AOF) |
| **Clustering** | Client-side | Server-side (Redis Cluster) |
| **Complexity** | Very simple | Feature-rich "Swiss Army Knife" |

---

## 💡 Summary

In the modern landscape, Redis is almost always the preferred choice because it can do everything Memcached does, plus a vast array of additional tasks. However, if you have a massive load consisting of simple session lookups that needs to be spread across many CPU cores, Memcached remains a proven and viable solution.