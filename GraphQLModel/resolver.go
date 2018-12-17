package GraphQLModel

import (
	"context"
	"sega/Base"
)

// ===============================================================================

type Resolver struct{}

// ===========================================================[Main]

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// ===========================================================[Mutation]

type mutationResolver struct{ *Resolver }

// ============================================[User]

func (r *mutationResolver) UpdateUser(ctx context.Context, Token Base.InputToken, User Base.NewAccountUser) (Base.CreateReturn, error) {
	panic("not implemented")
}

// ============================================[Account]

func (r *mutationResolver) CreateAccount(ctx context.Context, AccountIDPW Base.NewAccountIDPW, User Base.NewAccountUser) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) ChangePassword(ctx context.Context, Token Base.InputToken, OldPW Base.AccountPW, NewPW Base.AccountPW, ConfirmationPW Base.AccountPW) (Base.CreateReturn, error) {
	panic("not implemented")
}

// ============================================[CarID]

func (r *mutationResolver) AddCarID(ctx context.Context, Token Base.InputToken, InputCarNews Base.CarNews) (Base.CarIDReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateCarName(ctx context.Context, Token Base.InputToken, CarNameData Base.NewCarName) (Base.CreateReturn, error) {
	panic("not implemented")
}

// ============================================[Status]

func (r *mutationResolver) UpdateMonitor(ctx context.Context, InputMonitorData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateSecurity(ctx context.Context, InputSecurityData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddSecurity(ctx context.Context, InputSecurityData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

// ===========================================================[Query]

type queryResolver struct{ *Resolver }

// ============================================[User]

func (r *queryResolver) GetUser(ctx context.Context, ID string, Token string) (Base.Users, error) {
	panic("not implemented")
}

// ============================================[Account]

func (r *queryResolver) LogIn(ctx context.Context, ID string, Password string) (Base.LogInToken, error) {
	panic("not implemented")
}

func (r *queryResolver) LogOut(ctx context.Context, Token string) (Base.StatusData, error) {
	panic("not implemented")
}

func (r *queryResolver) CheckAccountHas(ctx context.Context, ID string) (Base.AccountHas, error) {
	panic("not implemented")
}

// ============================================[CarID]

func (r *queryResolver) GetCarID(ctx context.Context, ID string, Token string) ([]Base.CarData, error) {
	panic("not implemented")
}

func (r *queryResolver) DeleteCarID(ctx context.Context, ID string, Token string, CarID string) (Base.StatusData, error) {
	panic("not implemented")
}

// ============================================[Status]

func (r *queryResolver) GetMonitorStatus(ctx context.Context, ID string, Token string, SelectObject string) (Base.MonitorData, error) {
	panic("not implemented")
}

func (r *queryResolver) GetSecurityStatus(ctx context.Context, ID string, Token string, SelectObject string) (Base.SecurityData, error) {
	panic("not implemented")
}

func (r *queryResolver) GetTemporarilyToken(ctx context.Context, ID string, Token string) (Base.TemporarilyTokenData, error) {
	panic("not implemented")
}

// ===============================================================================