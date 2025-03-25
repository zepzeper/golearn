1. Parse and normalize the starting URL
2. Create a HTTP client with proper timeouts and redirect handling
3. Initialize the URL manager with the base domain
4. Create a buffered channel for page URLs to process
5. Create a results channel for link check results
6. Spawn a configurable number of worker goroutines
7. Add the starting URL to the page channel
8. Process pages concurrently:
   a. Fetch the page content
   b. Parse the HTML
   c. Extract all anchor tags
   d. For each link:
      i. Normalize the URL
      ii. Check if it's already visited
      iii. If same domain, add to page channel for processing
      iv. If external domain, just check if it's alive
9. Collect and report results
