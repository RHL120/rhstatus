# rhstatus
A status bar for dwm written in Go
## Configuration
### Applets
In order to add a new applet, add a new **Applet** struct to **Applets** in
applets/applets.go.
Applets are rendered on to the status bar by order (The top applet on the left
and the bottom applet on the right).
The **Applet** struct has a field **function**
which is a go function that takes in a bunch of arguments of any type and returns
a string. The returned string get rendered. To add a shell command applet use the
cmdApplet function. Use the audio applet as an example.
### Server
If you want to change the server port, change the **port** constant in api/api.go
### Commands
To add a new command, go to api/commands.go and add a **command** struct to the
**commands** map. the **command** struct contains 2 field **argCount** and
**function**. **function** is a function that takes a slice of strings and returns
an error, it will be called on command execution. **argCount** is the length of
the args slice expected by **function**. The length will be checked before
**function** is called so no need to check for that in **function**.
### Sleep time
In order to change the frequency of status bar refresh change the **sleepTime**
constant.

## Usage
### Server
In order to communicate with the status bar send commands to localhost : **port**
Commands start with the name of the command and then a list of arguments seperated
by spaces. After sneding the command send a newline. After that you should recive
either an error or a success message.
## Compatibility
I have tested this on Linux and OpenBSD
