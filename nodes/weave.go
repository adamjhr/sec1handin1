package nodes

import (
	"fmt"

	"github.com/adamjhr/sec1handin1/message"
	"github.com/adamjhr/sec1handin1/node"
)

func WeaveMain(self *node.Node) {
	m := <-self.Incoming
	fmt.Println(self.Name, ": received message from", m.Sender, "with pub_key", m.Pub_key, "and value", m.Value)

	newValue := m.Value * 2
	fmt.Println(self.Name, ": modifying message to have value", newValue)

	newMessage := message.Message{Sender: m.Sender, Pub_key: m.Pub_key, Value: newValue}

	for _, channel := range self.Outgoing {
		*channel <- newMessage
	}
}
