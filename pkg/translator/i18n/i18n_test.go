package i18n

import (
	"testing"

	"github.com/majedutd990/bookstore-api/pkg/translator"
	"github.com/majedutd990/bookstore-api/pkg/translator/messages"
)

func TestMessageBundle_Translate(t *testing.T) {
	type args struct {
		message  string
		language translator.Language
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate farsi",
			args: args{
				message:  messages.DBError,
				language: translator.Getlanguage("fa"),
			},
			want: "خطایی وجود دارد",
		}, {
			name: "translate english",
			args: args{
				message:  messages.UserNotFound,
				language: translator.Getlanguage("en"),
			},
			want: "user not found",
		}, {
			name: "message key not found",
			args: args{
				message:  "NoKeyFound",
				language: translator.Getlanguage("en"),
			},
			want: "NoKeyFound",
		},
	}
	translator, err := New("test/")
	if err != nil {
		t.Errorf("New() Error: %v\n", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translator.Translate(tt.args.language, tt.args.message)
			if got != tt.want {
				t.Errorf("Translate got=%v want=%v", got, tt.want)
			}
		})

	}
}

//  write for translate en ur self
