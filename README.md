# 🛠️ Golang RESTful API (Without Web Framework)

This is a simple RESTful API written in **Golang**, following these requirements:

- **MongoDB Integration**: Fetches data from a MongoDB collection.
- **In-Memory Database**: Supports creating (`POST`) and fetching (`GET`) data.
- **No Web Framework**: Uses only the standard `net/http` package.
- **Public GitHub Repository**: Clear instructions for configuration and running.

---

## 🚀 Features
- **POST `/mongo`** → Fetches data from MongoDB based on a filter.
- **POST `/memory`** → Adds a new user to an in-memory database.
- **GET `/memory`** → Retrieves all users from the in-memory database.

---

## 🛠️ Installation & Setup

### **1️⃣ Clone the Repository**

git clone https://github.com/Yhrone/rest-api_go.git
cd rest-api
