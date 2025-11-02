# 🎓 Course Service — Python + Flask + PostgreSQL Microservice

---

![Python](https://img.shields.io/badge/Python-3.11-blue?logo=python)
![Flask](https://img.shields.io/badge/Framework-Flask-lightgrey?logo=flask)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?logo=postgresql)
![Docker](https://img.shields.io/badge/Containerized-Docker-blue?logo=docker)
![Kubernetes](https://img.shields.io/badge/Orchestrated%20with-Kubernetes-326ce5?logo=kubernetes)
![License](https://img.shields.io/badge/License-MIT-yellow)

---

## 🧩 Overview

The **Course Service** is a lightweight, containerized **Python (Flask)** microservice that manages course-related data. It provides APIs to **create** and **retrieve** course information stored in a **PostgreSQL** database.

This service is part of a distributed microservices ecosystem (alongside the **Student Service**) and is designed with modularity, portability, and scalability in mind.

---

## 🏗️ Architectural Summary

```

+------------------------+
|     Course Service     |  --->  Python with Flask Framework
|------------------------|
|   /courses (POST)      |  --> Add new course
|   /courses (GET)       |  --> Retrieve all the courses
|   /courses/<id> (GET)  |  --> Retrieve student details
|   /                    |  --> Health check endpoint
+-----------+------------+
|
v
+-------------------+
|     PostgreSQL    | ---> Persistent data store
+-------------------+

```
The service is built using a **Flask REST API** that interacts with a **PostgreSQL** database via SQLAlchemy ORM.
It runs as a **containerized microservice** that can be deployed standalone or within a Kubernetes cluster.

At a high level:

1. The Flask app handles API requests and responses.
2. SQLAlchemy maps Course objects to relational tables in PostgreSQL.
3. Kubernetes manages deployment, service discovery, and autoscaling.
4. Persistent storage ensures PostgreSQL data durability via a **PersistentVolumeClaim (PVC)**.

---

## ⚙️ Technology Stack

| Layer                | Technology    | Purpose                                 |
| -------------------- | ------------- | --------------------------------------- |
| **Language**         | Python 3.11   | Application logic                       |
| **Framework**        | Flask         | REST API and web framework              |
| **ORM**              | SQLAlchemy    | Object-relational mapping to PostgreSQL |
| **Database**         | PostgreSQL 15 | Persistent course storage               |
| **Containerization** | Docker        | Environment encapsulation               |
| **Orchestration**    | Kubernetes    | Deployment, scaling, and management     |

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

---

## ☁️ Kubernetes Deployment Overview

The Kubernetes manifests define the complete lifecycle of the Course Service and its database.

### 🧩 `course-service.yaml`

* **Deployment:** Manages the Flask container lifecycle.
* **Service:** Exposes the API via a **ClusterIP**.
* **HPA:** Enables auto-scaling between 1–5 replicas based on CPU usage (50% target).

### 🗄️ `postgres.yaml`

* **Deployment:** Runs PostgreSQL 15 as a backend database.
* **PVC:** Provides persistent storage to retain data across restarts.
* **Service:** Exposes PostgreSQL internally to other microservices within the cluster.

---

## 🚀 Containerization

The service is fully Dockerized using a minimal **Python 3.11-slim** image for lightweight deployment.
The Dockerfile installs dependencies from `requirements.txt`, copies the application code, and exposes port `5000`.

This containerized design ensures:

* Consistent behavior across environments
* Faster deployments
* Isolation from host dependencies

---

## ⚡ Scalability & Reliability

To ensure smooth performance in cloud-native environments:

* **Horizontal Pod Autoscaler (HPA):** Dynamically scales pods based on CPU utilization.
* **Readiness & Liveness Probes:** Guarantee uptime and fault tolerance.
* **PersistentVolumeClaim (PVC):** Protects PostgreSQL data during pod restarts.

---

## 🧠 Key Design Principles

* **Separation of Concerns:** Application logic, database, and orchestration are isolated cleanly.
* **Cloud-Native:** Ready for deployment in Kubernetes with built-in scaling and resilience.
* **Lightweight:** Minimal dependencies for fast startup and efficient resource usage.
* **Consistent Interface:** Follows REST standards compatible with the Student Service.

---

## ⚖️ License

This project is licensed under the **MIT License**.  
You are free to use, modify, and distribute it for educational or commercial purposes.

---

## ❤️ Contributors

| Name | Role |
|------|------|
| [Siddharth Jain](https://github.com/2024mt03040-web) | Lead Developer |
| [Devesh Singh](https://github.com/2024mt03089-devesh) | Contributor |
| [KP Sahoo](https://github.com/Krushna-Prasad-Sahoo/) | Contributor & Maintainer |

---



> _“Microservices are not just about dividing code — they’re about dividing responsibilities.”_
