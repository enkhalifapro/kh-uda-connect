#!/bin/bash

# Variables
BROKER="localhost:9093"  # Change this to your broker's address if needed
TOPIC_NAME="example-topic"  # Replace with your desired topic name
PARTITIONS=1  # Number of partitions
REPLICATION_FACTOR=1  # Replication factor

# Create the Kafka topic
kafka-topics.sh --create --topic $TOPIC_NAME --bootstrap-server $BROKER --partitions $PARTITIONS --replication-factor $REPLICATION_FACTOR

# Check if the topic was created successfully
if [ $? -eq 0 ]; then
    echo "Topic '$TOPIC_NAME' created successfully!"
else
    echo "Failed to create topic '$TOPIC_NAME'."
fi
