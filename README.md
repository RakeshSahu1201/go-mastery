## Module 1: System Design & Architecture
### RESTful API Design

**1. Base Setup (In-Memory REST API)**
- **Concept:** Built a foundational HTTP server using Go's standard `net/http` library.
- **Key Takeaway:** Utilized Go 1.22+ method-based routing (`GET /users`) to handle requests without external frameworks, serving JSON from an in-memory struct slice.