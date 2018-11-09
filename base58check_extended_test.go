package base58_test

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/sammy00/base58"
)

func TestCheckEncodeX(t *testing.T) {
	testCases := []struct {
		version []byte
		input   string
		expect  string
	}{
		{
			[]byte{0x80},
			"A43A940577F4E97F5C4D39EB14FF083A98187C64EA7C99EF7CE460833959A519",
			"5K4caxezwjGCGfnoPTZ8tMcJBLB7Jvyjv4xxeacadhq8nLisLR2",
		},
		{
			[]byte{0x80},
			"C2C8036DF268F498099350718C4A3EF3984D2BE84618C2650F5171DCC5EB660A",
			"5KJ51SgxWaAYR13zd9ReMhJpwrcX47xTJh2D3fGPG9CM8vkv5sH",
		},
	}

	for i, c := range testCases {
		payload, _ := hex.DecodeString(c.input)

		if got := base58.CheckEncodeX(c.version, payload); got != c.expect {
			t.Fatalf("#%d invalid encoded string: got %s, expect %s", i, got,
				c.expect)
		}
	}
}

func TestCheckDecodeX(t *testing.T) {
	type expect struct {
		decoded string
		version []byte
		err     error
	}

	testCases := []struct {
		input      string
		versionLen int
		expect     expect
	}{
		{
			"5KN7MzqK5wt2TP1fQCYyHBtDrXdJuXbUzm4A9rKAteGu3Qi5CVR",
			1,
			expect{
				"CBF4B9F70470856BB4F40F80B87EDB90865997FFEE6DF315AB166D713AF433A5",
				[]byte{0x80},
				nil,
			},
		},
		{
			"5HtasZ6ofTHP6HCwTqTkLDuLQisYPah7aUnSKfC7h4hMUVw2gi5",
			1,
			expect{
				"09C2686880095B1A4C249EE3AC4EEA8A014F11E6F986D0B5025AC1F39AFBD9AE",
				[]byte{0x80},
				nil,
			},
		},
		{
			"cfrm38V8aXBn7JWA1ESmFMUn6erxeBGZGAxJPY4e36S9QWkzZKtaVqLNMgnifETYw7BPwWC9aPD",
			5,
			expect{
				"04bb458cef4fca5a974040f00103e572ae02b57cc3a84e78caac6ccf16867fa93e9d1e4a636766eef5dd0375d87a",
				[]byte{0x64, 0x3B, 0xF6, 0xA8, 0x9A},
				nil,
			},
		},
		{
			"cfrm38V8G4qq2ywYEFfWLD5Cc6msj9UwsG2Mj4Z6QdGJAFQpdatZLavkgRd1i4iBMdRngDqDs51",
			5,
			expect{
				"04494af136c40ea76fc501a00103fcbcd88fdce971a492ad3beeb99d0f3181470d48819039b7cee569e76acc4aa3",
				[]byte{0x64, 0x3B, 0xF6, 0xA8, 0x9A},
				nil,
			},
		},
		{"cfrm38", 5, expect{"", nil, base58.ErrInvalidFormat}},
		{
			"cfrm38V8G4qq2ywYEFfWLD5Cc6msj9UwsG2Mj4Z6QdGJAFQpdatZLavkgRd1i4iBMdRngDqDs52",
			5,
			expect{"", nil, base58.ErrChecksum},
		},
	}

	for i, c := range testCases {
		decoded, version, err := base58.CheckDecodeX(c.input, c.versionLen)

		if err != c.expect.err {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, err, c.expect.err)
		}

		if !bytes.Equal(version, c.expect.version) {
			t.Fatalf("#%d invalid version: got %x, expect %x", i, version,
				c.expect.version)
		}

		if decodedHex := hex.EncodeToString(
			decoded); !strings.EqualFold(decodedHex, c.expect.decoded) {
			t.Fatalf("#%d invalid decoded: got %s, expect %s", i, decodedHex,
				c.expect.decoded)
		}
	}
}
