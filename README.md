# SimpleFullStack

> 🚀 Full-Stack Web Application with PostgreSQL, Golang, and React  
> A fully Dockerized setup for an easy-to-run stack.

## 🛠 Tech Stack

<p align="center">
  <img src="https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang" />
  <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=white" alt="React" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>

This project demonstrates a full-stack setup with:

- **PostgreSQL** for database management 🗄️
- **Golang** for backend REST API development 🐹
- **React** for frontend UI 🌐
- **Docker** for containerization 🐳

## 🚀 Getting Started

### Prerequisites
Make sure you have the following installed on your system:

- **Docker** 🐳
- **Docker Compose** 📦

### Run the app!

1. **Clone the repository**

   ```bash
   git clone https://github.com/VictorTsokonov/SimpleFullStack.git
   cd SimpleFullStack
   
2. **Run the application with Docker Compose**
    ```bash
   docker-compose up --build
   
3. **Access the application**
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

4. **Test the API**

   You can test the API endpoints directly or through the frontend interface.
   The backend API serves endpoints for managing a list of games.

## 📖 Usage

### Frontend (React)
The frontend, built with React, provides a user-friendly interface for interacting with the API. You can add, delete, and view games in the library.

### Backend (Golang)
The backend provides a REST API with the following endpoints:

- `GET /games` - Retrieve a list of games
- `POST /game` - Add a new game
- `DELETE /game?id=<id>` - Delete a game by ID

### Database (PostgreSQL)
Data is stored in a PostgreSQL database container, ensuring persistence across sessions.

## 🤝 Contributing

1. **Fork the project**
2. **Create your feature branch** (`git checkout -b feature/AmazingFeature`)
3. **Commit your changes** (`git commit -m 'Add some AmazingFeature'`)
4. **Push to the branch** (`git push origin feature/AmazingFeature`)
5. **Open a pull request**

## 🙏 Acknowledgments

- **Docker** - Making containerization easy
- **Golang** - The Go Programming Language
- **PostgreSQL** - The World's Most Advanced Open Source Relational Database
- **React** - A JavaScript library for building user interfaces

