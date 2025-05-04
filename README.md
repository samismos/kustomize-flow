# kustomize-flow

kustomize-flow is a powerful tool for managing and visualizing **Kustomization** resources in any Kustomize-based project. It helps users explore and understand the relationships between different components in their configuration, providing **clarity** and **control** over complex setups.

## Features

- 📊 **Intuitive Visualization** – Gain insights into your Kustomization hierarchy and dependencies.
- ⚡ **Streamlined Management** – Easily explore, edit, and maintain Kustomization resources.
- 🐳 **Containerized Deployment** – Get started effortlessly with Docker Compose.

## Tech Stack

- **Backend:** [Go](https://golang.org/) & [Gin](https://gin-gonic.com/)
- **Frontend:** [React.js](https://react.dev/) & [Vite](https://vitejs.dev/)
- **Orchestration:** [Docker Compose](https://docs.docker.com/compose/)

## Prerequisites

Ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/samismos/kustomize-flow.git
cd kustomize-flow
```

### 2️⃣ Add repositories to visualize

Create the _/repos_ directory in the root of the project and populate it with repositories which you would like to visualize, in order for the backend to traverse their resources.

### 3️⃣ Start the Application

```sh
docker compose up --build
```

### 4️⃣ Access the Application UI

```sh
http://localhost:5173
```

## Project Structure

```sh
kustomize-flow/
│── backend/             # Go + Gin backend
│── frontend/            # React.js + Vite frontend
│── docker-compose.yaml  # Container orchestration
│── LICENSE.md           # Open-Source License
│── README.md            # Project documentation
│── .env                 # Environment variables
```

## Contributing
Contributions are welcome! Feel free to fork the repository and submit a pull request. Be sure to follow the project's coding guidelines.

## License
This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md) for more details.