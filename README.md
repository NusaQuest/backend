# NusaQuest : Turn River Cleanups into Concert Ticket NFTs. 🚀

## ✨ Overview

🌏 NusaQuest is an AI 🤖 and DAO 🧠-powered impact-to-earn platform on Lisk Sepolia 🛰️ that rewards you for joining river cleanups across Indonesia 🇮🇩. Collect trash before it reaches the ocean 🌊, earn NUSA tokens 💰, and redeem them for NFT concert tickets 🎫. With KTP-based KYC 🪪🔍, NusaQuest bridges real-world action 🌱 with the power of Web3 🌐.

## 🧩 Architecture

    ```
    ├── backend/
    │   ├── config/         # Configuration files (e.g., database configuration)
    │   ├── constants/      # Project-wide constant values (e.g., custom errors, custom success, etc)
    │   ├── controllers/    # Business logic handlers for each route
    │   │   └── helper/     # Helper functions used within controllers
    │   ├── handlers/       # Core functions handling incoming HTTP requests
    │   ├── middlewares/    # Middleware functions (e.g., database connection handling)
    │   ├── models/         # Database models or data schemas
    │   ├── output/         # JSON response structure for success or error results
    │   ├── router/         # Route definitions and API endpoints
    │   ├── utils/          # Utility functions (e.g., validator)
    │   ├── views/          # HTML templates
    │   ├── .env            # Environment variables
    │   ├── .gitignore      # Files and folders ignored by Git
    │   ├── Makefile        # Automation commands for building and running
    │   ├── go.mod          # Go module definitions (dependencies)
    │   ├── go.sum          # Hashes of module dependencies for reproducibility
    │   └── main.go         # Entry point of the backend application
    ```

## 🧭 How to Run

This project uses [GoFiber](https://gofiber.io/) and a custom `Makefile` for a smoother development experience.  
Just run `make <task>` without remembering long commands!

### 📦 1. Install Golang

#### 📥 Download & Install

Follow the official installation guide based on your OS:
🔗 https://go.dev/doc/install

#### ✅ Verify Installation

After installation, run the following command to confirm:

```bash
go version
```

### 📁 2. Clone Repository

```bash
> git clone https://github.com/NusaQuest/backend
> cd backend
```

### 📚 3. Install Dependencies

```bash
> make install
```

### 🧪 4. Run the Server Locally

```bash
> make run
```

## 🔐 .env Configuration

Before running deploy or verification commands, make sure your `.env` file is properly set up in the root directory.

```env
# 📦 MongoDB connection URI
MONGO_URI=mongodb+srv://your_user:your_pass@cluster.mongodb.net/

# 🗂️ Name of the database
DB_NAME=<DB_NAME>

# 🤖 OpenAI API key (for AI features)
OPENAI_API_KEY=sk-...
```

## 🤝 Contributors

- 🧑 Yobel Nathaniel Filipus :
  - 🐙 Github : [View Profile](https://github.com/yebology)
  - 💼 Linkedin : [View Profile](https://linkedin.com/in/yobelnathanielfilipus)
  - 📧 Email : [yobelnathaniel12@gmail.com](mailto:yobelnathaniel12@gmail.com)
