package protocplugin

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"regexp"
	"strings"
)

type HttpRule struct {
	*annotations.HttpRule
}

// PathTemplate returns a path template object parsed from the HTTP rule.
// The path template syntax is described in https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#path-template-syntax
// Verb is not supported
func (r *HttpRule) PathTemplate() *HttpRulePathTemplate {
	var pathPattern string
	switch p := r.GetPattern().(type) {
	default:
		return nil
	case *annotations.HttpRule_Get:
		pathPattern = p.Get
	case *annotations.HttpRule_Post:
		pathPattern = p.Post
	case *annotations.HttpRule_Put:
		pathPattern = p.Put
	case *annotations.HttpRule_Patch:
		pathPattern = p.Patch
	case *annotations.HttpRule_Delete:
		pathPattern = p.Delete
	case *annotations.HttpRule_Custom:
		pathPattern = p.Custom.Path
	}
	return &HttpRulePathTemplate{
		Segments: parsePathTemplateSegments(pathPattern),
	}
}

var (
	regexSingleAsterisk = regexp.MustCompile(`^\*`)
	regexDoubleAsterisk = regexp.MustCompile(`^\*\*`)
	regexLiteral        = regexp.MustCompile(`^(\*|\*\*|[^{/]+)`)
	regexVariable       = regexp.MustCompile(`^\{[^}]+}`)
)

func parsePathTemplateSegments(pathPattern string) (segments []*HttpRulePathTemplateSegment) {
	for pathPattern != "" {
		pathPattern = strings.TrimPrefix(pathPattern, "/") // skip leading slash

		if m := regexDoubleAsterisk.FindString(pathPattern); m != "" {
			segments = append(segments, &HttpRulePathTemplateSegment{Value: m})
			pathPattern = strings.TrimPrefix(pathPattern, m)
			continue
		}
		if m := regexSingleAsterisk.FindString(pathPattern); m != "" {
			segments = append(segments, &HttpRulePathTemplateSegment{Value: m})
			pathPattern = strings.TrimPrefix(pathPattern, m)
			continue
		}
		if m := regexLiteral.FindString(pathPattern); m != "" {
			segments = append(segments, &HttpRulePathTemplateSegment{Value: m})
			pathPattern = strings.TrimPrefix(pathPattern, m)
			continue
		}
		if m := regexVariable.FindString(pathPattern); m != "" {
			l, r, cut := strings.Cut(m[1:len(m)-1] /* trim '{' and '}' */, "=")
			if !cut {
				r = "*"
			}

			segments = append(segments, &HttpRulePathTemplateSegment{
				Value: m,
				Variable: &HttpRulePathTemplateVariable{
					FieldPath: strings.Split(l, "."),
					Segments:  parsePathTemplateSegments(r),
				},
			})
			pathPattern = strings.TrimPrefix(pathPattern, m)
			continue
		}

		panic("invalid path template syntax: " + pathPattern)
	}
	return segments
}

type HttpRulePathTemplate struct {
	Segments []*HttpRulePathTemplateSegment
}
type HttpRulePathTemplateSegment struct {
	Value    string
	Variable *HttpRulePathTemplateVariable
}
type HttpRulePathTemplateVariable struct {
	FieldPath []string
	Segments  []*HttpRulePathTemplateSegment
}
