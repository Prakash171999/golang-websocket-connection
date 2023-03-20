#1. cross-domain problems

This is a security restriction that prevents requests being made from one origin to another.
For example, it will prevent an https:// page hitting an http:// address because the protocol is different.
It will stop example.com calling another.com because it is a different domain.
It will stop www.example.com calling subdomain.example.com because it is a different sub domain.
And it will stop example.com:80 calling example.com:8080 because it is a different port.
It is possible to make cross-origin requests either using JSONP (if you trust the server!) or using a CORS request (Cross-Origin Resource Sharing), which both client and server must agree to (I can supply more details if you need it on this).


#2.

var upgrader = websocket.Upgrader{
	//Solve cross-domain problems	
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} //use default options

The above code is defining an instance of the Upgrader struct from the websocket package in Go. The Upgrader struct is used to upgrade a normal HTTP connection to a WebSocket connection.
The code is setting the CheckOrigin field of the Upgrader struct to a function that takes an http.Request object and returns a boolean value. This function is used to solve cross-domain problems that may arise when establishing a WebSocket connection. By setting it to true, the Upgrader allows connections from any origin.
The code is also using the default options for the Upgrader struct, meaning that no other options are being set. The Upgrader can be customized further by setting other fields such as ReadBufferSize and WriteBufferSize to specify the size of the read and write buffers used by the WebSocket connection.
