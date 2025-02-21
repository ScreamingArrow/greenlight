# Greenlight API Documentation

## 📌 Overview
The **Greenlight API** is a JSON-based web service designed for retrieving and managing movie information. It follows the principles outlined in the book *Let's Go Further* by **Alex Edwards** and is built using **Go** with best practices for RESTful API development.

For more details on the book, visit: [Let's Go Further](https://lets-go-further.alexedwards.net/).

## 🚀 Features
- **Rate Limiting** – Prevent abuse and manage traffic effectively.
- **User Authentication & Authorization** – Secure access with **JWT-based** authentication.
- **Full CRUD Operations** – Manage movies and users through a RESTful interface.
- **PostgreSQL Integration** – Store and retrieve data efficiently.
- **Structured JSON Logging** – Keep track of API activity for debugging and monitoring.
- **Robust Error Handling** – Handle API errors gracefully.
- **Metrics & Monitoring** – Collect and analyze performance metrics.

## 🛠️ Installation

1. **Clone the repository**
   ```sh
   git clone https://github.com/ScreamingArrow/greenlight.git
   cd greenlight-api
   ```
2. **Set up environment variables** (Create a `.envrc` file)
   ```sh
   GREENLIGHT_DB_DSN="your_database_dsn"
   SMTP_HOST="your_smtp_host"
   SMTP_PORT="your_smtp_port"
   SMTP_USERNAME="your_smtp_username"
   SMTP_PASSWORD="your_smtp_password"
   SMTP_SENDER="your_email_sender"
   ```
3. **Run database migrations**
   ```sh
   make db/migrations/up
   ```
4. **Start the API server**
   ```sh
   make run/api
   ```

