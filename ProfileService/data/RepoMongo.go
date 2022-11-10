package data

import "log"

type ProfileRepo struct {
	l *log.Logger
}

func mongoConcection(l *log.Logger) (*ProfileRepo, error) {

	//TODO

	return &ProfileRepo{l}, nil

}
