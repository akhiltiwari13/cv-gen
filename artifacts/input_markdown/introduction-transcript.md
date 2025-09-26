# Transcript

### Video Recording Transcript
Hello, I am Akhil, 
I am a Lead-Developer and Architect of Ultra-Low latency Institutional grade trading platforms for traditional and crypto based contracts.
I've spent over a decade developing such systems for a plethora of organizations spanning from global banks and brokerage houses on the ==sell side to hedge-funds and prop-shops on the buy side.

My entire career has centered on building **highly concurrent, multi-threaded, event-based real-time trading systems** where event coordination and thread-synchronization are critical for achieving microsecond-level performance.
I have developed/deployed and maintained event-driven architecture with dedicated threads for market data, order management, and risk calculations, all communicating via lock-free queues for the below firms:
- Edelweiss
- Morgan Stanley
- Gemini
- AlgoQuant


**Practical Real-time Trading Platform Implementations:**
- **AlgoQuant**: - **Morgan Stanley**: Designed thread pools for parallel order routing across multiple exchanges
- **Edelweiss**: Implemented Messaging based reactive-patterns with microsecond-level thread coordination
- **Gemini**: Managed concurrent access to order books using reader-writer locks with writer priority

**Thread Architecture & Synchronization:**
- Designed **thread-per-core architectures** with CPU pinning for deterministic latency
- Implemented **lock-free MPMC/SPSC queues** using atomic operations for inter-thread communication
- Built **wait-free algorithms** for market data distribution to multiple consumers
- Developed **custom spinlocks** optimized for low-contention scenarios
- Used **memory barriers and fences** (std::memory_order_acquire/release) for correct synchronization
- Implemented **hazard pointers** for safe memory reclamation in lock-free structures

