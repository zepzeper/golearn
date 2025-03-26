package checker

type linkChecker struct {
    URL string
    Status int
    isDead bool
    RefferelURL string
}

func NewLinkChecker(url string, status int, isDead bool, RefferelURL string) *linkChecker {
   return &linkChecker{
        URL: url,
        Status: status,
        isDead: isDead,
        RefferelURL: RefferelURL,
    } 
}

func (l *linkChecker) IsDead() bool  {
   return l.isDead 
}
