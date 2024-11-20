package main

import (
	"context"
	user "m2/ogen/user"
	"sync"
)

type usersService struct {
	users map[int]user.User
	id    int
	mux   sync.Mutex
}

func (p *usersService) AddUser(ctx context.Context, req *user.User) (*user.User, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.users[p.id] = *req
	p.id++
	return req, nil
}

func (p *usersService) DeleteUser(ctx context.Context, params user.DeleteUserParams) (user.DeleteUserRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

    user := p.users[params.ID]
	delete(p.users, params.ID)
	return user, nil
}

func (p *usersService) GetUser(ctx context.Context, params user.GetUserParams) (user.GetUserRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

    r := user.GetUserRes{
        
    }
	user, _ := p.users[params.ID]
	return &user, nil
}

func (p *usersService) PutUser(ctx context.Context, params user.PutUserParams) (user.PutUserRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	user := p.users[params.ID]
	p.users[params.ID] = user

	return nil
}

func (p *usersService) PostUser(ctx context.Context, req user.OptUser) (user.PostUserRes, error){
	p.mux.Lock()
	defer p.mux.Unlock()

	user := p.users[params.ID]
	p.users[params.ID] = user

	return nil
}

func (p *usersService) NewError(ctx context.Context, err error) *user.UserStatusCode {
    return nil
}
