
/**
 * Copyright 2019-present Facebook Inc. All rights reserved.
 *
 * This source code is licensed under the Apache 2.0 license found
 * in the LICENSE file in the root directory of this source tree.
 *
 * @format
 */


export default function(Prism) {
  Prism.languages.gotemplate = {
    'comment': [
      {
        pattern: /(^|[^\\])\/\*[\s\S]*?(?:\*\/|$)/,
        lookbehind: true,
        greedy: true
      },
      {
        pattern: /(^|[^\\:])\/\/.*/,
        lookbehind: true,
        greedy: true
      },
      /{{\/\*[\s\S]*\*\/}}/
    ],
    'string': {
      pattern: /(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1/,
      greedy: true
    },
    'boolean': /\b(?:_|iota|nil|true|false)\b/,
    'number': /(?:\b0x[a-f\d]+|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[-+]?\d+)?)i?/i,
    'operator': /[*\/%^!=]=?|\+[=+]?|-[=-]?|\|[=|]?|&(?:=|&|\^=?)?|>(?:>=?|=)?|<(?:<=?|=|-)?|:=|\.\.\./,
    'builtin': /\b(?:bool|byte|complex(?:64|128)|error|float(?:32|64)|rune|string|u?int(?:8|16|32|64)?|uintptr|append|cap|close|complex|copy|delete|imag|len|make|new|panic|print(?:ln)?|real|recover)\b/,
    'keyword': /\b(?:break|default|func|interface|select|case|map|struct|chan|else|goto|package|switch|const|fallthrough|if|range|type|continue|for|import|return|var|go|defer|bool|byte|complex64|complex128|float32|float64|int8|int16|int32|int64|string|uint8|uint16|uint32|uint64|int|uint|uintptr|rune|with|define|block|end)\b/,
  };
}