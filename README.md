# Linux and Programming Tasks

## Task 1: `chmod`
**Problem:**  
Users are unable to execute commands in Linux.  
Every command results in the error: `-bash: /usr/bin/ls: Permission denied`.

**Objective:**  
Fix the issue so that users can execute commands without errors.

---

## Task 2: `killByPid`
**Problem:**  
Users are complaining that the computer has become very slow.  
Upon checking the list of processes, you discover that Firefox is consuming all available memory.  
The user doesn’t mind if the browser process is terminated.

**Objective:**  
Terminate the Firefox process to free up memory.

---

## Task 3: `clearDisk`
**Problem:**  
The disk has suddenly run out of space.  

**Objective:**  
Find the files that are taking up the most disk space and delete them.

---

## Task 4: `childToParent`
**Objective:**  
Write a program demonstrating the following:

1. Creation of two child processes (more is optional).
2. Handling of two signals that require processing by the program.
   - For example, do not handle signal “9” (as it cannot be handled).

---

## Task 5: `clientServer`
**Objective:**  
Create two programs demonstrating a client-server interaction.

### Client:
1. Sends a message to the server (user input from the keyboard).
2. Waits for a response from the server immediately after sending the message.
3. Repeats in a loop: sends another user message and waits for a response.

### Server:
1. Waits for a message from the client.
2. Modifies the message by adding the prefix “Server: ”.
3. Sends the modified message back to the client.
4. Repeats the cycle: waits for the next client message, modifies it, and sends it back.

**Note:**  
The server should be able to handle multiple clients simultaneously (at least 2).
