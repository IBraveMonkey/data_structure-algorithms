# Basic Design Principles (KISS, DRY, YAGNI, etc.)

If SOLID is the foundation, then these principles represent the common sense that every developer should use daily.

---

## 💋 1. KISS (Keep It Simple, Stupid)

This is the most important principle. Code should be as simple as possible. If a task can be solved without a complex hierarchy of classes and patterns—solve it simply.

**Why is this important?**
- Simple code is easier to read.
- Simple code has fewer bugs.
- Simple code is easier for someone else to maintain.

---

## 🌵 2. DRY (Don't Repeat Yourself)

Every piece of knowledge (logic) should have a single, unambiguous representation within the system. If you find yourself copy-pasting code in three different places, create a function or a method.

**⚠️ Caution (Over-engineering):**
Sometimes "similar" code is not exactly the same. If you combine two different business processes into one function just because they both perform `a + b`, you create a rigid dependency. 
> *Sometimes copying code once is better than creating the wrong abstraction.*

---

## 🚫 3. YAGNI (You Ain't Gonna Need It)

Don't write code for features that "might be needed in the future." Implement only what is needed **now**.

**Example:**
- "Let's add support for 10 different currencies just in case we expand to the China market." ❌
- "We need to process payments in USD. Let's focus on USD only." ✅

---

## 😲 4. POLA (Principle of Least Astonishment)

Your code should behave predictably. If a method is named `GetUserDetails()`, it should not suddenly delete a user or change their password.

**Tips:**
- Use clear and descriptive names.
- Avoid hidden side effects.

---

## 🛡️ 5. Composition Over Inheritance

Instead of building long inheritance chains (e.g., `Animal -> Mammal -> Dog -> Bulldog`), it’s better to assemble objects from smaller parts (components).

**Analogy:**
- **Inheritance**: "A Bulldog *is a* Dog." (Rigid relationship).
- **Composition**: "A Robot Vacuum *has a* laser sensor and a motor." (Flexible assembly).

---

## ⚖️ How to maintain balance?

Sometimes these principles contradict each other. For example, **DRY** can make code more complex, potentially violating **KISS**. 
In such cases, prioritize as follows:
1. **KISS** (Simplicity first).
2. **YAGNI** (Don't do extra work).
3. Everything else.
