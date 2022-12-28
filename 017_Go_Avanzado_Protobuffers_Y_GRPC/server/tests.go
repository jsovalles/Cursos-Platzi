package server

import (
	"context"
	"github.com/jsovalles/grpc/models"
	"github.com/jsovalles/grpc/repository"
	"github.com/jsovalles/grpc/studentpb"
	"github.com/jsovalles/grpc/testpb"
	"io"
	"log"
	"time"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{
		repo: repo,
	}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		question := models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}
		err = s.repo.SetQuestion(question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		enrollment := models.Enrollment{
			TestId:    msg.TestId,
			StudentId: msg.StudentId,
		}
		err = s.repo.SetEnrollment(enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentsPerTest(req.TestId)
	if err != nil {
		return err
	}
	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err = stream.Send(student)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	questions, err := s.repo.GetQuestionsPerTest("t1")
	if err != nil {
		return err
	}
	i := 0
	var currentQuestion = models.Question{}
	for {
		if i < len(questions) {
			currentQuestion = questions[i]
		}

		if i <= len(questions)-1 {
			questionToSend := &testpb.Question{
				Id:       currentQuestion.Id,
				Question: currentQuestion.Question,
			}
			err := stream.Send(questionToSend)
			if err != nil {
				log.Printf("Error sending question: %v", err)
				return err
			}
			i++
		}
		answer, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving answer: %v", err)
			return err
		}
		log.Println("Answer for question:", currentQuestion.Question, "is", answer.GetAnswer())
	}
}
