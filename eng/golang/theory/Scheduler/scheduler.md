# 🗓 Go Scheduler — The Most Detailed Guide in Simple Terms

![Image 1](/assets/images/scheduler/image_1.png)

Let's design a Go scheduler from scratch — we'll start with the simplest and most understandable naive implementation, and then step by step we'll figure out what flaws it has, and come up with ways to solve them, gradually complicating the general model.

> [!TIP]
> This is one of the best ways to understand a complex system or concept — to go through the path of its step-by-step design. The system is complex, realizing it is very difficult, but we will break it down into simple steps that are very easy to understand. After that, the puzzle will come together in your head by itself, and the overall picture of the system will be just as simple and obvious to you.

---

## 1. 👋 Introduction: Why Study the Scheduler?

Concurrency is a very useful, but very complex thing. Usually, developers have to solve a lot of problems when working with it, but the Go language has reduced these problems to a minimum. Here, working with concurrent code is implemented very simply, clearly, and efficiently.

**Key entities that help with this:**
1. **Goroutines**, which can work concurrently and independently of each other.
2. **Scheduler**, which manages them.
3. **Channels**, which help exchange data between goroutines.

> [!NOTE]
> I also have a detailed video about the internal structure of channels — I also advise you to watch it to look at all this machinery from a different angle. In this article, we will analyze the scheduler and goroutines.

**The article is not for beginners — I assume that the reader is already familiar with Go. At a minimum:**
- Has gone through the Tour of Go.
- Understands how to work with goroutines, at least at the most basic level.

There is also a video version of this article. I tried to make every effort to express my thoughts as effectively as possible in both formats, so just choose the one that is closer to you. However, text and video have their pros and cons — some things are much easier to show visually in a video, and some concepts are better absorbed in text form. Therefore, if you want to learn the material as effectively as possible, I would advise you to watch the video first, and then read the article to consolidate.

By the way, the second half of the video is an additional practical part in which I dissect the scheduler using my utility **goschedviz**. There is no such thing in the current article, since it is quite difficult to format this as text. But if there are enough people interested, I am ready to present it in a separate article.

---

## 2. 🎯 Problem Statement

So, let's imagine that we are engineers from the Go Team, and we face an ambitious task — to teach our programs to work concurrently and, if possible, in parallel. Yes, the terms "concurrently" and "in parallel" are often confused, but later we will analyze what their essence is and what the difference is.

**Main requirements for our concurrently running programs:**
- All available processor cores must be used (if necessary).
- Cheapness in terms of performance and memory used — we want to run thousands, hundreds of thousands of tasks concurrently!
- The solution should be understandable and simple to use — the code should be intuitive.

> [!WARNING]
> I also draw attention to the fact that despite the volume of the article, I have to make simplifications in many cases to make the material easier to digest, and so that the article does not turn into a book. In some places I will draw your attention to this, and in some — not.

---

## 3. 📚 Brief Educational Support on Basics

In order to start designing the scheduler, we need an understanding of some basic concepts — the computer processor and its cores, the operating system scheduler and system threads, what concurrency and parallelism are. If you are familiar with all this, you can safely move on to the next section.

If it seems to you that after reading this section you have remembered little and have not realized everything — this is normal. In the course of the further narrative, I will regularly return to individual concepts from here, and gradually everything will sort itself out in your head, understanding will come, and then everything will be remembered by itself.

### 3.1 💻 How Computer Parallel Operation Works

As you probably understand, our computer can perform several actions simultaneously (in parallel) only if its processor has several cores.

![Image 2](/assets/images/scheduler/image_2.png)
_In the figure, we have two cores, which means that only two tasks can be executed simultaneously, the rest are waiting_

But there is good news — the processor is a very fast thing, so even with one core it can perform a huge number of actions almost simultaneously, quickly switching between them. For example, even on a computer with a single-core processor, you can "simultaneously" watch my videos on YouTube, write code, run it, and perform a number of other tasks. At any given moment, the processor will be performing only one task, but it will switch between them so quickly that it will seem to you like all at once.

![Image 3](/assets/images/scheduler/image_3.png)
_A single-core processor performs only one task at a time, but it seems to us that there are many._

### 3.2 🤔 Thoughts on Event Parallelism

In general, in a sense, parallelism is a function of time, or rather of a time interval. For example, you saw two events that seemed simultaneous to you. But in fact, they were executed with a difference of 0.02 ms. for our brain, these are simultaneous events, but, for example, for a slow-motion camera - no. If the events were executed with a difference of 2 seconds, then even we will notice that they are not simultaneous.

Thus, everything depends on the size of the observation "window". For example, if the operation execution speed is 10ms, and our window is 100ms, then we can say that 10 operations will be executed simultaneously during this time, and no matter in what order. If we take a period of 10ms, then operations will be performed sequentially.

![Image 4](/assets/images/scheduler/image_4.png)
_The camera saw these actions as sequential, and our eyes as simultaneous_

### 3.3 🔄 Concurrency vs Parallelism

Meanwhile, we have come closer to understanding the difference between parallel and concurrent execution.

- **Concurrency** — is about the design of our program. Namely, when we have several processes that can be executed independently of each other, no matter in what order.
- **Parallelism** — is already about program execution, namely — executing several tasks at one moment in time.

At the same time, independent (concurrent) processes can be executed both simultaneously (in parallel) and sequentially. For example, the processor can constantly switch between them, but at any given moment only one process is executed. However, it can switch at any time, and this should not affect the correctness of process execution in any way.

> [!NOTE]
> In other words, concurrency is when we work with many things at the same time, and parallelism is when we do many things at the same time.

In fact, this topic is complex and deep, and its discussion can result in a separate post — by the way, I already have such a post! In it, I explained these concepts in more detail, and also attached links to interesting materials for further study.

In any case, even if you have not fully realized the essence of these concepts now, this is normal — while reading the article, everything will gradually sort itself out in your head and become clearer.

### 3.4 🖥 Parallelism and Concurrency at the OS Level

And who deals with switching tasks and decides which of them will be executed on the core, and which will wait? The **operating system scheduler** deals with this, and it manages **threads**.

**Thread** - is a sequence of commands that are executed within a single process. in simple terms, this is the work that is performed on the processor or waiting for its turn.

Each thread can be in one of three states:
- **Executing**: executing right now on one of the cores.
- **Runnable**: ready for execution, waiting for its turn.
- **Waiting**: not ready for execution, as it is waiting for some event. For example, this may be due to an i/o operation, interaction with the OS (syscall), etc.

![Image 5](/assets/images/scheduler/image_5.png)
_System threads and their states_

### 3.5 ⚡️ Context Switching

The OS scheduler can switch threads at arbitrary times — disconnect executing ones from the processor core and replace them with **Runnable** threads. This is called **context switching**. We cannot influence the scheduler's decisions in any way, nor can we make any predictions regarding this — it makes all decisions independently. This approach is called **Preemptive multitasking**.

For comparison, the Go scheduler has **cooperative** multitasking — this means that goroutines (analogous to threads) themselves decide when to yield resources to their waiting comrades. True, with one caveat that appeared in Go v1.14. We will return to this later.

The most important thing to understand here: **threads are expensive**. Let's figure out why.

**First**, context switching takes quite a lot of time: need to update data in processor caches, save thread state, etc. The more threads we have in the Runnable state, the more often the context will switch, and the slower the program will work. For example, if there are too many switches, then in total they can take as much time as the execution of tasks themselves (or even more):

![Image 6](/assets/images/scheduler/image_6.png)
_Two examples of context switching between tasks. Good: switching happened only once; Bad: switching happened very often, due to which in total we spent about as much time on it as on completing tasks._

**Secondly**, each thread consumes a significant part of our computer's memory: the stack of each thread can often take up to a couple of megabytes (it stores local variables, function call chains, etc.)

> [!IMPORTANT]
> Now, knowing that threads are really expensive, we come to the understanding that we want to create as few of them as possible. The fewer threads, the less often the context switches, the less memory we spend. This is an important point that we will refer to very often in the future.

This is enough for understanding further material, but if you want deeper, then I have a more detailed post on the topic of threads and processes.

---

## 4. 🛠 Designing Our Own Scheduler for Go

So, finally we got to the most interesting part. Let me remind you, our main task is to teach our programs in Go concurrency and parallelism. Let's say we have code like this:

```go
func main() {
    task1() // run the first function and wait for its completion
    task2() // after that run the next one

    // ...
}
```

Functions `task1` and `task2` will be executed synchronously, that is, sequentially — `task2` will only start after `task1` completes. And we want `task2` to start without waiting for `task1` to complete, that is, asynchronously — that very concurrent execution. In terms of syntax, it looks very simple — just add `go` before the function call:

```go
func main() {
    go task1() // start the function and move on without waiting for its completion
    go task2() // start the next one and also move on

    // do some other work..
}
```

Thus, competing processes for us will be functions started with the `go` prefix. And since we have competition for resources — for processor time, we need some entity that will manage this business. Let's call such an entity - **Scheduler**. Here we met this guy!

What will he manage? Typically, the OS scheduler managed threads, let's steal this concept from him! Let's call our analogue of threads — **goroutines**. In a sense, these will be the same execution flows, but inside our program.

> [!NOTE]
> In other materials, you might also have seen the definitions of **kernel space** and **user space**. So — goroutine scheduling happens at the user space level, that is, they are managed by the Go scheduler (more precisely — Go Runtime), and thread scheduling is at the kernel space level, that is, they are managed by the OS.

 Since all our programs run inside the OS, we will still need threads — goroutines will run on them. That is, when we need to execute a goroutine, we attach a thread to it, and it will run in it.

But why do we need to introduce any goroutines at all if there are threads? Because, as we discussed above, threads do not always suit us in terms of work efficiency. We can build goroutines in such a way that they are much more lightweight in memory and, most importantly — so that switching between them is much cheaper than switching between threads. This is possible because they will all live inside one process (our program) and share memory access, and switching between them will occur without the participation of the OS kernel. In fact, there are many tricky optimizations and nuances that we will analyze, but in short — understanding all the limitations of threads, we will build the work of goroutines in such a way as to use these threads most profitably.

**In general, we can say that the Go scheduler is simply a way to optimize the use of system threads — optimize in terms of resource usage, as well as in terms of simplicity of code to work with them. And the further you dive into the article, the more you will come to understand this.**

So, let's summarize all this:
1.  The main resource for executing programs is **processor cores**.
2.  There are few cores, and they have a lot of work, so the concept of **threads** is introduced: threads run on cores.
3.  The OS scheduler manages threads and optimizes work with them so that cores do not stand idle without work.
4.  OS threads are big and scary — not because they were badly designed, but because at the OS level there are limitations and features of work, so few threads are available to us.
5.  There are few threads, and there is a lot of work inside our program, so the concept of **goroutines** is introduced: goroutines run on threads.
6.  The Go scheduler manages goroutines and optimizes work with them in such a way as to create the most efficient use of threads — create as few of them as possible and not allow them to stand idle.

Moving on, our goroutines will obviously need states. Let's borrow them also from threads:
- **Waiting** — the goroutine is not ready to start, as it is waiting for something.
- **Runnable** — ready to start as soon as a thread becomes free.
- **Executing** — executing on some thread.

Thus, we get the scheme:

![Image 7](/assets/images/scheduler/image_7.png)
_Goroutines and their states_

Does it remind you of anything? This is an exact copy of the OS scheduler scheme that I cited above, only the entity names differ. Indeed, in the most basic understanding, an analogy can be drawn here.

Next, we need to introduce two more entities:
1.  **Machine (M)** — will directly execute the goroutine.
2.  **Processor (P)** — will put goroutines (G) into the Machine.

> [!CAUTION]
> Do not confuse this Processor with the computer processor or CPU cores — they have similar names, but the essence is absolutely different (this is not my whim, this is an official term).

Now the picture is as follows:

![Image 8](/assets/images/scheduler/image_8.png)
_The Processor takes a goroutine and puts it into the Machine for execution_

Essentially, a **Machine** is just a system thread. That is, the Processor, placing a goroutine into the Machine, simply binds this goroutine to a thread. And since this is so, I will further call threads exactly threads, not machines, it is more convenient for me. I cited the term "Machine" here only so that you better understand the connection of my article with other materials.

By the way, these are the same **G, M and P** that you might have met in other articles and materials. And together they form the so-called model - **GMP**. It is precisely to it that we will come in the end.

So, we have a set of basic entities. How will we achieve concurrent and parallel execution with their help? Let's invent. We will act iteratively — first we will come up with the simplest naive solution, and then we will complicate it step by step, solving emerging problems.

### 4.1 1️⃣ Create a Thread for Each Call — 1:1 Model

![Image 9](/assets/images/scheduler/image_9.png)
_1:1 Model — a new thread is created for each new goroutine_

The simplest thing that can be invented is to create a separate thread for each goroutine, and after completion destroy these threads as unnecessary. That is, we pass all new goroutines to the Processor, it requests a new thread for each of them, and when the goroutine finishes work, its thread is utilized.

This is called — **1:1 Model**, that is, as many goroutines as threads. The solution is quite working — we will have both concurrency and parallelism, and we will definitely use all processor cores if necessary (because the OS scheduler will not allow threads to stand idle if there are free cores).

By the way, there is also a **1:N Model** — this is when all goroutines (or their analogues) always use only one OS thread. From the reasoning above, we understand that this does not suit us — after all, in this case we lose the opportunity to run anything in parallel. But it is useful to understand that it happens this way too.

![Image 10](/assets/images/scheduler/image_10.png)
_1:N Model — we always have only one thread, regardless of the number of goroutines_

Essentially, the 1:1 model delegates all the work to the OS scheduler. But does everything suit us here? Unfortunately, no. From the introductory educational program, we remember that creating and destroying threads is too expensive, and we want to minimize this. Let's think how.

### 4.2 2️⃣ Thread Pool

Okay, if creating and destroying threads is expensive, let's not destroy them, and instead of creating, if possible, reuse existing ones.
That is, as before, we will create new threads for goroutines, but now only when necessary. And when a thread becomes free (for example, a goroutine has finished its work), the Processor instead of destroying will put it into the pool to hold for other goroutines. The waiting thread will not perform work, which means it will not occupy the processor core.

![Image 11](/assets/images/scheduler/image_11.png)
_Three goroutines in different states: Running - executing on a thread, Runnable - takes a thread from the pool, Done - returns the used thread to the pool_

Great, we optimized our scheduler by getting rid of many expensive and unnecessary thread creation and destruction operations! Is everything good now?
No, it's still bad. Remember - not only creating threads is expensive, but also their very existence, and we want to create hundreds.. no, millions of goroutines! Alas, we cannot afford millions of threads. What will we do?

### 4.3 3️⃣ Limiting Thread Pool Size — M:N Model

If we do not want to create too many threads, it is obvious that we just need to limit their maximum number, let's do just that. The implementation scheme of such an approach is still quite simple — when a new goroutine starts, we do the following:

1.  Check if there is a free thread in the pool. If there is, take it to execute the goroutine.

    ![Image 12](/assets/images/scheduler/image_12.png)
    _There are free threads in the pool, taking from there_

2.  If not, check how many threads are currently in use. If less than the limit we set, then create a new thread and give it the goroutine. When the goroutine finishes work, send this thread to the pool.

    ![Image 13](/assets/images/scheduler/image_13.png)
    _There are no free threads in the pool. We see that earlier we created only 1 thread out of 8 possible, so we create a new one._

3.  If the number of threads has reached the limit, the goroutine will wait until one of the previously created ones becomes free, returning to the pool. As soon as it becomes free, we will pass the goroutine to it for work.

    ![Image 14](/assets/images/scheduler/image_14.png)
    _There are no free threads in the pool, and their limit is exhausted — other goroutines are currently executing on them. The current goroutine will have to wait._

Thus, we have the concept of **waiting goroutines**. But where will they wait? Let's just line them up in a queue, and call this queue **Global Run Queue (GRQ)**. Thus, every new goroutine that did not get a thread will be sent to the GRQ and wait for its hour.

![Image 15](/assets/images/scheduler/image_15.png)
_The goroutine did not get a thread, it is sent to wait in the Global Run Queue_

Since queues are different, it is worth clarifying — we will use a **FIFO queue (first in, first out)**, this is when those goroutines that came first are taken first. It turns out logically and honestly — the earlier the goroutine came, the earlier it receives a thread.
LIFO queue (last in, first out) would be less honest and sometimes problematic. For example, we could have a situation where the goroutine that arrived first is never executed because new goroutines are constantly arriving and executed instead of it. LIFO queues have their pros, but we won't dwell on this now.

#### 📏 Pool Size

Let's now think about what limit on the number of threads we want to set — that is, what size will our pool be? If we use too few threads, then all processor cores will not be involved — for example, if there are 8 cores and only 4 threads, then 4 cores will be idle. But we also don't want to create too many threads, because it is expensive and inefficient. So what value will be optimal?

The answer lies on the surface — we will create exactly as many threads as there are cores available to us, and then not a single core will be idle without work. Creating even more threads makes no sense, because then some of them will definitely stand idle waiting for a free core. That is, if we have 8 cores and 10 threads, then even in the best case, 2 threads will be idle.

![Image 16](/assets/images/scheduler/image_16.png)
_In the diagram, we have two CPU cores and, accordingly, two worker threads. Only two goroutines can be executed simultaneously, the rest are waiting for their turn._

> [!IMPORTANT]
> An important point: now that the number of threads is strictly limited, we can afford to assign a separate Processor (P) to each of them. That is, each Processor will now have its own thread. This does not affect the overall picture yet, but later working with each thread will become more complicated, and this approach will help us a lot.

Thus, we came to a real model called **M:N Threading**. It consists in the fact that we execute N goroutines on M threads.

By the way, we have already come closer to understanding the `runtime.GOMAXPROCS()` function — It sets the maximum number of Processors that our programs will use. That is, by default there will be exactly as many of them as there are cores available — `runtime.NumCPU()`, but if we want fewer, then we can set their exact number as follows:

```go
// Maximum number of processors = 2,
// regardless of the number of CPU cores
runtime.GOMAXPROCS(2)
```

Also, this function returns the current setting, and if you just want to see this value without changing it, you can call with argument 0:

```go
// The maximum number of processors will not change,
// and the function will simply return this value
n := runtime.GOMAXPROCS(0)
```

#### 🔒 Mutex for Global Run Queue

Meanwhile, we have another serious problem. If several parallel processes access a shared resource, then we need an access synchronization mechanism. Otherwise, it may happen that two Processors try to take one goroutine into work and break something — for example, both will start executing this goroutine.

Behavior where several concurrent processes gain access to a shared resource is called - **race condition**. In the general case, it is unpredictable and dangerous. It must not be allowed in any case.

To avoid this, we use the simplest synchronization primitive — **mutex (lock)**. If you are not familiar with it, I advise you to definitely get acquainted — this is a very simple but important mechanism that is very common. In simple terms, with its help, each Processor will be able to temporarily block the queue to work with it — in our case, to get a goroutine out of there. While it is blocked, no one else can interact with it. Thus, while one Processor takes a goroutine from the GRQ, the other Processors will have to wait.

![Image 17](/assets/images/scheduler/image_17.png)
_The Processor on the left cannot take a goroutine from the GRQ because the Processor on the right blocked it for the time of its work._

By the way, we are already starting to approach the real Go scheduler, it also:
- Has a Global Run Queue with a mutex.
- The number of worker threads matches the number of CPU cores.
- Each processor has its own thread.

You ask: Nikolay, do you really not like this scheme either? Everything works perfectly! Yes? No! Rather, not really. If our computer has only 2 cores, then everything is super, the scheme will work well. But what if there are more cores? For example, 16, 32.. 64? With the growth of the number of cores, the number of threads will also grow, but the queue remains one for all.
That is, even if we have a lot of Processors, only one of them can work with the queue at any given moment. Introduction of the mutex solved one of our problems, but created another — working with a shared resource (GRQ) slowed down significantly.

> [!TIP]
> By the way, here you can realize a very important point in any systems — in the design of any systems there are no free optimizations, these are almost always trade-offs: optimized CPU work, sacrificing memory; optimized memory, but increased load on CPU, etc.

And, as usual, I also have a more detailed post about this.

How will we solve this problem?

### 4.4 4️⃣ Local Queues

To solve the blocking problem.. Need to get rid of locks! And for this, we need to make sure that processors do not access the shared queue. To achieve this, we will give each Processor its own queue — **Local Run Queue (LRQ)**. Since only one processor will have access to each LRQ, blocking will not be required.

At the same time, we do not get rid of the GRQ with its mutex, it will still be useful to us — we will put goroutines that are not yet attached to any of the Processors there, and from there goroutines, ultimately, always get into the LRQ, and only after that are executed. We will return to this later.

Now our entire scheme looks like this:

![Image 18](/assets/images/scheduler/image_18.png)
_In this image we have CPU cores and each of them corresponds to its own thread, its own Processor and its own LRQ. Also depicted here is the GRQ, which is not attached to any of the Processors._

In the resulting scheme, we see goroutines in the Runnable state (located in the queue) and Running (executed by the Processor). Did we forget anything? And where are the Waiting goroutines? Let me remind you, these are those goroutines that are blocked by something and cannot be started. So where to put them? Maybe we need to start a separate queue in the Processor for them — Wait Queue?

In fact, the Processor will not have to watch waiting goroutines, other entities will do this. For example, if a goroutine is blocked due to reading or writing to a channel, then it will fall into the Wait Queue of this channel, and from there it will fall into the GRQ and LRQ. I have a very detailed video about channels, where I, among other things, analyze this mechanism. Similar mechanisms also exist for mutexes, timers and Network IO.

#### 📋 Queue Check Order

Now we have two types of queues - LRQ and GRQ. Accordingly, each Processor can take work in two places, which means we need to decide in what order it will do this.

It is obvious that the Processor will first look into the LRQ, because it is the fastest, we introduced it for this very reason. But if it is empty there, then next the Processor will look into the GRQ.

True, another problem may arise here — if the LRQ of all processors is constantly replenished, then Processors will always take work from there, and will never reach the GRQ. This is bad because goroutines in the global queue will start to stagnate, so we should sometimes check it out of turn. But how often? Obviously, once out of 61! That is, the algorithm for receiving a goroutine by the Processor is as follows:

1.  1/61 times we check the GRQ, and if there are goroutines there, we take from there.
2.  If not, check LRQ.
3.  If not there either, check GRQ.
4.  ... There will be more points here, but we'll talk about them later.

> [!NOTE]
> You ask — but why 1/61, what is the magic number? The most important thing here is — it is prime, and this helps to avoid synchronization of checks between different Ps, distributing them more evenly. In addition, it is not too big and not too small — that is, the GRQ will be checked not too often and not too rarely.

### 4.5 5️⃣ Work stealing

So, our scheduler is getting more complicated, but it works better and becomes more and more like a real one. However, it is still far from ideal. What problem worries us this time? You probably remember that from the very beginning we were very worried about the problem of idle cores. It would seem that we got rid of this problem a long time ago, but no. Look at this diagram:

![Image 19](/assets/images/scheduler/image_19.png)
_Goroutines are sad because they have to wait, although we have a free Processor. And we are sad because we have a whole processor core idle without work._

So, one Processor has a lot of work, and another has none at all. How to be? Let's teach Processors to steal work from each other! Namely, if the LRQ of one Processor is empty, then it will peek into the LRQ of any other Processor and take a goroutine from there. Or even better — let it take half of the goroutines from there at once, so as not to walk then once again.

![Image 20](/assets/images/scheduler/image_20.png)
_Processor #2 was out of work, so it stole half of the goroutines from the neighbor. Now he has something to do!_

This way we will automatically balance the work between all our Processors. And the full algorithm will now look like this:
1.  1/61 times check GRQ, and if there are goroutines there, then take from there.
2.  If not, check LRQ.
3.  If not there, try to steal from another Processor.
4.  If failed, check GRQ.
5.  ... Here we will check Network Poller, but more on that later.

### 4.6 6️⃣ Handoff

How good our Scheduler has become! But let's find a problem here too. What if the goroutine performs some blocking operation, for example a system call (**syscall**)? Then this goroutine will block a whole thread for us, and the processor core will again stand idle, despite the fact that we have Runnable goroutines in the queue.

![Image 21](/assets/images/scheduler/image_21.png)
_The goroutine performed a syscall and blocked the thread - the CPU core stands idle without work_

Once again I draw your attention — not only the goroutine with the Processor are blocked, but the thread itself is also blocked by the system call. That is, we cannot simply send this waiting goroutine to wait somewhere else, giving the thread some other work. Therefore, we will have to send them somewhere together — both the thread and the goroutine.

Here it is important to understand what a system call is in general (syscall). If you do not know — this is another must have topic for your piggy bank, be sure to figure it out deeper.

In the meantime, let's do this — imagine that you called a colleague and asked him a question. The question was difficult, the colleague thought deeply, and you are forced to wait. While he is thinking, you cannot do other useful things, you are blocked by waiting. You will be free as soon as the colleague comes up with an answer and reports it to you.
So, a call to a colleague — this is the syscall, and you are the thread that performs it.
So, if a thread is blocked by a system call, it can no longer perform work. In this state, it is not very useful to us, so we simply untie it from the Processor, create a new thread and bind it to the Processor. Such a mechanism is called - **handoff**.

![Image 22](/assets/images/scheduler/image_22.png)
_While thread #1 awaits a response from the system call, its Processor will be bound to another free thread_

Returning to the analogy with a question to a colleague — while you serve chatting with him, your boss asks to free up the work computer so that another colleague can work on it (yes, your company has hard times and not enough computers for everyone...)

Well, the mechanism is not bad, but it is quite expensive, since because of it new threads are constantly created, and we tried very hard to avoid this. Unfortunately, here it is inevitable — a large number of syscalls will always generate a large number of threads. However, we can optimize the process a little.

#### 👀 Sysmon

Some system calls are short-lived, that is, they block threads for a very short period of time. Creating a separate thread for them every time is unprofitable, so we will make such a tricky optimization:
- If we know that the system call will block the thread for a long time, then we immediately perform a handoff.
- In other cases, we will allow the thread to remain in a blocked state for some time, periodically checking if it has become free. If it exceeds a certain timeout (namely, 10ms), then we also start a handoff.

The process that performs these checks will run constantly in the background, and it is called - **Sysmon**.

![Image 23](/assets/images/scheduler/image_23.png)
_Sysmon monitors the time, and if the syscall drags on, it initiates a handoff_

By the way, on which thread will Sysmon itself run? Obviously, we cannot entrust this to threads that run goroutines, since they can fall asleep at any moment (and this is exactly what we are trying to track). And they have enough of their own affairs without that. It turns out that we need to start a separate thread for monitoring, which will perform its work independently of worker threads.

So, wait.. We agreed earlier that we limit the number of threads - there should not be more of them than the number of cores. Yes, this is still in force, but here it is worth clarifying that this agreement concerns only active threads processing goroutines. That is, this does not apply to special threads performing special tasks, as well as threads in the Waiting state. So, everything is fine here, there are no contradictions.

So, we untied the goroutine together with the thread and left them somewhere to wait for the return from the system call. What will we do when they wait? First of all, for optimization reasons, we want to return the goroutine to the processor from which it was taken. But only provided that it is now free, that is, not busy executing another goroutine. If it is busy, we will look for other free processors. If there are no free ones, we will send the goroutine to GRQ.

By the way, we can say that now we have moved from the **M:N Threading** model to the **M:P:N Threading** model. That is, we still have N goroutines and M threads, but also we have P processors. Of course, we introduced the processor entity a long time ago, but previously the total number of used threads and processors matched, but now no.

Handoff is a cool mechanism, but it still inevitably leads to the creation of new threads. Can we somehow optimize the work of the scheduler so as to resort to using it less often? Yes, we can - if the syscall is capable of executing asynchronously.

### 4.7 7️⃣ Network Poller

Blocking a thread during a system call is a limitation at the operating system level, that is, programmatically we cannot fight this in any way. But, fortunately, the OS itself usually provides mechanisms with which this limitation can be bypassed. Namely, these are mechanisms for performing asynchronous system calls, for example: epoll (Linux), kqueue (MacOS, BSD), IOCP (Windows). Most often, this concerns network operations. We will not go deeply into these mechanisms here, but if you have such a desire, you already know where to dig.

Let's figure out what the point is. Previously, our threads acted like this:
1.  System call is made.
2.  Thread blocked for waiting time for response.
3.  After receiving response, thread unblocked.
4.  Continuation of work.

The mechanisms listed above allow doing like this:
1.  Thread initiates system call and goes about its other business. System call will be registered in a special system, and we can return to it later.
2.  Periodically check if response for system call has arrived.
3.  Thus, work with system call occurs asynchronously.

![Image 24](/assets/images/scheduler/image_24.png)
_We register syscall in epoll, thanks to which our thread remained free and can execute other goroutines_

To make it even clearer, let's return to the example with a call to a colleague from the section about Handoff.
Imagine that instead of a phone call, you decided to send a colleague a message in the messenger. Now you do not need to waste time waiting, you just minimize the messenger and continue to do other things. You just need to periodically open the dialog box with a colleague to check for an answer. This is precisely the essence of the mechanisms described above. That is, you are the thread, your dialogue with a colleague is the goroutine, and the messenger is this very mechanism (epoll, etc.)
Unfortunately, not all syscalls can be executed asynchronously — imagine that your colleague cannot simply give an answer to a question, you need to actively discuss this topic. In this case, you simply cannot wait, you will be blocked anyway.

So, we got a cool tool, and now let's teach our scheduler to use it. First, we will need a component that will keep track of system calls that need to be checked for answers — let's call it **Network Poller (netpoller)**.

If a goroutine intends to perform a system call that can be executed asynchronously, we, instead of blocking the thread:
1.  Register the operation in the Network Poller.
2.  Transfer the goroutine to the Waiting state and pass it to the Netpoller.
3.  The Processor becomes free to execute other goroutines.

![Image 25](/assets/images/scheduler/image_25.png)
_The goroutine performed an asynchronous syscall, and we passed it to the Network Poller_

Once again I emphasize the main feature of this mechanism — waiting goroutines do not occupy a thread, and can be in this state as long as necessary, without slowing down the program in any way.

When the syscall ends, the goroutine switches back to the runnable state, that is, it is ready to continue work. When Processors run out of work, they themselves will contact the Netpoller to take free goroutines for themselves. Total, the work search algorithm for each Processor looks like this:
1.  1/61 times check GRQ, and if there are goroutines there, take from there.
2.  If not, check LRQ.
3.  If not there, try to steal from another Processor.
4.  If failed, check GRQ.
5.  Check Network Poller.

![Image 26](/assets/images/scheduler/image_26.png)
_The Processor ran out of work and went to the netpoller to see if there were any liberated goroutines there that could be taken back into work._

There is still a small problem here — it may turn out that Processors always have enough work, and they stop checking netpoller at all, or do it very rarely. This is not very fair to the goroutines waiting there, so we will teach Sysmon to make such checks in the background — if no one has addressed the netpoller for more than 10ms, Sysmon will do it. At the same time, it will move the liberated goroutines to the GRQ, and then the Processors themselves will perform balancing using the mechanisms we discussed above (remember that GRQ is checked out of turn once per 1/61 iteration).

Now let's emphasize once again that not all syscalls support asynchronous work. For example, network operations support, and file operations — no. How can this knowledge be used in practice? Very simply:

```go
// This will work via netpoller
conn, _ := net.Dial("tcp", "example.com:80")
data, _ := conn.Read(buf)

// And this will block the thread
file, _ := os.Open("bigfile.dat")
data, _ := file.Read(buf)
```

Therefore, when designing Go applications, it is important to understand this difference:
1.  For I/O operations with the network, you can safely use goroutines — they will be effectively processed through netpoller (there are also limitations here, but different ones).
2.  For operations with files, you need to be much more careful, as they will block threads.

> [!NOTE]
> In fact, in both cases we come to the conclusion that it is better for us to use a worker pool, but for different reasons and with different limitations.

Well, one problem less, so we move on and solve the next one.

### 4.8 8️⃣ Goroutine Execution Order

So, we came up with a very complex mechanism for distributing goroutines across Processors. It already works quite well, goroutines are distributed very efficiently. Now let's discuss in what order processors will execute their goroutines from the LRQ.
The simplest thing that can be invented here is simply to execute all goroutines in turn (the word "queue" will be dreamed of by you today!). That is, the Processor takes the first goroutine from the LRQ, executes it, then takes the next one, executes it, and so on. If our goroutines are short-lived, then this will work well.

Do you already understand what the problem is here? And what if we have long or even eternal goroutines? For example, nothing prevents us from starting a hundred goroutines that will periodically check something throughout the entire operation of the program. But since we have a limited number of cores and, accordingly, Processors — say, 8 pieces, only 8 will be executed from all these goroutines, the rest will wait forever.

![Image 27](/assets/images/scheduler/image_27.png)
_The current goroutine has been executing for a very long time, and the rest are forced to stand idle, waiting for it_

This is bad, because each goroutine must receive its share of processor time. How will we solve? Very simply — at certain points in time, the goroutine will check if it is time for it to interrupt and let others work.

It remains to decide at exactly what moments it will do this and how it will understand that it is time for it to rest. Let's start with the first — checks must be performed at those moments that are safe and optimal for eviction. Best of all, these points suit us:
- **Before function call (prologue)**: When a function is called, a stack frame is created (a dedicated area in memory where local variables, return address and other function data are stored). When interrupting at the moment of a function call, all necessary information is already saved in the frame, so the goroutine can be suspended without risk of losing data. When execution resumes, the goroutine will start from this function call, using the same variable values.
- **When performing blocking operations** (our favorite syscalls, timers, etc.): at these moments, the goroutine will still wait and stand idle, so let someone else work.

Okay, here the goroutine hit the check point. And by what sign will it understand that it is time for it to yield the Processor? Let's do it as simply and clumsily as possible — let's give each goroutine a **stackguard** flag. When we need to interrupt a goroutine, we will set this flag to a special value **stackPreempt**. At the next check, the goroutine will see this value and understand that it is time for it to transfer control to the scheduler.

Now let's figure out who and when will inform the goroutine about the need to interrupt. And here our old acquaintance — **Sysmon** — will help us again. Let's add one more task to him. In addition to checking blocked syscalls, he will also check the execution time of active goroutines. If any of them works too long (more than 10ms), Sysmon will set its flag `stackguard=stackPreempt` so that at the next check it interrupts.

![Image 28](/assets/images/scheduler/image_28.png)
_The goroutine checks the stackguard value set by Sysmon. If it equals stackPreempt, then it transfers control to the scheduler_

This approach to scheduling is called **cooperative multitasking** — this is when processes themselves decide when to give control to others. Unlike preemptive multitasking of the operating system, where a thread can be interrupted at any moment, here goroutines themselves control the moment of control transfer.
Is everything good here? In general, it is already quite good — in this form, the Go scheduler existed for a long time and was called cooperative. But all this time there was one well-known problem, which was finally solved only in Go version 1.14.

---

## 5. 🛑 Forced Eviction of Greedy Goroutines in Go v1.14

Let's imagine the following situation. Suppose we have only one Processor — this is easy to achieve using the `runtime.GOMAXPROCS` function, and it processes a goroutine. This goroutine actively counts something and does not want to transfer control — for example, performs some calculations in an infinite loop:

```go
func main() {
    // Set the number of processors = 1
    runtime.GOMAXPROCS(1)

    // launch a "greedy" goroutine that will capture the processor
    go func() {
        sum := 0
        for {
            sum++ // infinite loop without function calls
        }
    }()

    // Launch time.Sleep to transfer control to greedy goroutine
    // and not exit main() ahead of time
    time.Sleep(time.Second)
}
```

We have a serious problem — the goroutine blocked the entire Processor and is not going to give it back! It doesn't call functions, which means it doesn't check the stackguard, it doesn't perform blocking operations, it just steals all resources, and our Sysmon can't do anything about it.

![Image 29](/assets/images/scheduler/image_29.png)
_Sysmon asked the goroutine to yield the Processor by setting the stackguard flag, but the goroutine is busy working and does not check it_

How can we be? We cannot reach the goroutine in any way — no matter how we transmit messages to it, it simply does not read them, dealing with other work without stopping. We need some mechanism that will interrupt it itself from the outside. Fortunately, the OS provides us with such a mechanism — these are **signals**.

**Signals** — is a mechanism for notifying processes about any events at the level of the OS. The key moment for us — when sending a signal, the OS interrupts the execution of the process and starts a special handler that the process itself established. It is this handler that will save the state of the goroutine and return control to the scheduler.
As you might have guessed, I also have a detailed post on this topic.

In Unix systems, the `SIGURG` signal is used for this. It was chosen because it is rarely used in ordinary programs (initially intended to notify about urgent network events), which means we can safely use it for our purposes.
Windows, by the way, is a separate story, and I absolutely do not want to dive into it (there is not even a separate detailed post this time...). Just keep in mind that the implementation is somewhat different, but the essence is about the same.

Now that we have a suitable mechanism, we can do the following. If a goroutine does not voluntarily transfer control for too long (~10ms), then we send it a signal and interrupt it forcibly.

**Now we have two ways to evict goroutines:**
1.  **Main** — through checking `stackguard` at safe points. The goroutine itself decides when to stop.
2.  **Spare** — through operating system signals, if the goroutine does not want to stop itself.

True, the second method does not always work. In some parts of the code, interrupting a goroutine is dangerous — for example, during garbage collection execution or execution of system calls. In such cases, we assume relying only on the first method.

It is this mechanism that was introduced in Go version 1.14, and the scheduler became **preemptive-cooperative**. Most likely, you heard about this, but, perhaps, poorly understood or even forgot. But now you even know how it works.

> [!NOTE]
> By the way, do not confuse this mechanism with what we discussed in the section about handoff. There Sysmon also checks the thread running time (also with a timeout of 10ms), but with a different purpose — it checks if the thread has become free from the system call, and if not - performs a handoff. In that case, we are interested only in blocked threads, and here we are watching actively working goroutines that do not want to give up control voluntarily.

---

## 6. 🏁 Summary

Well, here we have built an effective scheduler — it uses available processor cores as efficiently as possible and at the same time minimizes the number of operating system threads.

And now it is very useful to look at and examine the final scheme entirely:

![Image 30](/assets/images/scheduler/image_30.png)
_General diagram of the internal structure of the Go scheduler_

Words of gratitude to Nikolai Tuzov! - link - https://habr.com/ru/articles/891426/
