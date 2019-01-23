# Mutex and Channel basics

### What is an atomic operation?
> Operations that run completely independently of any other processes. Must be performed entirely or not performed at all. 

### What is a semaphore?
> A variable that is almost like a flag.

### What is a mutex?
> A mutually exclusive flag. A gatekeeper that decides which thread get to work. The thread releases the mutex when done. (snakkepinne)

### What is the difference between a mutex and a binary semaphore?
> Binary semaphore is a signaling mechanism, while the mutex is a locking mechanism used to synchronize access to a resource. 

### What is a critical section?
> The part of a program that accesses shared resources. Only when a process is in its Critical Section can it be in a position to disrupt other processes.

### What is the difference between race conditions and data races?
 > Data races is a subset of all race conditions. Data-races happen when two or more threads access the same memory without proper locking which can lead to unexpected values.

 The Race condition term in general could also refer to e.g. threads that deadlock occasionally due to races in thread scheduling

### List some advantages of using message passing over lock-based synchronization primitives.
> Safer, simpler

### List some advantages of using lock-based synchronization primitives over message passing.
> Faster, used for very specific things
