
---

## ✅ **FAN-IN and FAN-OUT Interview Questions**

---

### 🔹 **Conceptual/Behavioral Questions (Theory)**

1. **What is the Fan-In pattern in concurrent programming?**
2. **What is the Fan-Out pattern in concurrent programming?**
3. **How does Fan-In help in reducing data processing bottlenecks?**
4. **How does Fan-Out increase concurrency in a program?**
5. **Compare Fan-In vs Fan-Out patterns.**
6. **What are the real-world scenarios where Fan-In is useful?**
7. **Give a real-world analogy of Fan-Out pattern.**
8. **What problems might arise if we don’t handle synchronization properly in Fan-In?**
9. **Why is it important to close channels in a Fan-In pattern?**
10. **How would you deal with channel leaks in a Fan-Out design?**
11. **What are the challenges with merging multiple input channels (Fan-In)?**
12. **How to broadcast one channel to multiple receivers (Fan-Out) in Golang?**
13. **What’s the difference between using `sync.WaitGroup` and closing channels in Fan-In?**
14. **What is the role of select statements in Fan-In/Fan-Out?**
15. **How do you implement graceful shutdown in a Fan-Out architecture?**

---

### 🔹 **Practical/Coding Interview Questions**

#### 🧩 **Fan-In Coding Questions**

1. ✅ **Write a Go program to merge two channels into a single output channel (Fan-In).**
2. ✅ **Implement a Fan-In with three producers and a single consumer.**
3. ✅ **How do you prevent goroutine leaks in a Fan-In pattern when one producer hangs?**
4. ✅ **Implement a timeout for a Fan-In channel using `context.WithTimeout`.**
5. ✅ **Merge N number of channels into one using `reflect.Select`.**
6. ✅ **Build a Fan-In pattern with signal propagation using a done channel.**
7. ✅ **Create a Fan-In pattern with dynamic producer registration.**

#### 🧩 **Fan-Out Coding Questions**

1. ✅ **Write a Go program where one input channel sends data to multiple goroutines (Fan-Out).**
2. ✅ **Build a worker pool with Fan-Out pattern to process items concurrently.**
3. ✅ **Distribute numbers from a channel to 3 goroutines and process them in parallel.**
4. ✅ **Implement rate limiting in a Fan-Out pattern.**
5. ✅ **Broadcast the same message to multiple consumers using Fan-Out (pub-sub style).**
6. ✅ **Fan-Out to multiple workers and use `sync.WaitGroup` to wait for all.**
7. ✅ **Use buffered channels to improve performance in a Fan-Out pattern.**

---

### 🔹 **Advanced Scenarios / Mixed Fan-In & Fan-Out**

1. ✅ **Design a concurrent pipeline using Fan-Out to process data and Fan-In to collect results.**
2. ✅ **Fan-Out tasks to workers, and Fan-In the processed results to a single channel.**
3. ✅ **Build a concurrent file processor using Fan-Out (for processing) and Fan-In (for result collation).**
4. ✅ **Implement error propagation from Fan-Out workers to Fan-In collector.**
5. ✅ **Create a multi-stage pipeline using Fan-Out and Fan-In patterns.**

---

## 🔁 Real-Time System Design Use-Cases (Discussion)

* How would you implement **log aggregation** using Fan-In?
* How would you build a **distributed web scraper** using Fan-Out?
* Design a **rate-limited API gateway** using Fan-Out worker pools.
* Build a **message queue consumer group** using Fan-Out pattern.
* In a **real-time analytics system**, how would you Fan-In multiple event sources?

---
