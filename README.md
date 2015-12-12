IRC Connector
--------------

IRCC makes it much cleaner to interact with IRC servers.
IRCC wraps the TCP connection and IRC protocol with HTTP and easy to parse JSON.

It comes in two flavors, CLI and HTTP server:
- The CLI flavor operates as a command line program. It communicates with JSON over STDIN and STDOUT.
- The HTTP server flavor operates as a web server. It receives commands on HTTP requests and responds with JSON webhooks to the originator.

IRCC is still in active development and is designed specifically to meet the needs of Chatterbox.
As work progresses, documentation on usage will be documented.
