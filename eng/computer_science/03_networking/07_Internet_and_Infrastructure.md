# 🌐 The Internet and Its Infrastructure


## 📑 Table of Contents
1. ["Network of Networks" Architecture](#internet-architecture)
2. [ISP Hierarchy (Tiers)](#internet-providers-and-exchange-points)
3. [BGP — The Glue of the Internet](#bgp-border-gateway-protocol)
4. [CDN (Content Delivery Network)](#cdn-content-delivery-networks)

---

The internet is not a monolith, but a "federation" of over 100,000 independent networks called **Autonomous Systems (AS)**. Each AS is an island (e.g., Google, Verizon, a University) that decides how to route traffic internally but negotiates with neighbors on how to pass data further.

---


## 1. 🏗️ Hierarchy and Economics: Who Pays Whom?

Connections on the internet are about both technology and money.

- **Transit**: A smaller network pays a larger one for access to the entire internet.
- **Peering**: Two networks (e.g., Netflix and an ISP) connect directly to exchange traffic for free. This benefits everyone: faster for users, cheaper for companies.
- **Tier 1 (The Elite)**: Global networks (Lumen, Telia, NTT) so massive that they don't pay anyone for transit. They "peer" with each other for free. Everyone else pays them.

---


## 2. 🗺️ BGP — The Navigator and Diplomat

BGP (Border Gateway Protocol) is the protocol through which Autonomous Systems exchange information: "You can reach these IP addresses through me."


### How BGP Chooses a Path
Unlike your home router, BGP doesn't choose a path based solely on speed, but on **policy and cost**:
1. First, use free **Peering**.
2. If no peering is available, use cheap **Transit**.
3. As a last resort, use expensive **Transit**.

> [!IMPORTANT]
> **BGP Convergence**: When a subsea cable breaks, the entire internet needs time (seconds to minutes) to "agree" on new routes. During this time, packets may "lag" or be lost.

---


## 3. 🌊 Physics: Cables on the Ocean Floor

99% of intercontinental traffic travels not via satellites (which are slow and expensive), but through **submarine fiber-optic cables**. These are thin strands on the ocean floor, protected by layers of steel and plastic. If a ship's anchor snags one, an entire country's internet could go down.

---


## 4. 🚀 CDN and Edge Computing

Why send a request to the USA when you can get an answer from a server in your city?
- **Anycast**: A unique L3/BGP magic where the same IP address is advertised from multiple locations worldwide. Your computer is automatically "pulled" to the physically closest server.
- **WAF (Web Application Firewall)**: CDNs (like Cloudflare) often check your requests for malicious activity before they ever reach your origin server.

---


## 🎯 Key Takeaways

- The internet is a **network of networks** connected by BGP.
- **Peering** makes the internet cheaper and faster.
- An **Autonomous System (AS)** is the building block of the global web.
- If BGP "breaks," parts of the world can become invisible to the rest of the internet.

<!-- QUIZ_START 
[
    {
        "question": "What is the name of the protocol often referred to as the 'GPS of the internet' because it is responsible for choosing the optimal route between Autonomous Systems (AS)?",
        "options": ["HTTP", "DNS", "BGP", "NAT"],
        "correctIndex": 2
    },
    {
        "question": "What is the primary task of Content Delivery Networks (CDN)?",
        "options": ["Generating random IP addresses", "Storing copies of content on servers physically close to the user", "Replacing the HTTP protocol with a faster one", "Encrypting all provider traffic"],
        "correctIndex": 1
    },
    {
        "question": "To which layer (Tier) of the hierarchy do global internet providers belong that own transoceanic cables and do not pay others for traffic?",
        "options": ["Tier 3", "Tier 2", "Tier 1", "Tier 0"],
        "correctIndex": 2
    }
]
QUIZ_END -->
