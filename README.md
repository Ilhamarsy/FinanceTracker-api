# Finance Tracker Application

Welcome to the Finance Tracker Application! This README provides a guide to help you get started with the application, configure environment variables, and understand the server endpoints.

## Getting Started

To get started with the Finance Tracker Application, follow these steps:

1. **Clone the Repository**

   First, clone the repository to your local machine:

   ```bash
   git clone https://github.com/Ilhamarsy/FinanceTracker-api.git
   ```
2. **Navigate to the Project Directory**
   Install the required dependencies using npm or yarn:

   ```bash
   cd finance-tracker-api
   ```
3. **Download Dependencies**
   Download the dependencies listed in the go.mod file:

   
   ```bash
   go mod tidy

   ```

4. **Configure Environment Variables**
   Rename the ``.env.example`` file to ``.env``:
   ```bash
   mv .env.example .env
   ```
   Open the ``.env`` file and update the values as needed. Ensure you configure the environment variables according to your setup
   ```bash
   DB_HOST=YOUR_DB_HOST
   DB_PORT=YOUR_DB_PORT
   DB_USER=YOUR_DB_USER
   DB_PASSWORD=YOUR_DB_PASSWORD
   DB_NAME=YOUR_DB_NAME
   JWT_SECRET=YOUR_JWT_SECRET
   PORT=YOUR_PORT
   ```

5. Start the Application
   Start the server:
   ```bash
   go run main.go
   ```

## Server Endpoint
The default endpoint for the web is:
```bash
http://localhost:8080/api
```

The endpoint are
- POST``/register``: To register new user
- POST``/login``: To login existing user

- POST``/category``: To add category
- GET``/categories``: To get data category

- POST``/income``: To add new income
- GET``/incomes``: To get data income
- PUT``/income/:id``: To update data income
- DELETE``/income/:id``: To delete income

- POST``/expense``: To add new expense
- GET``/expenses``: To get data expense
- PUT``/expense/:id``: To update data expense
- DELETE``/expense/:id``: To delete expense

- GET``/stats``: To get stats
- GET``/stats-yearly``: To get stats in a year

#
Thank you for using the Finance Tracker Application!
