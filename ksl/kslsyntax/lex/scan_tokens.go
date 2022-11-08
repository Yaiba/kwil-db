//line scan_tokens.rl:1

package lex

import (
	"bytes"

	"ksl"
)

// Code generated by scan_tokens.rl; DO NOT EDIT.

//line scan_tokens.go:15
var _ksltok_actions []byte = []byte{
	0, 1, 0, 1, 2, 1, 3, 1, 6,
	1, 7, 1, 8, 1, 9, 1, 10,
	1, 11, 1, 12, 1, 13, 1, 14,
	1, 17, 1, 18, 1, 19, 1, 20,
	1, 21, 1, 22, 1, 31, 1, 32,
	1, 33, 1, 34, 1, 35, 1, 36,
	1, 37, 1, 38, 1, 39, 1, 40,
	1, 41, 1, 42, 1, 43, 1, 44,
	1, 45, 1, 46, 1, 47, 1, 48,
	1, 49, 1, 50, 1, 51, 1, 52,
	1, 53, 1, 54, 2, 0, 1, 2,
	3, 4, 2, 3, 5, 2, 3, 15,
	2, 3, 16, 2, 3, 23, 2, 3,
	24, 2, 3, 25, 2, 3, 26, 2,
	3, 27, 2, 3, 28, 2, 3, 29,
	2, 3, 30,
}

var _ksltok_key_offsets []int16 = []int16{
	0, 0, 2, 7, 12, 16, 18, 23,
	27, 36, 37, 41, 43, 45, 57, 59,
	61, 63, 65, 67, 68, 70, 72, 74,
	120, 122, 123, 124, 127, 129, 130, 132,
	133, 138, 143, 148, 149, 157, 165, 174,
	183, 192, 201, 210, 219, 228, 237, 246,
	248, 250, 252, 266, 280, 282, 294, 296,
	298, 300, 312, 324, 326, 328,
}

var _ksltok_trans_keys []byte = []byte{
	48, 57, 46, 69, 101, 48, 57, 46,
	69, 101, 48, 57, 43, 45, 48, 57,
	48, 57, 45, 65, 90, 97, 122, 65,
	90, 97, 122, 10, 13, 95, 48, 57,
	65, 90, 97, 122, 10, 65, 90, 97,
	122, 128, 191, 128, 191, 10, 13, 128,
	191, 192, 223, 224, 239, 240, 247, 248,
	255, 128, 191, 128, 191, 128, 191, 128,
	191, 128, 191, 10, 128, 191, 128, 191,
	128, 191, 9, 10, 13, 32, 33, 34,
	35, 36, 42, 43, 45, 46, 47, 60,
	62, 92, 96, 102, 110, 116, 123, 125,
	0, 38, 39, 44, 48, 57, 58, 64,
	65, 90, 91, 93, 94, 95, 97, 122,
	124, 127, 192, 223, 224, 239, 240, 247,
	9, 32, 10, 10, 46, 48, 57, 48,
	57, 47, 10, 47, 10, 46, 69, 101,
	48, 57, 46, 69, 101, 48, 57, 46,
	69, 101, 48, 57, 60, 46, 95, 48,
	57, 65, 90, 97, 122, 46, 95, 48,
	57, 65, 90, 97, 122, 46, 95, 97,
	48, 57, 65, 90, 98, 122, 46, 95,
	108, 48, 57, 65, 90, 97, 122, 46,
	95, 115, 48, 57, 65, 90, 97, 122,
	46, 95, 101, 48, 57, 65, 90, 97,
	122, 46, 95, 117, 48, 57, 65, 90,
	97, 122, 46, 95, 108, 48, 57, 65,
	90, 97, 122, 46, 95, 108, 48, 57,
	65, 90, 97, 122, 46, 95, 114, 48,
	57, 65, 90, 97, 122, 46, 95, 117,
	48, 57, 65, 90, 97, 122, 128, 191,
	128, 191, 128, 191, 10, 13, 34, 92,
	128, 191, 192, 223, 224, 239, 240, 247,
	248, 255, 10, 13, 34, 92, 128, 191,
	192, 223, 224, 239, 240, 247, 248, 255,
	10, 13, 10, 13, 128, 191, 192, 223,
	224, 239, 240, 247, 248, 255, 128, 191,
	128, 191, 128, 191, 10, 13, 128, 191,
	192, 223, 224, 239, 240, 247, 248, 255,
	10, 13, 128, 191, 192, 223, 224, 239,
	240, 247, 248, 255, 128, 191, 128, 191,
	128, 191,
}

var _ksltok_single_lengths []byte = []byte{
	0, 0, 3, 3, 2, 0, 1, 0,
	3, 1, 0, 0, 0, 2, 0, 0,
	0, 0, 0, 1, 0, 0, 0, 22,
	2, 1, 1, 1, 0, 1, 2, 1,
	3, 3, 3, 1, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 0,
	0, 0, 4, 4, 2, 2, 0, 0,
	0, 2, 2, 0, 0, 0,
}

var _ksltok_range_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 2, 2,
	3, 0, 2, 1, 1, 5, 1, 1,
	1, 1, 1, 0, 1, 1, 1, 12,
	0, 0, 0, 1, 1, 0, 0, 0,
	1, 1, 1, 0, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1,
	1, 1, 5, 5, 0, 5, 1, 1,
	1, 5, 5, 1, 1, 1,
}

var _ksltok_index_offsets []int16 = []int16{
	0, 0, 2, 7, 12, 16, 18, 22,
	25, 32, 34, 37, 39, 41, 49, 51,
	53, 55, 57, 59, 61, 63, 65, 67,
	102, 105, 107, 109, 112, 114, 116, 119,
	121, 126, 131, 136, 138, 144, 150, 157,
	164, 171, 178, 185, 192, 199, 206, 213,
	215, 217, 219, 229, 239, 242, 250, 252,
	254, 256, 264, 272, 274, 276,
}

var _ksltok_indicies []byte = []byte{
	1, 0, 3, 5, 5, 4, 2, 3,
	5, 5, 6, 0, 7, 7, 6, 0,
	6, 0, 9, 10, 10, 8, 10, 10,
	8, 11, 12, 10, 10, 10, 10, 8,
	11, 8, 13, 13, 0, 15, 14, 16,
	14, 17, 17, 17, 19, 20, 21, 17,
	18, 18, 22, 19, 22, 20, 22, 18,
	23, 24, 23, 26, 25, 27, 25, 28,
	25, 30, 29, 31, 32, 33, 31, 34,
	35, 36, 34, 15, 37, 37, 38, 39,
	41, 15, 15, 34, 43, 44, 45, 46,
	47, 15, 34, 40, 34, 42, 34, 15,
	42, 15, 49, 50, 51, 48, 31, 31,
	52, 32, 53, 55, 36, 56, 57, 0,
	1, 0, 58, 53, 55, 59, 36, 61,
	59, 63, 5, 5, 40, 62, 3, 5,
	5, 6, 64, 3, 5, 5, 4, 65,
	66, 53, 67, 42, 42, 42, 42, 0,
	67, 13, 13, 13, 13, 68, 67, 42,
	70, 42, 42, 42, 69, 67, 42, 71,
	42, 42, 42, 69, 67, 42, 72, 42,
	42, 42, 69, 67, 42, 73, 42, 42,
	42, 69, 67, 42, 74, 42, 42, 42,
	69, 67, 42, 75, 42, 42, 42, 69,
	67, 42, 76, 42, 42, 42, 69, 67,
	42, 77, 42, 42, 42, 69, 67, 42,
	72, 42, 42, 42, 69, 15, 78, 16,
	78, 79, 78, 80, 80, 81, 82, 83,
	84, 85, 86, 83, 18, 87, 87, 87,
	88, 87, 19, 20, 21, 87, 18, 80,
	80, 89, 90, 90, 90, 19, 20, 21,
	90, 18, 18, 91, 24, 91, 92, 91,
	26, 93, 94, 95, 96, 97, 94, 27,
	26, 93, 98, 28, 30, 99, 98, 27,
	27, 100, 28, 100, 30, 100,
}

var _ksltok_trans_targs []byte = []byte{
	23, 28, 23, 3, 34, 4, 33, 5,
	23, 7, 8, 23, 9, 37, 23, 23,
	11, 50, 51, 14, 15, 16, 50, 50,
	17, 57, 57, 58, 20, 57, 21, 24,
	23, 25, 23, 23, 26, 27, 28, 29,
	32, 35, 36, 38, 42, 45, 23, 23,
	23, 47, 48, 49, 23, 23, 23, 23,
	1, 27, 30, 31, 23, 23, 23, 2,
	23, 23, 6, 10, 23, 23, 39, 40,
	41, 36, 43, 44, 36, 46, 23, 12,
	52, 50, 53, 50, 54, 55, 56, 50,
	13, 50, 50, 50, 18, 19, 57, 59,
	60, 61, 57, 22, 57,
}

var _ksltok_trans_actions []byte = []byte{
	83, 103, 77, 0, 103, 0, 106, 0,
	81, 0, 0, 51, 0, 115, 79, 55,
	0, 19, 88, 0, 0, 0, 23, 21,
	0, 35, 25, 94, 0, 33, 0, 0,
	41, 0, 43, 49, 0, 121, 121, 0,
	100, 5, 118, 118, 118, 118, 45, 47,
	53, 0, 5, 5, 57, 75, 71, 39,
	0, 100, 0, 0, 69, 37, 59, 0,
	63, 61, 0, 0, 65, 67, 118, 118,
	118, 109, 118, 118, 112, 118, 73, 0,
	0, 7, 91, 9, 0, 5, 5, 11,
	0, 13, 15, 17, 0, 0, 27, 0,
	97, 97, 29, 0, 31,
}

var _ksltok_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 0, 0, 0, 0, 0,
	0, 85, 0, 0, 0, 0,
}

var _ksltok_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 3, 0, 0, 0, 0, 0,
	0, 3, 0, 0, 0, 0,
}

var _ksltok_eof_trans []int16 = []int16{
	0, 1, 3, 1, 1, 1, 9, 9,
	9, 9, 1, 15, 15, 18, 23, 23,
	23, 24, 24, 26, 26, 26, 30, 0,
	53, 54, 55, 1, 1, 54, 55, 61,
	63, 65, 66, 54, 1, 69, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 79,
	79, 79, 0, 88, 90, 91, 92, 92,
	92, 0, 99, 101, 101, 101,
}

const ksltok_start int = 23
const ksltok_first_final int = 23
const ksltok_error int = 0

const ksltok_en_stringTemplate int = 50
const ksltok_en_heredocTemplate int = 57
const ksltok_en_main int = 23

//line scan_tokens.rl:16

func ScanTokens(data []byte, filename string, start ksl.Pos) []Token {
	stripData := stripUTF8BOM(data)
	start.Offset += len(data) - len(stripData)
	data = stripData

	f := &tokenAccum{
		Filename:  filename,
		Bytes:     data,
		Pos:       start,
		StartByte: start.Offset,
	}

//line scan_tokens.rl:164

	// Ragel state
	p := 0          // "Pointer" into data
	pe := len(data) // End-of-data "pointer"
	ts := 0
	te := 0
	act := 0
	eof := pe
	var stack []int
	var top int

	cs := ksltok_en_main

	var heredocs []heredocInProgress // stack of heredocs we're currently processing

//line scan_tokens.rl:187

	// Make Go compiler happy
	_ = ts
	_ = te
	_ = act
	_ = eof

	token := func(ty TokenType) {
		f.emitToken(ty, ts, te)
	}
	selfToken := func() {
		b := data[ts:te]
		if len(b) != 1 {
			// should never happen
			panic("selfToken only works for single-character tokens")
		}
		f.emitToken(TokenType(b[0]), ts, te)
	}

//line scan_tokens.go:292
	{
		top = 0
		ts = 0
		te = 0
		act = 0
	}

//line scan_tokens.go:300
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		_acts = int(_ksltok_from_state_actions[cs])
		_nacts = uint(_ksltok_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _ksltok_actions[_acts-1] {
			case 2:
//line NONE:1
				ts = p

//line scan_tokens.go:323
			}
		}

		_keys = int(_ksltok_key_offsets[cs])
		_trans = int(_ksltok_index_offsets[cs])

		_klen = int(_ksltok_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case data[p] < _ksltok_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _ksltok_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_ksltok_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case data[p] < _ksltok_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _ksltok_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_ksltok_indicies[_trans])
	_eof_trans:
		cs = int(_ksltok_trans_targs[_trans])

		if _ksltok_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_ksltok_trans_actions[_trans])
		_nacts = uint(_ksltok_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _ksltok_actions[_acts-1] {
			case 3:
//line NONE:1
				te = p + 1

			case 4:
//line scan_tokens.rl:112
				act = 2
			case 5:
//line scan_tokens.rl:130
				act = 4
			case 6:
//line scan_tokens.rl:127
				te = p + 1
				{
					top--
					cs = stack[top]
					{
						stack = stack[:len(stack)-1]
					}
					goto _again
				}
			case 7:
//line scan_tokens.rl:131
				te = p + 1
				{
					token(TokenBadUTF8)
				}
			case 8:
//line scan_tokens.rl:112
				te = p
				p--
				{
					ts--
					te++
					token(TokenQuotedLit)
				}
			case 9:
//line scan_tokens.rl:129
				te = p
				p--
				{
					token(TokenQuotedNewline)
				}
			case 10:
//line scan_tokens.rl:130
				te = p
				p--
				{
					token(TokenInvalid)
				}
			case 11:
//line scan_tokens.rl:131
				te = p
				p--
				{
					token(TokenBadUTF8)
				}
			case 12:
//line scan_tokens.rl:112
				p = (te) - 1
				{
					ts--
					te++
					token(TokenQuotedLit)
				}
			case 13:
//line scan_tokens.rl:131
				p = (te) - 1
				{
					token(TokenBadUTF8)
				}
			case 14:
//line NONE:1
				switch act {
				case 2:
					{
						p = (te) - 1

						ts--
						te++
						token(TokenQuotedLit)
					}
				case 4:
					{
						p = (te) - 1
						token(TokenInvalid)
					}
				}

			case 15:
//line scan_tokens.rl:107
				act = 7
			case 16:
//line scan_tokens.rl:137
				act = 8
			case 17:
//line scan_tokens.rl:81
				te = p + 1
				{
					topdoc := &heredocs[len(heredocs)-1]
					if topdoc.StartOfLine {
						maybeMarker := bytes.TrimSpace(data[ts:te])
						if bytes.Equal(maybeMarker, topdoc.Marker) {
							nls := te - 1
							nle := te
							te--
							if data[te-1] == '\r' {
								// back up one more byte
								nls--
								te--
							}
							token(TokenHeredocEnd)
							ts = nls
							te = nle
							token(TokenNewline)
							heredocs = heredocs[:len(heredocs)-1]
							top--
							cs = stack[top]
							{
								stack = stack[:len(stack)-1]
							}
							goto _again

						}
					}

					topdoc.StartOfLine = true
					token(TokenStringLit)
				}
			case 18:
//line scan_tokens.rl:137
				te = p + 1
				{
					token(TokenBadUTF8)
				}
			case 19:
//line scan_tokens.rl:107
				te = p
				p--
				{
					heredocs[len(heredocs)-1].StartOfLine = false
					token(TokenStringLit)
				}
			case 20:
//line scan_tokens.rl:137
				te = p
				p--
				{
					token(TokenBadUTF8)
				}
			case 21:
//line scan_tokens.rl:107
				p = (te) - 1
				{
					heredocs[len(heredocs)-1].StartOfLine = false
					token(TokenStringLit)
				}
			case 22:
//line NONE:1
				switch act {
				case 0:
					{
						cs = 0
						goto _again
					}
				case 7:
					{
						p = (te) - 1

						heredocs[len(heredocs)-1].StartOfLine = false
						token(TokenStringLit)
					}
				case 8:
					{
						p = (te) - 1
						token(TokenBadUTF8)
					}
				}

			case 23:
//line scan_tokens.rl:142
				act = 10
			case 24:
//line scan_tokens.rl:143
				act = 11
			case 25:
//line scan_tokens.rl:144
				act = 12
			case 26:
//line scan_tokens.rl:145
				act = 13
			case 27:
//line scan_tokens.rl:146
				act = 14
			case 28:
//line scan_tokens.rl:147
				act = 15
			case 29:
//line scan_tokens.rl:148
				act = 16
			case 30:
//line scan_tokens.rl:152
				act = 20
			case 31:
//line scan_tokens.rl:149
				te = p + 1
				{
					token(TokenDocComment)
				}
			case 32:
//line scan_tokens.rl:150
				te = p + 1
				{
					token(TokenComment)
				}
			case 33:
//line scan_tokens.rl:151
				te = p + 1
				{
					token(TokenNewline)
				}
			case 34:
//line scan_tokens.rl:152
				te = p + 1
				{
					selfToken()
				}
			case 35:
//line scan_tokens.rl:154
				te = p + 1
				{
					token(TokenLBrace)
				}
			case 36:
//line scan_tokens.rl:155
				te = p + 1
				{
					token(TokenRBrace)
				}
			case 37:
//line scan_tokens.rl:157
				te = p + 1
				{
					{
						stack = append(stack, 0)
						stack[top] = cs
						top++
						cs = 50
						goto _again
					}
				}
			case 38:
//line scan_tokens.rl:63
				te = p + 1
				{
					token(TokenHeredocBegin)
					marker := data[ts+2 : te-1]
					if marker[0] == '-' {
						marker = marker[1:]
					}
					if marker[len(marker)-1] == '\r' {
						marker = marker[:len(marker)-1]
					}

					heredocs = append(heredocs, heredocInProgress{
						Marker:      marker,
						StartOfLine: true,
					})

					{
						stack = append(stack, 0)
						stack[top] = cs
						top++
						cs = 57
						goto _again
					}
				}
			case 39:
//line scan_tokens.rl:160
				te = p + 1
				{
					token(TokenBadUTF8)
				}
			case 40:
//line scan_tokens.rl:161
				te = p + 1
				{
					token(TokenInvalid)
				}
			case 41:
//line scan_tokens.rl:141
				te = p
				p--

			case 42:
//line scan_tokens.rl:142
				te = p
				p--
				{
					token(TokenIntegerLit)
				}
			case 43:
//line scan_tokens.rl:143
				te = p
				p--
				{
					token(TokenFloatLit)
				}
			case 44:
//line scan_tokens.rl:144
				te = p
				p--
				{
					token(TokenNumberLit)
				}
			case 45:
//line scan_tokens.rl:147
				te = p
				p--
				{
					token(TokenQualifiedIdent)
				}
			case 46:
//line scan_tokens.rl:148
				te = p
				p--
				{
					token(TokenIdent)
				}
			case 47:
//line scan_tokens.rl:149
				te = p
				p--
				{
					token(TokenDocComment)
				}
			case 48:
//line scan_tokens.rl:150
				te = p
				p--
				{
					token(TokenComment)
				}
			case 49:
//line scan_tokens.rl:160
				te = p
				p--
				{
					token(TokenBadUTF8)
				}
			case 50:
//line scan_tokens.rl:161
				te = p
				p--
				{
					token(TokenInvalid)
				}
			case 51:
//line scan_tokens.rl:142
				p = (te) - 1
				{
					token(TokenIntegerLit)
				}
			case 52:
//line scan_tokens.rl:160
				p = (te) - 1
				{
					token(TokenBadUTF8)
				}
			case 53:
//line scan_tokens.rl:161
				p = (te) - 1
				{
					token(TokenInvalid)
				}
			case 54:
//line NONE:1
				switch act {
				case 10:
					{
						p = (te) - 1
						token(TokenIntegerLit)
					}
				case 11:
					{
						p = (te) - 1
						token(TokenFloatLit)
					}
				case 12:
					{
						p = (te) - 1
						token(TokenNumberLit)
					}
				case 13:
					{
						p = (te) - 1
						token(TokenBoolLit)
					}
				case 14:
					{
						p = (te) - 1
						token(TokenNullLit)
					}
				case 15:
					{
						p = (te) - 1
						token(TokenQualifiedIdent)
					}
				case 16:
					{
						p = (te) - 1
						token(TokenIdent)
					}
				case 20:
					{
						p = (te) - 1
						selfToken()
					}
				}

//line scan_tokens.go:725
			}
		}

	_again:
		_acts = int(_ksltok_to_state_actions[cs])
		_nacts = uint(_ksltok_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _ksltok_actions[_acts-1] {
			case 0:
//line NONE:1
				ts = 0

			case 1:
//line NONE:1
				act = 0

//line scan_tokens.go:743
			}
		}

		if cs == 0 {
			goto _out
		}
		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			if _ksltok_eof_trans[cs] > 0 {
				_trans = int(_ksltok_eof_trans[cs] - 1)
				goto _eof_trans
			}
		}

	_out:
		{
		}
	}

//line scan_tokens.rl:210

	// If we fall out here without being in a final state then we've
	// encountered something that the scanner can't match, which we'll
	// deal with as an invalid.
	if cs < ksltok_first_final {
		f.emitToken(TokenInvalid, ts, len(data))
	}

	// We always emit a synthetic EOF token at the end, since it gives the
	// parser position information for an "unexpected EOF" diagnostic.
	f.emitToken(TokenEOF, len(data), len(data))

	return f.Tokens
}
