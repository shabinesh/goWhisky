**TODO**

- Pass the python application import path the program

  ```
  $ ./pygo --app app.application
  ```

  
- A HTTP go server to handle traffic, The Go web handler invokes the
  python application callable with environ & start_response callable. 

- WSGI wrapper - Explore GIL &
  goroutine

# TODO
* Create a `wsgi.input` and `wsgi.error` streams that fits specs at PEP333
* Create the Go to handle socket connections.