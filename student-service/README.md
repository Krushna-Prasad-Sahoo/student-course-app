# 🎓 Student Service — NodeJS + Express + MongoDB Microsevice

![Node.js](https://img.shields.io/badge/Node.js-18.x-green?logo=node.js)
![Express.js](https://img.shields.io/badge/Express.js-4.x-lightgrey?logo=express)
![MongoDB](https://img.shields.io/badge/Database-MongoDB-green?logo=mongodb)
![Docker](https://img.shields.io/badge/Containerized-Docker-blue?logo=docker)
![Kubernetes](https://img.shields.io/badge/Orchestrated%20with-Kubernetes-326ce5?logo=kubernetes)
![License](https://img.shields.io/badge/License-MIT-yellow)

---

## 🧩 Overview

The **Student Service** is one of the core microservices in the **Student-Course-App** ecosystem.  It is responsible for handling **student-related data operations**, including creating and retrieving student profiles.  

Designed using **Node.js** and **Express**, it connects to a **MongoDB** database for persistent data storage.  The service is containerized with **Docker** and deployed to a **Kubernetes** cluster (via Minikube) for high availability and scalability.

---

## 🏗️ Architectural Summary

```

+------------------------+
|     Student Service    |  --->  Node.js + Express
|------------------------|
|   /students (POST)     |  --> Add new student
|   /students (GET)      |  --> Retrieve all the students details
|   /students/<id> (GET) |  --> Retrieve student details
|   /                    |  --> Health check endpoint
+-----------+------------+
|
v
+-------------------+
|     MongoDB       | ---> Persistent data store
+-------------------+

```

---

## 🧰 Technology Stack

| Component | Technology |
|------------|-------------|
| **Language** | Node.js (v18) |
| **Framework** | Express.js |
| **Database** | MongoDB |
| **Containerization** | Docker |
| **Orchestration** | Kubernetes (Minikube) |
| **Scaling** | K8s Horizontal Pod Autoscaler (HPA) |

---

## 📂 Directory Structure

```

student-service/
├── app.js                   # Core Express application
├── Dockerfile               # Docker build instructions
├── package.json             # Project metadata and dependencies
├── mongo.yaml               # MongoDB K8s deployment and service
├── student-service.yaml     # Student service deployment, service, and autoscaling
└── README.md                # Documentation

```

---

## ⚙️ Key Components

### 1️⃣ Application Logic (`app.js`)
- Uses **Express.js** to define HTTP routes.
- Connects to MongoDB through **Mongoose** ORM.
- Exposes endpoints to:
  - **Create student records**
  - **Retrieve all student records**
  - **Retrieve student records by ID**
  - **Health check route** for verification.
- Automatically connects to MongoDB using environment variable `MONGO_URI`.

---

### 2️⃣ Database Configuration (`mongo.yaml`)
- Deploys **MongoDB** as a Kubernetes **Deployment** with **PersistentVolumeClaim (PVC)**.
- Ensures data durability even when MongoDB pods restart.
- Exposes MongoDB internally through a **ClusterIP service** for secure communication.

---

### 3️⃣ Containerization (`Dockerfile`)
- Defines a lightweight **Node.js** image.
- Installs required dependencies from `package.json`.
- Sets up the app to run on port `3000`.
- Ensures consistent builds for all environments (local, dev, prod).

---

### 4️⃣ Kubernetes Deployment (`student-service.yaml`)
- Deploys the student service as a **Kubernetes Deployment**.
- Exposes it via a **NodePort Service** on port `3000`.
- Includes:
  - **Environment variables** for MongoDB connectivity.
  - **Resource requests/limits** for CPU and memory.
  - **Liveness & Readiness probes** for runtime health checks.
  - **Horizontal Pod Autoscaler (HPA)** that scales pods dynamically based on CPU utilization.

---

## 🔍 Core Features

✅ RESTful API design following microservice principles  
✅ Independent MongoDB persistence with schema validation  
✅ Fault tolerance with K8s health probes  
✅ Horizontal scalability via Kubernetes HPA  
✅ Modular design — easily replaceable or upgradable  
✅ Secure internal communication through service discovery in Kubernetes  

---

## 🧩 Role in Overall System

The **Student Service** acts as the **entry point for all student-related operations** in the larger **Student-Course-App** architecture.

It communicates with other services (like Enrollment and Course services) through internal REST calls inside the Kubernetes cluster.  
Each microservice manages its own data and logic, ensuring **loose coupling** and **independent scalability**.

---

## 🧠 Design Philosophy

This service follows a **polyglot microservice pattern**, meaning each microservice in the ecosystem is built using a different technology stack.  
For the Student Service:
- **Node.js** is chosen for its lightweight, asynchronous design.
- **MongoDB** offers flexibility for unstructured or evolving student data.
- The service is optimized for **containerized deployments** and **cloud-native scalability**.

---

## ⚖️ License

This project is licensed under the **MIT License**.  
You are free to use, modify, and distribute it for educational or commercial purposes.

---

## ❤️ Contributors

| Name | Role |
|------|------|
| [Tanuja RY](https://github.com/tanujary) | Lead Developer |
| [Devesh Singh](https://github.com/2024mt03089-devesh) | Contributor |
| [KP Sahoo](https://github.com/Krushna-Prasad-Sahoo/) | Contributor & Maintainer |

---

> _“Microservices are not just about dividing code — they’re about dividing responsibilities.”_
