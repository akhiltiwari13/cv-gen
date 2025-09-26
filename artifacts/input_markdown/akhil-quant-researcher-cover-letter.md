# **Akhil Tiwari**
Delhi, India | +91 8959815625 | [Email](mailto:akhiltiwari.13@gmail.com) | [GitHub](https://github.com/akhiltiwari13) | [LinkedIn](https://www.linkedin.com/in/akhiltiwari-13/) | [Website](https://quomptrade.com/)

---

## **Cover Letter**


Hello Team,

I am writing to express my strong interest in the Quantitative Researcher position at [Company Name]. With over 10 years of experience building ultra-low latency trading systems and a deep understanding of market microstructure gained from engineering perspective, I am excited to transition my technical expertise toward quantitative research and alpha generation.

**Unique Perspective from the Execution Side**

My extensive experience building trading infrastructure has given me an intimate understanding of how markets actually work at the microsecond level. At AlgoQuant, I architected systems achieving sub-microsecond tick-to-order latency, requiring deep analysis of market data patterns, order flow dynamics, and execution optimization. This hands-on experience with real market data at the most granular level provides me with unique insights into market inefficiencies and potential alpha opportunities that pure researchers might overlook.

At Morgan Stanley, while developing smart order routers for multiple exchanges (Bovespa, BMV, TSX) with ~27μs median latency, I implemented sophisticated algorithms for order placement optimization and developed wash trade detection logic. This required extensive statistical analysis of trading patterns and deep understanding of market impact models—skills directly transferable to quantitative research.

**Technical Foundation for Research**

My technical toolkit aligns perfectly with modern quantitative research requirements:
- **Advanced C++ and Python proficiency** for high-performance backtesting and simulation
- **Experience with massive data processing** through my work with tick-by-tick market data feeds
- **Statistical analysis capabilities** demonstrated through latency profiling and performance optimization
- **Distributed systems expertise** from building 16k TPS distributed ledgers, essential for large-scale research infrastructure

At Edelweiss, I built a comprehensive .pcap-based latency profiler that required statistical analysis of millions of network packets to identify performance bottlenecks—essentially conducting research to optimize trading strategies at the infrastructure level.

**Bridging Engineering and Research**

What sets me apart is my ability to not just understand theoretical models but also their practical implementation constraints. My experience spans the entire trading stack—from exchange connectivity to risk management systems—giving me a holistic view of how research translates into profitable trading strategies. This perspective would be invaluable in developing strategies that are not just theoretically sound but also practically implementable with real-world constraints like latency, market impact, and execution costs.

My work on blockchain protocols at Deqode, including designing a 16k TPS distributed ledger and implementing tokenomics models, demonstrates my ability to work with complex mathematical models and economic systems—skills directly applicable to developing trading strategies and market models.

I am particularly excited about [Company Name]'s [mention specific research area/approach if known] and believe my unique combination of deep market structure knowledge, technical expertise, and systems thinking would add significant value to your research team. I am eager to apply my engineering rigor to the challenge of discovering and exploiting market inefficiencies.

Thank you for considering my application. I would welcome the opportunity to discuss how my unique background can contribute to [Company Name]'s quantitative research efforts.

Sincerely,  
Akhil Tiwari
[Quomptrade](https://www.quomptrade.com/) 


### **Low-Level Performance Optimization & Kernel Tuning**

Throughout my career building ultra-low latency trading systems, I've demonstrated extensive systems programming expertise in C++:

**Linux Kernel Optimization (Edelweiss, AlgoQuant, Morgan Stanley)**
- Achieved **~200ns latency** at Edelweiss's DMA platform, requiring deep kernel-level optimizations including:
  - CPU affinity and core isolation to eliminate context switching
  - NUMA-aware memory allocation for optimal memory access patterns
  - Kernel bypass techniques for network I/O
  - Real-time scheduling with SCHED_FIFO/RR policies
  - IRQ affinity tuning to dedicate network interrupts to specific cores

**Network Stack & Packet Processing (Edelweiss)**
- Developed **.pcap-based latency profiler using libpcap++**, demonstrating:
  - Raw socket programming for packet capture at kernel level
  - Understanding of Linux networking stack internals
  - Bypass of standard network stack for ultra-low latency packet processing
  - Custom packet parsing and timestamping at nanosecond precision

**Memory Management & Cache Optimization (Multiple Roles)**
- Implemented **lock-free data structures** requiring deep understanding of:
  - Memory ordering and barriers (acquire-release semantics)
  - Cache coherency protocols (MESI/MOESI)
  - False sharing elimination through cache-line aligned allocations
  - Custom memory allocators for deterministic latency
  - Huge pages configuration for TLB optimization

**Systems-Level Debugging & Profiling**
- Extensive use of **Valgrind, GDB, and Tracy** for:
  - Kernel-space debugging and tracing
  - Performance counter analysis (cache misses, branch mispredictions)
  - System call tracing and optimization
  - Memory leak detection at system level

**Multi-Platform Systems Experience**
- Worked across **RHEL, Ubuntu 22.04, and Arch Linux**, showing:
  - Kernel module compatibility across distributions
  - System library optimization for different kernel versions
  - Build toolchain configuration (GCC 13.1, Clang) for optimal binary generation

**Real-Time Systems Programming (AlgoQuant)**
- Built **event-driven trading platform with sub-microsecond latency**:
  - Likely used kernel bypass technologies (DPDK/RDMA)
  - Implemented busy-polling to avoid kernel scheduling overhead
  - Zero-copy techniques between user and kernel space
  - Custom interrupt handling for minimal latency

**Distributed Systems Architecture (Deqode)**
- Designed **16k TPS distributed ledger**, requiring:
  - Inter-process communication optimization (shared memory, message passing)
  - Process synchronization primitives
  - System-level resource management across nodes
