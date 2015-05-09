package messages

import (
	"fmt"
	"sync"
)

type Messenger interface {
	Post(from, to, msg string) error
	Read(user string, n int) ([]Message, error)
}

type Message struct {
	From string
	Body string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s", m.From, m.Body)
}

type MessageMap struct {
	sync.Mutex
	userToMsgs map[string][]Message
}

func NewMessageMap() *MessageMap {
	return &MessageMap{userToMsgs: map[string][]Message{}}
}

func (m *MessageMap) Post(from, to, msg string) error {
	m.Lock()
	defer m.Unlock()
	m.userToMsgs[to] = append(m.userToMsgs[to], Message{from, msg})
	return nil
}

func (m *MessageMap) Read(user string, n int) ([]Message, error) {
	m.Lock()
	defer m.Unlock()
	var numMsgs int
	allMsgs, _ := m.userToMsgs[user]
	if len(allMsgs) > n {
		numMsgs = n
	} else {
		numMsgs = len(allMsgs)
	}
	msgs := make([]Message, numMsgs)
	copy(msgs, allMsgs[:numMsgs])
	// Remove read messages
	allMsgs = allMsgs[numMsgs:]
	m.userToMsgs[user] = allMsgs
	return msgs, nil
}
