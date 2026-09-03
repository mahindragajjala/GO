---

### ✅ **Built-in / Standard Errors**

1. `error` interface (basic error type)
2. `fmt.Errorf` (formatted error)
3. `errors.New` (basic static error)
4. `errors.Is` (error comparison)
5. `errors.As` (type assertion)
6. `errors.Unwrap` (extract wrapped error)

---

### ✅ **Custom Errors**

7. Struct-based custom errors (`type MyError struct`)
8. Custom error with additional fields (timestamp, code, etc.)
9. Custom error types implementing `Error() string`

---

### ✅ **Runtime Errors**

10. `nil` pointer dereference
11. Array/slice index out of range
12. Divide by zero
13. Invalid type assertions
14. Deadlocks (in goroutines/channels)
15. Stack overflow (due to infinite recursion)

---

### ✅ **Panic and Recover**

16. `panic()`
17. `recover()` (used to catch panics)

---

### ✅ **I/O and OS Errors**

18. File not found
19. Permission denied
20. EOF (end of file)
21. Network timeout
22. Connection refused
23. Disk full

---

### ✅ **Context Errors**

24. `context.DeadlineExceeded`
25. `context.Canceled`

---

### ✅ **Type-specific Errors**

26. `strconv.NumError` (conversion errors)
27. `json.SyntaxError` (JSON parsing error)
28. `json.UnmarshalTypeError`
29. `os.PathError`
30. `os.LinkError`
31. `os.SyscallError`
32. `net.OpError`
33. `net.DNSError`
34. `url.Error`

---

### ✅ **HTTP & Web Related Errors**

35. HTTP status code errors (e.g., 404, 500)
36. Invalid HTTP method or header
37. Request timeout
38. TLS handshake error

---
