# README

## Setup instructions

Hope you have ports 5001-5004 open, as that's what the project uses :)

The steps below detail how set up the project and run the hospital (server) and patient terminals.

You can set up the environment in one of two ways:

1. Manually open and run the terminals
2. Run script.sh to automatically open terminals and run start-up commands

---

#### Option 1: **Manually open and run terminals**

If you prefer to manually open each terminal, follow these steps:

0. **Prepare environment in the secu-handin2 folder**:
   - Create the development certificate:
   ```bash
   dotnet dev-certs https --trust
   ```
   - Build the project:
   ```bash
   dotnet build hospital_sol/hospital_sol.sln && dotnet build patient_sol/patient_sol.sln
   ```

1. **Open a terminal for the hospital server**:
   - Navigate to the `hospital_sol` folder:
     ```bash
     cd hospital_sol
     ```
   - Run the hospital server:
     ```bash
     dotnet run
     ```
   - The hospital server will start listening.
   - Combined code for copy-paste:
   ```bash
   cd hospital_sol && dotnet run
   ```

2. **Open three terminals for the patients**:
   - For each of the 3 remaining terminals, navigate to the `patient_sol` folder and run:
     ```bash
     cd patient_sol
     dotnet build patient_sol # only need to run this once
     dotnet run <ID> <DATA>
     ```
     Replace `<ID>` with 1, 2, or 3 (make sure each patient has a unique ID).
     Replace `<DATA>` with an integer between 1 and 15000.

3. **Execute**:
   - Press `Enter` in each patient terminal to distribute shares among the patients.
   - Press `Enter` again in each patient terminal to send the results to the hospital server.
   - You should see the final sum printed in the hospital terminal, confirming that the program has executed successfully.

---

#### Option 2: Run script.sh to automatically open terminals and run start-up commands****

If you prefer to automatically open the required terminals and run the start-up commands, you can use the provided script. This method will open the necessary terminals and run the commands for you.

1. **Run the setup script**:
   - Make the script executable:
     ```bash
     chmod +x script.sh
     ```
   - Execute the script:
     ```bash
     ./script.sh
     ```
   
   The script will:
   - Create a development **certificate**
   - **Build** the project using solutions hospital_sol.sln and patient_sol.sln ... 
   - Open a terminal for the **hospital server** and run `dotnet run`.
   - Open three terminals for the **patients**, each running the `dotnet run <ID> <DATA>` command.

2. **Execute**:
   - Once the script has opened the terminals, you just need to press `Enter` in each of the patient terminals to begin the share distribution process.
   - After that, press `Enter` again to send the results to the hospital server.
   - The hospital server will display the final aggregation once all results have been received.

---

### 3. Completing the Process

Once all the results have been sent, you will see the following output in the hospital terminal:

```
Received share <last share received>. Got 3 out of 3 shares.
Aggregation result: <3 * 15000 - patient1.Data - patient2.Data - patient3.Data>
```

IF you ran it using the script, or used 1, 2, 3 as the data inputs:
```
Aggregation result: 44994
```

Now, the program has finished executing, and you can (manually) close all terminals.
