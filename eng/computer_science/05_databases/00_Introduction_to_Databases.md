# 🗄️ Introduction to Databases

## 📑 Table of Contents
1. [What are Databases?](#-what-are-databases)
2. [Brief History of Development](#-brief-history-of-development)
3. [Main Types of Databases](#-main-types-of-databases)
4. [SQL vs NoSQL: Key Differences](#-sql-vs-nosql-key-differences)
5. [Where to Continue Learning?](#-where-to-continue-learning)

---

## 1. 🤔 What are Databases?

A **Database (DB)** is an organized collection of data that is stored and accessed electronically. It is a systematic storage of information that allows for efficient addition, modification, deletion, and retrieval of data.

A **DBMS (Database Management System)** is software that manages databases. It provides an interface between the database and end users or applications, ensuring data security, integrity, and efficiency.

> [!NOTE]
> **Analogy**:
> *   **DB** is like an archive of documents.
> *   **DBMS** is like a librarian who knows where everything is and can bring you the needed document upon request.
> *   **Query Language** is the way of communicating with the librarian (e.g., SQL).

---

## 2. 📜 Brief History of Development

1. **File Systems (1950s - 1960s)**: Simple data storage in text files.
   *Issues*: Difficult searching, data duplication, lack of relationships.

2. **Hierarchical and Network Models (1960s - 1970s)**: Data organized as trees or graphs.
   *Issues*: Difficulty changing structure, dependence on physical storage.

3. **Relational Databases (1970s, E.F. Codd)**: Data represented in tables with relationships.
   *Advantages*: De facto standard for 40+ years, strict structure, data integrity, SQL language.

4. **NoSQL (2000s)**: Response to Big Data and flexibility needs.
   *Advantages*: Scalability, flexible schema, high performance.

5. **NewSQL (2010s)**: Attempt to combine RDBMS and NoSQL benefits.
   *Goal*: Provide ACID guarantees and horizontal scalability.

---

## 3. 🏗️ Main Types of Databases

### Relational Databases (RDBMS)
- Data stored in tables with fixed schema
- Use SQL for queries
- Examples: PostgreSQL, MySQL, Oracle, SQLite

### NoSQL Databases
- Flexible or absent data schema
- Suitable for large volumes of data and high scalability
- Types: Document-oriented, Key-value, Columnar, Graph
- Examples: MongoDB, Redis, Cassandra, Neo4j

### OLAP (Online Analytical Processing)
- Optimized for analytical queries and reports
- Support complex queries on historical data
- Examples: Amazon Redshift, Google BigQuery, Apache Kylin

---

## 4. 🗣️ SQL vs NoSQL: Key Differences

### SQL (Relational Databases)
- **Strict Schema**: Tables have fixed structure
- **ACID Guarantees**: Atomicity, Consistency, Isolation, Durability
- **Query Language**: SQL (Structured Query Language)
- **Scaling**: Primarily vertical (increasing server power)

### NoSQL (Non-relational Databases)
- **Flexible Schema**: Structure can change dynamically
- **CAP Theorem**: Choice between Consistency, Availability, and Partition tolerance
- **Variety of Models**: Documents, Key-value, Graphs, Columns
- **Scaling**: Horizontal (adding more servers)

> [!TIP]
> **How to Choose?**
> *   For applications with well-defined data structure requiring integrity and transactions (e.g., banking systems) — choose SQL.
> *   For applications with rapidly changing requirements, large data volumes, and high scalability (e.g., social networks) — consider NoSQL.

## 🎯 Where to Continue Learning?

Now that you've familiarized yourself with the basics of databases, you can dive deeper into specific types:

- [Relational Databases (RDBMS)](./01_RDBMS/0.ACID.md) — Detailed study of relational databases, normalization, ACID properties, and SQL.
- [NoSQL Databases](./02_NoSQL/) — Exploration of various NoSQL models: document-oriented, key-value, columnar, and graph databases.
- [OLAP (Online Analytical Processing)](./03_OLAP/) — Analytical databases, multidimensional cubes, aggregations, and reporting.

<!-- QUIZ_START
[
    {
        "question": "Which of the following best describes the difference between a DB and a DBMS?",
        "options": ["DB is software and DBMS is data", "DB is data storage and DBMS is the program managing this data", "There is no difference", "DBMS is used only for NoSQL"],
        "correctIndex": 1
    },
    {
        "question": "Which type of database is best suited for applications requiring strict data integrity and transactions?",
        "options": ["NoSQL", "SQL", "OLAP", "Key-value"],
        "correctIndex": 1
    },
    {
        "question": "What does the acronym ACID mean in the context of databases?",
        "options": ["Atomicity, Consistency, Isolation, Durability", "Availability, Consistency, Integrity, Distribution", "Aggregation, Clustering, Indexing, Distribution", "Authentication, Confidentiality, Integrity, Detection"],
        "correctIndex": 0
    }
]
QUIZ_END -->