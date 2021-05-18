# Chat Server Using Go Final Project

## Introduction

**This project is heavily inspired by a source found on scotch.io (Ed Zynda).**
This app is a realtime chat server created using Go. It uses the Websockets protocol and its frontend is written using HTML, CSS, and VueJS.

Websockets is useful to enable a two-way communication between a client and a remote host to receive communication from each other. This allows the client to send messages back and forth to the server without having to refresh the page and enabling messages to appear real-time.

VueJS is a progressive framework that builds UI and it allows you to extend the capabilities of HTML code, mainly through encapsulation.

Another useful application used was Gravatar. This was used to encrypt emails to keep them private, but still use them as a unique identifier. 
Only needed to add a few lines to the js file:
```
        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
```
## Simple Chat
This server, called Simple Chat, has 3 files in a public directory and 1 file in the src directory. The public directory holds the HTML, CSS, and JS files and the src directory holds main.go.
In main.go, we imported the necessary packages, such as log, net/http, and websockets. Then we created a few global variables to connect the clients (using Websockets) and create a broadcast channel. in func main(), a simple file server was made that is able to view our public directory and is able to handle Websocket requests. Next, we created a few goroutines such as handleMessage and handleConnections that makes sure we can add clients to our local server.
In the public directory, the main things to note in the HTML file is that we needed to import all the libraries required such as Vue, jQuery, MD5 (for Gravatar). Also app.js is where the majority of the formatting and message parsing occurs. 

## User Interface
Run main.go in the terminal
```
go run main.go
```
Navigate to http://localhost:8000/
Add your email and username. Open another tab and navigate to the same website. Add another username and email. Now you can send messages locally.

![ChatServer2](https://user-images.githubusercontent.com/32800667/118669708-fa40f980-b7c3-11eb-8c96-0deb77af9815.png)

![ChatServer1](https://user-images.githubusercontent.com/32800667/118669720-fd3bea00-b7c3-11eb-9c31-ad076c499653.png)


## Resources
**Code**: https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets  
**VueJS**: https://vuejs.org/ (Navigate to "Why VUE.JS")  
**Go Tutorial**: https://tour.golang.org/welcome/1


