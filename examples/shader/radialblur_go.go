// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package main

var radialblur_go = []byte("// Copyright 2020 The Ebiten Authors\r\n//\r\n// Licensed under the Apache License, Version 2.0 (the \"License\");\r\n// you may not use this file except in compliance with the License.\r\n// You may obtain a copy of the License at\r\n//\r\n//     http://www.apache.org/licenses/LICENSE-2.0\r\n//\r\n// Unless required by applicable law or agreed to in writing, software\r\n// distributed under the License is distributed on an \"AS IS\" BASIS,\r\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\r\n// See the License for the specific language governing permissions and\r\n// limitations under the License.\r\n\r\n// +build ignore\r\n\r\npackage main\r\n\r\nvar Time float\r\nvar Cursor vec2\r\nvar ScreenSize vec2\r\n\r\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\r\n\tdir := normalize(position.xy - Cursor)\r\n\tclr := texture2At(texCoord)\r\n\r\n\tsamples := [10]float{\r\n\t\t-22, -14, -8, -4, -2, 2, 4, 8, 14, 22,\r\n\t}\r\n\t// TODO: Add len(samples)\r\n\tsum := clr\r\n\tfor i := 0; i < 10; i++ {\r\n\t\t// TODO: Consider the source region not to violate the region.\r\n\t\tsum += texture2At(texCoord + dir*samples[i]/texture2Size())\r\n\t}\r\n\tsum /= 10 + 1\r\n\r\n\tdist := distance(position.xy, Cursor)\r\n\tt := clamp(dist/256, 0, 1)\r\n\treturn mix(clr, sum, t)\r\n}\r\n")