The `os` package in Go is part of the **standard library** and provides **platform-independent interfaces** for interacting with the operating system. It includes **file handling, process management, environment variables, permissions, and system calls**.

**Platform-independent interfaces** - Platform-independent interfaces in Go mean that the same code works on multiple operating systems (Windows, Linux, macOS, etc.) without changes. The os package provides a common API to perform system operations (like file handling, env vars) so you don’t need to worry about OS-specific system calls.

---

## **1. What is the use of `os` package in Go?**

* To **work with files and directories** (create, read, write, delete).
* To **manage environment variables** (get, set, list).
* To **handle permissions** (file modes like `0777`).
* To **work with processes** (arguments, exit codes, signals).
* To **fetch system information** (hostname, user, temporary dirs).
* To **perform system-level operations** (pipes, symlinks).

---

## **2. Important Topics in `os` Package**

### **(A) File and Directory Operations**

* Creating, opening, reading, writing, closing files.
* Checking file existence and properties (`os.Stat`).
* Removing, renaming, and creating directories.

### **(B) Environment Variables**

* Accessing (`os.Getenv`)
* Setting (`os.Setenv`)
* Listing (`os.Environ`)

### **(C) Permissions and Modes**

* File modes (e.g., `os.FileMode`, `0644`, `0777`)
* Checking file types (directory, symlink).

### **(D) Process Management**

* Command-line arguments (`os.Args`)
* Process exit (`os.Exit`)
* Process IDs (`os.Getpid`, `os.Getppid`).

### **(E) Temporary Files and Directories**

* `os.TempDir()`
* `os.CreateTemp()`

### **(F) Error Handling**

* `os.IsNotExist()`, `os.IsExist()`, `os.IsPermission()`

### **(G) Working with Links**

* `os.Symlink()`
* `os.Readlink()`

---

## **3. Important Functions in `os` Package**

### **File Handling**

* `os.Create(name string)` – Creates a file (truncates if exists).
* `os.Open(name string)` – Opens file for reading.
* `os.OpenFile(name string, flag int, perm FileMode)` – Opens file with flags (read/write/append).
* `os.ReadFile(name string)` – Reads entire file into memory.
* `os.WriteFile(name string, data []byte, perm FileMode)` – Writes data to a file.
* `os.Remove(name string)` – Deletes a file.
* `os.Rename(oldpath, newpath string)` – Renames/moves file.
* `os.Stat(name string)` – Gets file info (size, mode, mod time).

---

### **Directory Handling**

* `os.Mkdir(name string, perm FileMode)` – Creates a directory.
* `os.MkdirAll(path string, perm FileMode)` – Creates nested directories.
* `os.RemoveAll(path string)` – Deletes directory and contents.

---

### **Environment Variables**

* `os.Getenv(key string)` – Get env variable value.
* `os.Setenv(key, value string)` – Set env variable.
* `os.Environ()` – Returns all environment variables as slice.

---

### **Process and Arguments**

* `os.Args` – Access command-line arguments.
* `os.Exit(code int)` – Exit program with code.
* `os.Getpid()` – Current process ID.
* `os.Getppid()` – Parent process ID.

---

### **Permissions and Modes**

* `os.Chmod(name string, mode FileMode)` – Change file permissions.
* `os.Chown(name string, uid, gid int)` – Change file owner.

---

### **Error Handling**

* `os.IsNotExist(err error)` – Checks if error is "file not found".
* `os.IsExist(err error)` – Checks if error is "file exists".
* `os.IsPermission(err error)` – Checks for permission errors.

---

### **Temporary Files**

* `os.TempDir()` – Returns default temporary directory path.
* `os.CreateTemp(dir, pattern string)` – Creates temporary file.

---

### **Symbolic Links**

* `os.Symlink(oldname, newname string)` – Create symbolic link.
* `os.Readlink(name string)` – Reads symbolic link target.

---

### **User/System Information**

* `os.Hostname()` – Returns hostname of the system.
* `os.Getwd()` – Returns current working directory.
* `os.Chdir(dir string)` – Change current working directory.

---

## **4. Code Examples**

### **Example 1: File Create and Write**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    file.WriteString("Hello, OS package in Go!\n")
    fmt.Println("File created and data written.")
}
```

---

### **Example 2: Read Environment Variable**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    user := os.Getenv("USER")
    fmt.Println("Current user is:", user)
}
```

---

### **Example 3: Command-Line Arguments**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Arguments:", os.Args)
}
```

---

### **Example 4: Check File Exists**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Stat("example.txt")
    if os.IsNotExist(err) {
        fmt.Println("File does not exist")
    } else {
        fmt.Println("File exists")
    }
}
```
