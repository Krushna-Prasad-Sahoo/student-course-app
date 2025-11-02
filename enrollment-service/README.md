
# 🎯 Enrollment Service — Go + MySQL Microservice

![Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Framework-Gin-lightgrey?logo=gin)
![MySQL](https://img.shields.io/badge/Database-MySQL-4479A1?logo=mysql)
![Docker](https://img.shields.io/badge/Containerized-Docker-blue?logo=docker)
![Kubernetes](https://img.shields.io/badge/Orchestrated%20with-Kubernetes-326ce5?logo=kubernetes)
![License](https://img.shields.io/badge/License-MIT-yellow)

---

## 🧩 Overview

The **Enrollment Service** acts as the **core integration point** within the *Student-Course-App* ecosystem.
Developed in **Go** using the **Gin web framework** and backed by **MySQL**, this service manages the relationship between students and courses — ensuring that only valid student and course records are linked.

It embodies **microservice autonomy**, **high performance**, and **reliable data validation** by communicating directly with the `student-service` and `course-service` APIs.

---

## 🧠 Workflow Description

1. A user or external client requests enrollment of a student in a course.
2. The service fetches student details from the Student Service (MongoDB backend).
3. It fetches course details from the Course Service (PostgreSQL backend).
4. Enrollment data (student ID and course ID) is stored in the MySQL database.
5. The system exposes REST endpoints to retrieve enrollment details per student.

This coordination ensures all services remain independent yet interoperable.

---

## ⚙️ Technology Stack

| Component            | Technology Used       | Purpose                                    |
| -------------------- | --------------------- | ------------------------------------------ |
| **Language**         | Go (1.21+)            | High-performance, concurrent service logic |
| **Framework**        | Gin                   | RESTful API routing and handling           |
| **Database**         | MySQL                 | Persistent storage for enrollment records  |
| **Containerization** | Docker                | Consistent build and runtime environment   |
| **Orchestration**    | Kubernetes (Minikube) | Deployment, service discovery, and scaling |

---

## 🧭 Architecture Summary Diagram

          ┌──────────────────┐         ┌──────────────────┐
          │  Student Service │         │  Course Service  │
          │ (Node.js + Mongo)│         │(Python + Postgre)│
          └─────────┬────────┘         └─────────┬────────┘
                    │                            │
                    │           Fetch            |
                    |       Student & Course     │
                    │        Information         │
                    ▼                            ▼
              ┌─────────────────────────────────────────┐
              │           Enrollment Service            │
              │           (Go + MySQL + Gin)            │
              └─────────────────────────────────────────┘
                                 │
                                 ▼
                        ┌─────────────────┐
                        │  MySQL Database │
                        │  (Enrollments)  │
                        └─────────────────┘


**Explanation:**

* The **Enrollment Service** validates both student and course information before creating a record.
* It acts as a **bridge** between the other two microservices, maintaining the logical integrity of enrollments.

---

## 🗂️ Directory Structure

```
enrollment-service/
├── main.go                  # Main Go application entry point
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksum file
├── Dockerfile               # Docker build configuration
├── enrollment-service.yaml  # Kubernetes deployment manifest
├── mysql.yaml               # MySQL deployment and PVC configuration
└── README.md                # Service documentation
```

---

## 🧩 Database Schema

The **MySQL** database maintains clean, referentially validated enrollment records:

| Field         | Type                              | Description                                   |
| ------------- | --------------------------------- | --------------------------------------------- |
| `id`          | INT (Primary Key, Auto Increment) | Unique enrollment record ID                   |
| `student_id`  | INT                               | References a student from the student service |
| `course_id`   | INT                               | References a course from the course service   |
| `enrolled_at` | DATETIME                          | Timestamp when enrollment occurred            |

> Each enrollment record represents a successful, validated relationship between a student and a course.

---

## ☁️ Kubernetes Deployment

### 🧩 `enrollment-service.yaml`

Defines the deployment, service, and autoscaling configuration for the Enrollment Service.
Includes:

* **Deployment** with replicas and rolling updates.
* **Service (ClusterIP)** for internal communication.
* **Readiness & Liveness Probes** for health monitoring.
* **HPA (Horizontal Pod Autoscaler)** for dynamic scaling based on CPU usage.

### 🗄️ `mysql.yaml`

Sets up:

* MySQL Deployment + Service
* Persistent Volume & Claim (PVC)
* Configurable environment variables for credentials and initialization scripts

---

✅ **In essence:**
The **Enrollment Service** is the *heart of coordination* in the Student-Course-App ecosystem — ensuring that every enrollment is legitimate, every linkage is validated, and every operation scales seamlessly in a modern, cloud-native way.

---

## 📜 License

This project is licensed under the **MIT License** — you’re free to use, modify, and distribute it with proper attribution.

---

## ❤️ Contributors

| Name | Role |
|------|------|
| [Aarya Nanndaann Singh M N](https://github.com/Aarya5122) | Lead Developer |
| [Devesh Singh](https://github.com/2024mt03089-devesh) | Contributor |
| [KP Sahoo](https://github.com/Krushna-Prasad-Sahoo/) | Contributor & Maintainer |

---



> _“Microservices are not just about dividing code — they’re about dividing responsibilities.”_
