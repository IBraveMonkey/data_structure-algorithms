/*
Hash Table

What is it?
A Hash Table is a data structure that implements the "associative array" abstract data type, which maps keys to values. It uses a hash function to calculate an index into an array of buckets or slots, from which the desired value can be found. This allows for efficiency in carrying out insertion, search, and deletion of elements.

Why is it needed?
- Fast data access by key (average O(1) for operations).
- Efficient storage of key-value pairs.
- Grouping and matching of data.
- Counting the frequency of elements.

What's the point?
- Converting a key into an array index using a hash function.
- Storing values in an array at the calculated indices.
- Handling collisions (when different keys result in the same index).
- Fast access to data by key.

When to use?
- When fast access to data by key is required.
- For counting the frequency of elements.
- For checking if an element exists in a set.
- For grouping related data.

How does it work?
- The key is converted into an index using a hash function.
- The value is stored in an array at that index.
- During a search, the key is hashed again, and the value is retrieved by the index.
- Collisions are handled using various methods (chaining, open addressing).

### Complexity

| Operation | Average (O) | Worst (O) | Space Complexity (O) |
|:---|:---:|:---:|:---:|
| Insertion | O(1) | O(n) | O(1) |
| Search | O(1) | O(n) | O(1) |
| Deletion | O(1) | O(n) | O(1) |
| Storage | — | — | O(n) |

*The worst case O(n) occurs when there are many collisions, such as when all keys map to the same bucket.

How to know if a problem fits Hash Table?
- You need to find elements by key quickly.
- You need to count the frequency of elements.
- You need to check if an element has been encountered before.
- You need to map associated data.

Hash Table in Go (map):

In Go, hash tables are implemented as a built-in data structure — the map.
- map[KeyType]ValueType — hash table declaration.
- Keys must be comparable types (int, string, struct, etc.).
- Values can be of any type.
- Go automatically manages sizing and collisions.
- No guarantee of element order during iteration.
- Not thread-safe (requires synchronization for concurrent access).

Examples of using Hash Table:
See the example.go file.
*/

package hash_table
