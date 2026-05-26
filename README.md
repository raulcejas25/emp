
### "Drug Categories API"

This is a simple, realistic backend task similar to internal tools.

#### 🔹 Backend (Go)

You'll work with a provided Go skeleton file.

**What you'll implement:**

- A simple HTTP endpoint (`/categories`)
- In-memory data storage (no database)
- Logic to:
  - Return a list of categories (GET)
  - Create a new category (POST)
  - Basic input validation

**Data model:**

- `id` (int)
- `name` (string)
- `description` (string)

#### 🧠 Important Notes

You do **not** need:

- A real database
- Authentication
- Frameworks or external libraries

We are evaluating:

👉 Clarity  
👉 Correctness  
👉 Practical backend thinking

---

## 🏗 System Design Overview

You'll also discuss a backend system design scenario.

**Example context:**

- Managing **Drug Categories and rules**
- Propagating changes to other systems
- Ensuring reliability and consistency

#### 🔹 What You May Be Asked

- How to model data (PostgreSQL)
- How APIs interact with the system
- How changes propagate to other services
- How to design for **scalability and reliability**

👉 You may discuss concepts like:

- Events or messaging systems (Kafka)
- Data consistency
- Failure handling
- System evolution over time

---

## 🎯 What We're Looking For

### Backend Fundamentals

- Can you write a simple, working API?
- Can you handle JSON input/output?
- Can you structure code clearly?

### Engineering Thinking

- Do you keep solutions simple?
- Can you explain your decisions?
- Do you consider edge cases?

### System Design Awareness

- Do you understand how systems interact?
- Can you reason about scalability and reliability?
- Are you aware of asynchronous/event-driven systems?

### Communication

- Do you explain your thinking clearly?
- Do you ask questions when needed?

---

## 🧠 How to Prepare

To feel comfortable, we recommend:

- Reviewing basic **Go HTTP handlers** (`net/http`)
- Practicing simple **REST endpoints** (GET/POST)
- Understanding **JSON parsing in Go**
- Reviewing basic **PostgreSQL concepts** (tables, indexes)
- Familiarizing yourself with **event-driven systems (Kafka basics)**:
  - What is a producer/consumer?
  - Why use events?

👉 You do **NOT** need deep expertise — just solid fundamentals.

---

## 🚀 Final Notes

- Think aloud — we care about your reasoning.
- Ask clarifying questions if needed.
- Don't over-engineer.
- If you get stuck, explain what you're thinking — that's completely okay.


