

# 🎓 Student-Course-App — Microservices Ecosystem

![Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)
![Python](https://img.shields.io/badge/Python-3.11-blue?logo=python)
![Node.js](https://img.shields.io/badge/Language-Node.js-43853d?logo=node.js)
![Docker](https://img.shields.io/badge/Containerized-Docker-blue?logo=docker)
![Kubernetes](https://img.shields.io/badge/Orchestrated-Kubernetes-326ce5?logo=kubernetes)
![License](https://img.shields.io/badge/License-MIT-yellow)

---

## 🧩 Overview

**Student-Course-App** is a cloud-native microservices ecosystem built to manage students, courses, and enrollments.
It consists of three independent services:

1. **Student Service** — manages student data (Node.js + MongoDB)
2. **Course Service** — manages courses (Python + Flask + PostgreSQL)
3. **Enrollment Service** — manages student-course enrollments (Go + Gin + MySQL)

Each service exposes REST APIs, is containerized with Docker, and can be orchestrated via Kubernetes. Together, they ensure **modularity, scalability, and reliable coordination** between students and courses.

---

## 🏗️ Architecture Diagram

#### Simple Application Architeture Diagram:

```mermaid
flowchart LR
    A[External Client / User / Postman]

    subgraph Student_Service["Student Service"]
        S_DB[(MongoDB)]
    end

    subgraph Course_Service["Course Service"]
        C_DB[(PostgreSQL)]
    end

    subgraph Enrollment_Service["Enrollment Service"]
        E_DB[(MySQL)]
    end

    %% External access
    A --> Student_Service
    A --> Course_Service
    A --> Enrollment_Service

    %% Internal communication for validation
    Enrollment_Service --> Student_Service
    Enrollment_Service --> Course_Service

    %% Databases
    Student_Service --> S_DB
    Course_Service --> C_DB
    Enrollment_Service --> E_DB

```

#### Detailed Architeture Diagram showing all Workflows:

```mermaid
flowchart LR
    subgraph External Client
        A[User / Postman / API Client]
    end

    subgraph Student_Service["Student Service (Node.js + MongoDB)"]
        S1[/GET /students/]
        S2[/POST /students/]
        S3[/GET /students/:id/]
        DB1[(MongoDB)]
    end

    subgraph Course_Service["Course Service (Python + PostgreSQL)"]
        C1[/GET /courses/]
        C2[/POST /courses/]
        C3[/GET /courses/:id/]
        DB2[(PostgreSQL)]
    end

    subgraph Enrollment_Service["Enrollment Service (Go + MySQL)"]
        E1[/GET /enrollments/]
        E2[/POST /enroll/]
        E3[/GET /enrollments/student/:id/]
        DB3[(MySQL)]
    end

    %% External client can access all services
    A --> S1
    A --> S2
    A --> S3
    A --> C1
    A --> C2
    A --> C3
    A --> E1
    A --> E2
    A --> E3

    %% Enrollment service communicates with Student and Course Services for validation
    E2 --> S1
    E2 --> C1

    %% Databases
    S1 --> DB1
    S2 --> DB1
    S3 --> DB1

    C1 --> DB2
    C2 --> DB2
    C3 --> DB2

    E1 --> DB3
    E2 --> DB3
    E3 --> DB3

```

#### Detailed Architecture When Project Is Onboarded into Kubernetes:

```mermaid
flowchart TD
    %% Minikube Node
    subgraph Minikube_Node["Minikube Node"]
        direction TB
        
        %% Student Service Pods
        subgraph Student_Pods["Student Pods (with HPA)"]
            SS_Pod1["student-service-pod-1"]
            SS_Pod2["student-service-pod-2"]
        end
        SS_Service["NodePort: student-service"]

        %% Course Service Pods
        subgraph Course_Pods["Course Pods (with HPA)"]
            CS_Pod1["course-service-pod-1"]
            CS_Pod2["course-service-pod-2"]
        end
        CS_Service["NodePort: course-service"]

        %% Enrollment Service Pods
        subgraph Enrollment_Pods["Enrollment Pods (with HPA)"]
            ES_Pod1["enrollment-service-pod-1"]
            ES_Pod2["enrollment-service-pod-2"]
        end
        ES_Service["NodePort: enrollment-service"]

        %% Databases (internal only)
        MongoDB["MongoDB - PVC - ClusterIP Service (internal)"]
        PostgreSQL["PostgreSQL - PVC - ClusterIP Service (internal)"]
        MySQL["MySQL - PVC - ClusterIP Service (internal)"]

        %% Connections
        SS_Service --> SS_Pod1
        SS_Service --> SS_Pod2
        SS_Pod1 --> MongoDB
        SS_Pod2 --> MongoDB

        CS_Service --> CS_Pod1
        CS_Service --> CS_Pod2
        CS_Pod1 --> PostgreSQL
        CS_Pod2 --> PostgreSQL

        ES_Service --> ES_Pod1
        ES_Service --> ES_Pod2
        ES_Pod1 --> MySQL
        ES_Pod2 --> MySQL

        %% Internal service calls
        ES_Pod1 --> SS_Service
        ES_Pod1 --> CS_Service
        ES_Pod2 --> SS_Service
        ES_Pod2 --> CS_Service
    end

    %% External access via NodePort
    User[User / Postman / Client] --> SS_Service
    User --> CS_Service
    User --> ES_Service    
```

---

## ⚙️ Technology Stack

| Layer                  | Technology                  | Purpose                                          |
| ---------------------- | --------------------------- | ------------------------------------------------ |
| **Student Service**    | Node.js + MongoDB           | Stores and manages student records               |
| **Course Service**     | Python + Flask + PostgreSQL | Stores and manages courses                       |
| **Enrollment Service** | Go + Gin + MySQL            | Manages enrollments between students and courses |
| **Containerization**   | Docker                      | Environment encapsulation                        |
| **Orchestration**      | Kubernetes                  | Deployment, scaling, and service discovery       |

---

## 📂 Directory Structure

```
root/
├── docker-compose.yaml
├── student-service/
│   ├── app.js
│   ├── Dockerfile
│   ├── package.json
│   ├── mongo.yaml
│   └── student-service.yaml
├── course-service/
│   ├── main.py
│   ├── Dockerfile
│   ├── requirements.txt
│   ├── postgres.yaml
│   └── course-service.yaml
├── enrollment-service/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── mysql.yaml
│   └── enrollment-service.yaml
└── README.md
```

---

## 🚀 Setup Guide

### 1️⃣ Local Docker Compose

1. Ensure **Docker daemon is running**.
2. In the project root, run:

```bash
docker compose up --build
```

3. All services (MongoDB, PostgreSQL, MySQL + microservices) will start locally.

---

### 2️⃣ Minikube / Kubernetes Deployment

1. Point Docker CLI to Minikube:

```bash
eval $(minikube docker-env)
```

2. Build service images locally:

```bash
docker build -t student-service:latest ./student-service
docker build -t course-service:latest ./course-service
docker build -t enrollment-service:latest ./enrollment-service
```

3. Apply database manifests:

```bash
kubectl apply -f student-service/mongo.yaml
kubectl apply -f course-service/postgres.yaml
kubectl apply -f enrollment-service/mysql.yaml
```

4. Deploy services:

```bash
kubectl apply -f student-service/student-service.yaml
kubectl apply -f course-service/course-service.yaml
kubectl apply -f enrollment-service/enrollment-service.yaml
```

5. Verify pods and services:

```bash
kubectl get pods
kubectl get svc
```

6. Access services via `<minikube_ip>:<forwarded_port>` in Postman or curl.

---

### 3️⃣ Load Testing / HPA Verification

Simulate CPU load to observe **Horizontal Pod Autoscaler** behavior:

```bash
seq 10000 | xargs -n1 -P20 -I{} curl -s -o /dev/null http://localhost:5001/courses
```

---

## 🧪 API Testing

Replace `<minikube_ip>` and `<port>` with your actual IP/port.

---

### **Student Service**

**1. Health Check**

```bash
curl -X GET 'http://<minikube_ip>:<student_service_port>/'
```

**2. Add Students (POST /students)**

```bash
curl -X POST -H "Content-Type: application/json" -d '{ "name": "Krushna", "email": "krushna@kr.com"}' 'http://<minikube_ip>:<student_service_port>/students'
```

**3. Fetch All Students (GET /students)**

```bash
curl -X GET 'http://<minikube_ip>:<student_service_port>/students' | jq
```

**4. Fetch Single Student (GET /students/:id)**

```bash
curl -X GET 'http://<minikube_ip>:<student_service_port>/students/<student_id>'
```

---

### **Course Service**

**1. Health Check**

```bash
curl -X GET 'http://<minikube_ip>:<course_service_port>/'
```

**2. Add Courses (POST /courses)**

```bash
curl -X POST -H "Content-Type: application/json" -d '{ "title": "Physics", "instructor": "Bob", "credits": 4}' 'http://<minikube_ip>:<course_service_port>/courses'
```

**3. Fetch All Courses (GET /courses)**

```bash
curl -X GET 'http://<minikube_ip>:<course_service_port>/courses' | jq
```

**4. Fetch Single Course (GET /courses/:id)**

```bash
curl -X GET 'http://<minikube_ip>:<course_service_port>/courses/<course_id>'
```

---

### **Enrollment Service**

**1. Health Check**

```bash
curl -X GET 'http://<minikube_ip>:<enrollment_service_port>/'
```

**2. Enroll Student (POST /enroll)**

```bash
curl -X POST -H "Content-Type: application/json" -d '{ "student_id": 1, "course_id": 2}' 'http://<minikube_ip>:<enrollment_service_port>/enroll'
```

**3. Fetch All Enrollments (GET /enrollments)**

```bash
curl -X GET 'http://<minikube_ip>:<enrollment_service_port>/enrollments'
```

**4. Fetch Enrollments for a Student (GET /enrollments/student/:id)**

```bash
curl -X GET 'http://<minikube_ip>:<enrollment_service_port>/enrollments/student/<student_id>'
```

---

## ⚡ Scalability & Reliability

* **Horizontal Pod Autoscaler (HPA)** for CPU-based dynamic scaling
* **Readiness & Liveness Probes** for health monitoring
* **PersistentVolumeClaims (PVCs)** ensure data durability across pod restarts

---

## ❤️ Contributors

| Name                                                      | Role                                |
| --------------------------------------------------------- | ----------------------------------- |
| [Siddharth Jain](https://github.com/2024mt03040-web)      | Lead Developer (Course Service)     |
| [Tanuja RY](https://github.com/tanujary)                  | Lead Developer (Student Service)    |
| [Aarya Nanndaann Singh M N](https://github.com/Aarya5122) | Lead Developer (Enrollment Service) |
| [Devesh Singh](https://github.com/2024mt03089-devesh)     | Contributor (Docker & GitHub Action)|
| [KP Sahoo](https://github.com/Krushna-Prasad-Sahoo/)      | Contributor (K8s) & Maintainer      |

---

## ⚖️ License

This project is licensed under the **MIT License** — you may freely use, modify, and distribute it.

---

> *“Microservices are not just about dividing code — they’re about dividing responsibilities.”*

---
