// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottltest"
)

func hello[K any]() (ExprFunc[K], error) {
	return func(ctx context.Context, tCtx K) (interface{}, error) {
		return "world", nil
	}, nil
}

func Test_newGetter(t *testing.T) {
	tests := []struct {
		name string
		val  value
		ctx  interface{}
		want interface{}
	}{
		{
			name: "string literal",
			val: value{
				String: ottltest.Strp("str"),
			},
			want: "str",
		},
		{
			name: "float literal",
			val: value{
				Literal: &mathExprLiteral{
					Float: ottltest.Floatp(1.2),
				},
			},
			want: 1.2,
		},
		{
			name: "int literal",
			val: value{
				Literal: &mathExprLiteral{
					Int: ottltest.Intp(12),
				},
			},
			want: int64(12),
		},
		{
			name: "bytes literal",
			val: value{
				Bytes: (*byteSlice)(&[]byte{1, 2, 3, 4, 5, 6, 7, 8}),
			},
			want: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "nil literal",
			val: value{
				IsNil: (*isNil)(ottltest.Boolp(true)),
			},
			want: nil,
		},
		{
			name: "bool literal",
			val: value{
				Bool: (*boolean)(ottltest.Boolp(true)),
			},
			want: true,
		},
		{
			name: "path expression",
			val: value{
				Literal: &mathExprLiteral{
					Path: &Path{
						Fields: []Field{
							{
								Name: "name",
							},
						},
					},
				},
			},
			want: "bear",
		},
		{
			name: "function call",
			val: value{
				Literal: &mathExprLiteral{
					Converter: &converter{
						Function: "Hello",
					},
				},
			},
			want: "world",
		},
		{
			name: "enum",
			val: value{
				Enum: (*EnumSymbol)(ottltest.Strp("TEST_ENUM_ONE")),
			},
			want: int64(1),
		},
		{
			name: "empty list",
			val: value{
				List: &list{
					Values: []value{},
				},
			},
			want: []any{},
		},
		{
			name: "string list",
			val: value{
				List: &list{
					Values: []value{
						{
							String: ottltest.Strp("test0"),
						},
						{
							String: ottltest.Strp("test1"),
						},
					},
				},
			},
			want: []any{"test0", "test1"},
		},
		{
			name: "int list",
			val: value{
				List: &list{
					Values: []value{
						{
							Literal: &mathExprLiteral{
								Int: ottltest.Intp(1),
							},
						},
						{
							Literal: &mathExprLiteral{
								Int: ottltest.Intp(2),
							},
						},
					},
				},
			},
			want: []any{int64(1), int64(2)},
		},
		{
			name: "float list",
			val: value{
				List: &list{
					Values: []value{
						{
							Literal: &mathExprLiteral{
								Float: ottltest.Floatp(1.2),
							},
						},
						{
							Literal: &mathExprLiteral{
								Float: ottltest.Floatp(2.4),
							},
						},
					},
				},
			},
			want: []any{1.2, 2.4},
		},
		{
			name: "bool list",
			val: value{
				List: &list{
					Values: []value{
						{
							Bool: (*boolean)(ottltest.Boolp(true)),
						},
						{
							Bool: (*boolean)(ottltest.Boolp(false)),
						},
					},
				},
			},
			want: []any{true, false},
		},
		{
			name: "byte slice list",
			val: value{
				List: &list{
					Values: []value{
						{
							Bytes: (*byteSlice)(&[]byte{1, 2, 3, 4, 5, 6, 7, 8}),
						},
						{
							Bytes: (*byteSlice)(&[]byte{9, 8, 7, 6, 5, 4, 3, 2}),
						},
					},
				},
			},
			want: []any{[]byte{1, 2, 3, 4, 5, 6, 7, 8}, []byte{9, 8, 7, 6, 5, 4, 3, 2}},
		},
		{
			name: "path expression",
			val: value{
				List: &list{
					Values: []value{
						{
							Literal: &mathExprLiteral{
								Path: &Path{
									Fields: []Field{
										{
											Name: "name",
										},
									},
								},
							},
						},
					},
				},
			},
			ctx:  "bear",
			want: []any{"bear"},
		},
		{
			name: "function call",
			val: value{
				List: &list{
					Values: []value{
						{
							Literal: &mathExprLiteral{
								Converter: &converter{
									Function: "Hello",
								},
							},
						},
					},
				},
			},
			want: []any{"world"},
		},
		{
			name: "nil slice",
			val: value{
				List: &list{
					Values: []value{
						{
							IsNil: (*isNil)(ottltest.Boolp(true)),
						},
						{
							IsNil: (*isNil)(ottltest.Boolp(true)),
						},
					},
				},
			},
			want: []any{nil, nil},
		},
		{
			name: "heterogeneous slice",
			val: value{
				List: &list{
					Values: []value{
						{
							String: ottltest.Strp("test0"),
						},
						{
							Literal: &mathExprLiteral{
								Int: ottltest.Intp(1),
							},
						},
					},
				},
			},
			want: []any{"test0", int64(1)},
		},
	}

	functions := map[string]interface{}{"Hello": hello[interface{}]}

	p := NewParser(
		functions,
		testParsePath,
		testParseEnum,
		component.TelemetrySettings{},
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := p.newGetter(tt.val)
			assert.NoError(t, err)

			tCtx := tt.want

			if tt.ctx != nil {
				tCtx = tt.ctx
			}

			val, _ := reader.Get(context.Background(), tCtx)
			assert.Equal(t, tt.want, val)
		})
	}

	t.Run("empty value", func(t *testing.T) {
		_, err := p.newGetter(value{})
		assert.Error(t, err)
	})
}
