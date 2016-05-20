# golang sample send email package

[![GoDoc](https://godoc.org/github.com/gcloudplatform/email?status.png)](https://godoc.org/github.com/gcloudplatform/email)

requirement go 1.6 (mime.QEncoding.Encode(charset, s string))

## Feature
- send html format email
- send email with `attachment`

## Usage

Install
    go get github.com/gcloudplatform/email

```go
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

//detail see email_test.go
//ref https://github.com/scorredoira/email
```