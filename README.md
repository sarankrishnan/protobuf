# protobuf

Protobuf is a simple project with the Google protocol buffers. 


It has a sender and receiver program. These two programs talks via TCP socket. Serialization happens using the Google protocol buffers. 
`msg.proto` defines the template of the data type to be serialized. Code generated after compiling the .proto file should be imported by both (sender and receiver) program. 

## Build

To build this project, execute the *build.sh* shell script. Following compilers are the pre-requites for successful compilation. 
* Golang compiler
* protoc
* protoc-gen-go

You should see two binaries (sender and receiver) in the path project cloned directory. 

## Run

Invoke receiver binary with Host:Port argument of your choice. Next, invoke the sender binary with Host:Port on which the receiver listens. 

## Output

Receiver binary should print `index:10 status:"Hello"`
The program hard codes this data. Please customize as per your needs. 
