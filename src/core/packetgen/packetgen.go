// MIT License

// Copyright (c) [2022] [Arriven (https://github.com/Arriven)]

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package packetgen [allows sending customized tcp/udp traffic. Inspired by https://github.com/bilalcaliskan/syn-flood]
package packetgen

import (
	"fmt"

	"github.com/google/gopacket"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

type Packet struct {
	Link        gopacket.LinkLayer
	Network     gopacket.NetworkLayer
	Transport   gopacket.TransportLayer
	Application gopacket.ApplicationLayer
}

type LayerConfig struct {
	Type string
	Data map[string]interface{}
}

type PacketConfig struct {
	Link        LayerConfig
	Network     LayerConfig
	Transport   LayerConfig
	Application LayerConfig
}

func (c PacketConfig) Build() (result Packet, err error) {
	if result.Link, err = BuildLinkLayer(c.Link); err != nil {
		return Packet{}, err
	}

	if result.Network, err = BuildNetworkLayer(c.Network); err != nil {
		return Packet{}, err
	}

	if result.Transport, err = BuildTransportLayer(c.Transport, result.Network); err != nil {
		return Packet{}, err
	}

	if result.Application, err = BuildApplicationLayer(c.Application); err != nil {
		return Packet{}, err
	}

	return result, nil
}

func (p Packet) Serialize(payloadBuf gopacket.SerializeBuffer) (err error) {
	return SerializeLayers(payloadBuf, p.Link, p.Network, p.Transport, p.Application)
}

func (p Packet) IPV4() (ipHeader *ipv4.Header, err error) {
	if p.Network == nil {
		return nil, fmt.Errorf("no network layer present in packet")
	}

	ipHeaderBuf := gopacket.NewSerializeBuffer()
	if err = Serialize(ipHeaderBuf, p.Network); err != nil {
		return nil, err
	}

	return ipv4.ParseHeader(ipHeaderBuf.Bytes())
}

func (p Packet) IPV6() (ipHeader *ipv6.Header, err error) {
	ipHeaderBuf := gopacket.NewSerializeBuffer()

	if err = Serialize(ipHeaderBuf, p.Network); err != nil {
		return nil, err
	}

	return ipv6.ParseHeader(ipHeaderBuf.Bytes())
}
