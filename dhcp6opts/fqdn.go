package dhcp6opts

import (
	"github.com/mdlayher/dhcp6/internal/buffer"
)

type FQDN struct {
	Flags uint8
	Name  string
}

const (
	FQDNFlagS = 1
	FQDNFlagO = 2
	FQDNFlagN = 4
)

func (f *FQDN) UnmarshalBinary(b []byte) error {
	buf := buffer.New(b)
	f.Flags = buf.Read8()
	var err error
	f.Name, err = unmarshalDNSName(buf)
	return err
}

func (f FQDN) MarshalBinary() ([]byte, error) {
	buf := buffer.New(nil)
	buf.Write8(f.Flags)
	buf.WriteBytes([]byte(f.Name))
	return buf.Data(), nil
}

func fqdnFlagStr(f uint8) (r string) {
	if f&FQDNFlagS != 0 {
		r += "S"
	}
	if f&FQDNFlagO != 0 {
		r += "O"
	}
	if f&FQDNFlagN != 0 {
		r += "N"
	}
	return r
}

func (f FQDN) String() string {
	return f.Name + ":" + fqdnFlagStr(f.Flags)
}
