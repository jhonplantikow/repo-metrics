
  

# Documentation: Commit Analysis Project
### **Project Overview**

This project processes and analyzes commit data from a CSV file to calculate and rank repositories by their activity scores. It generates two output files:

-  **`rejected_cases.txt`**: Contains invalid rows with explanations of the errors.

-  **`output.txt`**: Contains valid rows sorted by their activity score in descending order.

  

### **How to Run**

1. Clone the project repository.

2. Navigate to the project directory in your terminal.

3. Set up the necessary environment variables using the `.env` file.

4. Execute the following command:

  

```bash

go  run  .\cmd\main.go  startup  APP_ENV=local.env

```

  

This command initializes the application using the `local.env` configuration file.

  

---

  

### **File Structure**

-  **Input File**:

The input file is a CSV located at:

`assets/in/commits.csv`

  

Example format:

```

timestamp,username,repository,files,additions,deletions
1610969774,user123,repo123,3,10,5

```

  

-  **Output Files**:

The program generates the following files in the directory `assets/out/`:

1.  **`rejected_cases.txt`**: Contains rows that failed validation, along with error messages explaining why they were rejected.

2.  **`output.txt`**: Contains valid rows, sorted by activity score in descending order.

  

---

  

### **Score Calculation**

The activity score for each repository is calculated using the following formula:

  
```text
Score = CommitScore + 0.5 * Additions + 0.5 * Deletions + 0.2 * Files
```

  

Where:

-  **CommitScore**: A constant value (default is `1`) assigned to every commit.

-  **Additions**: Number of lines added in the commit.

-  **Deletions**: Number of lines removed in the commit.

-  **Files**: Number of files modified in the commit.

  

The program aggregates the scores for all commits in a repository to compute its total activity score.

  

---

  

### **Top 10 Repositories**

After processing, the top 10 repositories with the highest scores are displayed in the terminal and written to the `output.txt` file. Example output:

  

```

Repository: repo476, Score: 1821267.38
Repository: repo260, Score: 571665.00
Repository: repo920, Score: 328947.81
Repository: repo795, Score: 283194.16
Repository: repo161, Score: 207054.09
Repository: repo1143, Score: 194406.61
Repository: repo518, Score: 176195.44
Repository: repo1185, Score: 151120.91
Repository: repo1243, Score: 140042.42
Repository: repo250, Score: 120840.69

```
---
### **Error Handling**

Rows that fail validation are logged in the `rejected_cases.txt` file with detailed error messages. For example:

```
1610969774,user123,,3,10,5,error:invalid data: repository must not be empty
1610969774,user123,repo123,0,10,5,error:invalid data: files must be greater than 0
1610969774,user123,repo123,2,invalid,5,error:failed to sanitize additions: strconv.Atoi: parsing "invalid": invalid syntax
1610969774,user123,repo123,2,10,invalid,error:failed to sanitize deletions: strconv.Atoi: parsing "invalid": invalid syntax
```
---
### **Configuration**

The application relies on environment variables defined in the `.env` file. Example configuration:
```env
IN_REPOS=assets/in/commits.csv
OUT_REJECTED=assets/out/rejected_cases.txt
OUT_REPOS=assets/out/output.txt
APP_ENV=local
```
---
### **Notes**
- Ensure that the input file follows the required CSV format.
- The application automatically creates the output directories and files if they do not exist.


For any questions or issues, please contact me. 🚀