# ToDo CLI App

![image](https://github.com/vinay-chin/todocli/assets/109643166/5e6bc1dd-9398-47a9-914e-12ab5a33ac16)

This is a simple command-line interface (CLI) ToDo app built with Gorm ORM for MySQL and Cobra framework. It allows you to manage your tasks efficiently through the command line.

## Features

- **Add Tasks:** Easily add new tasks with titles and descriptions.
- **List Tasks:** View a list of all tasks along with their details.
- **Mark as Finished:** Mark tasks as finished when they are completed.
- **Database Integration:** Utilizes Gorm ORM to connect to a MySQL database for data persistence.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/vinay-chin/todo-cli.git
    ```

2. Navigate to the project directory:

    ```bash
    cd todo-cli
    ```

3. Build the executable:

    ```bash
    go build .
    ```

4. Run the app:

    ```bash
    ./todocli
    ```

## Usage

- **Adding a Task:**

    ```bash
    ./todocli add-task
    ```

- **Listing Tasks:**

    ```bash
    ./todocli view-tasks
    ```

- **Marking as Finished:**

    ```bash
    ./todocli update-task
    ```

## Configuration

Update the database connection details in the `cmd/task.go` file:

# License

Todo App CLI is licensed under the Creative Commons Attribution 4.0 International License. To view a copy of this license, visit [http://creativecommons.org/licenses/by/4.0/](http://creativecommons.org/licenses/by/4.0/).

You are free to:

- **Share:** Copy and redistribute the material in any medium or format.
- **Adapt:** Remix, transform, and build upon the material for any purpose, even commercially.

Under the following terms:

- **Attribution:** You must give appropriate credit, provide a link to the license, and indicate if changes were made. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

No additional restrictions — You may not apply legal terms or technological measures that legally restrict others from doing anything the license permits.

