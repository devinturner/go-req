
## Workflow

1. Client created
2. Client performs GET request and returns Response
3. Parser created
4. Parser decodes response and prints string

Client
  + RequestHeaders map[string]string
  - Get(string) Response

Response
  + StatusCode int
  + Headers map[string]string
  + Body io.ReadCloser

Parser
  - Parse(Response) error
