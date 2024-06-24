
Hashcash is used because the algorithm is popular and lightweight. If there are any specific requirements, other algorithms can be used. ChatGPT will provide a list.

What was simplified and can be improved:

- We do not use context. The server cannot be stopped, only killed.
- The challenge and result verification are performed within a single connection. It can be done across different connections, but then a cache for storing challenges would be needed (Redis would be a good fit for this).
- Use slog and other tech improvments