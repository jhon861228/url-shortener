# 🌐 URL Shortener

Welcome to the **URL Shortener** project! This application provides a simple way to shorten long URLs, making them easier to share and manage. The project is composed of a powerful backend built with Golang and a sleek frontend developed with Astro.

## 🛠️ Tech Stack

### Infrastructure
- **Cloud:** AWS
- **IAC:** Terraform

### Backend
- **Language:** Go (Golang)
- **Framework:** Custom Go-based REST API
- **Database:** DynamoDB (NoSQL)
- **Deployment:** AWS Lambda

### Frontend
- **Framework:** Astro.js
- **Styling:** Tailwind CSS
- **Build Tool:** Vite

## 🚀 Features

- **URL Shortening:** Shorten long URLs with a simple API call.
- **Redirection:** Automatically redirect users from the short URL to the original URL.
- **Responsive UI:** User-friendly interface that works seamlessly across devices.

## 📂 Project Structure
```
url-shortener/
├─ backend/
│  ├─ cmd/
│  ├─ e2e/
│  ├─ internal/
│  │  ├─ shared/
│  │  ├─ shorturl/
│  │  │  ├─ app/
│  │  │  ├─ core/
│  │  │  ├─ infrastructure/
├─ frontend/
```


## 📝 Setup and Installation

### Infrastructure
1. **Terraform init:**
    ```bash
    terraform init
    ```
2. Configure credentials in ./aws/credentials and set profile as personal (in my case)

3. **Terraform Plan**
    ```bash
    terraform plan
    ```
3. **Terraform Apply**
    ```bash
    terraform apply
    ```


### Backend

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener/backend
   ```
2. **Install dependencies:**
    ```bash
   cd backend && go mod tidy
   ```
3. **Clone the repository:**
    ```bash
   cd backend && ./build.sh
   ```

