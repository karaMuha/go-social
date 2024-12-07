package authtoken

type ITokenProvider interface {
	GenerateToken(userId string) (string, error)
	VerifyToken(token string) (any, error)
	GetUserIDFromToken(parsedToken any) (string, error)
}
