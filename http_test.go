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
		{"TestHttpCli", args{"GET", "https://jsonplaceholder.typicode.com/posts/1", nil, "", 5 * time.Second, true, ""}, "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}", false},
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

func TestHttpDelete(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		body    string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpDelete(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpDelete() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpGet(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		{"TestHttpCli", args{"https://jsonplaceholder.typicode.com/posts/1", nil, 5 * time.Second}, "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpGet(tt.args.url, tt.args.headers, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpHead(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpHead(tt.args.url, tt.args.headers, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpHead() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpOptions(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpOptions(tt.args.url, tt.args.headers, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpOptions() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpPatch(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		body    string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpPatch(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpPatch() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpPost(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		body    string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpPost(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpPost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpPut(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		body    string
		timeout time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpPut(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpPut() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsDelete(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		body       string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsDelete(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsDelete() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsGet(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsGet(tt.args.url, tt.args.headers, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsHead(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsHead(tt.args.url, tt.args.headers, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsHead() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsOptions(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsOptions(tt.args.url, tt.args.headers, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsOptions() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsPatch(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		body       string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsPatch(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsPatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsPatch() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsPost(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		body       string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsPost(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsPost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestHttpsPut(t *testing.T) {
	type args struct {
		url        string
		headers    map[string]string
		body       string
		timeout    time.Duration
		caCertPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := HttpsPut(tt.args.url, tt.args.headers, tt.args.body, tt.args.timeout, tt.args.caCertPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpsPut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("HttpsPut() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
