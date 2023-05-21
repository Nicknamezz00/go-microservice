/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package netutil

import (
	"net"
	"strings"
)

func GetLocalIPv4() (ip string) {
	interfaces, err := net.Interfaces()
	net.InterfaceAddrs()
	if err != nil {
		return
	}
	if len(interfaces) == 2 {
		for _, f := range interfaces {
			if strings.Contains(f.Name, "lo") {
				continue
			}
			addrs, err := f.Addrs()
			if err != nil {
				return
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ipStr := ipnet.IP.String()
						if !strings.Contains(ipStr, ":") && ipStr != "127.0.0.1" {
							ip = ipStr
						}
					}
				}
			}
		}
	}
	for _, f := range interfaces {
		if strings.Contains(f.Name, "lo") {
			continue
		}
		addrs, err := f.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipStr := ipnet.IP.String()
					if !strings.Contains(ipStr, ":") && ipStr != "127.0.0.1" && isIntranetIpv4(ipStr) {
						ip = ipStr
					}
				}
			}
		}
	}
	return
}

func isIntranetIpv4(ip string) bool {
	if strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "169.254.") ||
		strings.HasPrefix(ip, "172.") ||
		strings.HasPrefix(ip, "10.30.") ||
		strings.HasPrefix(ip, "10.31.") {
		return true
	}
	return false
}
