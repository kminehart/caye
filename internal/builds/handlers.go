package builds

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/kminehart/caye"
)

func (s *Server) StartBuild(ctx context.Context, req *caye.StartBuildRequest) (*caye.StartBuildResponse, error) {
	log.Println("Starting build...")
	id := uuid.New().String()
	s.builds[id] = caye.Build{
		Status: caye.Status_RUNNING,
		Steps:  map[string]*caye.BuildStep{},
	}
	return &caye.StartBuildResponse{
		Id: id,
	}, nil
}

func (s *Server) EndBuild(ctx context.Context, req *caye.EndBuildRequest) (*caye.EndBuildResponse, error) {
	log.Println("Finished build")
	status := caye.Status_DONE
	if req.GetError() {
		status = caye.Status_FAILED
	}

	build, ok := s.builds[req.GetId()]
	if !ok {
		return nil, errors.New("not found")
	}

	build.Status = status

	s.builds[req.GetId()] = build
	return &caye.EndBuildResponse{}, nil
}

func (s *Server) StartRun(ctx context.Context, req *caye.StartRunRequest) (*caye.StartRunResponse, error) {
	log.Println("Client reported it started running a task...")
	id := uuid.New().String()
	s.builds[req.GetBuildId()].Steps[id] = &caye.BuildStep{
		Status: caye.Status_RUNNING,
	}

	return &caye.StartRunResponse{
		Id: id,
	}, nil
}

func (s *Server) EndRun(ctx context.Context, req *caye.EndRunRequest) (*caye.EndRunResponse, error) {
	log.Println("Client reported it has finished running a task...")
	status := caye.Status_DONE
	if req.GetError() {
		status = caye.Status_FAILED
	}

	build, ok := s.builds[req.GetBuildId()]
	if !ok {
		return nil, errors.New("not found")
	}

	run, ok := build.Steps[req.GetRunId()]
	if !ok {
		return nil, errors.New("not found")
	}

	run.Status = status

	s.builds[req.GetBuildId()].Steps[req.GetRunId()] = run

	return &caye.EndRunResponse{}, nil
}

func (s *Server) List(ctx context.Context, req *caye.ListBuildsRequest) (*caye.ListBuildsResponse, error) {
	b := make([]*caye.Build, len(s.builds))
	i := 0
	for _, v := range s.builds {
		b[i] = &v
		i++
	}

	return &caye.ListBuildsResponse{
		Builds: b,
	}, nil
}
