#!/bin/bash

curr_dir=$(pwd)

# Step 1: Check for required software dependencies (like dotnet)
if ! command -v dotnet &> /dev/null
then
    echo "dotnet not found. Please install .NET SDK 8.0.0 first."
    exit 1
fi

# Step 2: Create the certificate
echo "Creating development certificate..."
dotnet dev-certs https --trust


# Step 3: Running the patients
for i in {1..3}
do
    echo "Running Patient $i with input $i"
    gnome-terminal -- bash -c "cd $(pwd)/patient_sol && dotnet build patient_sol.sln && dotnet run $i $i; exec bash" &
done

# Step 4: Running the hospital server
echo "Running the hospital server in the background..."
cd $(pwd)/hospital_sol && dotnet build hospital_sol.sln && dotnet run

# Instructions for user to execute the steps
echo "Setup complete! Please interact with the patient terminals as described in the readme."
