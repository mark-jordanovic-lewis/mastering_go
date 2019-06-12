# Mastering Go
Just trying to get good at a pretty cool language. Here's some stuff to copy and paste.

### A router
```go
router = httprouter.Router()

router.GET('some/route', respond200JSON(`{"put": "stuff here, config is always handy"}`)

func respond200JSON(inJson string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf(  `{"success": %s}`, inJson)))
	}
}

serverCloseChannel = runServer(&router, ":9000")
defer func() { serverCloseChannel<-true }()
```

### A server
```go
func runServer(router *http.Handler, port string) chan bool {
	var closer = make(chan bool, 1)
	go func() {
		var server = http.Server{
			Addr:           port,
			Handler:        *router,
			ReadTimeout:    2 * time.Second,
			WriteTimeout:   2 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Error occurred running test server: ", err)
			return
		}
		select {
		case <-closer:
			if err := server.Close(); err != nil {
				fmt.Println("Error occurred closing test server: ", err)
			}
			time.Sleep(50*time.Millisecond)
			return
		}
	}()
	time.Sleep(50*time.Millisecond)
	return closer
}
```
