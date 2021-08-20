package dhcp6opts

import (
	"strings"

	"github.com/mdlayher/dhcp6/internal/buffer"
)

type DNSNames []string

func (d *DNSNames) UnmarshalBinary(b []byte) error {
	buf := buffer.New(b)
	var result []string
	for buf.Len() > 0 {
		v, err := unmarshalDNSName(buf)
		if err != nil {
			return err
		}
		result = append(result, v)
	}
	*d = result
	return nil
}

func unmarshalDNSName(buf *buffer.Buffer) (string, error) {
	var v []string
	for {
		l := buf.Read8()
		if l == 0 {
			break
		}
		b2 := make([]byte, int(l))
		buf.ReadBytes(b2)
		v = append(v, string(b2))
	}
	return strings.Join(v, "."), nil
}

func (d DNSNames) MarshalBinary() ([]byte, error) {
	buf := buffer.New(nil)
	for _, v := range d {
		for _, n := range strings.Split(v, ".") {
			buf.Write8(uint8(len(n)))
			buf.WriteBytes([]byte(n))
		}
		buf.Write8(0)
	}
	return buf.Data(), nil
}
