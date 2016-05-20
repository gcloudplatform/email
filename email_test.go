package email

import (
	"fmt"
	"io/ioutil"
	"net/mail"
	"testing"
)

func Test_Send(t *testing.T) {
	m := New("smtpip:25", "password", &mail.Address{
		Address: "account@hostname",
		Name:    "account",
	})

	var to []*mail.Address
	to = append(to, &mail.Address{
		Address: "any@someone",
		Name:    "any",
	})

	err := m.Send("email title", "<h2>email content</h2>", to)
	if err != nil {
		t.Fatalf("send error: %s", err)
	} else {
		t.Log("send ok.")
	}
}

func Test_Boundary(t *testing.T) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d: \"%s\"\n", i, genBoundary(28))
	}
	fmt.Println()
}

func Test_Attach(t *testing.T) {
	m := New("smtpip:25", "password", &mail.Address{
		Address: "account@hostname",
		Name:    "account",
	})

	var to []*mail.Address
	to = append(to, &mail.Address{
		Address: "any@someone",
		Name:    "any",
	})

	attach := make([]*Attachment, 0)
	attach = append(attach, &Attachment{
		Filename: "file.pdf",
		Inline:   false,
	})
	if f1, err := ioutil.ReadFile("/go/src/github.com/gcloudplatform/email/file.pdf"); err != nil {
		t.Fatalf("read file.pdf error: %s", err)
	} else {
		attach[0].Data = f1
	}

	err := m.SendWithAttachment("email title", "<h2>email content</h2>", to, attach)
	if err != nil {
		t.Fatalf("send error: %s", err)
	} else {
		t.Log("send ok.")
	}
}
