package protocplugin

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/api/annotations"
	"testing"
)

func TestHttpRule_PathTemplate(t *testing.T) {
	tests := []struct {
		name string
		sut  HttpRule
		want *HttpRulePathTemplate
	}{
		{
			name: "Get(/v1/{name=messages/*})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/{name=messages/*}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "{name=messages/*}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"name"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "messages"}, {Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Post(/v1/{name=messages/*})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Post{Post: "/v1/{name=messages/*}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "{name=messages/*}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"name"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "messages"}, {Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Put(/v1/{name=messages/*})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Put{Put: "/v1/{name=messages/*}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "{name=messages/*}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"name"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "messages"}, {Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Patch(/v1/{name=messages/*})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Patch{Patch: "/v1/{name=messages/*}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "{name=messages/*}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"name"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "messages"}, {Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Delete(/v1/{name=messages/*})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Delete{Delete: "/v1/{name=messages/*}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "{name=messages/*}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"name"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "messages"}, {Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Get(/v1/messages/{message_id})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/messages/{message_id}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "messages"},
				{Value: "{message_id}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"message_id"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Get(/v1/messages/{message_id}/{sub.subfield})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/messages/{message_id}/{sub.subfield}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "messages"},
				{Value: "{message_id}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"message_id"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "*"}},
				}},
				{Value: "{sub.subfield}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"sub", "subfield"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "*"}},
				}},
			},
			},
		},
		{
			name: "Get(/v1/messages/{message_id}/subs/{sub.subfield=/sub/**})",
			sut:  HttpRule{HttpRule: &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/messages/{message_id}/subs/{sub.subfield=/sub/**}"}}},
			want: &HttpRulePathTemplate{Segments: []*HttpRulePathTemplateSegment{
				{Value: "v1"},
				{Value: "messages"},
				{Value: "{message_id}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"message_id"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "*"}},
				}},
				{Value: "subs"},
				{Value: "{sub.subfield=/sub/**}", Variable: &HttpRulePathTemplateVariable{
					FieldPath: []string{"sub", "subfield"},
					Segments:  []*HttpRulePathTemplateSegment{{Value: "sub"}, {Value: "**"}},
				}},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.PathTemplate()
			assert.Equal(t, tt.want, got)
		})
	}
}
