package email_html

import "testing"

func TestGenerateHtml(t *testing.T) {
	type args struct {
		userName  string
		url       string
		emailType EmailType
	}
	tests := []struct {
		name          string
		args          args
		wantEmailBody string
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			name:          "jonesjedi@gmail.com",
			args:          args{userName: "jonesjedi@gmail.com", url: "https://onb.io", emailType: 1},
			wantEmailBody: "",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmailBody, err := GenerateHtml(tt.args.userName, tt.args.url, "zh-CN", tt.args.emailType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHtml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEmailBody != tt.wantEmailBody {
				t.Errorf("GenerateHtml() = %v, want %v", gotEmailBody, tt.wantEmailBody)
			}
		})
	}
}
