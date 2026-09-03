### Socket programming
***Socket programming is the way programs communicate with each other over a network using sockets.***


***A socket is a communication endpoint created by the operating system that allows programs to talk over a network.***


***Socket = IP address + Port number + Protocol (TCP/UDP)***


**Server**


IP        : 192.168.1.10


Port      : 8080


Protocol  : TCP


This forms a TCP socket


**client**


IP        : 192.168.1.20


Port      : 54012 (random)


Protocol  : TCP


This forms another socket


(Client IP, Client Port, Server IP, Server Port) - ***This is called a socket pair.***


- listener, _ := net.Listen("tcp", ":8080") - ***This creates a listening socket***


- conn, _ := listener.Accept() - ***This creates a new socket per client***



This is a core socket-design concept:


👉 you need a simple protocol between client and server.

***A socket only sends bytes, not “functions”.***


So we do this:


- Client sends:   **COMMAND + DATA**


- Server does:   **Read command → choose function → execute → respond**

***This is called message-based dispatch.***

Client sends text like:

## How can a single socket server call different functions based on what the client sends?

