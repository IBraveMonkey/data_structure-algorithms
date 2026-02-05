# SOLID: The Five Pillars of Quality Code 🏛️

SOLID is an acronym for five principles of object-oriented programming and design that help make code more understandable, flexible, and maintainable.

---

## 1. 🎯 S: Single Responsibility Principle (SRP)
> A class (or function) should have only one reason to change.

**In simple terms:** A single component should be responsible for only one task. If a class calculates salaries, saves them to the database, and sends emails, it becomes a "God Object" that is liable to break with any change.

#### ❌ Bad:
The `User` class is doing too much.
```go
type User struct {
    Name string
}

func (u *User) SaveToDB() { /* ... */ }
func (u *User) SendEmail() { /* ... */ } // A User should not know how to send emails!
```

#### ✅ Good:
Separate the logic.
```go
type User struct {
    Name string
}

type UserRepository struct {} // Handles database operations
func (r *UserRepository) Save(u *User) {}

type EmailService struct {} // Handles email operations
func (s *EmailService) Send(u *User) {}
```

---

## 2. 🔓 O: Open/Closed Principle (OCP)
> Software entities should be open for extension but closed for modification.

**In simple terms:** You should be able to add new features without changing existing (and already tested) code.

#### ❌ Bad:
To add a new type of notification, you must modify `SendNotification`.
```go
func SendNotification(msg string, tool string) {
    if tool == "email" { /* ... */ }
    if tool == "sms" { /* ... */ } // You’ll have to add an 'if' here every time
}
```

#### ✅ Good:
Use an interface.
```go
type Notifier interface {
    Notify(msg string)
}

// Now we can create a new struct without touching the main code
type TelegramNotifier struct {}
func (t TelegramNotifier) Notify(msg string) { /* ... */ }
```

---

## 3. 🔄 L: Liskov Substitution Principle (LSP)
> Objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program.

**In simple terms:** If you have a function that accepts a "Bird," it should also work with an "Ostrich," even if the ostrich doesn't fly. If the architecture breaks when you replace one subclass with another, the principle is violated.

#### ❌ Typical Error:
Rectangle and Square. A Square is mathematically a rectangle, but if a Square's `SetWidth` method also changes its height, it will break logic that expects the behavior of a standard Rectangle.

---

## 4. ✂️ I: Interface Segregation Principle (ISP)
> Clients should not be forced to depend on methods they do not use.

**In simple terms:** It is better to have many small, specialized interfaces than one huge "universal" interface.

#### ❌ Bad:
```go
type Worker interface {
    Work()
    Eat() // A robot doesn't need to eat! ❌
}
```

#### ✅ Good:
```go
type Workable interface { Work() }
type Eatable interface { Eat() }

type Robot struct {} // Implements only Workable
```

---

## 5. 💉 D: Dependency Inversion Principle (DIP)
> High-level modules should not depend on low-level modules. Both should depend on abstractions.

**In simple terms:** Do not tie yourself to specific implementations (e.g., Postgres, Redis). Instead, depend on interfaces. (We covered this in detail in the DI/IoC section).

---

> [!TIP]
> **SOLID** is not a strict law but a guideline. Sometimes, for the sake of simplicity (KISS), it's okay to slightly deviate from SOLID in small projects.
