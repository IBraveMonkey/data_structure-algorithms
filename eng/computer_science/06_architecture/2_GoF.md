# 📚 GoF (Gang of Four) — Design Patterns

## 📑 Table of Contents
1. [What is it?](#what-is-it)
2. [🏗️ Creational](#creational)
3. [🔌 Structural](#structural)
4. [🧠 Behavioral](#behavioral)
5. [🎯 Summary: Which to choose?](#summary-which-to-choose)

---

## 1. 🤔 What is it?

**Gang of Four (GoF)** refers to the four authors of the legendary book "Design Patterns" (1994). They described 23 classic solutions to common problems in OOP.

> [!IMPORTANT]
> A pattern is not specific code you can copy, but a **concept** for solving a problem. It’s like a recipe: the steps are the same, but the ingredients (programming language) may differ.

---

## 2. 🏗️ Creational

These patterns deal with **object creation** mechanisms. They help make a system independent of how its objects are created, composed, and represented.

### 👑 Singleton

Guarantees that a class has **only one** instance and provides a global point of access to it.

*   **🛠️ Problem**: You need exactly one instance of something in the system (e.g., a config, a logger, or a database connection pool) and want it to be accessible from anywhere.
*   **💡 Solution**: Hide the constructor and create a static method that always returns the same object.
*   **📦 Example**: 
    ```go
    // In Go, Singleton is often implemented using sync.Once
    var instance *Database
    var once sync.Once

    func GetInstance() *Database {
        once.Do(func() {
            instance = &Database{} // Will be created only the FIRST time
        })
        return instance
    }
    ```
*   **⚖️ Pros**: Control over a shared resource, memory savings.
*   **⚠️ Cons**: Global state (hard to test), violates the Single Responsibility Principle.

---

### FACTORY METHOD 🏭

Defines an interface for creating an object but lets subclasses decide which class to instantiate.

*   **🛠️ Problem**: You don't know beforehand which types of objects you'll need in the future, or you want to allow the library to be easily extended.
*   **💡 Solution**: Instead of calling `new` directly, you call a special "factory" method.
*   **📦 Example**: A "Logistics" application. 
    - Base factory: `Logistics` with a `CreateTransport()` method.
    - Subclass `RoadLogistics` creates a `Truck`.
    - Subclass `SeaLogistics` creates a `Ship`.
*   **✅ When to use**: When your code should work with various product types, but you want to hide the details of their creation.

---

### 🏗️ Abstract Factory

Allows you to create **families** of related objects without being tied to specific classes of those objects.

*   **🛠️ Problem**: You have sets of interfaces (e.g., "Furniture": Chair + Table). These sets come in different styles: "Victorian", "Modern". You need to ensure all objects in an order belong to the same style.
*   **💡 Solution**: Create an abstract furniture factory that knows how to create a Chair and a Table. Specific implementations (ModernFactory, VictorianFactory) create items in their respective styles.
*   **📦 Example**: A UI library. One factory creates buttons and checkboxes for macOS, another for Windows.

---

### 👷 Builder

Separates the construction of a complex object from its representation, so that the same process can create different representations.

*   **🛠️ Problem**: A massive constructor with 15 parameters (many optional) — an anti-pattern known as a "Telescoping Constructor."
*   **💡 Solution**: Move object creation to a separate builder class that assembles the object part by part.
*   **📦 Example**: Building a burger.
    `NewBurgerBuilder().AddCheese().AddBacon().NoOnions().Build()`
*   **✅ When to use**: When object creation involves many steps or when an object needs to exist in multiple configurations.

---

## 3. 🔌 Structural

These patterns explain how to assemble objects and classes into larger structures while keeping them flexible and efficient.

### 🔌 Adapter

Allows objects with incompatible interfaces to work together.

*   **🛠️ Problem**: You have an old class that outputs data in XML, and a new service that only understands JSON.
*   **💡 Solution**: Create an intermediate object (Adapter) that wraps one interface and makes it understandable to the other.
*   **📦 Example**: A `MicroUSB -> USB-C` adapter or a currency converter.

---

### 🎭 Decorator

Dynamically adds new responsibilities to an object without changing its original code.

*   **🛠️ Problem**: You need to extend a class's functionality, but inheritance isn't an option (or would lead to a "combinatorial explosion" of subclasses). For example, notifications: via Telegram, Email, or both.
*   **💡 Solution**: Place the object inside a "wrapper" that implements the same interface and adds its own behavior before or after calling the main method.
*   **📦 Example**: A logging system. You can "wrap" a basic logger with an `EncryptingLogger` (to encrypt logs) or a `FileLogger` (to write to a file).
*   **✅ When to use**: When you cannot use inheritance to extend behavior.

---

### 🏛️ Facade

Provides a simple (high-level) interface to a complex system of classes, a library, or a framework.

*   **🛠️ Problem**: Your system consists of 20 subsystems, and a client needs to call 15 methods in a specific order just to upload a video.
*   **💡 Solution**: Create a "Facade" — a single class with an `UploadVideo()` method that handles all the internal subsystems.
*   **📦 Example**: A "1-click order" button in an online store. It hides inventory checks, payment, and logistics.

---

### 🛡️ Proxy

Acts as a substitute for the real object, allowing something to be performed before or after a request reaches it.

*   **🛠️ Problem**: The real object is "heavy" (uses a lot of memory) or requires access control (permissions).
*   **💡 Solution**: The client communicates with the Proxy. The Proxy can: cache results, check permissions, or load the real object only when it is actually needed (**Lazy Loading**).
*   **📦 Example**: A caching proxy for heavy database queries.

---

### 🌲 Composite

Allows you to group objects into a tree-like structure and work with them as if they were a single object.

*   **🛠️ Problem**: You have a "Box" that can contain "Items" or other "Boxes." You need to calculate the total price.
*   **💡 Solution**: Both the item and the box implement the same `GetPrice()` interface. The box simply sums the prices of its contents.
*   **📦 Example**: A file system (Files and Folders). A Folder is a composite; a File is a leaf.

---

## 4. 🧠 Behavioral

These patterns deal with algorithms and the **assignment of responsibilities** between objects.

### 📡 Observer

Defines a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.

*   **🛠️ Problem**: You have a "Button" object and want 10 different actions to occur in different parts of the program when it's clicked, but you don't want to "hardwire" these actions into the Button itself.
*   **💡 Solution**: The object (Subject) maintains a list of subscribers. When an event occurs, it iterates through the list and calls a method on each subscriber.
*   **📦 Example**: YouTube. You (Observer) subscribe to a channel (Subject). When a new video is posted, all subscribers receive a notification.
*   **✅ When to use**: When changing one object requires changing others, and you don't know beforehand how many there are.

---

### 🗺️ Strategy

Defines a family of similar algorithms and makes them interchangeable at runtime.

*   **🛠️ Problem**: You have a massive `switch` or a bunch of `if-else` blocks to select a payment method (Visa, PayPal, Crypto). The code becomes unreadable.
*   **💡 Solution**: Extract each algorithm into a separate class. The main class simply stores a reference to the current "strategy" and calls it.
*   **📦 Example**: A navigator. Today I'm driving (one strategy), tomorrow I’m walking (another). The goal is the same—to arrive—but the methods differ.

---

### 🗣️ Command

Turns a request into a stand-alone object. This allows you to pass requests as parameters, queue them, or log them.

*   **🛠️ Problem**: You want to implement an "Undo" button or a task queue.
*   **💡 Solution**: Instead of calling a method directly, you create a `Command` object that knows which method to call. The command can store state for an "Undo" operation.
*   **📦 Example**: A remote control button. Each button is a "Command" object (`VolumeUp`, `Mute`). A waiter in a restaurant writes an order on a slip (this is a Command) and takes it to the kitchen.

---

### 🧱 Template Method

Defines the skeleton of an algorithm in a base class but lets subclasses override specific steps of the algorithm without changing its structure.

*   **🛠️ Problem**: You have two similar processes (e.g., parsing CSV and JSON). Both: 1. Open the file, 2. Read the data (differently), 3. Close the file.
*   **💡 Solution**: Create a method in the base class where steps 1 and 3 are implemented, while step 2 is abstract (to be implemented by children).
*   **📦 Example**: Making a drink: boil water -> brew the base (tea/coffee) -> pour into a cup -> add extras (lemon/sugar). The skeleton is fixed; the details vary.

---

### 🚦 State

Allows an object to alter its behavior when its internal state changes. The object will appear to change its class.

*   **🛠️ Problem**: A massive `switch` inside every method that depends on a status (`if status == "Closed" { ... }`).
*   **💡 Solution**: Create classes for each state (`OpenState`, `ClosedState`). The object simply delegates the work to the current state.
*   **📦 Example**: An ATM. Its behavior (dispense cash, issue receipt) depends entirely on whether a card is inserted, a PIN is entered, and if there is enough money.

---

### ⛓️ Chain of Responsibility

Allows passing requests along a chain of handlers.

*   **🛠️ Problem**: Request processing consists of several stages (authorization, logging, validation). If any stage fails, we don't proceed.
*   **💡 Solution**: Each handler contains a reference to the next. If it cannot handle the request or needs to pass it on, it calls the next handler.
*   **📦 Example**: Support service. First, a robot speaks -> then a junior specialist -> then a senior manager. If the robot solves the problem, the chain stops.

---

## 🎯 Summary: Which to choose?

1.  Need to create a complex object? — **Creational** (Builder, Factory).
2.  Need to connect incompatible interfaces or simplify structure? — **Structural** (Adapter, Facade).
3.  Need to coordinate communication between objects or switch algorithms on the fly? — **Behavioral** (Observer, Strategy).