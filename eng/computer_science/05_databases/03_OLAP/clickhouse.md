# 📊 ClickHouse


## 📑 Table of Contents
1. [What is it? (OLAP vs. OLTP)](#what-is-it-olap-vs-oltp)
2. [Rows vs. Columns](#rows-vs-columns)
3. [MergeTree (The Heart of ClickHouse)](#mergetree-the-heart-of-clickhouse)

---


## 1. 🤔 What is it? (OLAP vs. OLTP)

**ClickHouse** is a column-oriented DBMS designed for **OLAP** (Online Analytical Processing).


### OLTP (Postgres, MySQL)
*   **Goal**: To quickly insert an order or update a balance. Optimized for operations on **individual rows**.
*   **Sample Query**: `SELECT * FROM users WHERE id = 1`


### OLAP (ClickHouse)
*   **Goal**: High-performance analytics. For example, "What was our total revenue across all regions last year?" Optimized for processing **billions of rows** across a handful of columns.
*   **Sample Query**: `SELECT sum(price) FROM orders WHERE date > '2023-01-01'`

> [!NOTE]
> ClickHouse is incredibly fast. A query that might take 10 minutes to execute in Postgres can often be completed in ClickHouse in under 0.1 seconds.

---


## 2. 🏛️ Rows vs. Columns

Consider a table structure like this: `| Date | ID | Price |`


### Row-Oriented DBMS (e.g., MySQL)
Stores data sequentially, row by row:
`[2023-01-01, 1, 100], [2023-01-01, 2, 200], ...`
To calculate the total of the `Price` column, the disk must read **EVERY** part of the row (including the `Date` and `ID` fields, which are irrelevant to the query).


### Column-Oriented DBMS (e.g., ClickHouse)
Stores each column in its own separate file:
*   **Date File**: `[2023-01-01, 2023-01-01, ...]`
*   **ID File**: `[1, 2, ...]`
*   **Price File**: `[100, 200, ...]`

To calculate the total of the `Price` column, ClickHouse reads **only the Price file**.
*   **I/O Efficiency**: The system reads significantly less data (often 100x less).
*   **Compression**: Storing data of the same type (like numbers) together allows for much more effective compression.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    subgraph Row_Header ["Row-Oriented"]
        R1[Data from all columns is interleaved]
    end
    
    subgraph Col_Header ["Column-Oriented"]
        C1[File: Prices]
        C2[File: Dates]
        C3[File: IDs]
    end
    
    Query[Query: SUM Price] --> R1
    Query --> C1
    
    style C1 fill:#bbf,stroke:#333,stroke-width:2px
    style R1 fill:#ffc,stroke:#333



linkStyle default stroke:#009688,stroke-width:2px;




```

---


## 3. 🌲 MergeTree (The Heart of ClickHouse)

The most robust and widely used table engine in ClickHouse is **MergeTree**.
It functions similarly to an LSM-Tree (Log-Structured Merge Tree).

1.  **Insertion**: Data is written in small, sorted "parts."
2.  **Background Processing**: ClickHouse continuously merges these small parts into larger, sorted files in the background.

> [!WARNING]
> ClickHouse is **not** designed for frequent, single-row insertions (`INSERT INTO ... VALUES (1)`).
> You must write data in batches of 1,000 to 10,000 rows. Otherwise, the server will be overwhelmed by the background merge operations.

---


## 💡 Summary

If you need to generate complex reports, build dashboards, or calculate metrics across terabytes of log data, **ClickHouse** is your best tool. Avoid using it as a primary transactional database for your application backend.

<!-- QUIZ_START 
[
    {
        "question": "What is the primary conceptual difference between ClickHouse and PostgreSQL?",
        "options": ["ClickHouse is a NoSQL database", "ClickHouse uses columnar storage, while PostgreSQL is row-oriented", "ClickHouse does not support SQL", "ClickHouse only works in the cloud"],
        "correctIndex": 1
    },
    {
        "question": "Why is it recommended to write data to ClickHouse in large batches?",
        "options": ["To save disk space", "To avoid creating too many small 'parts' and overwhelming the background merge process", "It is a requirement of the HTTP protocol", "To allow enough time for data encryption"],
        "correctIndex": 1
    },
    {
        "question": "Which ClickHouse table engine is the primary and most powerful for general use?",
        "options": ["Log", "Memory", "MergeTree", "Kafka"],
        "correctIndex": 2
    }
]
QUIZ_END -->