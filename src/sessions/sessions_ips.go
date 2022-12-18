package sessions

import (
	"github.com/sandrolain/go-utilities/pkg/netutils"
)

func (s *Session) GetAllowedIPs() SessionIPs {
	return s.AllowedIPs
}

func (s *Session) HasAllowedIP(ip SessionIP) (bool, error) {
	for _, value := range s.AllowedIPs {
		ok, err := netutils.NetworkContainsIP(string(value), string(ip))
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func (s *Session) SetAllowedIPs(ips SessionIPs) {
	s.AllowedIPs = ips
}

func (s *Session) AddAllowedIPs(ips SessionIPs) {
	s.AllowedIPs = append(s.AllowedIPs, ips...)
}

func (s *Session) DeleteAllowedIP(ip SessionIP) {
	ips := make(SessionIPs, 0)
	for _, value := range s.AllowedIPs {
		if value != ip {
			ips = append(ips, value)
		}
	}
	s.AllowedIPs = ips
}

func (s *Session) ResetAllowedIPs() {
	s.AllowedIPs = make(SessionIPs, 0)
}
