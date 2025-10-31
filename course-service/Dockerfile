# Base image
FROM python:3.11-slim

# Set working dir
WORKDIR /app

# Copy dependencies
COPY requirements.txt .

# Install packages
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

# Copy app code
COPY . .

# Expose port
EXPOSE 5000

# Run app
CMD ["python", "app.py"]
