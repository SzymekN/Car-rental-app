package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/SzymekN/CRUD/pkg/controller"
	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/producer"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	UnimplementedUserapiServer
}

var (
	port = os.Getenv("GRPC_PORT")
)

func (s *server) DeleteUser(c context.Context, in *UserRequest) (*UserReply, error) {

	u, err := controller.DeleteUserCassandra(int(in.GetId()))
	k, msg := "", "userapi_grpc.users"
	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	if err != nil {
		k = strconv.Itoa(int(in.GetId()))
		msg += "[" + k + "] GRPC delete error: incorrect parameters: "
		return &UserReply{Message: msg, Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
	}

	k = strconv.Itoa(int(in.GetId()))
	msg += "[" + k + "] GRPC delete succesful: "

	return &UserReply{Message: msg, Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
}

func (s *server) PutUser(c context.Context, in *UserPut) (*UserReply, error) {

	u := model.User{Id: int(in.GetId()), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: int(in.GetAge())}

	err := controller.UpdateUserCassandra(int(in.GetId()), u)
	k, msg := "", "userapi_grpc.users"
	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	if err != nil {
		k = strconv.Itoa(int(in.GetId()))
		msg += "[" + k + "] GRPC PUT error: " + err.Error()
		return &UserReply{Message: msg, Id: in.GetId()}, nil
	}

	k = strconv.Itoa(int(in.GetId()))
	msg += "[" + k + "] GRPC PUT succesfull: "
	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, nil
}
func (s *server) PostUser(c context.Context, in *UserPost) (*UserReply, error) {
	u := model.User{Id: int(in.GetId()), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: int(in.GetAge())}
	err := controller.SaveUserCassandra(u)

	k, msg := "", "userapi_grpc.users"
	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	if err != nil {
		k = strconv.Itoa(int(in.GetId()))
		msg += "[" + k + "] GRPC POST error: " + err.Error()
		return &UserReply{Message: msg, Id: in.GetId()}, nil
	}

	k = strconv.Itoa(int(in.GetId()))
	msg += "[" + k + "] GRPC POST succesfull: "
	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, err

}

func (s *server) GetUsers(in *UserRequest, stream Userapi_GetUsersServer) error {
	// users := []model.User{}
	users, _ := controller.GetUsersCassandra()
	k, msg := "all", "userapi_grpc.users"
	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	for _, u := range users {

		if err := stream.Send(&UserReply{Id: int64(u.Id), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}); err != nil {
			msg += "[" + k + "] GRPC GetUsers error " + err.Error()
			return err
		}
	}

	msg += "[" + k + "] GRPC GetUsers error succesfull"
	return nil
}
func (s *server) GetUser(c context.Context, in *UserRequest) (*UserReply, error) {

	u, err := controller.GetUserByIdCassandra(int(in.GetId()))
	k, msg := "all", "userapi_grpc.users"
	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	if err != nil {
		k = strconv.Itoa(int(in.GetId()))
		msg += "[" + k + "] GRPC GetById error: " + err.Error()
		return &UserReply{Message: msg, Id: in.GetId()}, nil
	}

	k = strconv.Itoa(int(in.GetId()))
	msg += "[" + k + "] GRPC GetById succesfull: " + err.Error()
	return &UserReply{Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
}

func CreateGRPCServer() {

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterUserapiServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
