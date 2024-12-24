package utilize

import (
	"testing"
	"time"
)

func TestHttpCli(t *testing.T) {
	type args struct {
		method             string
		url                string
		headers            map[string]string
		body               string
		timeout            time.Duration
		insecureSkipVerify bool
		caCertPath         string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestHttpCli", args{"GET", "https://jsonplaceholder.typicode.com/posts/1", nil, "", 5 * time.Second, false, ""}, "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HttpCli(tt.args.method, tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout, tt.args.insecureSkipVerify, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpCli() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HttpCli() got = %v, want %v", got, tt.want)
			}
		})
	}
}
