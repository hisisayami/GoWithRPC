# inventory

1. command to generate protobuf files
#protoc --go_out=. --go-grpc_out=. ./proto/inventory.proto
(Note: to run this command need to be in inventory/internal directory eg: /Users/rijanprajapati/Documents/GITHUB/GoProject/inventory/internal)

 2. initilizing go mod

 #go mod init example.com/go-inventory-grpc

 3. generate protobuf
protoc ./internal/proto/inventory.proto --go_out=./internal/endpoint --go_opt=paths=source_relative --go-grpc_out=./internal/endpoint --go-grpc_opt=paths=source_relative ./internal/proto/inventory.proto

(Note: to run this command need to be in /inventory directory eg: /Users/rijanprajapati/Documents/GITHUB/GoProject/inventory)

 4. Now, let’s get the module/ library needed for this project.
#go get entgo.io/ent/cmd/ent
#go get -u github.com/gorilla/mux

5. Initialize the Entity
#go run entgo.io/ent/cmd/ent init {{User}}
#User is table name

6. after adding field in the ent file use this command to generate ent again with added column or field

#go generate ./ent

7. install protoc 

# brew install protoc

8. install grpc

# brew install grpc
9. install go 

# brew install hg
# brew install go

1) Create Directories

mkdir $HOME/Go
mkdir -p $HOME/Go/src/github.com/user
2) Setup your paths

export GOPATH=$HOME/Go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
3) Install Go

brew install go

4) "go get" the basics

go get golang.org/x/tools/cmd/godoc




Error ANd steps to resolve it 

protoc-gen-go-grpc: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.

to resolve this 

run this command -> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

another steps to fix this error 






#Ent
RESOURCE: https://entgo.io/docs/getting-started/


--Installation--
go get entgo.io/ent/cmd/ent

--Setup A Go Environment--
go mod init <project>

--Create Your First Schema--
go run entgo.io/ent/cmd/ent init User

--Run go generate from the root directory of the project--
go generate ./ent

--Create Your First Entity--
create a new ent.Client using db, ie: pg admin, sqllite

--Query Your Entities--
ent generates a package for each entity schema that contains its
 predicates, default values, validators and additional information 
 about storage elements (column names, primary keys, etc).

--Add Your First Edge (Relation)--
go run entgo.io/ent/cmd/ent init Car Group


 export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


  # go mod tody for import issue