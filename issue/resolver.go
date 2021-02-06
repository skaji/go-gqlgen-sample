package issue

import "context"

type User struct {
	Name     string
	Location string
}

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{}
}

type queryResolver struct{}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return []*User{
		&User{Name: "User1", Location: "America"},
		&User{Name: "User2", Location: "America"},
		&User{Name: "User3", Location: "America"},
		&User{Name: "User4", Location: "America"},
		&User{Name: "User5", Location: "America"},
		&User{Name: "User6", Location: "Japan"},
		&User{Name: "User7", Location: "Japan"},
		&User{Name: "User8", Location: "Japan"},
		&User{Name: "User9", Location: "Japan"},
		&User{Name: "User10", Location: "Japan"},
	}, nil
}

func (r *Resolver) User() UserResolver {
	return &userResolver{}
}

type userResolver struct{}

func (r *userResolver) Friends(ctx context.Context, u *User, filter Filter) ([]*User, error) {
	return nil, nil
}
