
# Resume Generator with Golang and Docker

This project generates synthetic resumes using Golang and runs inside a Docker container.

## Prerequisites

- Docker installed on your machine
- Go (Golang) installed on your machine (for local development)

## Project Structure

```
resume-generator/
├── Dockerfile       # Docker configuration file
├── main.go          # Go application code to generate resumes
└── README.md        # This file
```

## Getting Started

### Step 1: Build the Docker Image

To build the Docker image, run the following command in the project directory:

```bash
docker build -t resume-generator .
```

This will build the Docker image with the Go application inside.

### Step 2: Running the Docker Container

To run the Docker container and execute the Go application, you have two options:

#### Option 1: Run Without Volume Mounting

This option runs the container using the binary built during the image build process:

```bash
docker run resume-generator
```

#### Option 2: Run With Volume Mounting

If you'd like to mount the current directory (for example, to save the `resume.json` file directly to your host machine), use the following steps:

1. **Run the container interactively**:
   ```bash
   docker run -it -v $(pwd):/app resume-generator /bin/sh
   ```

2. **Rebuild the binary inside the container**:
   After starting the container, rebuild the Go binary inside the mounted directory:
   ```bash
   go build -o main .
   ```

3. **Run the application**:
   Now, execute the binary:
   ```bash
   ./main
   ```

   This will run the application and either print the synthetic resume to the console or save it as a file (depending on the application logic).

### Step 3: Copy Generated Resume (Optional)

If the Go application saves the resume to a file (e.g., `resume.json`), you can copy the file from the container to your local machine using:

```bash
docker cp <container_id>:/app/resume.json .
```

This will copy the `resume.json` file from the container to your current directory.
