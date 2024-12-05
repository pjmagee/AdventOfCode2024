# Advent of Code 2024

This repository contains my solutions to the **2024 Advent of Code** challenges.  
I’ll be solving the puzzles using a variety of programming languages, including:

- Go
- Python
- C++
- C#
- Rust

**Note:** Not every puzzle will have solutions in all languages. Balancing time constraints and complexity may lead to some puzzles being skipped or solved in a limited set of languages.

---

## 📂 Repository Structure

The solutions are organized by programming language in the `src` directory.  
Each subdirectory corresponds to a specific language and contains solutions for the days completed in that language.

---

## 📜 Input Data Policy

In compliance with the guidelines on the Advent of Code website, **input data is not included in this repository**. Input data is downloaded dynamically at runtime using a **Dagger Pipeline**, leveraging the session cookie provided to the `--session` argument.

---

Here’s a revised version of your text with improved structure, clarity, and formatting:

---

## 🛠 Developing Solutions

To ensure consistency and streamline the CI/CD pipeline, solutions should adhere to a standardized structure and execution process.

### General Rules

1. **Main File:** Each language directory must include a `main` file to run the solution for a specific day.
2. **Function Signature:** Solutions must be implemented as a function that:
  - Accepts the input data as a string.
  - Returns the result as a string.
3. **Input and Output:** Solutions should:
  - Read input data from `stdin`.
  - Write the output to `stdout`.
4. **Command Line Argument:** The challenge day should be provided as a command-line argument.

### Command Syntax

#### Bash
```bash
[command] [day] < [input file]
```

#### PowerShell
```powershell
Get-Content -Path input.txt | [command] [day]
```

### Example Commands for Different Languages

- **Go:**
  ```bash
  go run main.go 1 < input.txt
  ```

- **Python:**
  ```bash
  python main.py 1 < input.txt
  ```

- **C# (.NET):**
  ```bash
  dotnet run --project src/csharp -- 1 < input.txt
  ```

- **CMake:**
  ```bash
  cmake --build build --target run -- -d 1 < input.txt
  ```

---

This version is more concise, organized, and easier to follow while maintaining all the essential details.

## 🌤️ Running Solutions in the Cloud

Solutions can be run in the cloud using the **Dagger Engine**. Refer to GitHub Actions workflows in the `.github/workflows` directory for examples of how to run solutions in the cloud.

## 🛠️ Running Solutions Locally

### 1. Using Dagger

You can run solutions using **Dagger functions** with the following command:

```bash
dagger call run --lang=LANG --day=DAY
```

#### Input Data Handling
- **If input data is already available:**  
  Ensure the `inputs/` folder contains the input data for the day you are running. This folder will be uploaded to the Dagger Engine during execution.
- **If input data is missing:**  
  The pipeline will automatically fetch the input data from the Advent of Code website using your session cookie.

#### Storing Your Session Cookie
Save your Advent of Code session cookie in a file (e.g., `session.txt`) and run the following command to provide it to the pipeline:

```bash
dagger call --session=file:./session.txt
```

### 2. Using Local Development Tooling

You can also run solutions directly using the build tools of the respective programming language.

#### Example: Running Go Solutions
To run the Go solutions for a specific day, use the following commands:
```bash
cd src/go
go run main.go DAY < INPUT.txt
```

For other languages, use the appropriate build and execution commands as required for the language.

---

## 🤝 Contribution and Feedback

This is primarily a personal project, but feedback and contributions are welcome!  
Feel free to fork the repository or submit issues and pull requests for improvements or alternative solutions.

---

## 📜 License

This repository is licensed under the [GPL-3.0 license](LICENSE). Feel free to adapt or reuse the solutions, but don’t forget to credit appropriately!

---

Enjoy Advent of Code 2024! 🎄✨