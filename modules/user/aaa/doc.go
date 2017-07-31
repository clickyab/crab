package aaa

// TODO : remove me when you are done

type User struct {
}

type Manager struct {
}

func (m *Manager) FindUserByToken(token string) (*User, error) {
	return nil, nil
}

func NewAaaManager() *Manager {
	return &Manager{}
}
