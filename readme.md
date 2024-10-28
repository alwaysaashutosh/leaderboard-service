# Leaderboard-Service

The Leaderboard-Service is a backend application designed for efficient user ranking management, handling submissions, and tracking ranks across various scopes. It supports a range of API endpoints that allow for updating, and querying user ranks and top scores.

## Getting Started

Follow these steps to set up and run the Leaderboard-Service.

### Prerequisites
- **Go** - Make sure Go is installed on your machine.
- **PostgreSQL** - This project uses PostgreSQL for data storage, so ensure it’s up and running.
- **Docker** (optional) - To run the application in a Docker container.

### Steps to Run the Project

1. **Clone the Project**

   ```bash
   $ git clone https://github.com/alwaysaashutosh/leaderboard-service.git 
   ```

2. **Update the Configuration File**

   - Edit the configuration file at ```conf/app.yaml``` as needed.
   - To automate table creation, set ```autoMigrate = true```, which will enable GORM auto-migration to create tables from the GORM model schema.

3. **Build and Run the Service**
   ```bash
   $ go build .

   $ ./leaderboard-service
   ```
4. **Building and Running with Docker**

   - Ensure the configuration file is updated in ```conf/app.yaml```, including database URL and other settings.

   - Build the Docker Image
     ``` bash
     $ docker build -t leaderboard-service .
     ```
     
     ``` bash
     $ docker run -d -p 8080:8080 --name leaderboard-container leaderboard-service
     ```
   
   - The service should now be running in Docker and accessible at http://localhost:8080.

5. **Swagger**
    - Swagger will be available at http://localhost:8080/leaderboard-service/swagger/index.html#/

6. **Database**
   - This service uses **PostgreSQL**. Ensure that your PostgreSQL instance is running and accessible via the connection URL provided in the configuration file.


