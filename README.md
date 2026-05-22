# Interview Preparation Guide

> Hello there!
>
> We want to set you up for success in your upcoming technical interview. To do that, we've prepared this short guide so you know exactly what to expect and how to prepare.
>
> Please take a few minutes to review it — there are no surprises, and the goal is for you to feel comfortable and confident showcasing your skills.
>
> **Format:** 1 Live Technical Interview (Single Session)
>
> **Duration:** ~60–90 minutes
>
> **Primary Focus:** Backend engineering (Go + APIs + system design)
>
> *The interview is designed to reflect real, day-to-day backend work: building services, handling data, designing APIs, and reasoning about system behavior and scalability.*
>

---

## ⚙️ What to Expect

You'll participate in **one live technical interview**, combining:

- Short coding exercises
- A practical backend implementation
- A system design discussion

**Format:** Live coding + discussion (single session)

**Duration:** ~60–90 minutes

**Style:** Collaborative, practical, no trick questions

---

## 🎯 Focus Areas

- Backend development with **Go**
- API design (REST or GraphQL concepts)
- Data handling and basic validation
- Understanding of **databases (PostgreSQL basics)**
- Awareness of **event-driven systems (Kafka or similar)**
- Clear reasoning and communication

---

## 💻 Environment Setup (Important)

You'll be coding live, so please come prepared with:

- Your preferred **IDE or code editor**
- A working **Go environment** (able to run a simple HTTP server)
- Ability to run code locally (no setup during the interview)

👉 We recommend testing this beforehand:

- Run a simple Go file (`go run main.go`)
- Create a basic HTTP handler
- Print/log output locally

---

## 💻 Live Coding Overview

You'll work on a small backend exercise:

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


