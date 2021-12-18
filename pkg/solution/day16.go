package solution

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Packet struct {
	version int64
	typeID  int64
	value   int64
	packets []Packet
}

func toBinaryString(s string) string {
	var sb strings.Builder
	m := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	for _, r := range s {
		sb.WriteString(m[r])
	}
	return sb.String()
}

func parseLiteral(s string, idx int) (l string, nextIdx int) {
	b := s[idx : idx+1]
	if b == "0" {
		return s[idx+1 : idx+1+4], idx + 1 + 4
	} else {
		var sb strings.Builder
		sb.WriteString(s[idx+1 : idx+1+4])
		rest, nextIdx := parseLiteral(s, idx+1+4)
		sb.WriteString(rest)
		return sb.String(), nextIdx
	}
}

func parsePacket(s string, idx int) (p Packet, nextIdx int) {
	version := s[idx : idx+3]
	typeID := s[idx+3 : idx+3+3]
	p.version, _ = strconv.ParseInt(version, 2, 32)
	p.typeID, _ = strconv.ParseInt(typeID, 2, 32)
	if p.typeID == 4 {

		literal, nIdx := parseLiteral(s, idx+3+3)
		p.value, _ = strconv.ParseInt(literal, 2, 64)
		return p, nIdx
	} else {
		l := s[idx+6 : idx+6+1]
		if l == "0" {
			// total lenght in bits for subpackets
			lb := s[idx+7 : idx+7+15]
			lbInt, _ := strconv.ParseInt(lb, 2, 32)
			endIdx := idx + 7 + 15 + int(lbInt)

			spNextIdx := idx + 7 + 15
			for {
				var sp Packet
				sp, spNextIdx = parsePacket(s, spNextIdx)
				p.packets = append(p.packets, sp)
				if spNextIdx >= endIdx {
					break
				}
			}
			return p, endIdx
		} else {
			// total number of subpackets
			np := s[idx+7 : idx+7+11]
			npInt, _ := strconv.ParseInt(np, 2, 32)
			var sp Packet
			spNextIdx := idx + 7 + 11
			for i := 0; i < int(npInt); i++ {
				sp, spNextIdx = parsePacket(s, spNextIdx)
				p.packets = append(p.packets, sp)
			}
			return p, spNextIdx
		}
	}
}

func (s *Solution) Day16Part1(fn string) (ret int) {
	rawHex := toLines(fn)[0]
	bs := toBinaryString(rawHex)
	p, _ := parsePacket(bs, 0)

	var versionSum func(p Packet) int64
	versionSum = func(p Packet) int64 {
		s := p.version
		for _, sp := range p.packets {
			s += versionSum(sp)
		}
		return s
	}
	return int(versionSum(p))
}

func (s *Solution) Day16Part2(fn string) (ret int) {
	rawHex := toLines(fn)[0]
	bs := toBinaryString(rawHex)
	p, _ := parsePacket(bs, 0)

	var value func(p Packet) int64
	reduce := func(arr []Packet, iv int64, op func(spValue int64, cur int64) int64) int64 {
		for _, s := range arr {
			v := value(s)
			iv = op(v, iv)
		}
		return iv
	}

	value = func(p Packet) int64 {
		if p.typeID == 0 {
			// sum
			return reduce(p.packets, 0, func(spValue, cur int64) int64 {
				return spValue + cur
			})
		} else if p.typeID == 1 {
			// product
			return reduce(p.packets, 1, func(spValue, cur int64) int64 {
				return spValue * cur
			})
		} else if p.typeID == 2 {
			// min
			return reduce(p.packets, math.MaxInt64, func(spValue, cur int64) int64 {
				if spValue < cur {
					return spValue
				} else {
					return cur
				}
			})
		} else if p.typeID == 3 {
			// max
			return reduce(p.packets, 0, func(spValue, cur int64) int64 {
				if spValue > cur {
					return spValue
				} else {
					return cur
				}
			})
		} else if p.typeID == 4 {
			return p.value
		} else if p.typeID == 5 {
			// gt
			if value(p.packets[0]) > value(p.packets[1]) {
				return 1
			} else {
				return 0
			}
		} else if p.typeID == 6 {
			// lt
			if value(p.packets[0]) < value(p.packets[1]) {
				return 1
			} else {
				return 0
			}
		} else if p.typeID == 7 {
			// eq
			if value(p.packets[0]) == value(p.packets[1]) {
				return 1
			} else {
				return 0
			}
		} else {
			panic(fmt.Sprintf("%d not supported", p.typeID))
		}
	}

	return int(value(p))
}
